package gfile

import (
	"bufio"
	"io"
	"os"
)

func ReadFileToBytes(filePath string) ([]byte, error) {
	return os.ReadFile(filePath)
}

func ReadFileToString(filePath string) (string, error) {
	content, err := os.ReadFile(filePath)
	return string(content), err
}

func ReadLines(filePath string, callback func(line string)) {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}

		// callback
		callback(string(line))
	}
}

func WriteBytesToFile(filePath string, data []byte) error {
	return os.WriteFile(filePath, data, os.ModePerm)
}

func WriteStringToFile(filePath string, data string) error {
	return os.WriteFile(filePath, []byte(data), os.ModePerm)
}
