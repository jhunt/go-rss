package rss

import (
	"encoding/xml"
)

type Feed struct {
	XMLName          xml.Name `xml:"rss"                 json:"-"`
	Version          string   `xml:"version,attr"        json:"version"`
	ContentNamespace string   `xml:"xmlns:content,attr"  json:"-"`

	Channel          Channel  `json:"channel"`
}

type Content struct {
	XMLName xml.Name `xml:"content:encoded" json:"-"`
	Content string   `xml:",cdata"          json:"content"`
}

type Image struct {
	XMLName xml.Name `xml:"image"             json:"-"`
	Url     string   `xml:"url"               json:"url"`
	Title   string   `xml:"title"             json:"title"`
	Link    string   `xml:"link"              json:"link"`
	Width   int      `xml:"width,omitempty"   json:"width,omitempty"`
	Height  int      `xml:"height,omitempty"  json:"height,omitempty"`
}

type TextInput struct {
	XMLName     xml.Name `xml:"textInput"    json:"-"`
	Title       string   `xml:"title"        json:"title"`
	Description string   `xml:"description"  json:"description"`
	Name        string   `xml:"name"         json:"name"`
	Link        string   `xml:"link"         json:"link"`
}

type Channel struct {
	XMLName        xml.Name `xml:"channel"                  json:"-"`
	Title          string   `xml:"title"                    json:"title"`
	Link           string   `xml:"link"                     json:"link"`
	Description    string   `xml:"description"              json:"description"`
	Language       string   `xml:"language,omitempty"       json:"language,omitempty"`
	Copyright      string   `xml:"copyright,omitempty"      json:"copyright,omitempty"`
	ManagingEditor string   `xml:"managingEditor,omitempty" json:"managingEditor,omitempty"`
	WebMaster      string   `xml:"webMaster,omitempty"      json:"webMaster,omitempty"`
	PubDate        string   `xml:"pubDate,omitempty"        json:"pubDate,omitempty"`
	LastBuildDate  string   `xml:"lastBuildDate,omitempty"  json:"lastBuildDate,omitempty"`
	Category       string   `xml:"category,omitempty"       json:"category,omitempty"`
	Generator      string   `xml:"generator,omitempty"      json:"generator,omitempty"`
	Docs           string   `xml:"docs,omitempty"           json:"docs,omitempty"`
	Cloud          string   `xml:"cloud,omitempty"          json:"cloud,omitempty"`
	Ttl            int      `xml:"ttl,omitempty"            json:"ttl,omitempty"`
	Rating         string   `xml:"rating,omitempty"         json:"rating,omitempty"`
	SkipHours      string   `xml:"skipHours,omitempty"      json:"skipHours,omitempty"`
	SkipDays       string   `xml:"skipDays,omitempty"       json:"skipDays,omitempty"`

	Image     *Image     `json:"image,omitempty"`
	TextInput *TextInput `json:"textInput,omitempty"`

	Items []*Item `xml:"item" json:"items"`
}

type Item struct {
	XMLName     xml.Name `xml:"item"                json:"-"`
	Title       string   `xml:"title"               json:"title"`
	Link        string   `xml:"link"                json:"link"`
	Description string   `xml:"description"         json:"description"`
	Author      string   `xml:"author,omitempty"    json:"author,omitempty"`
	Category    string   `xml:"category,omitempty"  json:"category,omitempty"`
	Comments    string   `xml:"comments,omitempty"  json:"comments,omitempty"`
	Guid        string   `xml:"guid,omitempty"      json:"guid,omitempty"`
	PubDate     string   `xml:"pubDate,omitempty"   json:"pubDate,omitempty"`
	Source      string   `xml:"source,omitempty"    json:"source,omitempty"`

	Content   *Content   `json:"content,omitempty"`
	Enclosure *Enclosure `json:"enclosure,omitempty"`
}

type Enclosure struct {
	XMLName xml.Name `xml:"enclosure"   json:"-"`
	Url     string   `xml:"url,attr"    json:"url"`
	Length  string   `xml:"length,attr" json:"length"`
	Type    string   `xml:"type,attr"   json:"type"`
}
