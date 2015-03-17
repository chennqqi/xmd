package main

import (
"flag"
"fmt"
"os"
"io/ioutil"
"archive/zip"
// "path/filepath"
// "strings"
"io"
// "log"
"encoding/xml"
"bytes"
)

const (
	gitConfigPath = ".git/config"
	commitCommandPath = ".git/hooks/xmd-commit"
	zipTmpOutput = "tmp/out/"
)

var (
	commitCommand = []byte("echo 'xmd was here'\n")
)

var usage = `Usage: xmd [options...]

Options:
  init  setup current git repo to use xmd.
  to    convert file from docx to xmd.
  from  convert file from xmd to docx.
`

type WordDocument struct {
	XMLName   xml.Name   `xml:"document"`
	Background Background `xml:"background"`
	Body Body `xml:"body"`
}
type Body struct {
	Paragraph P `xml:"p"`
	SectPr SectPr `xml:"sectPr"`
}
type SectPr struct {
	PgSz PgSz `xml:"pgSz"`
	PgMar PgMar `xml:"pgMar"`
	PgNumType PgNumType `xml:"pgNumType"`
}
type PgNumType struct {
	Start string `xml:"start,attr"`
}
type PgMar struct {
	Left string `xml:"left,attr"`
	Right string `xml:"right,attr"`
	Top string `xml:"top,attr"`
	Bottom string `xml:"bottom,attr"`
}

type PgSz struct {
	W string`xml:"w,attr"`
	H string `xml:"h,attr"`

}
type P struct {
	RsidP string `xml:"rsidP,attr"`
	RsidRPr string `xml:"rsidRPr,attr"`
	RsidR string `xml:"rsidR,attr"`
	RsidDel string `xml:"rsidDel,attr"`
	RsidRDefault string `xml:"rsidRDefault,attr"`
	PPR PPR `xml:"pPr"`
	R R `xml:"r"`


}
type PPR struct {
	ContextualSpacing ContextualSpacing `xml:"contextualSpacing"`
}

type ContextualSpacing struct {
	Val string `xml:"val,attr"`
}

type R struct {
	RsidRPr string `xml:"rsidRPr,attr"`
	RsidR string `xml:"rsidR,attr"`
	RsidDel string `xml:"rsidDel,attr"`
	RPr RPr `xml:"rPr"`
	T T `xml:"t"`
}

type RPr struct {
	Rtl Rtl `xml:"rtl"`
}
type Rtl struct {
	Val string `xml:"val,attr"`
}

type T struct {
	Space string `xml:"space,attr"`
	Name string `xml:",chardata"`
}
type Background struct {
	Color string `xml:"color,attr"`
}

func main(){

	flag.Usage = func() {
		fmt.Fprint(os.Stderr, usage)
	}

	flag.Parse()


	if flag.NArg() < 1 {
		fmt.Println("no args...")
		usageAndExit("")
	} else {
		if flag.Arg(0) == "init" {
			setupGit()
		} else if flag.Arg(0) == "to" {
			// fmt.Println("to")			
			unzip(flag.Arg(1))
		} else if flag.Arg(0) == "from" {
			//TODO
		} else {
			usageAndExit("")
		}
	}
}

func unzip(src string) error {
	if len(src) > 0 {

		fileExists, err := exists(src)
		check(err)
		if fileExists {
			r, err := zip.OpenReader(src)
			check(err)
		    defer r.Close()
		    for _, f := range r.File {

		    	if f.Name == "word/document.xml"{
		    		fmt.Println(f.Name)
		    		rc, err := f.Open()
		            check(err)
		            defer rc.Close()

					buf := bytes.NewBuffer(nil)
					io.Copy(buf, rc)
					s := string(buf.Bytes())

					// fmt.Println(s)

					v := new(WordDocument)
					err = xml.Unmarshal([]byte(s), v)
					if err != nil {
				        fmt.Printf("error: %v", err)
				    } else {
				    	fmt.Printf("%#v\n", v)
				    }
				    
		    	}          
		    }
	    } else {
	    	fmt.Println(src+" not found")
	    }
	} else {
		fmt.Println("no file provided")
	}
    return nil
    
}

func setupGit(){
	configExists, err := exists(gitConfigPath)
	check(err)
	if configExists {
		//write commit hook
		err := writeBytesToFile(commitCommandPath,commitCommand)
		if err == nil {
			fmt.Println("written commit hook")
		}
		check(err)

		//TODO write pull/clone hook

	} else {
		fmt.Println("its not a valid git repo")
	}
}

func writeBytesToFile(path string, bytes []byte) error {
	err := ioutil.WriteFile(path, bytes, 0644)
    return err
}

func haveArg(a string) bool {
	if stringInSlice(a, flag.Args()) {
		return true
	} 
	return false
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func usageAndExit(message string) {
	if message != "" {
		fmt.Fprintf(os.Stderr, message)
		fmt.Fprintf(os.Stderr, "\n\n")
	}
	flag.Usage()
	fmt.Fprintf(os.Stderr, "\n")
	os.Exit(1)
}

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil { return true, nil }
	if os.IsNotExist(err) { return false, nil }
	return false, err
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}
