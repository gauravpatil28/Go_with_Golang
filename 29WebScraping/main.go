package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

type Quote struct {
	Quote  string
	Author string
}

func main() {

	quotes := []Quote{}

	c := colly.NewCollector( //this is the collector which initialises a collector which allows you to collect data from the given domain
		colly.AllowedDomains("quotes.toscrape.com"),
		// this allowed domain allows you to make domains whitelist so that we can
		// access it when we try
	)

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36")
		/*Whenever anyone tries to scrape a site the site will not allow everyone hence we give us as agent and give the agent address
		in the next parameter whcih can be found on google. hence it will allow us to get the info we need*/
		fmt.Println("Visiting", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Response Code", r.StatusCode)
		// response code mhanje je codes astat na site che te i.e 404 page not found, 200 everything good, ani baki bhar pur
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Error", err.Error())
	})

	c.OnHTML(".quote", func(h *colly.HTMLElement) {
		div := h.DOM
		/*Basically he kay karta ki purna quote cha class uschalta ani mag tyachat apan DOM vaprun ek div initialise karu shakto
		  jo ki tya class madhun iterate karel ani mag aplyala je hava aahe to class with the help of 'Find' method anun deil
		  and te kuthlya format madhe hava aahe te hi sangu shakto tyala.
		*/
		quote := div.Find(".text").Text()
		author := div.Find(".author").Text()

		q := Quote{
			Quote:  quote,
			Author: author,
		}

		quotes = append(quotes, q)
		//fmt.Printf("Quote : %s \n Given By : %s \n\n\n", quote, author)
	})

	// c.OnHTML(".text", func(h *colly.HTMLElement) {
	// 	fmt.Println("Quote :", h.Text)
	// })

	// /*Hya dogha functions madhe kay hoil ki apan je kahi tya pahilya param madhe dila aahe '.' chya madhe te tya given
	//   site var jaun to class shodel ani tya class cha data gheun yeil ani display karel
	// */

	// c.OnHTML(".author", func(h *colly.HTMLElement) {
	// 	fmt.Println("Author : ", h.Text)
	// })

	c.Visit("http://quotes.toscrape.com")

	fmt.Println(quotes)
}
