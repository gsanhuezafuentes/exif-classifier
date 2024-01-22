package exif_reader

import (
	"fmt"
	"github.com/dsoprea/go-exif/v3"
	exifcommon "github.com/dsoprea/go-exif/v3/common"
	"io"
	"os"
	"time"
)

const (
	DATETIME    = "DateTime"
	ORIENTATION = "Orientation"
	LENS_MODEL  = "LensModel"
)

const (
	HORIZONTAL = "Horizontal"
	VERTICAL   = "Vertical"
	UNKNOWN    = "Unknown"
)

const DATEFORMAT = "2006:01:02 15:04:05"

type ExifData struct {
	CreatedTime time.Time
	Orientation string
	LensModel   string
}

type ExifReader interface {
	Read(filename string) (*ExifData, error)
}

type ExifPrinter interface {
	PrintExif(imagePath string, out io.Writer) error
}

type DefaultExifReader struct{}

func New() DefaultExifReader {
	return DefaultExifReader{}
}

// Read reads the exif data and return the values of specific tags
func (r DefaultExifReader) Read(imagePath string) (*ExifData, error) {
	index, err := getIndex(imagePath)
	if err != nil {
		return nil, err
	}

	date, err := getTagValue(index.RootIfd, DATETIME)
	if err != nil {
		return nil, err
	}

	dateTime, err := time.Parse(DATEFORMAT, date.(string))
	if err != nil {
		return nil, err
	}

	orientation, err := getTagValue(index.RootIfd, ORIENTATION)
	if err != nil {
		return nil, err
	}

	lens, err := getTagValue(index.Lookup["IFD/Exif"], LENS_MODEL)
	if err != nil {
		return nil, err
	}

	return &ExifData{
		CreatedTime: dateTime,
		Orientation: humanizeOrientation(orientation.([]uint16)[0]),
		LensModel:   lens.(string),
	}, nil
}

// humanizeOrientation map the integer extracted from Orientation Exif property to a string representation (Vertical, Horizontal, Unknown).
func humanizeOrientation(orientation uint16) string {
	switch orientation {
	case 1, 2, 3, 4:
		return HORIZONTAL
	case 5, 6, 7, 8:
		return VERTICAL
	default:
		return UNKNOWN
	}
}

// PrintExif Print tags in the image file in the out buffer.
func (r DefaultExifReader) PrintExif(imagePath string, out io.Writer) error {
	index, err := getIndex(imagePath)
	if err != nil {
		return err
	}

	rootIfd := index.RootIfd
	cb := func(ifd *exif.Ifd, ite *exif.IfdTagEntry) error {
		return printTagEntry(ite, out)
	}

	err = rootIfd.EnumerateTagsRecursively(cb)

	return nil
}

// getTagValue find in the ifd the tag by name and return the associated value.
func getTagValue(ifd *exif.Ifd, name string) (any, error) {
	results, err := ifd.FindTagWithName(name)

	ite := results[0]

	valueRaw, err := ite.Value()
	if err != nil {
		return nil, err
	}

	return valueRaw, nil
}

// getIndex Scan the image and return the index which is an object used to query exif data.
func getIndex(imagePath string) (*exif.IfdIndex, error) {
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

	return &index, nil
}

// printTagEntry print the tag entry in out buffer.
func printTagEntry(ite *exif.IfdTagEntry, out io.Writer) error {
	valueRaw, err := ite.Value()
	if err != nil {
		return err
	}
	fmt.Fprintf(
		out,
		"IFD-PATH=[%s] ID=(0x%04x) NAME=[%s] COUNT=(%d) TYPE=[%s] VALUE=[%v]\n",
		ite.IfdPath(),
		ite.TagId(),
		ite.TagName(),
		ite.UnitCount(),
		ite.TagType(),
		valueRaw,
	)
	return nil
}
