package main

import (
	"fmt"
	"github.com/gsanhuezafuentes/exif-classifier/logger"
	"os"
)

type GroupCmd struct {
	Date        bool   `help:"Sort images by date"`
	Lens        bool   `help:"Sort images by lens info"`
	Orientation string `help:"Sort images by orientation info (horizontal|vertical)"`
	Path        string `arg:"" help:"Location folder of images. by default use the date option" optional:"" type:"existingdir"`
}

func (r *GroupCmd) Run() error {
	if r.Path == "" {
		currentDir, err := os.Getwd()
		if err != nil {
			logger.GetLogger().Errorf("%+v\n", r)
			return err
		}
		logger.GetLogger().Debugf("Using Getwd() %s", currentDir)
		r.Path = currentDir
	}
	fmt.Printf("%+v\n", r)
	return nil
}
