package exif_reader

import (
	"fmt"
	"github.com/dsoprea/go-exif/v3"
	exifcommon "github.com/dsoprea/go-exif/v3/common"
	"io"
	"os"
	"time"
)

type ExifData struct {
	CreatedTime time.Time
	Orientation string
	LensModel   string
}

type ExifReader interface {
	Read(filename string) (ExifData, error)
}

type DefaultExifReader struct{}

// Read reads the exif data and return the values of specific tags
func (r DefaultExifReader) Read(imagePath string) (*ExifData, error) {
	file, err := os.Open(imagePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	rawExif, err := exif.SearchAndExtractExif(data)
	if err != nil {
		return nil, err
	}

	im, err := exifcommon.NewIfdMappingWithStandard()
	if err != nil {
		return nil, err
	}

	ti := exif.NewTagIndex()

	_, index, err := exif.Collect(im, ti, rawExif)
	if err != nil {
		return nil, err
	}

	results, err := index.RootIfd.FindTagWithName("DateTime")
	if err != nil {
		return nil, err
	}

	ite := results[0]

	valueRaw, err := ite.Value()
	if err != nil {
		return nil, err
	}

	fmt.Printf("Datetime %+v\n", valueRaw)

	results, err = index.RootIfd.FindTagWithName("Orientation")
	if err != nil {
		return nil, err
	}

	ite = results[0]

	valueRaw, err = ite.Value()
	if err != nil {
		return nil, err
	}

	fmt.Printf("Orientation %+v\n", valueRaw)

	results, err = index.Lookup["IFD/Exif"].FindTagWithName("LensModel")

	ite = results[0]

	valueRaw, err = ite.Value()
	if err != nil {
		return nil, err
	}

	fmt.Printf("Orientation %+v\n", valueRaw)

	return nil, nil
}

// PrintExif Print the tags in the image file
func (r DefaultExifReader) PrintExif(imagePath string) error {
	file, err := os.Open(imagePath)
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	rawExif, err := exif.SearchAndExtractExif(data)
	if err != nil {
		return err
	}

	im, err := exifcommon.NewIfdMappingWithStandard()
	if err != nil {
		return err
	}

	ti := exif.NewTagIndex()

	_, index, err := exif.Collect(im, ti, rawExif)
	if err != nil {
		return err
	}

	rootIfd := index.RootIfd
	cb := func(ifd *exif.Ifd, ite *exif.IfdTagEntry) error {
		valueRaw, err := ite.Value()
		if err != nil {
			return err
		}
		fmt.Printf("IFD-PATH=[%s] ID=(0x%04x) NAME=[%s] COUNT=(%d) TYPE=[%s] VALUE=[%v]\n", ite.IfdPath(), ite.TagId(), ite.TagName(), ite.UnitCount(), ite.TagType(), valueRaw)
		return nil
	}

	err = rootIfd.EnumerateTagsRecursively(cb)

	return nil
}
