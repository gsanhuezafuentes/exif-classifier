package organize

import (
	"errors"
	"fmt"
	"github.com/gsanhuezafuentes/exif-classifier/exif_reader"
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
	if err := r.checkFilesExist(); err != nil {
		return err
	}
	if err := r.scanExifData(); err != nil {
		return err
	}

	for _, image := range r.cache {
		creationTime := image.CreatedTime
		date := creationTime.Format(DATE_FORMAT)
		newPath := filepath.Join(filepath.Dir(image.path), date, filepath.Base(image.path))
		fmt.Printf("Move files from [%s] to [%s]\n", image.path, newPath)
		image.path = newPath
		fmt.Println(image.path)
	}
	return nil
}

func (r *DefaultOrganizer) OrganizeImgsByLens() error {
	if err := r.checkFilesExist(); err != nil {
		return err
	}
	if err := r.scanExifData(); err != nil {
		return err
	}
	for _, image := range r.cache {
		lens := image.LensModel
		newPath := filepath.Join(filepath.Dir(image.path), lens, filepath.Base(image.path))
		fmt.Printf("Move files from [%s] to [%s]\n", image.path, newPath)
		image.path = newPath
	}
	return nil
}

func (r *DefaultOrganizer) OrganizeImgsByOrientation() error {
	if err := r.checkFilesExist(); err != nil {
		return err
	}
	if err := r.scanExifData(); err != nil {
		return err
	}
	for _, image := range r.cache {
		orientation := image.Orientation
		newPath := filepath.Join(filepath.Dir(image.path), orientation, filepath.Base(image.path))
		fmt.Printf("Move files from [%s] to [%s]\n", image.path, newPath)
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
