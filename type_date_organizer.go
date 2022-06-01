package main

import (
	"os"
	"path/filepath"
	"time"
)

type TypeDateOrganizer struct {
	FolderSuffix string
}

func (t TypeDateOrganizer) Organize(file os.FileInfo, rootLocation string) error {
	// get extension from file
	origLocation := rootLocation + SeparatorString + file.Name()
	extension := filepath.Ext(origLocation)

	// safely set extension
	if len(extension) < 1 {
		extension = "uunknown"
	}

	// create file if not exists
	extensionLoc := rootLocation + SeparatorString + extension[1:] + t.FolderSuffix
	if err := CreateDirIfNotExists(extensionLoc); err != nil {
		return err
	}

	// see if folder exists for today under extension
	dateString := time.Now().Format("2006-02-01-Monday")
	dateLocation := extensionLoc + SeparatorString + dateString
	if err := CreateDirIfNotExists(dateLocation); err != nil {
		return err
	}

	// move file to dir
	newFileLoc := dateLocation + SeparatorString + file.Name()
	return os.Rename(origLocation, newFileLoc)
}
