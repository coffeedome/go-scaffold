package main

import (
	"fmt"
	"io/fs"
	"os"
)

/*
	CreateDir creates a directory with dirName and permissions
*/
func CreateDir(dirName string, permissions fs.FileMode) error {
	err := os.Mkdir(dirName, permissions)

	if err != nil && !os.IsExist(err) {
		return err
	}

	return nil
}

/*
	CreateFile creates a file
*/
func CreateFile(fileName string) (*os.File, error) {
	file, err := os.Create(fmt.Sprintf("scaffolding/%s.go", fileName))

	if err != nil {
		return nil, err
	}

	return file, nil
}
