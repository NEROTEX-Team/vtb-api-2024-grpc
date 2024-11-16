package antivirus

import (
	"bytes"
	"errors"
	"io"
	"net"
	"os"
	"strings"
	"testing"
	"time"
)

type mockConn struct {
	io.Reader
	io.Writer
	Closed bool
}

func (m *mockConn) Close() error {
	m.Closed = true
	return nil
}

func (m *mockConn) LocalAddr() net.Addr                { return nil }
func (m *mockConn) RemoteAddr() net.Addr               { return nil }
func (m *mockConn) SetDeadline(t time.Time) error      { return nil }
func (m *mockConn) SetReadDeadline(t time.Time) error  { return nil }
func (m *mockConn) SetWriteDeadline(t time.Time) error { return nil }

func TestScanFile_AntivirusEnabled_CleanFile(t *testing.T) {
	scanner := &Scanner{
		address:      "localhost:3310",
		network:      "tcp",
		timeout:      10 * time.Second,
		useAntivirus: true,
		dialer: func(network, address string, timeout time.Duration) (net.Conn, error) {
			response := "OK\n"
			mockConn := &mockConn{
				Reader: strings.NewReader(response),
				Writer: &bytes.Buffer{},
			}
			return mockConn, nil
		},
	}

	tmpFile, err := os.CreateTemp("", "cleanfile")
	if err != nil {
		t.Fatalf("Не удалось создать временный файл: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	err = scanner.ScanFile(tmpFile.Name())

	if err != nil {
		t.Errorf("Ожидалась отсутствие ошибки, получено: %v", err)
	}
}

func TestScanFile_AntivirusEnabled_InfectedFile(t *testing.T) {
	scanner := &Scanner{
		address:      "localhost:3310",
		network:      "tcp",
		timeout:      10 * time.Second,
		useAntivirus: true,
	}

	scanner.dialer = func(network, address string, timeout time.Duration) (net.Conn, error) {
		response := "stream: Eicar-Test-Signature FOUND\n"
		mockConn := &mockConn{
			Reader: strings.NewReader(response),
			Writer: &bytes.Buffer{},
		}
		return mockConn, nil
	}

	tmpFile, err := os.CreateTemp("", "infectedfile")
	if err != nil {
		t.Fatalf("не удалось создать временный файл: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	err = scanner.ScanFile(tmpFile.Name())

	if err == nil {
		t.Errorf("ожидалась ошибка для зараженного файла, получено nil")
	} else if !errors.Is(err, ErrFileInfected) {
		t.Errorf("ожидалась ошибка ErrFileInfected, получено: %v", err)
	}
}

func TestScanFile_AntivirusDisabled(t *testing.T) {
	scanner := &Scanner{
		useAntivirus: false,
	}

	err := scanner.ScanFile("anyfile")

	if err != nil {
		t.Errorf("ожидалось отсутствие ошибки при отключенном антивирусе, получено: %v", err)
	}
}
