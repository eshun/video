package request

import (
	"net/http"
	"time"
	"bytes"
	"io"
	"github.com/PuerkitoBio/goquery"
)

func Get(url string) (*goquery.Document) {
	client := http.Client{Timeout: 5 * time.Second}
	resp, error := client.Get(url)
	defer resp.Body.Close()
	if error != nil {
		panic(error)
	}

	var buffer [512]byte
	result := bytes.NewBuffer(nil)
	for {
		n, err := resp.Body.Read(buffer[0:])
		result.Write(buffer[0:n])
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
	}

	doc, err := goquery.NewDocumentFromReader(result)
	if err != nil {
		panic(error)
	}

	return doc
}