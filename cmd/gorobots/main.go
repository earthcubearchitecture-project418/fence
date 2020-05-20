package main

import (
	"fmt"

	"github.com/jimsmart/grobotstxt"
)

func main() {

	// Contents of robots.txt file.
	robotsTxt := `
    # robots.txt with restricted area

    User-agent: *
    Disallow: /members/*

    Sitemap: http://example.net/sitemap.xml
`

	// Target URI.
	uri := "http://example.net/members/index.html"

	// Is bot allowed to visit this page?
	ok := grobotstxt.AgentAllowed(robotsTxt, "FooBot/1.0", uri)

	fmt.Println(ok)

	sitemaps := grobotstxt.Sitemaps(robotsTxt)

	fmt.Println(sitemaps)

}
