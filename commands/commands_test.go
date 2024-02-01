package commands

import (
	"github.com/gsanhuezafuentes/exif-classifier/logger"
	"github.com/gsanhuezafuentes/exif-classifier/mocks/github.com/gsanhuezafuentes/exif-classifier/commands"
	"github.com/gsanhuezafuentes/exif-classifier/mocks/github.com/gsanhuezafuentes/exif-classifier/exif_reader"
	"github.com/gsanhuezafuentes/exif-classifier/mocks/github.com/gsanhuezafuentes/exif-classifier/organize"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func getCmdContext() CmdContext {
	return CmdContext{
		Logger: logger.GetLogger(),
		Stdout: os.Stdout,
	}
}

func TestGroupCmd_Run(t *testing.T) {
	organizer := organize.NewMockOrganizer(t)
	fileOperation := commands.NewMockGroupCmdFileOperation(t)
	ctx := GroupCmdContext{
		CmdContext:    getCmdContext(),
		Organizer:     organizer,
		FileOperation: fileOperation,
	}
	fakePath := "/home/"

	t.Run("Default directory", func(t *testing.T) {
		cmd := GroupCmd{}

		imagesPath := []string{"file1", "file2"}

		fileOperation.EXPECT().GetCurrentDirectory().Return(fakePath, nil).Once()
		fileOperation.EXPECT().GetImageFilesPathFromDirectory(fakePath).Return(imagesPath, nil).Once()
		organizer.EXPECT().SetImagesPath(imagesPath).Once()
		organizer.EXPECT().OrganizeImgsByDate().Return(nil).Once()

		err := cmd.Run(&ctx)

		assert.NoError(t, err)
		fileOperation.AssertExpectations(t)
		organizer.AssertExpectations(t)
	})

	t.Run("With directory", func(t *testing.T) {
		cmd := GroupCmd{Path: fakePath}

		imagesPath := []string{"file1", "file2"}

		fileOperation.EXPECT().GetImageFilesPathFromDirectory(fakePath).Return(imagesPath, nil).Once()
		organizer.EXPECT().SetImagesPath(imagesPath).Once()
		organizer.EXPECT().OrganizeImgsByDate().Return(nil).Once()

		err := cmd.Run(&ctx)

		assert.NoError(t, err)
		fileOperation.AssertExpectations(t)
		organizer.AssertExpectations(t)
	})

	t.Run("With --lens", func(t *testing.T) {
		cmd := GroupCmd{
			Lens: true,
			Path: fakePath,
		}

		imagesPath := []string{"file1", "file2"}

		fileOperation.EXPECT().GetImageFilesPathFromDirectory(fakePath).Return(imagesPath, nil).Once()
		organizer.EXPECT().SetImagesPath(imagesPath).Once()
		organizer.EXPECT().OrganizeImgsByLens().Return(nil).Once()

		err := cmd.Run(&ctx)

		assert.NoError(t, err)
		fileOperation.AssertExpectations(t)
		organizer.AssertExpectations(t)
	})

	t.Run("With --orientation", func(t *testing.T) {
		cmd := GroupCmd{
			Orientation: true,
			Path:        fakePath,
		}

		imagesPath := []string{"file1", "file2"}

		fileOperation.EXPECT().GetImageFilesPathFromDirectory(fakePath).Return(imagesPath, nil).Once()
		organizer.EXPECT().SetImagesPath(imagesPath).Once()
		organizer.EXPECT().OrganizeImgsByOrientation().Return(nil).Once()

		err := cmd.Run(&ctx)

		assert.NoError(t, err)
		fileOperation.AssertExpectations(t)
		organizer.AssertExpectations(t)
	})

	t.Run("With all flags", func(t *testing.T) {
		cmd := GroupCmd{
			Date:        true,
			Orientation: true,
			Lens:        true,
			Path:        fakePath,
		}

		imagesPath := []string{"file1", "file2"}

		fileOperation.EXPECT().GetImageFilesPathFromDirectory(fakePath).Return(imagesPath, nil).Once()
		organizer.EXPECT().SetImagesPath(imagesPath).Once()
		organizer.EXPECT().OrganizeImgsByDate().Return(nil).Once()
		organizer.EXPECT().OrganizeImgsByLens().Return(nil).Once()
		organizer.EXPECT().OrganizeImgsByOrientation().Return(nil).Once()
		err := cmd.Run(&ctx)

		assert.NoError(t, err)
		fileOperation.AssertExpectations(t)
		organizer.AssertExpectations(t)
	})
}

func TestPrintExifCmd_Run(t *testing.T) {
	exifPrinter := exif_reader.NewMockExifPrinter(t)
	ctx := PrintExifCmdContext{
		CmdContext:  getCmdContext(),
		ExifPrinter: exifPrinter,
	}

	path := "C://Images/File.dng"

	cmd := PrintExifCmd{
		Path: path,
	}

	exifPrinter.EXPECT().PrintExif(path, os.Stdout).Return(nil)

	assert.Nil(t, cmd.Run(&ctx))
	exifPrinter.AssertExpectations(t)
}
