package std

import (
	"archive/zip"
	"encoding/base64"
	"fmt"
	"io"
	"os"
	"os/user"
	"path/filepath"
	"time"
)

func DecodeBase64Encoded(base64encoded []byte) ([]byte, error) {
	return base64.StdEncoding.DecodeString(string(base64encoded))
}

func GetCurrentTimeInUnixMilli() int64 {
	currentTime := time.Now()
	return currentTime.UnixMilli()
}

func ConvertIntegerToBoolean(integer int) bool {
	return integer > 0
}

func Unzip(source string, destination string) error {

	zipReader, _ := zip.OpenReader(source)
	for _, file := range zipReader.Reader.File {
		zippedFile, err := file.Open()
		if err != nil {
			return err
		}
		defer zippedFile.Close()

		extractedFilePath := filepath.Join(
			destination,
			file.Name,
		)

		if file.FileInfo().IsDir() {
			os.MkdirAll(extractedFilePath, file.Mode())
		} else {
			outputFile, err := os.OpenFile(
				extractedFilePath,
				os.O_WRONLY|os.O_CREATE|os.O_TRUNC,
				file.Mode(),
			)
			if err != nil {
				return err
			}
			defer outputFile.Close()

			_, err = io.Copy(outputFile, zippedFile)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func WriteCurrentPID(filePath string) error {
	pid := os.Getpid()

	pidBytes := []byte(fmt.Sprintf("%d", pid))

	err := os.WriteFile(filePath, pidBytes, 0777)

	return err
}

func CopyFile(fromLocation string, toLocation string) error {
	var err error

	original, err := os.Open(fromLocation)
	if err != nil {
		return err
	}

	new, err := os.Create(toLocation)
	if err != nil {
		original.Close()
		return err
	}

	_, err = io.Copy(new, original)

	original.Close()
	new.Close()
	return err
}

func CurrentUsername() (string, error) {

	var username string

	currentUser, err := user.Current()
	if err == nil {
		username = currentUser.Username
	}

	return username, err
}

func ExecutableCurrentDirectory() (string, error) {
	currentExecutable, err := os.Executable()
	if err == nil {
		return filepath.Dir(currentExecutable), err
	}
	return EmptyString, nil
}
