package xmd

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"log"
	"os"
)

const (
	TmpDir      = "tmpx"
	TemplateDir = "template"
)

// EmptyDocument returns a valid Document without any text or paragraphs
func EmptyDocument() Document {

	eDoc := Document{

		XMLName: xml.Name{
			Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main",
			Local: "document"},
		W:   "http://schemas.openxmlformats.org/wordprocessingml/2006/main",
		A:   "http://schemas.openxmlformats.org/drawingml/2006/main",
		C:   "http://schemas.openxmlformats.org/drawingml/2006/chart",
		Dgm: "http://schemas.openxmlformats.org/drawingml/2006/diagram",
		Lc:  "http://schemas.openxmlformats.org/drawingml/2006/lockedCanvas",
		M:   "http://schemas.openxmlformats.org/officeDocument/2006/math",
		Mc:  "http://schemas.openxmlformats.org/markup-compatibility/2006",
		O:   "urn:schemas-microsoft-com:office:office",
		Pic: "http://schemas.openxmlformats.org/drawingml/2006/picture",
		R:   "http://schemas.openxmlformats.org/officeDocument/2006/relationships",
		Sl:  "http://schemas.openxmlformats.org/schemaLibrary/2006/main",
		V:   "urn:schemas-microsoft-com:vml",
		W10: "urn:schemas-microsoft-com:office:word",
		Wne: "http://schemas.microsoft.com/office/word/2006/wordml",
		Wp:  "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing",
		Wpg: "http://schemas.microsoft.com/office/word/2010/wordprocessingGroup",
		Wps: "http://schemas.microsoft.com/office/word/2010/wordprocessingShape",
		// Background: Background{Color: "FFFFFF"},
		Body: Body{
			Paragraphs: []P{},
			SectPr: SectPr{
				PgSz: PgSz{
					W: "12240",
					H: "15840",
				},
				PgMar: PgMar{
					Left:   "1440",
					Right:  "1440",
					Top:    "1440",
					Bottom: "1440"},
				PgNumType: PgNumType{
					Start: "1",
				},
			},
		},
	}
	return eDoc
}

// NewParagraph returns a P with the input text injexted
func NewParagraph(text string) P {
	return P{
		// RsidP:        "00000000",
		// RsidRPr:      "00000000",
		// RsidR:        "00000000",
		// RsidDel:      "00000000",
		// RsidRDefault: "00000000",
		// PPR: PPR{
		// ContextualSpacing: ContextualSpacing{
		// Val: "0",
		// },
		// },
		R: R{
			// RsidRPr: "00000000",
			// RsidR:   "00000000",
			// RsidDel: "00000000",
			RPr: RPr{
				Rtl: Rtl{
					Val: "0",
				},
			},
			T: T{
				Space: "preserve",
				Text:  text,
			},
		},
	}
}

// ToX is the init function for the md to x conversion
func ToX(src string, outDir string) error {

	if len(src) < 1 {
		return ErrNoInputFile
	}

	fileExists, err := Exists(src)
	Check(err)
	if fileExists {

		doc := EmptyDocument()

		file, err := os.Open(src)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			// output := blackfriday.MarkdownCommon(scanner.Bytes())

			// fmt.Println("line: " + string(output))
			p := NewParagraph(scanner.Text())
			// xml.*SetUseNullEndTag(true)*

			// p = p.MakeItBold()
			// p = p.MakeItItalic()
			// p = p.MakeItUnderlined()

			doc = doc.AddParagraph(p)
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

		CopyDir(TemplateDir, TmpDir)

		//todo push new file
		// slcDlcB, _ := xml.Marshal(doc)
		byteXML, err := xml.MarshalIndent(doc, "", `   `)
		Check(err)
		response := xml.Header + string(byteXML)

		fmt.Println(response)
		//
		// s := strings.Replace(string(slcDlcB), "<b></b>", "<b/>", -1)
		// s = strings.Replace(string(slcDlcB), "<i></i>", "<i/>", -1)
		// s = strings.Replace(string(slcDlcB), "<u></u>", "<u/>", -1)

		// outBytes := []byte(s)

		WriteBytesToFile(TmpDir+"/word/document.xml", []byte(response))

		ZipFolder(TmpDir, outDir)

		RemoveDir(TmpDir)

	} else {
		return ErrFileDoestNotExist
	}

	return nil
}
