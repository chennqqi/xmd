package xmd

import (
"os"
"io/ioutil"
)

func Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil { return true, nil }
	if os.IsNotExist(err) { return false, nil }
	return false, err
}

func WriteBytesToFile(path string, bytes []byte) error {
	err := ioutil.WriteFile(path, bytes, 0644)
    return err
}

func Check(e error) {
    if e != nil {
        panic(e)
    }
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
