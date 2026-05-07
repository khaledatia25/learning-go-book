package link

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)

// Link represents a link (<a href="...">) in an HTML
// document.
type Link struct {
	Href string
	Text string
}

func getHref(node *html.Node) string {
	for _, v := range node.Attr {
		if v.Key == "href" {
			return v.Val
		}
	}
	return ""
}

func getText(node *html.Node) string {
	if node.Type == html.TextNode {
		return node.Data
	}
	if node.Type != html.ElementNode {
		return ""
	}
	var text string
	for ch := node.FirstChild; ch != nil; ch = ch.NextSibling {
		text += getText(ch)
	}
	return strings.Join(strings.Fields(text), " ")
}

func parse(node *html.Node) []Link {
	if node == nil {
		return nil
	}
	var links []Link
	for ch := node.FirstChild; ch != nil; ch = ch.NextSibling {
		if ch.Type == html.ElementNode && ch.Data == "a" {
			links = append(links, Link{
				Href: getHref(ch),
				Text: getText(ch),
			})
		} else {
			tmpLinks := parse(ch)
			if tmpLinks != nil {
				links = append(links, parse(ch)...)
			}
		}
	}
	return links
}

// Parse will take in an HTML document and will return a
// slice of links parsed from it.
func Parse(reader io.Reader) ([]Link, error) {
	doc, err := html.Parse(reader)
	if err != nil {
		return nil, err
	}
	return parse(doc), nil
}
