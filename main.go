package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/PuerkitoBio/goquery"
)

func main() {
	getWebPageData()
}

func getWebPageData() {
	res, err := http.Get("http://yenerpolat.com/blog/")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("Hata Kodu: %d %s", res.StatusCode, res.Status)
	}

	// Html içeriğini yükledik.
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Aşağıda belirtilen class isimlerine göre sayfa üzerinde arama yaptık
	doc.Find(".article-content .entry-title").Each(func(i int, s *goquery.Selection) {
		// Getirilen her veriyi title a aktardık.
		title := s.Find("a").Text()
		fmt.Printf("Data %d: %s\n", i, title)
	})

}
