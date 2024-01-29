package main

import (
	"github.com/gsanhuezafuentes/exif-classifier/mocks/github.com/gsanhuezafuentes/exif-classifier/exif_reader"
	"github.com/gsanhuezafuentes/exif-classifier/mocks/github.com/gsanhuezafuentes/exif-classifier/organize"
	"github.com/gsanhuezafuentes/exif-classifier/mocks/io"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestGroupCmd_Run(t *testing.T) {
	output := io.NewMockWriter(t)
	organizer := organize.NewMockOrganizer(t)
	ctx := Context{
		ProgramOutput: output,
		Organizer:     organizer,
	}

	t.Run("Default directory", func(t *testing.T) {
		cmd := GroupCmd{}
		err := cmd.Run(ctx)

		assert.NoError(t, err)
	})
}

func TestPrintExifCmd_Run(t *testing.T) {
	output := io.NewMockWriter(t)
	exifPrinter := exif_reader.NewMockExifPrinter(t)
	ctx := Context{
		ProgramOutput: output,
		ExifPrinter:   exifPrinter,
	}

	path := "C://Images/File.dng"

	cmd := PrintExifCmd{
		Path: path,
	}

	output.EXPECT().Write(mock.Anything).Return(1, nil)
	exifPrinter.EXPECT().PrintExif(path, output).Return(nil)

	assert.Nil(t, cmd.Run(ctx))
	output.AssertCalled(t, "Write", mock.Anything)
	exifPrinter.AssertExpectations(t)
}
