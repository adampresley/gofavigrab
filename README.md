Go Favicon Grabber
==================

Go Favicon Grabber is a little component that, when given a URL, will grab the favicon. It does this by trying to location the **<link>** tag's URL.

Installation
------------
To install open your terminal and execute the following.

```bash
go get github.com/adampresley/gofavigrab
```

Usage
-----
First step is to include the component in your project.

```go
import (
   "github.com/adampresley/gofavigrab/parser"
   "github.com/adampresley/gofavigrab/downloader"
)
```

Once you have imported the component there are three steps to grab a favicon.

0. Read HTML from some location
0. Attempt to parse and retrieve the URL to the favicon from the HTML source
0. Download the favicon

Here is a sample of this in action.

```go
package main

import (
   "io/ioutil"
   "log"
   "net/http"
   "os"

   "github.com/adampresley/gofavigrab/downloader"
   "github.com/adampresley/gofavigrab/parser"
)

func main() {
   /*
    * Let's test by downloading the favicon from my website.
    */
   client := &http.Client{}
   response, err := client.Get("http://adampresley.com")
   if err != nil {
      log.Println("Unable to perform GET to website:", err)
      return
   }

   if response.StatusCode != 200 {
      log.Println("Error getting website HTML")
      return
   }

   html, err := ioutil.ReadAll(response.Body)
   if err != nil {
      log.Println("Problem reading the HTTP body content:", err)
      return
   }

   /*
    * Create a parser to try and get the favicon URL
    */
   htmlParser := parser.NewHTMLParser(string(html))
   url, err := htmlParser.GetFaviconURL()
   if err != nil {
      log.Println("A URL could not be found:", err)
      return
   }

   log.Println("URL located:", url)

   faviconDownloader := downloader.NewFaviconDownloader(htmlParser)
   favicon, err := faviconDownloader.Download("http://adampresley.com")
   if err != nil {
      log.Println("Can't download it!", err)
      return
   }

   ioutil.WriteFile("favicon.ico", favicon, os.ModePerm)
   log.Println("Success!")
}
```


License
-------
The MIT License (MIT)

Copyright (c) 2015 Adam Presley

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.


