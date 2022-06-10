package utils

import (
	"crypto/sha1"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/go_server/constant"
	"github.com/go_server/logger"
)

func GetContentType(seeker io.ReadSeeker) (string, error) {

	buff := make([]byte, 512)
	_, err := seeker.Seek(0, io.SeekStart)
	if err != nil {
		return "", err
	}
	bytesRead, err := seeker.Read(buff)
	if err != nil && err != io.EOF {
		return "", err
	}
	buff = buff[:bytesRead]
	return http.DetectContentType(buff), nil
}

func WriteDataToFile(fileHeader *multipart.FileHeader, file multipart.File) error {
	ext := strings.Split(fileHeader.Filename, ".")[1]
	h := sha1.New()
	io.Copy(h, file)
	fname := fmt.Sprintf("%x", h.Sum(nil)) + "." + ext
	wd, err := os.Getwd()
	if err != nil {
		logger.Debug(err.Error())
		return err
	}

	path := filepath.Join(wd, constant.UploadedFile, fname)
	if _, err := os.Stat(filepath.Join(wd, constant.UploadedFile)); os.IsNotExist(err) {
		os.Mkdir(filepath.Join(wd, constant.UploadedFile), os.ModePerm)
	}
	newFile, err := os.Create(path)
	if err != nil {
		logger.Debug(err.Error())
		return err
	}
	defer newFile.Close()
	file.Seek(0, 0)
	_, err = io.Copy(newFile, file)
	if err != nil {
		logger.Debug(err.Error())
		return err
	}
	return nil
}
