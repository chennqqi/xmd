package xmd

import (
	"encoding/xml"
)

type Document struct {
	XMLName xml.Name `xml:"w:document"`

	W   string `xml:"xmlns:w,attr"`
	A   string `xml:"xmlns:a,attr"`
	C   string `xml:"xmlns:c,attr"`
	Dgm string `xml:"xmlns:dgm,attr"`
	Lc  string `xml:"xmlns:lc,attr"`
	M   string `xml:"xmlns:m,attr"`
	Mc  string `xml:"xmlns:mc,attr"`
	O   string `xml:"xmlns:o,attr"`
	Pic string `xml:"xmlns:pic,attr"`
	R   string `xml:"xmlns:r,attr"`
	Sl  string `xml:"xmlns:sl,attr"`
	V   string `xml:"xmlns:v,attr"`
	W10 string `xml:"xmlns:w10,attr"`
	Wne string `xml:"xmlns:wne,attr"`
	Wp  string `xml:"xmlns:wp,attr"`
	Wpg string `xml:"xmlns:wpg,attr"`
	Wps string `xml:"xmlns:wps,attr"`

	// Background Background `xml:"background"`
	Body Body `xml:"w:body"`
}

// AddParagraph adds a pargraph to
func (d Document) AddParagraph(p P) Document {
	d.Body.Paragraphs = append(d.Body.Paragraphs, p)
	return d
}

type Background struct {
	Color string `xml:"w:color,attr"`
}
type Body struct {
	Paragraphs []P    `xml:"w:p"`
	SectPr     SectPr `xml:"w:sectPr"`
}
type SectPr struct {
	PgSz      PgSz      `xml:"w:pgSz"`
	PgMar     PgMar     `xml:"w:pgMar"`
	PgNumType PgNumType `xml:"w:pgNumType"`
}
type PgNumType struct {
	Start string `xml:"w:start,attr"`
}
type PgMar struct {
	Left   string `xml:"w:left,attr"`
	Right  string `xml:"w:right,attr"`
	Top    string `xml:"w:top,attr"`
	Bottom string `xml:"w:bottom,attr"`
}

type PgSz struct {
	W string `xml:"w:w,attr"`
	H string `xml:"w:h,attr"`
}
type P struct {
	// RsidP        string `xml:"w:rsidP,attr"`
	// RsidRPr      string `xml:"w:rsidRPr,attr"`
	// RsidR        string `xml:"w:rsidR,attr"`
	// RsidDel      string `xml:"w:rsidDel,attr"`
	// RsidRDefault string `xml:"w:rsidRDefault,attr"`
	PPR PPR `xml:"w:pPr"`
	R   R   `xml:"w:r"`
}

// func (p P) MakeItBold() P {
// 	// p.R.RPr.Bold = true
// 	return p
// }
// func (p P) MakeItItalic() P {
// 	// p.R.RPr.Italic = true
// 	return p
// }
// func (p P) MakeItUnderlined() P {
// 	// p.R.RPr.Underlined = true
// 	return p
// }

type PPR struct {
	// ContextualSpacing ContextualSpacing `xml:"contextualSpacing"`
	// <w:rFonts w:ascii="Helvetica" w:hAnsi="Helvetica" w:cs="Helvetica" />
	//  <w:sz w:val="24" />
	//  <w:sz-cs w:val="24" />
}

type ContextualSpacing struct {
	Val string `xml:"w:val,attr"`
}

type R struct {
	RsidRPr string `xml:"w:rsidRPr,attr"`
	RsidR   string `xml:"w:rsidR,attr"`
	RsidDel string `xml:"w:rsidDel,attr"`
	RPr     RPr    `xml:"w:rPr"`
	T       T      `xml:"w:t"`
}

type RPr struct {
	Rtl Rtl `xml:"rtl"`
	// Bold       string `xml:"b"`
	// Italic     string `xml:"i"`
	// Underlined string `xml:"u"`
}
type Rtl struct {
	Val string `xml:"val,attr"`
}

type T struct {
	Space string `xml:"xml:space,attr"`
	Text  string `xml:",chardata"`
}
