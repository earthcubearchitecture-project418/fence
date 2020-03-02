package fence

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
	"github.com/mafredri/cdp/devtool"
)

// try to move to https://github.com/mafredri/cdp
// in a headlessv2 version of this function

func headless(urlloc string) (string, error) {

	log.Println("Headless call in flight")

	ctx := context.Background()
	devTools := devtool.New("http://headless:9222")
	pt, err := devTools.Get(ctx, devtool.Page)
	if err != nil {
		pt, err = devTools.Create(ctx)
		if err != nil {
			log.Print(err)
		}
	}

	ctxt, cancel := chromedp.NewRemoteAllocator(context.Background(), pt.WebSocketDebuggerURL)
	defer cancel()

	// Default
	// Create context and headless chrome instances
	// ctxt, cancel := chromedp.NewContext(context.Background())
	// defer cancel()

	// var c string
	// err := chromedp.Run(ctxt, domprocess(urlloc, &c))
	// if err != nil {
	// 	return "", err
	// }

	var c string
	if err := chromedp.Run(ctxt,
		chromedp.Navigate(urlloc),
		chromedp.Sleep(5*time.Second),
		//chromedp.WaitVisible(`#jsonld`, chromedp.ByID), // 	<script id="schemaorg" type="application/ld+json">
		//chromedp.WaitVisible(`script[type="application/ld+json"]`, chromedp.ByQuery), // 	<script id="schemaorg" type="application/ld+json">
		//chromedp.WaitVisible(`script#jsonld"]`, chromedp.ByQuery), // 	<script id="schemaorg" type="application/ld+json">

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

func getDebugURL() string {
	resp, err := http.Get("http://headless:9222/json/version")
	if err != nil {
		log.Print(err)
	}

	var result map[string]interface{}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Print(err)
	}
	return result["webSocketDebuggerUrl"].(string)
}
