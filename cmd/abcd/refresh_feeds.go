package main

import (
	"fmt"
	"net/http"
	"reflect"
	"time"
	"regexp"
	"strings"

	"github.com/yhat/scrape"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

const feed_index = "http://www.abc.net.au/services/rss/programs.htm"

func RefreshFeeds()(feeds []Feed){
	resp, err := http.Get("http://www.abc.net.au/services/rss/programs.htm")
	if err != nil {
		panic(err)
	}
	root, err := html.Parse(resp.Body)
	if err != nil {
		panic(err)
	}
	feeds_data := scrape.FindAll(root, scrape.ByClass("program"))
	for i, feed_data := range feeds_data {
		if i <= 2 {
			fmt.Println("FeedData")
		}
		feed := Feed{}
		feed.Id = uint(i)
		v := reflect.ValueOf(feed)
		t := v.Type()
		// @TODO GORO - pop off fields, async reading of shared mem - provided no resizing of fields?
		for j := 0; j < v.NumField(); j++ {
			field := t.Field(j).Name
			scrapeFor(field, feed_data, &feed)
		}
		feeds = append(feeds, feed)
	}
	return feeds
}

func scrapeFor(fName string, feed_data *html.Node, fp *Feed){
	switch fName {
	case "Name":
		h3, _ := scrape.Find(feed_data, scrape.ByTag(atom.H3))
		a,_ := scrape.Find(h3, scrape.ByTag(atom.A))
		name := scrape.Text(a)
		fp.Name = name
	case "Description":
		p, _ := scrape.Find(feed_data, scrape.ByTag(atom.P))
		desc := scrape.Text(p)
		fp.Description = desc
	case "URI":
		c, _ := scrape.Find(feed_data, scrape.ByClass("rss"))
		uri := scrape.Attr(c, "href")
		fp.URI = uri
	case "Subjects":
		// CamelCase -> Subject
		var subjects []string
		class_names := strings.Split(scrape.Attr(feed_data, "class"), " ")
		for _, v := range class_names {
			m, _ := regexp.MatchString("([A-Z]+[a-z])+", v)
			if v == "TV" || m {
				subjects = append(subjects, v)
			}
		}
		fp.Subjects = subjects
	case "Networks":
		// snake_case -> Network
		var networks []string
		class_names := strings.Split(scrape.Attr(feed_data, "class"), " ")
		for _, v := range class_names {
			m, _ := regexp.MatchString("([a-z]+_)+[a-z]+", v)
			if m {
				networks = append(networks, v)
			}
		}
		fp.Networks = networks
	case "Last_contact_datetime":
		t := time.Now().String()
		fp.Last_contact_datetime = t
	}
}
