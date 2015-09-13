package parser

import (
	"errors"
	"log"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

type HTMLParser struct {
	Content string
}

func (parser *HTMLParser) GetFaviconURL() (string, error) {
	tokenizer := html.NewTokenizer(strings.NewReader(parser.Content))

	var tokenType html.TokenType
	var hasAttributes bool

	var attributeKey []byte
	var attributeValue []byte
	var hasMoreAttributes bool

	hasFavicon := false

	for {
		tokenType = tokenizer.Next()

		if tokenType == html.ErrorToken {
			log.Println("An error occurred while parsing HTML:", tokenizer.Err())
			break
		}

		if tokenType == html.StartTagToken || tokenType == html.SelfClosingTagToken {
			_, hasAttributes = tokenizer.TagName()

			if hasAttributes {
				for {
					attributeKey, attributeValue, hasMoreAttributes = tokenizer.TagAttr()

					if string(attributeKey) == "ref" || string(attributeKey) == "rel" {
						if strings.Contains(string(attributeValue), "shortcut") || strings.Contains(string(attributeValue), "icon") {
							hasFavicon = true
						}
					}

					if string(attributeKey) == "href" && hasFavicon {
						// This only returns the raw value. Need to account for non-normalized URLs
						return string(attributeValue), nil
					}

					if !hasMoreAttributes {
						break
					}
				}
			}
		}
	}

	return "", errors.New("URL not found")
}

func NewHTMLParser(content string) *HTMLParser {
	return &HTMLParser{
		Content: content,
	}
}

func (parser *HTMLParser) NormalizeURL(baseURL, partialURL string) (string, error) {
	result := ""

	parsedBase, err := url.Parse(baseURL)
	if err != nil {
		return result, err
	}

	parsedURL, err := url.Parse(partialURL)
	if err != nil {
		return result, err
	}

	resolvedURL := parsedBase.ResolveReference(parsedURL)
	result = resolvedURL.String()

	return result, nil
}
