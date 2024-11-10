package antivirus

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

type Scanner struct {
	address      string
	network      string
	timeout      time.Duration
	useAntivirus bool
}

func NewScanner(address, network string, timeout time.Duration, useAntivirus bool) *Scanner {
	return &Scanner{
		address:      address,
		network:      network,
		timeout:      timeout,
		useAntivirus: useAntivirus,
	}
}

func (s *Scanner) ScanFile(filePath string) error {
	if !s.useAntivirus {
		return nil
	}

	conn, err := net.DialTimeout(s.network, s.address, s.timeout)
	if err != nil {
		return fmt.Errorf("Ошибка соединения сервиса антивируса: %w", err)
	}
	defer conn.Close()

	_, err = fmt.Fprintf(conn, "INSTREAM\n")
	if err != nil {
		return fmt.Errorf("Ошибка отправки INSTREAM команды: %w", err)
	}

	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("Ощибка открытия файла: %w", err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	buf := make([]byte, 1024*64)

	for {
		n, err := reader.Read(buf)
		if n > 0 {
			sizeBuf := []byte{
				byte(n >> 24),
				byte(n >> 16),
				byte(n >> 8),
				byte(n),
			}
			if _, err := conn.Write(sizeBuf); err != nil {
				return fmt.Errorf("Ошибка чтения размера чанка: %w", err)
			}

			if _, err := conn.Write(buf[:n]); err != nil {
				return fmt.Errorf("Ошибка записи даты чанка: %w", err)
			}
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("Ошибка чтения файла: %w", err)
		}
	}

	zero := []byte{0, 0, 0, 0}
	if _, err := conn.Write(zero); err != nil {
		return fmt.Errorf("Ощибка записи нулевого чанка: %w", err)
	}

	response, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		return fmt.Errorf("Ошибка чтения ответа: %w", err)
	}

	if response == "" {
		return fmt.Errorf("Пустой ответ от сервиса антивируса")
	}

	if responseContainsVirus(response) {
		return fmt.Errorf("ФАЙЛ ЗАРАЖЕН: %s", response)
	}

	return nil
}

func responseContainsVirus(response string) bool {
	return response != "ОК\n" && response != "stream: OK\n"
}
