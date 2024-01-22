package main

import (
	"fmt"
	"os"
)

type GroupCmd struct {
	Date        bool   `help:"Sort images by date"`
	Lens        bool   `help:"Sort images by lens info"`
	Orientation string `help:"Sort images by orientation info (horizontal|vertical)"`
	Path        string `arg:"" help:"Location folder of images. by default use the date option" optional:"" type:"existingdir"`
}

func (r *GroupCmd) Run(ctx Context) error {
	if r.Path == "" {
		currentDir, err := os.Getwd()
		if err != nil {
			ctx.Logger.Errorf("%+v\n", r)
			return err
		}
		ctx.Logger.Debugf("Using Getwd() %s", currentDir)
		r.Path = currentDir
	}
	fmt.Fprintf(ctx.ProgramOutput, "%+v\n", r)

	exifData, err := ctx.ExifReader.Read(r.Path)
	fmt.Fprintf(ctx.ProgramOutput, "%+v\n", exifData)
	return err
}

type PrintExifCmd struct {
	Path string `arg:"" help:"Location folder of images. by default use the date option" optional:"" type:"existingfile"`
}

func (r *PrintExifCmd) Run(ctx Context) error {
	if r.Path == "" {
		currentDir, err := os.Getwd()
		if err != nil {
			ctx.Logger.Errorf("%+v\n", r)
			return err
		}
		ctx.Logger.Debugf("Using Getwd() %s", currentDir)
		r.Path = currentDir
	}
	fmt.Fprintf(ctx.ProgramOutput, "%+v\n", r)

	ctx.ExifPrinter.PrintExif(r.Path, ctx.ProgramOutput)
	return nil
}
