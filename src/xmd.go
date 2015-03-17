package main

import (
"flag"
"fmt"
"os"
"archive/zip"
"io"
"encoding/xml"
"bytes"
xmd "./xmd"
)

const (
	zipTmpOutput = "tmp/out/"
)

var usage = `Usage: xmd [options...]

Options:
  init  setup current git repo to use xmd.
  to    convert file from docx to xmd.
  from  convert file from xmd to docx.
`



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
			xmd.SetupGit()
		} else if flag.Arg(0) == "to" {
			// fmt.Println("to")			
			from(flag.Arg(1))
		} else if flag.Arg(0) == "from" {
			//TODO
		} else {
			usageAndExit("")
		}
	}
}

func from(src string) error {
	if len(src) > 0 {

		fileExists, err := xmd.Exists(src)
		xmd.Check(err)
		if fileExists {
			r, err := zip.OpenReader(src)
			xmd.Check(err)
		    defer r.Close()
		    for _, f := range r.File {

		    	if f.Name == "word/document.xml"{
		    		fmt.Println(f.Name)
		    		rc, err := f.Open()
		            xmd.Check(err)
		            defer rc.Close()

					buf := bytes.NewBuffer(nil)
					io.Copy(buf, rc)
					s := string(buf.Bytes())


					v := new(xmd.WordDocument)
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


func usageAndExit(message string) {
	if message != "" {
		fmt.Fprintf(os.Stderr, message)
		fmt.Fprintf(os.Stderr, "\n\n")
	}
	flag.Usage()
	fmt.Fprintf(os.Stderr, "\n")
	os.Exit(1)
}


