package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"strings"

	"github.com/samchentw/weather-icon-download/tools"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
)

func main() {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	sourceUrl := "https://www.cwb.gov.tw"
	var res string
	var err error

	err = chromedp.Run(ctx,
		chromedp.Navigate(sourceUrl+`/V8/C/K/Weather_Icon.html`),
		chromedp.OuterHTML(`.wrapper`, &res, chromedp.NodeVisible, chromedp.ByQuery),
	)

	if err != nil {
		log.Fatal(err)
	}

	dom, err := goquery.NewDocumentFromReader(strings.NewReader(strings.TrimSpace(res)))

	if err != nil {
		log.Fatal(err)
	}

	tools.CreateFileOrRead("早上")
	tools.CreateFileOrRead("晚上")

	dom.Find("img").Each(func(i int, s *goquery.Selection) {
		url, _ := s.Attr("src")
		alt, _ := s.Attr("alt")

		fmt.Printf("Review %d: %s\n", i, strings.TrimSpace(url))
		fileName := alt + ".svg"
		fullUrl := sourceUrl + url

		c := math.Mod(float64(i), 2)
		err := tools.DownloadFile(fullUrl, fileName, c == 0)
		if err != nil {
			log.Fatal(err)
		}
	})
	//  log.Println(dom.Html())
}
