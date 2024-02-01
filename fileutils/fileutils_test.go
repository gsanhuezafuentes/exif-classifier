package fileutils

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
)

func TestGetCurrentDirectory(t *testing.T) {
	currentDir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	result, err := GetCurrentDirectory()
	if result != currentDir {
		t.Fatalf("The current directory are different. Expected [%s] get [%s]", currentDir, result)
	}
}

func TestGetImageFilesPathFromDirectory(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "test_get_images")
	defer func() {
		err := os.RemoveAll(tempDir)
		if err != nil {
			fmt.Println("Error removing temporary directory:", err)
		}
	}()
	assert.NoError(t, err)

	files := []string{"file1.dng", "file2.jpg", "file3.arw", "file4.ARW", "file5.png", "file6.txt"}
	for _, file := range files {
		newFile, err := os.Create(filepath.Join(tempDir, file))
		newFile.Close()
		assert.NoError(t, err)
	}

	resultFiles, err := GetImageFilesPathFromDirectory(tempDir)
	assert.NoError(t, err)

	assert.Len(t, resultFiles, 5, "Expected 5 files but got %d. %s", len(resultFiles), resultFiles)
}
