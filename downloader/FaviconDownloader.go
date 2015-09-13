package downloader

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/adampresley/gofavigrab/parser"
)

type FaviconDownloader struct {
	htmlParser *parser.HTMLParser
}

func (download *FaviconDownloader) Download(baseURL string) ([]byte, error) {
	var result []byte
	var rawURL string
	var resolvedURL string
	var err error

	rawURL, err = download.htmlParser.GetFaviconURL()
	if err != nil {
		log.Println("Error getting favicon URL:", err)
		return result, err
	}

	resolvedURL, err = download.htmlParser.NormalizeURL(baseURL, rawURL)
	if err != nil {
		log.Println("Error normalizing URL:", err)
		return result, err
	}

	log.Println(resolvedURL)

	client := &http.Client{}
	response, err := client.Get(resolvedURL)
	if err != nil {
		log.Println("Error perform HTTP get to retrieve favicon:", err)
		return result, err
	}

	if response.StatusCode != 200 {
		log.Println("The return code from getting the favicon was not 200")
		return result, errors.New("HTTP call to get favicon failed")
	}

	result, err = ioutil.ReadAll(response.Body)
	return result, nil
}

func NewFaviconDownloader(htmlParser *parser.HTMLParser) *FaviconDownloader {
	return &FaviconDownloader{
		htmlParser: htmlParser,
	}
}
