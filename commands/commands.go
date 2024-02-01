package commands

import (
	"fmt"
	"github.com/gsanhuezafuentes/exif-classifier/exif_reader"
	"github.com/gsanhuezafuentes/exif-classifier/fileutils"
	"github.com/gsanhuezafuentes/exif-classifier/logger"
	"github.com/gsanhuezafuentes/exif-classifier/organize"
	"io"
)

type GroupCmd struct {
	Date        bool   `help:"Sort images by date"`
	Lens        bool   `help:"Sort images by lens info"`
	Orientation bool   `help:"Sort images by orientation"`
	Path        string `arg:"" help:"Location folder of images. by default use the date option" optional:"" type:"existingdir"`
}

type CmdContext struct {
	Logger logger.Logger
	Stdout io.Writer
}

type GroupCmdContext struct {
	CmdContext
	Organizer     organize.Organizer
	FileOperation GroupCmdFileOperation
}

type GroupCmdFileOperation interface {
	GetCurrentDirectory() (string, error)
	GetImageFilesPathFromDirectory(string) ([]string, error)
}

func (r *GroupCmd) Run(ctx *GroupCmdContext) error {
	if r.Path == "" {
		directory, err := ctx.FileOperation.GetCurrentDirectory()
		ctx.Logger.Debugf("Using Getwd() %s", directory)
		if err != nil {
			return err
		}
		r.Path = directory
	}
	fmt.Fprintf(ctx.Stdout, "%+v\n", r)

	files, err := ctx.FileOperation.GetImageFilesPathFromDirectory(r.Path)
	if err != nil {
		return err
	}

	fmt.Fprintf(ctx.Stdout, "%s\n", files)

	ctx.Organizer.SetImagesPath(files)

	if !(r.Date || r.Lens || r.Orientation) {
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

type DefaultCmdOperation struct{}

func (r DefaultCmdOperation) GetCurrentDirectory() (string, error) {
	return fileutils.GetCurrentDirectory()
}

func (r DefaultCmdOperation) GetImageFilesPathFromDirectory(s string) ([]string, error) {
	return fileutils.GetImageFilesPathFromDirectory(s)
}

type PrintExifCmd struct {
	Path string `arg:"" help:"File path of the image" type:"existingfile"`
}

type PrintExifCmdContext struct {
	CmdContext
	ExifPrinter exif_reader.ExifPrinter
}

func (r *PrintExifCmd) Run(ctx *PrintExifCmdContext) error {
	fmt.Fprintf(ctx.Stdout, "%+v\n", r)

	ctx.ExifPrinter.PrintExif(r.Path, ctx.Stdout)
	return nil
}
