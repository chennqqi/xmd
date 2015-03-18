package xmd

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

var (
	//ErrNoInputFile user provided no file to `tox` or `tomd`
	ErrNoInputFile       = errors.New("no input specified")
	ErrFileDoestNotExist = errors.New("file does not exist")
)

// Exists tests is a parh exists
func Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// WriteBytesToFile writes a byte arry to file
func WriteBytesToFile(path string, bytes []byte) error {
	err := ioutil.WriteFile(path, bytes, 0644)
	return err
}

// Check tests an error and sends a panic if its valid
func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func CopyFile(source string, dest string) (err error) {
	sourcefile, err := os.Open(source)
	if err != nil {
		return err
	}

	defer sourcefile.Close()

	destfile, err := os.Create(dest)
	if err != nil {
		return err
	}

	defer destfile.Close()

	_, err = io.Copy(destfile, sourcefile)
	if err == nil {
		sourceinfo, err := os.Stat(source)
		if err != nil {
			err = os.Chmod(dest, sourceinfo.Mode())
		}

	}

	return
}

func CopyDir(source string, dest string) (err error) {

	// get properties of source dir
	sourceinfo, err := os.Stat(source)
	if err != nil {
		return err
	}

	// create dest dir

	err = os.MkdirAll(dest, sourceinfo.Mode())
	if err != nil {
		return err
	}

	directory, _ := os.Open(source)

	objects, err := directory.Readdir(-1)

	for _, obj := range objects {

		sourcefilepointer := source + "/" + obj.Name()

		destinationfilepointer := dest + "/" + obj.Name()

		if obj.IsDir() {
			// create sub-directories - recursively
			err = CopyDir(sourcefilepointer, destinationfilepointer)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			// perform copy
			err = CopyFile(sourcefilepointer, destinationfilepointer)
			if err != nil {
				fmt.Println(err)
			}
		}

	}
	return
}

func RemoveDir(dir string) error {
	err := os.RemoveAll(TmpDir)
	Check(err)
	return nil
}

// func haveArg(a string) bool {
// 	if stringInSlice(a, flag.Args()) {
// 		return true
// 	}
// 	return false
// }

// func stringInSlice(a string, list []string) bool {
// 	for _, b := range list {
// 		if b == a {
// 			return true
// 		}
// 	}
// 	return false
// }
