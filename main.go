package main

import (
	"github.com/alecthomas/kong"
	"github.com/gsanhuezafuentes/exif-classifier/exif_reader"
	"github.com/gsanhuezafuentes/exif-classifier/logger"
	"github.com/gsanhuezafuentes/exif-classifier/organize"
	"os"
)

func setLoggerOutput(cli *CLI) error {
	if cli.Debug {
		logger.GetLogger().SetOutput(os.Stderr)
	}
	return nil
}

func setLoggerLevel(cli *CLI) error {
	switch cli.LogLevel {
	case "debug":
		logger.GetLogger().SetLogLevel(logger.DEBUG)
	case "info":
		logger.GetLogger().SetLogLevel(logger.INFO)
	case "warning":
		logger.GetLogger().SetLogLevel(logger.WARNING)
	case "error":
		logger.GetLogger().SetLogLevel(logger.WARNING)
	}
	return nil

}

type CLI struct {
	Debug     bool         `short:"D" help:"Enable debug mode"`
	LogLevel  string       `short:"l" default:"info" help:"Set the logging level (debug|info|warning|error)" enum:"debug,info,warning,error"`
	Group     GroupCmd     `cmd:"" help:"Move images based in a specific exif attribute"`
	PrintExif PrintExifCmd `cmd:"" help:"Print exif tag of the images"`
}

func main() {
	exifReader := exif_reader.New()

	cli := CLI{}
	ctx := kong.Parse(&cli,
		kong.Name("exif-classifier"),
		kong.Description("An application to move images to folder using the exif data"),
		kong.UsageOnError(),
		kong.ConfigureHelp(kong.HelpOptions{
			Compact: true,
		}),
		kong.Vars{
			"version": "0.0.1",
		},
		kong.Bind(
			&GroupCmdContext{
				CmdContext:    CmdContext{Logger: logger.GetLogger(), ProgramOutput: os.Stdout},
				Organizer:     organize.NewDefaultOrganizer(exif_reader.New(), os.Rename),
				FileOperation: DefaultCmdOperation{},
			},
			&PrintExifCmdContext{
				CmdContext: CmdContext{
					Logger:        logger.GetLogger(),
					ProgramOutput: os.Stdout,
				},
				ExifPrinter: exifReader,
			},
		),
	)

	setLoggerOutput(&cli)
	setLoggerLevel(&cli)
	ctx.FatalIfErrorf(ctx.Run())
}
