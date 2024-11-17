package antivirus

import (
	"io"
	"net/http"
	"os"
)

func (s *Scanner) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !s.useAntivirus {
			next.ServeHTTP(w, r)
			return
		}

		file, _, err := r.FormFile("file")
		if err != nil {
			http.Error(w, "не удалось получить файл для сканирования", http.StatusBadRequest)
			return
		}
		defer file.Close()

		tmpFile, err := os.CreateTemp("", "uploaded-file-*")
		if err != nil {
			http.Error(w, "не удалось создать временный файл", http.StatusInternalServerError)
			return
		}
		defer os.Remove(tmpFile.Name())
		defer tmpFile.Close()

		if _, err := io.Copy(tmpFile, file); err != nil {
			http.Error(w, "не удалось сохранить файл для сканирования", http.StatusInternalServerError)
			return
		}

		if err := s.ScanFile(tmpFile.Name()); err != nil {
			http.Error(w, "файл заражен или сканирование не удалось: "+err.Error(), http.StatusBadRequest)
			return
		}

		if _, err := tmpFile.Seek(0, io.SeekStart); err != nil {
			http.Error(w, "не удалось обработать файл", http.StatusInternalServerError)
			return
		}
		r.Body = io.NopCloser(tmpFile)

		next.ServeHTTP(w, r)
	})
}
