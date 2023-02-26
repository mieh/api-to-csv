package output

import (
	"os"
	"path/filepath"
)

func ToFile(csv string, outputPath string, fileName string) error {
	filePath := filepath.Join(outputPath, fileName)

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(csv)
	if err != nil {
		return err
	}

	return nil
}
