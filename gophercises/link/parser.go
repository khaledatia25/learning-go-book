package link

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)

// Link represents a lik (<a href="...">) in an HTML
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
	if node == nil || node.Data == "scripts" || node.Data == "style" {
		return ""
	}
	if node.Type == html.TextNode {
		return strings.TrimSpace(node.Data)
	}
	var text []string
	for ch := node.FirstChild; ch != nil; ch = ch.NextSibling {
		text = append(text, getText(ch))
	}
	return strings.Join(strings.Fields(strings.Join(text, " ")), " ")
}

func parse(node *html.Node) []Link {
	if node == nil {
		return []Link{}
	}
	var links []Link
	for ch := node.FirstChild; ch != nil; ch = ch.NextSibling {
		if ch.Type == html.ElementNode && ch.Data == "a" {
			links = append(links, Link{
				Href: getHref(ch),
				Text: getText(ch),
			})
		} else {
			links = append(links, parse(ch)...)
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
