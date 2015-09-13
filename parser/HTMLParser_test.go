package parser

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestHTMLParser(t *testing.T) {
	Convey("The HTML Parser", t, func() {
		Convey("Can be created with HTML content", func() {
			content := "<html><head><link ref=\"shortcut icon\" href=\"favicon.ico\" /></head><body><p>Testing</p></body></html>"
			expected := &HTMLParser{
				Content: content,
			}

			actual := NewHTMLParser(content)
			So(actual, ShouldResemble, expected)
		})

		Convey("Can grab a URL from a self-closing link tag of type shortcut", func() {
			content := "<html><head><link ref=\"shortcut icon\" href=\"favicon.ico\" /></head><body><p>Testing</p></body></html>"
			parser := NewHTMLParser(content)

			expected := "favicon.ico"
			actual, _ := parser.GetFaviconURL()
			So(actual, ShouldEqual, expected)
		})

		// Should not see this, but there is a lot of bad HTML out there...
		Convey("Can grab a URL from a link tag of type shortcut", func() {
			content := "<html><head><link ref=\"shortcut icon\" href=\"favicon.ico\"></link></head><body><p>Testing</p></body></html>"
			parser := NewHTMLParser(content)

			expected := "favicon.ico"
			actual, _ := parser.GetFaviconURL()
			So(actual, ShouldEqual, expected)
		})

		Convey("Can normalize a relative URL", func() {
			expected := "http://localhost/favicon.ico"
			parser := NewHTMLParser("")

			actual, _ := parser.NormalizeURL("http://localhost", "favicon.ico")
			So(actual, ShouldEqual, expected)
		})
	})
}
