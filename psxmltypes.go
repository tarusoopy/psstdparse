package main

import "encoding/xml"

type AI struct {
	XMLName xml.Name `xml:"AI,omitempty" json:"AI,omitempty"`
	Flag    bool     `xml:",chardata" json:",omitempty"`
}

type AV struct {
	XMLName xml.Name `xml:"AV,omitempty" json:"AV,omitempty"`
	string  string   `xml:",chardata" json:",omitempty"`
}

type I64 struct {
	XMLName xml.Name `xml:"I64,omitempty" json:"I64,omitempty"`
	AttrN   string   `xml:"N,attr"  json:",omitempty"`
	Flag    bool     `xml:",chardata" json:",omitempty"`
}

type MS struct {
	XMLName xml.Name `xml:"MS,omitempty" json:"MS,omitempty"`
	I64     *I64     `xml:"http://schemas.microsoft.com/powershell/2004/04 I64,omitempty" json:"I64,omitempty"`
	PR      *PR      `xml:"http://schemas.microsoft.com/powershell/2004/04 PR,omitempty" json:"PR,omitempty"`
}

type Nil struct {
	XMLName xml.Name `xml:"Nil,omitempty" json:"Nil,omitempty"`
}

type Obj struct {
	XMLName   xml.Name `xml:"Obj,omitempty" json:"Obj,omitempty"`
	AttrRefId string   `xml:"RefId,attr"  json:",omitempty"`
	AttrS     string   `xml:"S,attr"  json:",omitempty"`
	MS        *MS      `xml:"http://schemas.microsoft.com/powershell/2004/04 MS,omitempty" json:"MS,omitempty"`
	TN        *TN      `xml:"http://schemas.microsoft.com/powershell/2004/04 TN,omitempty" json:"TN,omitempty"`
}

type Objs struct {
	XMLName     xml.Name `xml:"Objs,omitempty" json:"Objs,omitempty"`
	AttrVersion string   `xml:"Version,attr"  json:",omitempty"`
	Attrxmlns   string   `xml:"xmlns,attr"  json:",omitempty"`
	Obj         *Obj     `xml:"http://schemas.microsoft.com/powershell/2004/04 Obj,omitempty" json:"Obj,omitempty"`
	S           []*S     `xml:"http://schemas.microsoft.com/powershell/2004/04 S,omitempty" json:"S,omitempty"`
}

type PC struct {
	XMLName xml.Name `xml:"PC,omitempty" json:"PC,omitempty"`
	Number  int8     `xml:",chardata" json:",omitempty"`
}

type PI struct {
	XMLName xml.Name `xml:"PI,omitempty" json:"PI,omitempty"`
	Number  int8     `xml:",chardata" json:",omitempty"`
}

type PR struct {
	XMLName xml.Name `xml:"PR,omitempty" json:"PR,omitempty"`
	AttrN   string   `xml:"N,attr"  json:",omitempty"`
	AI      *AI      `xml:"http://schemas.microsoft.com/powershell/2004/04 AI,omitempty" json:"AI,omitempty"`
	AV      *AV      `xml:"http://schemas.microsoft.com/powershell/2004/04 AV,omitempty" json:"AV,omitempty"`
	Nil     *Nil     `xml:"http://schemas.microsoft.com/powershell/2004/04 Nil,omitempty" json:"Nil,omitempty"`
	PC      *PC      `xml:"http://schemas.microsoft.com/powershell/2004/04 PC,omitempty" json:"PC,omitempty"`
	PI      *PI      `xml:"http://schemas.microsoft.com/powershell/2004/04 PI,omitempty" json:"PI,omitempty"`
	SD      *SD      `xml:"http://schemas.microsoft.com/powershell/2004/04 SD,omitempty" json:"SD,omitempty"`
	SR      *SR      `xml:"http://schemas.microsoft.com/powershell/2004/04 SR,omitempty" json:"SR,omitempty"`
	T       []*T     `xml:"http://schemas.microsoft.com/powershell/2004/04 T,omitempty" json:"T,omitempty"`
}

type S struct {
	XMLName xml.Name `xml:"S,omitempty" json:"S,omitempty"`
	AttrS   string   `xml:"S,attr"  json:",omitempty"`
	string  string   `xml:",chardata" json:",omitempty"`
}

type SD struct {
	XMLName xml.Name `xml:"SD,omitempty" json:"SD,omitempty"`
}

type SR struct {
	XMLName xml.Name `xml:"SR,omitempty" json:"SR,omitempty"`
	Number  int8     `xml:",chardata" json:",omitempty"`
}

type T struct {
	XMLName xml.Name `xml:"T,omitempty" json:"T,omitempty"`
	string  string   `xml:",chardata" json:",omitempty"`
}

type TN struct {
	XMLName   xml.Name `xml:"TN,omitempty" json:"TN,omitempty"`
	AttrRefId string   `xml:"RefId,attr"  json:",omitempty"`
	T         []*T     `xml:"http://schemas.microsoft.com/powershell/2004/04 T,omitempty" json:"T,omitempty"`
}
