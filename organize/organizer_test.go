package organize

import (
	er "github.com/gsanhuezafuentes/exif-classifier/exif_reader"
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
	"time"
)

var testFileNames = []string{"file1.dng", "file2.jpg", "file3.arw", "file4.arw"}
var creationTime = []string{"2024-02-01", "2024-02-02", "2024-02-02", "2024-02-03"}
var orientation = []string{"Horizontal", "Horizontal", "Vertical", "Unknown"}
var lens = []string{"Sony 55-210", "Sigma 16mm", "Tamrom 150-500mm", "Tamrom 150-500mm"}

func createTestFiles(t *testing.T, tempDir string, filepaths []string) (files []string) {
	for _, file := range filepaths {
		fullpath := filepath.Join(tempDir, file)
		newFile, err := os.Create(fullpath)
		newFile.Close()
		files = append(files, fullpath)
		assert.NoError(t, err)
	}
	return
}

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

	files := createTestFiles(t, tempDir, testFileNames)

	var mockValues = make(map[string]*er.ExifData)
	for i, file := range files {
		date, err := time.Parse(DATE_FORMAT, creationTime[i])
		assert.NoError(t, err)
		mockValues[file] = &er.ExifData{
			CreatedTime: date,
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

func TestDefaultOrganizer_OrganizeImgsByDate_EmptyDate(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "test_organizer")
	defer os.RemoveAll(tempDir)
	assert.NoError(t, err)

	files := createTestFiles(t, tempDir, testFileNames)

	var mockValues = make(map[string]*er.ExifData)
	for _, file := range files {
		mockValues[file] = &er.ExifData{}
	}

	exifReader := &mockExifReader{
		MockResult: mockValues,
	}
	organizer := NewDefaultOrganizer(exifReader, os.Rename)

	organizer.SetImagesPath(files)

	assert.NoError(t, organizer.OrganizeImgsByDate())

	for key, _ := range mockValues {
		_, err := os.Stat(filepath.Join(filepath.Dir(key), filepath.Base(key)))
		if os.IsNotExist(err) {
			t.Fatal("File was not classified.")
		}
	}
}

func TestDefaultOrganizer_OrganizeImgsByOrientation(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "test_organizer")
	defer os.RemoveAll(tempDir)
	assert.NoError(t, err)

	files := createTestFiles(t, tempDir, testFileNames)

	var mockValues = make(map[string]*er.ExifData)
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

	files := createTestFiles(t, tempDir, testFileNames)

	var mockValues = make(map[string]*er.ExifData)
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

func TestDefaultOrganizer_All(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "test_organizer")
	defer os.RemoveAll(tempDir)
	assert.NoError(t, err)

	files := createTestFiles(t, tempDir, testFileNames)

	var mockValues = make(map[string]*er.ExifData)
	for i, file := range files {
		date, err := time.Parse(DATE_FORMAT, creationTime[i])
		assert.NoError(t, err)
		mockValues[file] = &er.ExifData{
			CreatedTime: date,
			Orientation: orientation[i],
			LensModel:   lens[i],
		}
	}

	exifReader := &mockExifReader{
		MockResult: mockValues,
	}
	organizer := NewDefaultOrganizer(exifReader, os.Rename)

	organizer.SetImagesPath(files)

	assert.NoError(t, organizer.OrganizeImgsByDate())
	assert.NoError(t, organizer.OrganizeImgsByOrientation())
	assert.NoError(t, organizer.OrganizeImgsByLens())

	for key, exifData := range mockValues {
		date := exifData.CreatedTime.Format(DATE_FORMAT)
		_, err := os.Stat(filepath.Join(filepath.Dir(key), date, exifData.Orientation, exifData.LensModel, filepath.Base(key)))
		if os.IsNotExist(err) {
			t.Fatal("File was not classified.")
		}
	}
}
