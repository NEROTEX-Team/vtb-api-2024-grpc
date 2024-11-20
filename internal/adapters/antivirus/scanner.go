package antivirus

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
	"time"
)

type dialerFunc func(network, address string, timeout time.Duration) (net.Conn, error)

var ErrFileInfected = errors.New("file is infected")

type Scanner struct {
	address      string
	network      string
	timeout      time.Duration
	useAntivirus bool
	dialer       dialerFunc
}

func NewScanner(address, network string, timeout time.Duration, useAntivirus bool) *Scanner {
	return &Scanner{
		address:      address,
		network:      network,
		timeout:      timeout,
		useAntivirus: useAntivirus,
		dialer:       net.DialTimeout,
	}
}

func (s *Scanner) UseAntivirus() bool {
	return s.useAntivirus
}

func (s *Scanner) ScanFile(filePath string) error {
	if !s.useAntivirus {
		return nil
	}

	conn, err := s.dialer(s.network, s.address, s.timeout)
	if err != nil {
		return fmt.Errorf("ошибка соединения: %w", err)
	}
	defer conn.Close()

	_, err = fmt.Fprintf(conn, "INSTREAM\n")
	if err != nil {
		return fmt.Errorf("ошибка отправки запроса: %w", err)
	}

	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("ошибка открытия файла: %w", err)
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
				return fmt.Errorf("ошибка чтения размера чанка: %w", err)
			}

			if _, err := conn.Write(buf[:n]); err != nil {
				return fmt.Errorf("ошибка записи даты чанка: %w", err)
			}
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("ошибка чтения файла: %w", err)
		}
	}

	zero := []byte{0, 0, 0, 0}
	if _, err := conn.Write(zero); err != nil {
		return fmt.Errorf("ошибка записи нулевого чанка: %w", err)
	}

	response, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		return fmt.Errorf("ошибка чтения ответа: %w", err)
	}

	if response == "" {
		return fmt.Errorf("пустой ответ от антивируса")
	}
	if responseContainsVirus(response) {
		return fmt.Errorf("%w: %s", ErrFileInfected, response)
	}

	return nil
}

func responseContainsVirus(response string) bool {
	response = strings.TrimSpace(response)
	return response != "OK" && response != "stream: OK"
}
