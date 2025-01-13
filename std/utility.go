package std

import (
	"archive/zip"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"os"
	"os/user"
	"path/filepath"
	"time"

	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
)

func DecodeBase64(encoded []byte) ([]byte, error) {
	decoded := make([]byte, base64.StdEncoding.DecodedLen(len(encoded)))
	n, err := base64.StdEncoding.Decode(decoded, encoded)
	return decoded[:n], err
}

func GetCurrentTimeInUnixMilli() int64 {
	return time.Now().UnixMilli()
}

func GetCurrentTimeInUnixNano() int64 {
	return time.Now().UnixNano()
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

	err := os.WriteFile(filePath, pidBytes, 0755)

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
	return "", nil
}

func Ping(ip string) (bool, error) {

	c, err := icmp.ListenPacket("ip4:icmp", "0.0.0.0")
	if err != nil {
		return false, err
	}
	defer c.Close()

	wm := icmp.Message{
		Type: ipv4.ICMPTypeEcho, Code: 0,
		Body: &icmp.Echo{
			ID: os.Getpid() & 0xffff, Seq: 1,
			Data: []byte(""),
		},
	}

	wb, err := wm.Marshal(nil)
	if err != nil {
		return false, err
	}

	if _, err := c.WriteTo(wb, &net.IPAddr{IP: net.ParseIP(ip)}); err != nil {
		return false, err
	}

	rb := make([]byte, 1500)
	err = c.SetReadDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		return false, err
	}

	n, _, err := c.ReadFrom(rb)
	if err != nil {
		return false, err
	}

	rm, err := icmp.ParseMessage(ipv4.ICMPTypeEchoReply.Protocol(), rb[:n])
	if err != nil {
		return false, err
	}

	if rm.Type == ipv4.ICMPTypeEchoReply {
		return true, nil
	}

	return false, nil
}

// UUID generates a random UUID of the specified length
// The length should be a multiple of 2
func UUID(length int) (string, error) {
	bytes := make([]byte, length/2)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

// ToHashCode takes an interface{} value that must be marshable into JSON,
// marshals it into JSON, and then computes the SHA-256 hash of the resulting
// JSON byte slice. It returns the hash as a byte slice and any error encountered
// during the marshaling process.
//
// Parameters:
//   - value: The input value of type interface{} to be hashed.
//
// Returns:
//   - []byte: The SHA-256 hash of the JSON representation of the input value.
//   - error: An error if the JSON marshaling fails, otherwise nil.
func ToHashCode(value interface{}) ([]byte, error) {

	data, err := json.Marshal(value)
	if err != nil {
		return nil, err
	}

	hash := sha256.Sum256(data)

	return hash[:], nil
}
