package cmd

import (
	"archive/zip"
	"fmt"
	"github.com/mingrammer/cfmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/urfave/cli/v2"
)

var initCmd = &cli.Command{
	Name: "init",
	Action: func(c *cli.Context) error {
		if c.NArg() != 1 {
			return fmt.Errorf("please provide exactly one argument")
		}
		//  check if directory name already exists
		wd, err := os.Getwd()
		if err != nil {
			return err
		}

		name := strings.ToLower(c.Args().First())
		dir := filepath.Join(wd, name)
		_, err = os.Stat(dir)
		zip := filepath.Join(wd, "test.zip")
		if !os.IsNotExist(err) {
			return fmt.Errorf("target directory %s already exists", dir)
		}
		// Download template project (https://github.com/System-Glitch/goyave-template/archive/master.zip)

		err = spin("Downloading template project",
			"Downloaded template project",
			func() error {
				// return downloadFile(zip, "https://github.com/System-Glitch/goyave-template/archive/master.zip")
				return downloadFile(zip, "https://github.com/Wulfheart/goyave-template/archive/master.zip")
			})
		if err != nil {
			return err
		}
		// unzip template
		err = spin("Unzipping template project",
			"Unzipped template project",
			func() error {
				return unzip(zip, wd)
			})
		if err != nil {
			return err
		}

		// rename folder
		err = spin("Cleaning folder up",
			"Cleaned folder up",
			func() error {
				err = os.Rename(filepath.Join(wd, "goyave-template-master"), dir)
				if err != nil {
					return err
				}

				err = os.Remove(zip)
				if err != nil {
					return err
				}
				return nil
			})
		if err != nil {
			return err
		}

		// rename goyave_template to name in files
		err = spin("Renaming",
			"Renamed",
			func() error {
				return filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
					if err != nil {
						return err
					}

					if info.IsDir() {
						return nil
					}
					// Assuming that here are only files
					ext := filepath.Ext(path)
					matched, err := regexp.MatchString("(.go)|(.mod)|(.json)", ext)
					if err != nil {
						return err
					}
					if matched {
						read, err := ioutil.ReadFile(path)
						if err != nil {
							return err
						}

						newContents := strings.Replace(string(read), "goyave_template", name, -1)

						err = ioutil.WriteFile(path, []byte(newContents), 0)
						if err != nil {
							return err
						}
					}
					return nil
				})
			})
		// copy config.example.json to config.json
		err = spin("Creating config", "Created config", func() error {
			data, err := ioutil.ReadFile(filepath.Join(dir, "config.example.json"))
			if err != nil {
				return err
			}
			// Write data to dst
			err = ioutil.WriteFile(filepath.Join(dir, "config.json"), data, 0644)
			if err != nil {
				return err
			}
			return nil
		})
		if err != nil {
		    return err
		}

		_, _ = cfmt.Successln("Initialized project at ", dir)
		return nil
	},
	Usage:     "Initializes a new project",
	UsageText: "Text is a little bit longer",
}

// reference https://golangcode.com/download-a-file-from-a-url/
func downloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

// reference https://stackoverflow.com/questions/20357223/easy-way-to-unzip-file-with-golang
func unzip(src, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer func() {
		if err := r.Close(); err != nil {
			panic(err)
		}
	}()

	os.MkdirAll(dest, 0755)

	// Closure to address file descriptors issue with all the deferred .Close() methods
	extractAndWriteFile := func(f *zip.File) error {
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer func() {
			if err := rc.Close(); err != nil {
				panic(err)
			}
		}()

		path := filepath.Join(dest, f.Name)

		if f.FileInfo().IsDir() {
			os.MkdirAll(path, f.Mode())
		} else {
			os.MkdirAll(filepath.Dir(path), f.Mode())
			f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer func() {
				if err := f.Close(); err != nil {
					panic(err)
				}
			}()

			_, err = io.Copy(f, rc)
			if err != nil {
				return err
			}
		}
		return nil
	}

	for _, f := range r.File {
		err := extractAndWriteFile(f)
		if err != nil {
			return err
		}
	}

	return nil
}
