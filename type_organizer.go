package main

import (
	"os"
	"path/filepath"
)

type TypeOrganizer struct {
	FolderSuffix string
}

func (t TypeOrganizer) Organize(file os.FileInfo, rootLocation string) error {
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

	// move file to dir
	newFileLoc := extensionLoc + SeparatorString + file.Name()
	return os.Rename(origLocation, newFileLoc)
}
