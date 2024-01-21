package main

import (
	"fmt"
	exif_reader "github.com/gsanhuezafuentes/exif-classifier/exif-reader"
	"os"
)

type GroupCmd struct {
	Date        bool   `help:"Sort images by date"`
	Lens        bool   `help:"Sort images by lens info"`
	Orientation string `help:"Sort images by orientation info (horizontal|vertical)"`
	Path        string `arg:"" help:"Location folder of images. by default use the date option" optional:"" type:"existingfile"`
}

func (r *GroupCmd) Run(ctx Context) error {
	if r.Path == "" {
		currentDir, err := os.Getwd()
		if err != nil {
			ctx.logger.Errorf("%+v\n", r)
			return err
		}
		ctx.logger.Debugf("Using Getwd() %s", currentDir)
		r.Path = currentDir
	}
	fmt.Fprintf(ctx.output, "%+v\n", r)

	_, err := exif_reader.DefaultExifReader{}.Read(r.Path)
	return err
}

type PrintExifCmd struct {
	Path string `arg:"" help:"Location folder of images. by default use the date option" optional:"" type:"existingfile"`
}

func (r *PrintExifCmd) Run(ctx Context) error {
	if r.Path == "" {
		currentDir, err := os.Getwd()
		if err != nil {
			ctx.logger.Errorf("%+v\n", r)
			return err
		}
		ctx.logger.Debugf("Using Getwd() %s", currentDir)
		r.Path = currentDir
	}
	fmt.Fprintf(ctx.output, "%+v\n", r)

	exif_reader.DefaultExifReader{}.PrintExif(r.Path)
	return nil
}
