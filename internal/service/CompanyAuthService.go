package service

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
)

var allowedTypes = map[string]bool{
	".jpg":  true,
	".png":  true,
	".jpeg": true,
}

func SaveImage(file *multipart.FileHeader) (string, error) {
	if file.Size > 5*1024*1024 {
		return "", errors.New("文件大小超过5MB限制")
	}

	ext := strings.ToLower(filepath.Ext(file.Filename))
	if !allowedTypes[ext] {
		return "", errors.New("不支持的文件类型")
	}

	uuidStr := uuid.New().String()
	subDir := filepath.Join("uploads", uuidStr[:2])
	err := os.MkdirAll(subDir, os.ModePerm)
	if err != nil {
		return "", err
	}

	savePath := filepath.Join(subDir, uuidStr+ext)
	if err := SaveUploadedFile(file, savePath); err != nil {
		return "", fmt.Errorf("保存文件失败: %v", err)
	}
	return filepath.Join(uuidStr[:2], uuidStr+ext), nil
}

func SaveUploadedFile(file *multipart.FileHeader, dst string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer func(src multipart.File) {
		err := src.Close()
		if err != nil {
			fmt.Println("Error closing file:", err)
		}
	}(src)

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer func(out *os.File) {
		err := out.Close()
		if err != nil {
			fmt.Println("Error closing file:", err)
		}
	}(out)

	_, err = io.Copy(out, src)
	return err
}
