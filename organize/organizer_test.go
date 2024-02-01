package organize

import (
	er "github.com/gsanhuezafuentes/exif-classifier/exif_reader"
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
	"time"
)

type mockExifReader struct {
	MockResult map[string]*er.ExifData
}

func (r *mockExifReader) Read(filepath string) (*er.ExifData, error) {
	return r.MockResult[filepath], nil
}

func TestDefaultOrganizer_OrganizeImgsByDate(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "test_organizer")
	defer os.RemoveAll(tempDir)
	assert.NoError(t, err)

	var files []string
	for _, file := range []string{"file1.dng", "file2.jpg", "file3.arw"} {
		fullpath := filepath.Join(tempDir, file)
		newFile, err := os.Create(fullpath)
		newFile.Close()
		files = append(files, fullpath)
		assert.NoError(t, err)
	}

	var mockValues = make(map[string]*er.ExifData)
	start := time.Date(2024, 2, 1, 18, 11, 0, 0, time.UTC)
	for i, file := range files {
		mockValues[file] = &er.ExifData{
			CreatedTime: start.Add(time.Duration(i * int(24*time.Hour))),
		}
	}

	exifReader := &mockExifReader{
		MockResult: mockValues,
	}
	organizer := NewDefaultOrganizer(exifReader, os.Rename)

	organizer.SetImagesPath(files)

	assert.NoError(t, organizer.OrganizeImgsByDate())

	for key, exifData := range mockValues {
		folder := exifData.CreatedTime.Format(DATE_FORMAT)
		_, err := os.Stat(filepath.Join(filepath.Dir(key), folder, filepath.Base(key)))
		if os.IsNotExist(err) {
			t.Fatal("File was not classified.")
		}
	}
}

func TestDefaultOrganizer_OrganizeImgsByOrientation(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "test_organizer")
	defer os.RemoveAll(tempDir)
	assert.NoError(t, err)

	var files []string
	for _, file := range []string{"file1.dng", "file2.jpg", "file3.arw"} {
		fullpath := filepath.Join(tempDir, file)
		newFile, err := os.Create(fullpath)
		newFile.Close()
		files = append(files, fullpath)
		assert.NoError(t, err)
	}

	var mockValues = make(map[string]*er.ExifData)
	orientation := []string{"Horizontal", "Vertical", "Unknown"}
	for i, file := range files {
		mockValues[file] = &er.ExifData{
			Orientation: orientation[i],
		}
	}

	exifReader := &mockExifReader{
		MockResult: mockValues,
	}
	organizer := NewDefaultOrganizer(exifReader, os.Rename)

	organizer.SetImagesPath(files)

	assert.NoError(t, organizer.OrganizeImgsByOrientation())

	for key, exifData := range mockValues {
		folder := exifData.Orientation
		_, err := os.Stat(filepath.Join(filepath.Dir(key), folder, filepath.Base(key)))
		if os.IsNotExist(err) {
			t.Fatal("File was not classified.")
		}
	}
}

func TestDefaultOrganizer_OrganizeImgsByLens(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "test_organizer")
	defer os.RemoveAll(tempDir)
	assert.NoError(t, err)

	var files []string
	for _, file := range []string{"file1.dng", "file2.jpg", "file3.arw"} {
		fullpath := filepath.Join(tempDir, file)
		newFile, err := os.Create(fullpath)
		newFile.Close()
		files = append(files, fullpath)
		assert.NoError(t, err)
	}

	var mockValues = make(map[string]*er.ExifData)
	lens := []string{"Sony 55-210", "Sigma 16mm", "Tamrom 150-500mm"}
	for i, file := range files {
		mockValues[file] = &er.ExifData{
			LensModel: lens[i],
		}
	}

	exifReader := &mockExifReader{
		MockResult: mockValues,
	}
	organizer := NewDefaultOrganizer(exifReader, os.Rename)

	organizer.SetImagesPath(files)

	assert.NoError(t, organizer.OrganizeImgsByLens())

	for key, exifData := range mockValues {
		folder := exifData.LensModel
		_, err := os.Stat(filepath.Join(filepath.Dir(key), folder, filepath.Base(key)))
		if os.IsNotExist(err) {
			t.Fatal("File was not classified.")
		}
	}
}
