package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	doc, err := goquery.NewDocument("https://www.rememberthemilk.com/services/api/methods.rtm")
	if err != nil {
		fmt.Printf("init error: %s\n", err)
		return
	}

	doc.Find("body > main > section > article > div > ul > li > a").Each(func(i int, s *goquery.Selection) {
		href, ok := s.Attr("href")
		if ok {
			api, err := goquery.NewDocument("https://www.rememberthemilk.com" + href)
			if err != nil {
				fmt.Fprint(os.Stderr, err)
				os.Exit(1)
			}

			name := api.Find("body > main > section > article > h1").Text()
			arguments := api.Find("body > main > section > article > div.api-method > dl").First()
			argElements := arguments.Find("dt")

			var params []string
			if len(argElements.Nodes) > 0 {
				argElements.Each(func(i int, s *goquery.Selection) {
					param := s.Find("code").First().Text()
					if param != "api_key" {
						params = append(params, param)
					}
				})
			}

			fmt.Printf("%s %s\n", name, strings.Join(params, ","))

			time.Sleep(3 * time.Second)
			return
		}
	})
}
