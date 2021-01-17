package config

import (
	"io/ioutil"
	"os"
)

// LoadFile loads []byte content from a file
func LoadFile(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	content, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	return content, nil
}
