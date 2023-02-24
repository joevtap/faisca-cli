package scaffolder

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

type Tree map[string]Dir

type Dir struct {
	Dirs  map[string]interface{}
	Files []File
}

type File struct {
	Name    string
	Content *template.Template
	Data    interface{}
}

func (t Tree) Scaffold() error {
	for key, value := range t {
		dirPath := filepath.Join(".", key)

		err := os.MkdirAll(dirPath, 0755)
		if err != nil {
			return err
		}
		err = value.Scaffold(dirPath)
		if err != nil {
			return err
		}
	}

	return nil
}

func (d Dir) Scaffold(parentDirPath string) error {
	for key, value := range d.Dirs {
		subDirPath := filepath.Join(parentDirPath, key)

		subDir, ok := interface{}(value).(Dir)
		if !ok {
			return fmt.Errorf("%v is not a directory", value)
		}

		err := os.MkdirAll(subDirPath, 0755)
		if err != nil {
			return err
		}

		err = subDir.Scaffold(subDirPath)
		if err != nil {
			return err
		}
	}

	err := createFiles(d.Files, parentDirPath)
	if err != nil {
		return err
	}

	return nil
}

func createFiles(files []File, parentDirPath string) error {
	for _, file := range files {
		err := createFile(file, parentDirPath)
		if err != nil {
			return err
		}
	}

	return nil
}

func createFile(file File, parentDirPath string) error {

	filePath := filepath.Join(".", parentDirPath, file.Name)

	err := os.MkdirAll(filepath.Dir(filePath), 0755)
	if err != nil {
		return err
	}

	newFile, err := os.Create(filePath)
	if err != nil {
		return err
	}

	defer newFile.Close()

	if file.Content == nil {
		return nil
	}

	if file.Data == nil {
		return nil
	}

	err = file.Content.Execute(newFile, file.Data)
	if err != nil {
		return err
	}

	return nil
}
