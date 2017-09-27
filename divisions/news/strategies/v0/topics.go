package v0

import (
	"bytes"
	"fmt"
	"net/http"
	"regexp"

	"github.com/yhat/scrape"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

var (
	topicsForCategoryByLocationUrl = "http://www.abc.net.au/news/topics/location/%letter"
	topicsForCategoryBySubjectUrl  = "http://www.abc.net.au/news/topics/subject/%letter"

	topic_id_to_nameRE = regexp.MustCompile("-")
	letterRE           = regexp.MustCompile("%letter")

	BASE_HACK = "https://www.abc.net.au/news/topics/"
)

type Topic struct {
	topic_id string
	Url      string
}

func (t Topic) CapitalizedName() string {
	pretty := topic_id_to_nameRE.ReplaceAllString(t.topic_id, " ")
	fst := bytes.ToUpper([]byte(pretty[0:1]))
	return string(fst) + pretty[1:len(pretty)]
}

func ListTopicsWithURL(url string, client http.Client, l string) ([]Topic, error) {
	var topics []Topic

	resp, err := client.Get(url)
	if err != nil {
		return topics, err
	}

	root, err := html.Parse(resp.Body)
	if err != nil {
		return topics, err
	}

	subject_cont, ok := scrape.Find(root, scrape.ByTag(atom.H2))
	if !ok {
		return topics, scrapeError
	}

	/*
	 * TODO
	 * Bug in libraries. Not allowing userland to walk tree correctly.
	 * UL reference is a giant hack to get this to work
	 *
	 */

	ul := subject_cont.NextSibling.NextSibling
	ul_data := scrape.Text(ul)

	tokens := bytes.Split([]byte(ul_data), []byte(" "))

	for _, n := range tokens {
		topic_id := string(n)
		t := Topic{
			topic_id: topic_id,
			Url:      BASE_HACK + topic_id,
		}
		topics = append(topics, t)
	}

	return topics, nil
}

func ListTopicsForLetterByLocation(client http.Client, l string) ([]Topic, error) {
	var topics []Topic

	if len(l) > 1 {
		return topics, fmt.Errorf("Please provide a 1 character string as arg, not: %s", l)
	}
	letter := string(bytes.ToLower([]byte(l[0:1])))

	url := letterRE.ReplaceAllString(topicsForCategoryByLocationUrl, letter)

	return ListTopicsWithURL(url, client, letter)
}

func ListTopicsForLetterBySubject(client http.Client, l string) ([]Topic, error) {
	var topics []Topic

	if len(l) > 1 {
		return topics, fmt.Errorf("Please provide a 1 character string as arg, not: %s", l)
	}
	letter := string(bytes.ToLower([]byte(l[0:1])))
	url := letterRE.ReplaceAllString(topicsForCategoryBySubjectUrl, letter)

	return ListTopicsWithURL(url, client, letter)
}
