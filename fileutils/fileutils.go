package fileutils

import (
	"os"
	"path/filepath"
	"strings"
)

var validExtension = map[string]struct{}{
	".png": {},
	".jpg": {},
	".dng": {},
	".arw": {},
}

func GetCurrentDirectory() (string, error) {
	currentDir, err := os.Getwd()

	if err != nil {
		return "", err
	}

	return currentDir, nil
}

func GetImageFilesPathFromDirectory(path string) ([]string, error) {
	files, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	fileNames := make([]string, 0, len(files))

	for _, file := range files {
		fileExtension := strings.ToLower(filepath.Ext(file.Name()))
		if _, exist := validExtension[fileExtension]; exist && !file.IsDir() {
			fileNames = append(fileNames, filepath.Join(path, file.Name()))
		}
	}

	return fileNames, nil
}
