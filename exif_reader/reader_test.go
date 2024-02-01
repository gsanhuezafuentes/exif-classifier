package exif_reader

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

const testImageWithoutOrientation = "../assests/_DSC3711_without_orientation.jpg"

func getTestImage(imagePath string) string {
	dir, _ := os.Getwd()
	return filepath.Join(dir, imagePath)
}

func normalizeLineBreaks(s string) string {
	return strings.ReplaceAll(s, "\n\t", "\n")
}

func TestDefaultExifReader_Read(t *testing.T) {
	ExifReader := New()

	data, err := ExifReader.Read(getTestImage(testImageWithoutOrientation))
	assert.NoError(t, err)

	time, err := time.Parse(DATEFORMAT, "2024:02:01 00:09:58")
	assert.NoError(t, err)

	assert.Equal(t, data.CreatedTime, time)
	assert.Equal(t, data.Orientation, UNKNOWN)
	assert.Equal(t, data.LensModel, "E 150-500mm F5-6.7 A057")
}

func TestDefaultExifReader_PrintExif(t *testing.T) {
	buffer := &bytes.Buffer{}
	ExifReader := New()

	err := ExifReader.PrintExif(getTestImage(testImageWithoutOrientation), buffer)
	assert.NoError(t, err)

	assert.Equal(t, normalizeLineBreaks(`IFD-PATH=[IFD] ID=(0x010f) NAME=[Make] COUNT=(5) TYPE=[ASCII] VALUE=[SONY]
	IFD-PATH=[IFD] ID=(0x0110) NAME=[Model] COUNT=(10) TYPE=[ASCII] VALUE=[ILCE-6400]
	IFD-PATH=[IFD] ID=(0x011a) NAME=[XResolution] COUNT=(1) TYPE=[RATIONAL] VALUE=[[{70 1}]]
	IFD-PATH=[IFD] ID=(0x011b) NAME=[YResolution] COUNT=(1) TYPE=[RATIONAL] VALUE=[[{70 1}]]
	IFD-PATH=[IFD] ID=(0x0128) NAME=[ResolutionUnit] COUNT=(1) TYPE=[SHORT] VALUE=[[2]]
	IFD-PATH=[IFD] ID=(0x0131) NAME=[Software] COUNT=(51) TYPE=[ASCII] VALUE=[Adobe Photoshop Lightroom Classic 13.0.2 (Windows)]
	IFD-PATH=[IFD] ID=(0x0132) NAME=[DateTime] COUNT=(20) TYPE=[ASCII] VALUE=[2024:02:01 00:09:58]
	IFD-PATH=[IFD/Exif] ID=(0x829a) NAME=[ExposureTime] COUNT=(1) TYPE=[RATIONAL] VALUE=[[{1 1000}]]
	IFD-PATH=[IFD/Exif] ID=(0x829d) NAME=[FNumber] COUNT=(1) TYPE=[RATIONAL] VALUE=[[{63 10}]]
	IFD-PATH=[IFD/Exif] ID=(0x8822) NAME=[ExposureProgram] COUNT=(1) TYPE=[SHORT] VALUE=[[3]]
	IFD-PATH=[IFD/Exif] ID=(0x8827) NAME=[ISOSpeedRatings] COUNT=(1) TYPE=[SHORT] VALUE=[[100]]
	IFD-PATH=[IFD/Exif] ID=(0x8830) NAME=[SensitivityType] COUNT=(1) TYPE=[SHORT] VALUE=[[2]]
	IFD-PATH=[IFD/Exif] ID=(0x8832) NAME=[RecommendedExposureIndex] COUNT=(1) TYPE=[LONG] VALUE=[[100]]
	IFD-PATH=[IFD/Exif] ID=(0x9000) NAME=[ExifVersion] COUNT=(4) TYPE=[UNDEFINED] VALUE=[0231]
	IFD-PATH=[IFD/Exif] ID=(0x9003) NAME=[DateTimeOriginal] COUNT=(20) TYPE=[ASCII] VALUE=[2024:01:07 14:34:10]
	IFD-PATH=[IFD/Exif] ID=(0x9004) NAME=[DateTimeDigitized] COUNT=(20) TYPE=[ASCII] VALUE=[2024:01:07 14:34:10]
	IFD-PATH=[IFD/Exif] ID=(0x9010) NAME=[OffsetTime] COUNT=(7) TYPE=[ASCII] VALUE=[-03:00]
	IFD-PATH=[IFD/Exif] ID=(0x9011) NAME=[OffsetTimeOriginal] COUNT=(7) TYPE=[ASCII] VALUE=[-03:00]
	IFD-PATH=[IFD/Exif] ID=(0x9012) NAME=[OffsetTimeDigitized] COUNT=(7) TYPE=[ASCII] VALUE=[-03:00]
	IFD-PATH=[IFD/Exif] ID=(0x9201) NAME=[ShutterSpeedValue] COUNT=(1) TYPE=[SRATIONAL] VALUE=[[{9965784 1000000}]]
	IFD-PATH=[IFD/Exif] ID=(0x9202) NAME=[ApertureValue] COUNT=(1) TYPE=[RATIONAL] VALUE=[[{5310704 1000000}]]
	IFD-PATH=[IFD/Exif] ID=(0x9203) NAME=[BrightnessValue] COUNT=(1) TYPE=[SRATIONAL] VALUE=[[{26764 2560}]]
	IFD-PATH=[IFD/Exif] ID=(0x9204) NAME=[ExposureBiasValue] COUNT=(1) TYPE=[SRATIONAL] VALUE=[[{0 10}]]
	IFD-PATH=[IFD/Exif] ID=(0x9205) NAME=[MaxApertureValue] COUNT=(1) TYPE=[RATIONAL] VALUE=[[{1188 256}]]
	IFD-PATH=[IFD/Exif] ID=(0x9207) NAME=[MeteringMode] COUNT=(1) TYPE=[SHORT] VALUE=[[3]]
	IFD-PATH=[IFD/Exif] ID=(0x9208) NAME=[LightSource] COUNT=(1) TYPE=[SHORT] VALUE=[[10]]
	IFD-PATH=[IFD/Exif] ID=(0x9209) NAME=[Flash] COUNT=(1) TYPE=[SHORT] VALUE=[[16]]
	IFD-PATH=[IFD/Exif] ID=(0x920a) NAME=[FocalLength] COUNT=(1) TYPE=[RATIONAL] VALUE=[[{1510 10}]]
	IFD-PATH=[IFD/Exif] ID=(0xa001) NAME=[ColorSpace] COUNT=(1) TYPE=[SHORT] VALUE=[[1]]
	IFD-PATH=[IFD/Exif] ID=(0xa20e) NAME=[FocalPlaneXResolution] COUNT=(1) TYPE=[RATIONAL] VALUE=[[{83662979 32768}]]
	IFD-PATH=[IFD/Exif] ID=(0xa20f) NAME=[FocalPlaneYResolution] COUNT=(1) TYPE=[RATIONAL] VALUE=[[{83662979 32768}]]
	IFD-PATH=[IFD/Exif] ID=(0xa210) NAME=[FocalPlaneResolutionUnit] COUNT=(1) TYPE=[SHORT] VALUE=[[3]]
	IFD-PATH=[IFD/Exif] ID=(0xa300) NAME=[FileSource] COUNT=(1) TYPE=[UNDEFINED] VALUE=[0x00000003]
	IFD-PATH=[IFD/Exif] ID=(0xa301) NAME=[SceneType] COUNT=(1) TYPE=[UNDEFINED] VALUE=[0x00000001]
	IFD-PATH=[IFD/Exif] ID=(0xa401) NAME=[CustomRendered] COUNT=(1) TYPE=[SHORT] VALUE=[[0]]
	IFD-PATH=[IFD/Exif] ID=(0xa402) NAME=[ExposureMode] COUNT=(1) TYPE=[SHORT] VALUE=[[0]]
	IFD-PATH=[IFD/Exif] ID=(0xa403) NAME=[WhiteBalance] COUNT=(1) TYPE=[SHORT] VALUE=[[1]]
	IFD-PATH=[IFD/Exif] ID=(0xa404) NAME=[DigitalZoomRatio] COUNT=(1) TYPE=[RATIONAL] VALUE=[[{16 16}]]
	IFD-PATH=[IFD/Exif] ID=(0xa405) NAME=[FocalLengthIn35mmFilm] COUNT=(1) TYPE=[SHORT] VALUE=[[226]]
	IFD-PATH=[IFD/Exif] ID=(0xa406) NAME=[SceneCaptureType] COUNT=(1) TYPE=[SHORT] VALUE=[[0]]
	IFD-PATH=[IFD/Exif] ID=(0xa408) NAME=[Contrast] COUNT=(1) TYPE=[SHORT] VALUE=[[0]]
	IFD-PATH=[IFD/Exif] ID=(0xa409) NAME=[Saturation] COUNT=(1) TYPE=[SHORT] VALUE=[[0]]
	IFD-PATH=[IFD/Exif] ID=(0xa40a) NAME=[Sharpness] COUNT=(1) TYPE=[SHORT] VALUE=[[0]]
	IFD-PATH=[IFD/Exif] ID=(0xa432) NAME=[LensSpecification] COUNT=(4) TYPE=[RATIONAL] VALUE=[[{1500 10} {5000 10} {50 10} {67 10}]]
	IFD-PATH=[IFD/Exif] ID=(0xa434) NAME=[LensModel] COUNT=(24) TYPE=[ASCII] VALUE=[E 150-500mm F5-6.7 A057]
	IFD-PATH=[IFD] ID=(0x010f) NAME=[Make] COUNT=(5) TYPE=[ASCII] VALUE=[SONY]
	IFD-PATH=[IFD] ID=(0x0110) NAME=[Model] COUNT=(10) TYPE=[ASCII] VALUE=[ILCE-6400]
	IFD-PATH=[IFD] ID=(0x011a) NAME=[XResolution] COUNT=(1) TYPE=[RATIONAL] VALUE=[[{70 1}]]
	IFD-PATH=[IFD] ID=(0x011b) NAME=[YResolution] COUNT=(1) TYPE=[RATIONAL] VALUE=[[{70 1}]]
	IFD-PATH=[IFD] ID=(0x0128) NAME=[ResolutionUnit] COUNT=(1) TYPE=[SHORT] VALUE=[[2]]
	IFD-PATH=[IFD] ID=(0x0131) NAME=[Software] COUNT=(51) TYPE=[ASCII] VALUE=[Adobe Photoshop Lightroom Classic 13.0.2 (Windows)]
	IFD-PATH=[IFD] ID=(0x0132) NAME=[DateTime] COUNT=(20) TYPE=[ASCII] VALUE=[2024:02:01 00:09:58]
	IFD-PATH=[IFD/Exif] ID=(0x829a) NAME=[ExposureTime] COUNT=(1) TYPE=[RATIONAL] VALUE=[[{1 1000}]]
	IFD-PATH=[IFD/Exif] ID=(0x829d) NAME=[FNumber] COUNT=(1) TYPE=[RATIONAL] VALUE=[[{63 10}]]
	IFD-PATH=[IFD/Exif] ID=(0x8822) NAME=[ExposureProgram] COUNT=(1) TYPE=[SHORT] VALUE=[[3]]
	IFD-PATH=[IFD/Exif] ID=(0x8827) NAME=[ISOSpeedRatings] COUNT=(1) TYPE=[SHORT] VALUE=[[100]]
	IFD-PATH=[IFD/Exif] ID=(0x8830) NAME=[SensitivityType] COUNT=(1) TYPE=[SHORT] VALUE=[[2]]
	IFD-PATH=[IFD/Exif] ID=(0x8832) NAME=[RecommendedExposureIndex] COUNT=(1) TYPE=[LONG] VALUE=[[100]]
	IFD-PATH=[IFD/Exif] ID=(0x9000) NAME=[ExifVersion] COUNT=(4) TYPE=[UNDEFINED] VALUE=[0231]
	IFD-PATH=[IFD/Exif] ID=(0x9003) NAME=[DateTimeOriginal] COUNT=(20) TYPE=[ASCII] VALUE=[2024:01:07 14:34:10]
	IFD-PATH=[IFD/Exif] ID=(0x9004) NAME=[DateTimeDigitized] COUNT=(20) TYPE=[ASCII] VALUE=[2024:01:07 14:34:10]
	IFD-PATH=[IFD/Exif] ID=(0x9010) NAME=[OffsetTime] COUNT=(7) TYPE=[ASCII] VALUE=[-03:00]
	IFD-PATH=[IFD/Exif] ID=(0x9011) NAME=[OffsetTimeOriginal] COUNT=(7) TYPE=[ASCII] VALUE=[-03:00]
	IFD-PATH=[IFD/Exif] ID=(0x9012) NAME=[OffsetTimeDigitized] COUNT=(7) TYPE=[ASCII] VALUE=[-03:00]
	IFD-PATH=[IFD/Exif] ID=(0x9201) NAME=[ShutterSpeedValue] COUNT=(1) TYPE=[SRATIONAL] VALUE=[[{9965784 1000000}]]
	IFD-PATH=[IFD/Exif] ID=(0x9202) NAME=[ApertureValue] COUNT=(1) TYPE=[RATIONAL] VALUE=[[{5310704 1000000}]]
	IFD-PATH=[IFD/Exif] ID=(0x9203) NAME=[BrightnessValue] COUNT=(1) TYPE=[SRATIONAL] VALUE=[[{26764 2560}]]
	IFD-PATH=[IFD/Exif] ID=(0x9204) NAME=[ExposureBiasValue] COUNT=(1) TYPE=[SRATIONAL] VALUE=[[{0 10}]]
	IFD-PATH=[IFD/Exif] ID=(0x9205) NAME=[MaxApertureValue] COUNT=(1) TYPE=[RATIONAL] VALUE=[[{1188 256}]]
	IFD-PATH=[IFD/Exif] ID=(0x9207) NAME=[MeteringMode] COUNT=(1) TYPE=[SHORT] VALUE=[[3]]
	IFD-PATH=[IFD/Exif] ID=(0x9208) NAME=[LightSource] COUNT=(1) TYPE=[SHORT] VALUE=[[10]]
	IFD-PATH=[IFD/Exif] ID=(0x9209) NAME=[Flash] COUNT=(1) TYPE=[SHORT] VALUE=[[16]]
	IFD-PATH=[IFD/Exif] ID=(0x920a) NAME=[FocalLength] COUNT=(1) TYPE=[RATIONAL] VALUE=[[{1510 10}]]
	IFD-PATH=[IFD/Exif] ID=(0xa001) NAME=[ColorSpace] COUNT=(1) TYPE=[SHORT] VALUE=[[1]]
	IFD-PATH=[IFD/Exif] ID=(0xa20e) NAME=[FocalPlaneXResolution] COUNT=(1) TYPE=[RATIONAL] VALUE=[[{83662979 32768}]]
	IFD-PATH=[IFD/Exif] ID=(0xa20f) NAME=[FocalPlaneYResolution] COUNT=(1) TYPE=[RATIONAL] VALUE=[[{83662979 32768}]]
	IFD-PATH=[IFD/Exif] ID=(0xa210) NAME=[FocalPlaneResolutionUnit] COUNT=(1) TYPE=[SHORT] VALUE=[[3]]
	IFD-PATH=[IFD/Exif] ID=(0xa300) NAME=[FileSource] COUNT=(1) TYPE=[UNDEFINED] VALUE=[0x00000003]
	IFD-PATH=[IFD/Exif] ID=(0xa301) NAME=[SceneType] COUNT=(1) TYPE=[UNDEFINED] VALUE=[0x00000001]
	IFD-PATH=[IFD/Exif] ID=(0xa401) NAME=[CustomRendered] COUNT=(1) TYPE=[SHORT] VALUE=[[0]]
	IFD-PATH=[IFD/Exif] ID=(0xa402) NAME=[ExposureMode] COUNT=(1) TYPE=[SHORT] VALUE=[[0]]
	IFD-PATH=[IFD/Exif] ID=(0xa403) NAME=[WhiteBalance] COUNT=(1) TYPE=[SHORT] VALUE=[[1]]
	IFD-PATH=[IFD/Exif] ID=(0xa404) NAME=[DigitalZoomRatio] COUNT=(1) TYPE=[RATIONAL] VALUE=[[{16 16}]]
	IFD-PATH=[IFD/Exif] ID=(0xa405) NAME=[FocalLengthIn35mmFilm] COUNT=(1) TYPE=[SHORT] VALUE=[[226]]
	IFD-PATH=[IFD/Exif] ID=(0xa406) NAME=[SceneCaptureType] COUNT=(1) TYPE=[SHORT] VALUE=[[0]]
	IFD-PATH=[IFD/Exif] ID=(0xa408) NAME=[Contrast] COUNT=(1) TYPE=[SHORT] VALUE=[[0]]
	IFD-PATH=[IFD/Exif] ID=(0xa409) NAME=[Saturation] COUNT=(1) TYPE=[SHORT] VALUE=[[0]]
	IFD-PATH=[IFD/Exif] ID=(0xa40a) NAME=[Sharpness] COUNT=(1) TYPE=[SHORT] VALUE=[[0]]
	IFD-PATH=[IFD/Exif] ID=(0xa432) NAME=[LensSpecification] COUNT=(4) TYPE=[RATIONAL] VALUE=[[{1500 10} {5000 10} {50 10} {67 10}]]
	IFD-PATH=[IFD/Exif] ID=(0xa434) NAME=[LensModel] COUNT=(24) TYPE=[ASCII] VALUE=[E 150-500mm F5-6.7 A057]
	`), buffer.String(),
	)
}
