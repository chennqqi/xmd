package xmd

import (
"encoding/xml"
)

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