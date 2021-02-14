package file

import (
	"github.com/tombenke/axon-go/common/log"
	"io/ioutil"
	"os"
)

// LoadFile loads []byte content from a file
func LoadFile(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return []byte(""), err
	}
	defer file.Close()
	content, err := ioutil.ReadAll(file)
	if err != nil {
		return []byte(""), err
	}

	log.Logger.Debugf("Load file from '%s'", path)
	return content, nil
}
