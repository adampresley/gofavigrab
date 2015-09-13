package downloader

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/adampresley/gofavigrab/parser"

	. "github.com/smartystreets/goconvey/convey"
)

func TestDownloader(t *testing.T) {
	Convey("Downloader", t, func() {
		Convey("Can download favicon based on URL from link tag", func() {
			content := "<html><head><link ref=\"shortcut icon\" href=\"favicon.ico\" /></head><body><p>Testing</p></body></html>"
			fakeContentBytes := bytes.NewBufferString("fake favicon bytes")

			htmlParser := parser.NewHTMLParser(content)
			faviconDownload := NewFaviconDownloader(htmlParser)

			fakeServer := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
				writer.WriteHeader(200)
				writer.Header().Set("Content-Type", "image/x-ico")
				writer.Write(fakeContentBytes.Bytes())
			}))

			defer fakeServer.Close()

			actual, err := faviconDownload.Download(fakeServer.URL)
			log.Println(err)
			So(actual, ShouldResemble, fakeContentBytes.Bytes())
		})
	})
}
