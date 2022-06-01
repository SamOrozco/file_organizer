package main

import (
	"os"
	"path/filepath"
	"time"
)

type DateTypeOrganizer struct {
	FolderSuffix string
}

func (d DateTypeOrganizer) Organize(file os.FileInfo, rootLocation string) error {
	// get extension from file
	origLocation := rootLocation + SeparatorString + file.Name()
	extension := filepath.Ext(origLocation)

	// see if folder exists for today under root
	dateString := time.Now().Format("2006-02-01-Monday")
	dateLocation := rootLocation + SeparatorString + dateString
	if err := CreateDirIfNotExists(dateLocation); err != nil {
		return err
	}

	// safely set extension
	if len(extension) < 1 {
		extension = "uunknown"
	}

	// create file if not exists
	extensionLoc := dateLocation + SeparatorString + extension[1:] + d.FolderSuffix
	if err := CreateDirIfNotExists(extensionLoc); err != nil {
		return err
	}

	// move file to dir
	newFileLoc := extensionLoc + SeparatorString + file.Name()
	return os.Rename(origLocation, newFileLoc)
}
