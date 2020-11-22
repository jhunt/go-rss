package rss

import (
	"encoding/xml"
	"io"
)

func Read(in io.Reader) (Feed, error) {
	var f Feed
	return f, xml.NewDecoder(in).Decode(&f)
}

func ReadClose(in io.ReadCloser) (Feed, error) {
	f, err := Read(in)
	in.Close()
	return f, err
}

func ReadBytes(b []byte) (Feed, error) {
	var f Feed
	return f, xml.Unmarshal(b, &f)
}

func ReadString(s string) (Feed, error) {
	return ReadBytes([]byte(s))
}
