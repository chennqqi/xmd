package xmd

import (
"encoding/xml"
)

type Document struct {
	XMLName xml.Name `xml:"document"`

	W string `xml:"w,attr"`
    A string `xml:"a,attr"`
    C string `xml:"c,attr"`
    Dgm string `xml:"dgm,attr"`
    Lc string `xml:"lc,attr"`
    M string `xml:"m,attr"`
    Mc string `xml:"mc,attr"`
    O string `xml:"o,attr"`
    Pic string `xml:"pic,attr"`
    R string `xml:"r,attr"`
    Sl string `xml:"sl,attr"`
    V string `xml:"v,attr"`
    W10 string `xml:"w10,attr"`
    Wne string `xml:"wne,attr"`
    Wp string `xml:"wp,attr"`
    Wpg string `xml:"wpg,attr"`
    Wps string `xml:"wps,attr"`

	Background Background `xml:"background"`
	Body Body `xml:"body"`
}
type Background struct {
	Color string `xml:"color,attr"`
}
type Body struct {
	Paragraphs []P `xml:"p"`
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
	Text string `xml:",chardata"`
}





