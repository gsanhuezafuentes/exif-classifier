package organize

import (
	"errors"
	"github.com/gsanhuezafuentes/exif-classifier/exif_reader"
	"os"
	"path/filepath"
)

const DATE_FORMAT = "2006-01-02"

type Organizer interface {
	// SetImagesPath save the images path to use it with the other methods.
	SetImagesPath([]string)
	// OrganizeImgsByDate move the images to specific folder based in the date.
	OrganizeImgsByDate() error
	// OrganizeImgsByLens move the images to folder named with the name of the lens used.
	OrganizeImgsByLens() error
	// OrganizeImgsByOrientation move the images to folder named with the orientation of the image.
	OrganizeImgsByOrientation() error
}

type DefaultOrganizer struct {
	ExifReader exif_reader.ExifReader
	cache      []*exifDataWithPath
	files      []string
	move       func(string, string) error
}

type exifDataWithPath struct {
	*exif_reader.ExifData
	path string
}

func NewDefaultOrganizer(reader exif_reader.ExifReader, move func(oldpath, newpath string) error) Organizer {
	organizer := &DefaultOrganizer{ExifReader: reader, move: move}
	return organizer
}

func (r *DefaultOrganizer) SetImagesPath(files []string) {
	r.files = files
}

func (r *DefaultOrganizer) OrganizeImgsByDate() error {
	return r.organizeFiles(func(exif *exifDataWithPath) string {
		creationTime := exif.CreatedTime
		if creationTime.IsZero() {
			return ""
		}
		date := creationTime.Format(DATE_FORMAT)
		return date
	})
}

func (r *DefaultOrganizer) OrganizeImgsByLens() error {
	return r.organizeFiles(func(exif *exifDataWithPath) string {
		return exif.LensModel
	})
}

func (r *DefaultOrganizer) OrganizeImgsByOrientation() error {
	return r.organizeFiles(func(exif *exifDataWithPath) string {
		return exif.Orientation
	})
}

func (r *DefaultOrganizer) organizeFiles(extractor func(exifData *exifDataWithPath) string) error {
	if err := r.checkFilesExist(); err != nil {
		return err
	}
	if err := r.scanExifData(); err != nil {
		return err
	}
	for _, image := range r.cache {
		property := extractor(image)
		newPath := filepath.Join(filepath.Dir(image.path), property, filepath.Base(image.path))
		if err := createFolderIfNotExist(newPath); err != nil {
			return err
		}
		if err := os.Rename(image.path, newPath); err != nil {
			return err
		}
		image.path = newPath
	}
	return nil
}

func (r *DefaultOrganizer) scanExifData() error {
	if len(r.cache) != 0 {
		return nil
	}
	exifData, err := getExifDataOfImages(r.files, r.ExifReader)
	if err != nil {
		return err
	}
	r.cache = exifData
	return nil
}

func (r *DefaultOrganizer) checkFilesExist() error {
	if len(r.files) == 0 {
		return errors.New("files has not been initialized. Use SetImagesPath to add the images to modify")
	}
	return nil
}

func createFolderIfNotExist(destPath string) error {
	destDir := filepath.Dir(destPath)
	if _, err := os.Stat(destDir); os.IsNotExist(err) {
		if err := os.MkdirAll(destDir, os.ModePerm); err != nil {
			return err
		}
	}
	return nil
}

func getExifDataOfImages(files []string, reader exif_reader.ExifReader) ([]*exifDataWithPath, error) {
	exifData := make([]*exifDataWithPath, 0, len(files))
	for _, file := range files {
		data, err := reader.Read(file)
		if err != nil {
			return nil, err
		}
		exifData = append(exifData, &exifDataWithPath{ExifData: data, path: file})
	}

	return exifData, nil
}
