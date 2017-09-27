package v0

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"regexp"
)

type Category struct {
	Name string
	feed_id
}

type CategoryResult struct {
	XMLName xml.Name `xml:"rss"`
	Version string   `xml:"version,attr"`
	NS      string   `xml:"xmlns:dc,attr"`
	Media   string   `xml:"xmlns:media,attr"`
	Channel struct {
		XMLName        xml.Name `xml:"channel"`
		Title          string   `xml:"title"`
		Category       string   `xml:"category"`
		Link           string   `xml:"link"`
		Language       string   `xml:"language"`
		Copyright      string   `xml:"copyright"`
		PubDate        string   `xml:"pubDate"`
		LastBuildDate  string   `xml:"lastBuildDate"`
		Docs           string   `xml:"docs"`
		ManagingEditor string   `xml:"managingEditor"`
		Image          struct {
			XMLName xml.Name `xml:"image"`
			Title   string   `xml:"title"`
			Url     string   `xml:"url"`
			Link    string   `xml:"link"`
		}
		Items []struct {
			Title       string `xml:"title"`
			Link        string `xml:"link"`
			Description string `xml:"description"`
			DcCreator   string `xml:"dc:creator"`
			PubDate     string `xml:"pubDate"`
			Guid        string `xml:"guid"`
			Category    string `xml:"category"`
			MediaGroup  []struct {
				MediaDescription string `xml:"media:description"`
				MediaThumbnail   string `xml:"media:thumbnail"`
				MediaContent     []struct {
					Content string `xml:"url,attr"`
					Medium  string `xml:"medium,attr"`
					Type    string `"xml:type,attr"`
					Width   string `"xml:width,attr"`
					Height  string `"xml:height,attr"`
				} `xml:"media:content"`
			} `xml:"media:group"`
		} `xml:"item"`
	}
}

func (c Category) getFeed(client http.Client) ([]Report, error) {
	var reports []Report
	var q CategoryResult
	category_url := categoryRE.ReplaceAllString(string(categoryFeedTemplate), string(c.feed_id))

	resp, err := client.Get(category_url)
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
	categoryFeedTemplate = feed_template("http://www.abc.net.au/news/feed/%feed_id/rss.xml")
	categoryRE           = regexp.MustCompile("%feed_id")
	JustIn               = Category{
		"Just In",
		feed_id("51120"),
	}
	TopStories = Category{
		"Top Stories",
		feed_id("45910"),
	}
	World = Category{
		"World",
		feed_id("52278"),
	}
	Australia = Category{
		"Australia",
		feed_id("46182"),
	}
	Business = Category{
		"Business",
		feed_id("51892"),
	}
	Entertainment = Category{
		"Entertainment",
		feed_id("46800"),
	}
	Sport = Category{
		"Sport",
		feed_id("45924"),
	}
	TheDrum = Category{
		"The Drum",
		feed_id("1054578"),
	}
	Categories = []Category{
		JustIn,
		TopStories,
		World,
		Australia,
		Business,
		Entertainment,
		Sport,
		TheDrum,
	}
)
