package xmd

import (
	"archive/zip"
	"encoding/xml"
	"fmt"
	"io"
	// "encoding/json"
	"bytes"
)

func ToMD(src string, out string) error {
	if len(src) > 0 {

		fileExists, err := Exists(src)
		Check(err)
		if fileExists {
			r, err := zip.OpenReader(src)
			Check(err)
			defer r.Close()
			for _, f := range r.File {

				if f.Name == "word/document.xml" {
					fmt.Println(f.Name)
					rc, err := f.Open()
					Check(err)
					defer rc.Close()

					buf := bytes.NewBuffer(nil)
					io.Copy(buf, rc)
					s := string(buf.Bytes())

					v := new(Document)
					err = xml.Unmarshal([]byte(s), v)
					if err != nil {
						fmt.Printf("error: %v", err)
					} else {
						fmt.Printf("%#v\n", v)
						paragraphs := v.Body.Paragraphs
						for i := range paragraphs {
							fmt.Println(paragraphs[i].R.T.Text)
						}

						slcDlcB, _ := xml.Marshal(v)

						fmt.Println(string(slcDlcB))

						//TODO write out
						if len(out) > 0 {
							// text:=v.Body.Paragraphs[0].R.T.Text
							// fmt.Printf("%#v\n", text)
							// err := WriteBytesToFile(out, []byte(text+"\n"))
							// Check(err)

						}
					}

				}
			}
		} else {
			fmt.Println("'" + src + "' not found")
		}
	} else {
		fmt.Println("no file provided")
	}
	return nil
}
