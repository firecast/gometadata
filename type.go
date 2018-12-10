package gometadata

import (
	"io"
	"net/http"
	"net/url"

	"github.com/PuerkitoBio/goquery"
)

// Metadata represents the metadata of an HTML document
type Metadata struct {
	Document    *goquery.Document
	Title       string
	Description string
	Image       string
}

// NewMetadata is a Metadata constructor that takes a string URL as argument.
// It loads the specified document, parses it, and stores the metadata
func NewMetadata(url string) (*Metadata, error) {
	document, e := goquery.NewDocument(url)
	if e != nil {
		return nil, e
	}
	return newMetadata(document), nil
}

// NewMetadataFromReader returns a Metadata from an io.Reader.
func NewMetadataFromReader(r io.Reader, url *url.URL) (*Metadata, error) {
	document, e := goquery.NewDocumentFromReader(r)
	document.Url = url
	if e != nil {
		return nil, e
	}
	return newMetadata(document), nil
}

// NewMetadataFromResponse is another Metadata constructor that takes an http response as argument.
func NewMetadataFromResponse(res *http.Response) (*Metadata, error) {
	document, e := goquery.NewDocumentFromResponse(res)
	if e != nil {
		return nil, e
	}
	return newMetadata(document), nil
}

// Private constructor, make sure all fields are correctly filled.
func newMetadata(document *goquery.Document) *Metadata {
	m := &Metadata{Document: document}
	m.SetTitle()
	m.SetDescription()
	m.SetImage()
	return m
}
