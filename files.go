package main

import "os"

func CreateDirIfNotExists(filePath string) error {
	if !FileExists(filePath) {
		return CreateDir(filePath)
	}
	return nil
}

func FileExists(filePath string) bool {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false
	}
	return true
}

func CreateDir(filePath string) error {
	return os.Mkdir(filePath, os.ModePerm)
}

func CreateFile(filePath string) error {
	_, err := os.Create(filePath)
	return err
}
