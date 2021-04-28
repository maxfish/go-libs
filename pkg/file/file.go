package file

import (
	"encoding/json"
	"io"
	"io/fs"
	"path/filepath"
)

func Read(fileSys fs.FS, filePath string, fileName string) ([]byte, error) {
	fullPath := filepath.Join(filePath, fileName)
	file, err := fileSys.Open(fullPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	fileContent, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return fileContent, nil
}

func ReadStructFromJSON(fileSys fs.FS, filePath string, fileName string, data interface{}) error {
	content, err := Read(fileSys, filePath, fileName)
	if err != nil {
		return err
	}

	err = json.Unmarshal(content, data)
	if err != nil {
		return err
	}

	return nil
}

