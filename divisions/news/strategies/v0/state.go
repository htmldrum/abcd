package v0

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"regexp"
)

type State struct {
	Name string
	feed_id
}

type StateResult CategoryResult

func (s State) getFeed(client http.Client) ([]Report, error) {
	var reports []Report
	var q StateResult
	state_url := stateRE.ReplaceAllString(string(stateFeedTemplate), string(s.feed_id))

	resp, err := client.Get(state_url)
	if err != nil {
		return reports, err
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return reports, err
	}
	err = xml.Unmarshal(b, &q)
	if err != nil {
		return reports, err
	}

	for _, i := range q.Channel.Items {
		r := Report{
			Title:       i.Title,
			Link:        i.Link,
			Description: i.Description,
			DcCreator:   i.DcCreator,
			PubDate:     i.PubDate,
			Guid:        i.Guid,
			Category:    i.Category,
		}
		reports = append(reports, r)
	}

	return reports, nil
}

var (
	stateFeedTemplate = feed_template("http://www.abc.net.au/news/feed/%feed_id/rss.xml")
	stateRE           = regexp.MustCompile("%feed_id")
	NewSouthWales     = State{
		"New South Wales",
		feed_id("52498"),
	}
	Victoria = State{
		"Victoria",
		feed_id("54242"),
	}
	Queensland = State{
		"Queensland",
		feed_id("50990"),
	}
	WesternAustralia = State{
		"Western Australia",
		feed_id("52764"),
	}
	SouthAustralia = State{
		"South Australia",
		feed_id("54702"),
	}
	Tasmania = State{
		"Tasmania",
		feed_id("50042"),
	}
	AustralianCapitalTerritory = State{
		"Australian Capital Territory",
		feed_id("48320"),
	}
	NorthernTerritory = State{
		"Northern Territory",
		feed_id("53408"),
	}
	States = []State{NewSouthWales, Victoria, Queensland, WesternAustralia, SouthAustralia, Tasmania, AustralianCapitalTerritory, NorthernTerritory}
)
