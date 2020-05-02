package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func changeWorkingDirectoryToRoot() error {
	err := spin("Searching for goyave root",
		"Goyave root found",
		func() error {
			cwd, err := os.Getwd()
			if err != nil {
				return err
			}

			err = moveUp(cwd)
			if err != nil {
				return err
			}

			return nil
		})

	return err

}

func moveUp(curDir string) error {
	ex, err := exists(filepath.Join(curDir, ".goyaveroot"))
	if err != nil {
		return err
	}
	if ex {
		err = os.Chdir(curDir)
		if err != nil {
			return err
		}
	} else {
		if strings.HasSuffix(curDir, string(os.PathSeparator)) {
			return fmt.Errorf(".goyaveroot not found")
		}
		err = moveUp(filepath.Dir(curDir))
		if err != nil {
			return err
		}
	}
	return nil
}

func exists(filename string) (bool, error) {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false, nil
	}
	return !info.IsDir(), err
}
