package gometadata

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// SetTitle adds the title to the metadata object
func (m *Metadata) SetTitle() {
	var title string

	// Get title from og: properties
	m.Document.Find("meta").Each(func(i int, s *goquery.Selection) {
		if name, _ := s.Attr("property"); strings.EqualFold(name, "og:title") {
			titleContent, _ := s.Attr("content")
			title = titleContent
			return
		}
	})

	if title != "" {
		m.Title = title
		return
	}

	// Set document <title> tag text
	m.Title = m.Document.Find("title").Text()
}

// SetDescription adds the description to the metadata object
func (m *Metadata) SetDescription() {
	var description string

	// Get title from og: properties
	m.Document.Find("meta").Each(func(i int, s *goquery.Selection) {
		if property, _ := s.Attr("property"); strings.EqualFold(property, "og:description") {
			descContent, _ := s.Attr("content")
			description = descContent
			return
		}

		if name, _ := s.Attr("name"); strings.EqualFold(name, "description") {
			descContent, _ := s.Attr("content")
			description = descContent
		}
	})

	m.Description = description
}

// SetImage adds the image to the metadata object
func (m *Metadata) SetImage() {
	url := m.Document.Url

	var image string
	if url != nil {
		image = url.Scheme + "://" + url.Host + "/favicon.ico"
	}

	// Get title from og: properties
	m.Document.Find("meta").Each(func(i int, s *goquery.Selection) {
		if property, _ := s.Attr("property"); strings.EqualFold(property, "og:image") {
			imageContent, _ := s.Attr("content")
			image = imageContent
			return
		}
	})

	m.Image = image
}
