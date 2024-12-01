package services

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

var EmailNotFound = errors.New("Email не найден")

type JsonStorage struct {
	filePath string
}

func newJsonStorageService(filePath string) *JsonStorage {
	if _, err := os.Stat(filePath); errors.Is(err, os.ErrNotExist) {
		if _, err := os.Create(filePath); err != nil {
			log.Printf("os.Create: %v", err)
		}
	}

	return &JsonStorage{
		filePath: filePath,
	}
}

func (service *JsonStorage) SaveEmailHash(ctx context.Context, email string, hash string) error {
	var globalError error
	file, err := os.OpenFile(service.filePath, os.O_RDWR, 0666)
	if err != nil {
		return fmt.Errorf("os.Open: %w", err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			globalError = fmt.Errorf("file.Close: %w", err)
		}
	}()

	m := make(map[string]string)
	if err := json.NewDecoder(file).Decode(&m); err != nil && !errors.Is(err, io.EOF) {
		return fmt.Errorf("json.NewDecoder.Decode: %w", err)
	}

	m[hash] = email

	if err := file.Truncate(0); err != nil {
		return fmt.Errorf("file.Truncate: %w", err)
	}

	if _, err := file.Seek(0, 0); err != nil {
		return fmt.Errorf("file.Seek: %w", err)
	}

	if err := json.NewEncoder(file).Encode(m); err != nil {
		return fmt.Errorf("json.NewEncoder.Encode %w", err)
	}

	return globalError
}

func (service *JsonStorage) EmailByHash(ctx context.Context, hash string) (string, error) {
	var globalError error
	file, err := os.Open(service.filePath)
	if err != nil {
		return "", fmt.Errorf("os.Open: %w", err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			globalError = fmt.Errorf("file.Close: %w", err)
		}
	}()

	var m map[string]string
	if err := json.NewDecoder(file).Decode(&m); err != nil && !errors.Is(err, io.EOF) {
		return "", fmt.Errorf("json.NewDecoder.Decode: %w", err)
	}

	email, ok := m[hash]
	if !ok {
		return "", EmailNotFound
	}

	return email, globalError
}

func (service *JsonStorage) DeleteHash(ctx context.Context, hash string) error {
	var globalError error
	file, err := os.OpenFile(service.filePath, os.O_RDWR, 0666)
	if err != nil {
		return fmt.Errorf("os.Open: %w", err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			globalError = fmt.Errorf("file.Close: %w", err)
		}
	}()

	var m map[string]string
	if err := json.NewDecoder(file).Decode(&m); err != nil {
		return fmt.Errorf("json.NewDecoder.Decode: %w", err)
	}

	if _, ok := m[hash]; !ok {
		return EmailNotFound
	}

	delete(m, hash)

	if err := file.Truncate(0); err != nil {
		return fmt.Errorf("file.Truncate: %w", err)
	}

	if _, err := file.Seek(0, 0); err != nil {
		return fmt.Errorf("file.Seek: %w", err)
	}

	if err := json.NewEncoder(file).Encode(m); err != nil {
		return fmt.Errorf("json.NewEncoder.Encode: %w", err)
	}

	return globalError
}
