package fence

import (
	"context"
	"log"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
)

func headless(urlloc string) (string, error) {

	log.Println("Headless call in flight")

	// Create context and headless chrome instances
	ctxt, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// var c string
	// err := chromedp.Run(ctxt, domprocess(urlloc, &c))
	// if err != nil {
	// 	return "", err
	// }

	var c string
	if err := chromedp.Run(ctxt,
		chromedp.Navigate(urlloc),
		// chromedp.WaitVisible("#logo_homepage_link"),
		chromedp.Sleep(2*time.Second),

		chromedp.OuterHTML("html", &c),
	); err != nil {
		log.Fatalf("Failed getting body: %v", err)
	}

	log.Println(c)

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(c))
	if err != nil {
		// handler error
	}

	var jsonld string
	if err == nil {
		doc.Find("script").Each(func(i int, s *goquery.Selection) {
			val, _ := s.Attr("type")
			if val == "application/ld+json" {
				err = isValid(s.Text())
				if err != nil {
					log.Printf("ERROR: At %s JSON-LD is NOT valid: %s", urlloc, err)
				}
				jsonld = s.Text()
			}
		})
	}

	return jsonld, err

}

func domprocess(targeturl string, res *string) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(targeturl),
		chromedp.Text(`#head`, res, chromedp.ByID),
	}
}
