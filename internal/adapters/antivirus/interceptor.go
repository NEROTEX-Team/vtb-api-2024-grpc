package antivirus

import (
	"context"
	"os"
	"reflect"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Scanner) UnaryServerInterceptor(fileFieldName string) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		if !s.useAntivirus {
			return handler(ctx, req)
		}

		reqValue := reflect.ValueOf(req)
		if reqValue.Kind() != reflect.Ptr {
			return nil, status.Error(codes.InvalidArgument, "Запрос должен быть указателем")
		}
		reqElem := reqValue.Elem()
		fileField := reqElem.FieldByName(fileFieldName)
		if !fileField.IsValid() {
			return nil, status.Errorf(codes.InvalidArgument, "Поле %s не найдено в запросе", fileFieldName)
		}
		if fileField.Kind() != reflect.Slice || fileField.Type().Elem().Kind() != reflect.Uint8 {
			return nil, status.Errorf(codes.InvalidArgument, "Поле %s должно быть []byte", fileFieldName)
		}

		fileData := fileField.Bytes()

		tmpFile, err := os.CreateTemp("", "uploaded-file-*")
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Не удалось создать временный файл: %v", err)
		}
		defer os.Remove(tmpFile.Name())
		defer tmpFile.Close()

		if _, err := tmpFile.Write(fileData); err != nil {
			return nil, status.Errorf(codes.Internal, "Не удалось записать данные файла: %v", err)
		}

		if err := s.ScanFile(tmpFile.Name()); err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "Файл заражен или сканирование не удалось: %v", err)
		}

		return handler(ctx, req)
	}
}
