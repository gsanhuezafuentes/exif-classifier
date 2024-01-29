package main

import (
	"fmt"
	"github.com/gsanhuezafuentes/exif-classifier/fileutils"
)

type GroupCmd struct {
	Date        bool   `help:"Sort images by date"`
	Lens        bool   `help:"Sort images by lens info"`
	Orientation bool   `help:"Sort images by orientation"`
	Path        string `arg:"" help:"Location folder of images. by default use the date option" optional:"" type:"existingdir"`
}

func (r *GroupCmd) Run(ctx Context) error {
	if r.Path == "" {
		directory, err := fileutils.GetCurrentDirectory()
		ctx.Logger.Debugf("Using Getwd() %s", directory)
		if err != nil {
			return err
		}
		r.Path = directory
	}
	fmt.Fprintf(ctx.ProgramOutput, "%+v\n", r)

	files, err := fileutils.GetImageFilesPathFromDirectory(r.Path)
	fmt.Fprintf(ctx.ProgramOutput, "%s\n", files)

	ctx.Organizer.SetImagesPath(files)

	if !(r.Date && r.Lens && r.Orientation) {
		r.Date = true
	}

	if r.Date {
		err = ctx.Organizer.OrganizeImgsByDate()
		if err != nil {
			return err
		}
	}

	if r.Orientation {
		err = ctx.Organizer.OrganizeImgsByOrientation()
		if err != nil {
			return err
		}
	}

	if r.Lens {
		err = ctx.Organizer.OrganizeImgsByLens()
		if err != nil {
			return err
		}
	}
	return err
}

type PrintExifCmd struct {
	Path string `arg:"" help:"File path of the image" type:"existingfile"`
}

func (r *PrintExifCmd) Run(ctx Context) error {
	fmt.Fprintf(ctx.ProgramOutput, "%+v\n", r)

	ctx.ExifPrinter.PrintExif(r.Path, ctx.ProgramOutput)
	return nil
}
