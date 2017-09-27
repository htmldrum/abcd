package v0

import (
	"fmt"
	"net/http"

	"github.com/yhat/scrape"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

var (
	feedUrl     = "http://www.abc.net.au/news/feeds/rss/"
	topicsUrl   = "http://www.abc.net.au/news/topics/"
	podcastsUrl = "http://www.abc.net.au/news/feeds/"

	scrapeError = fmt.Errorf("Node not found")

	defaultProgramName = "Full Program"
)

type Podcast struct {
	Name        string
	Website     string
	Description string
	Programs    []PodcastProgram
}

type PodcastProgram struct {
	Name     string
	Url      string
	Provider string // iTunes || RSS
}

func ListPodcasts(client http.Client) ([]Podcast, error) {
	var (
		podcasts  []Podcast
		p_ndx     = -1 // Pre-increment on 0-indexed array
		i         = 0  // Counter
		start_ndx = 4  // Scraping value, podcasts dont start until after the 4th element
	)

	resp, err := client.Get(podcastsUrl)
	if err != nil {
		return podcasts, err
	}

	root, err := html.Parse(resp.Body)
	if err != nil {
		return podcasts, err
	}

	p_cont, ok := scrape.Find(root, scrape.ByClass("placed"))
	if !ok {
		return podcasts, scrapeError
	}

	for nxt := p_cont.FirstChild; nxt != nil; nxt = nxt.NextSibling {
		if i > start_ndx { // Podcasts dont start until 5th node
			if node_is_h2(nxt) {
				podcast_name := scrape.Text(nxt)

				if podcast_name == "Terms of Use" {
					break
				}

				website, err := key_in_attr(nxt.FirstChild, "href")
				if err != nil {
					return podcasts, fmt.Errorf("Unable to find url for %s", podcast_name)
				}

				p := Podcast{
					Name:        podcast_name,
					Website:     website,
					Description: scrape.Text(nxt.NextSibling),
				}
				podcasts = append(podcasts, p)
				p_ndx += 1
			}

			if node_is_ul(nxt) {
				switch current_show_name := podcasts[p_ndx].Name; current_show_name {
				case "7.30":
				case "Background Briefing":
				case "The Business":
				case "Insiders":
				case "Lateline":
				case "Law Report":
				case "Offsiders":
				case "Pacific Beat":
				case "Radio National Breakfast":
				case "Rural News":
					url, err := key_in_attr(nxt.FirstChild.FirstChild, "href")
					if err != nil {
						return podcasts, fmt.Errorf("Unable to find url for %s", current_show_name)
					}
					pp := PodcastProgram{
						Name:     defaultProgramName,
						Url:      url,
						Provider: "RSS",
					}
					podcasts[p_ndx].Programs = append(podcasts[p_ndx].Programs, pp)
				case "Countrywide":
				case "DC Washup":
				case "Dishonourable Members":
					url1, err := key_in_attr(nxt.FirstChild.FirstChild, "href")
					if err != nil {
						return podcasts, fmt.Errorf("Unable to find url1 for %s", current_show_name)
					}

					var url2 string

					if nxt.FirstChild.NextSibling.FirstChild.DataAtom == atom.P {
						url2, err = key_in_attr(nxt.FirstChild.NextSibling.FirstChild.FirstChild, "href")
						if err != nil {
							return podcasts, fmt.Errorf("Unable to find url2 for %s", current_show_name)
						}
					} else {
						url2, err = key_in_attr(nxt.FirstChild.NextSibling.FirstChild, "href")
						if err != nil {
							return podcasts, fmt.Errorf("Unable to find url2 for %s", current_show_name)
						}
					}

					pp1 := PodcastProgram{
						Name:     defaultProgramName,
						Url:      url1,
						Provider: "RSS",
					}
					pp2 := PodcastProgram{
						Name:     defaultProgramName,
						Url:      url2,
						Provider: "iTunes",
					}
					podcasts[p_ndx].Programs = append(podcasts[p_ndx].Programs, pp1, pp2)

				case "AM":
				case "Correspondents Report":
				case "PM":
				case "The World Today":

					url1, err := key_in_attr(nxt.FirstChild.FirstChild, "href")
					if err != nil {
						return podcasts, fmt.Errorf("Unable to find url1 for %s", current_show_name)
					}
					url2, err := key_in_attr(nxt.FirstChild.NextSibling.FirstChild, "href")
					if err != nil {
						return podcasts, fmt.Errorf("Unable to find url2 for %s", current_show_name)
					}
					pp1 := PodcastProgram{
						Name:     "Full Program",
						Url:      url1,
						Provider: "RSS",
					}
					pp2 := PodcastProgram{
						Name:     "Individual Stories",
						Url:      url2,
						Provider: "RSS",
					}
					podcasts[p_ndx].Programs = append(podcasts[p_ndx].Programs, pp1, pp2)

				case "Country Hour":
					for li := nxt.FirstChild; li != nil; li = li.NextSibling {
						// The Western Australia cast needs special handling
						var p *html.Node
						if li.FirstChild.DataAtom == atom.P {
							p = li.FirstChild
						} else {
							p = li
						}

						url1, err := key_in_attr(p.FirstChild.NextSibling, "href")
						if err != nil {
							return podcasts, fmt.Errorf("Unable to find url1 for %s", current_show_name)
						}
						url2, err := key_in_attr(p.FirstChild.NextSibling.NextSibling.NextSibling, "href")
						if err != nil {
							return podcasts, fmt.Errorf("Unable to find url2 for %s", current_show_name)
						}
						pp1 := PodcastProgram{
							Name:     p.Data,
							Url:      url1,
							Provider: "RSS",
						}
						pp2 := PodcastProgram{
							Name:     p.Data,
							Url:      url2,
							Provider: "iTunes",
						}
						podcasts[p_ndx].Programs = append(podcasts[p_ndx].Programs, pp1, pp2)
					}
				}
			}
		}
		i += 1
	}

	return podcasts, err
}

func node_is_p(node *html.Node) bool {
	return node.DataAtom == atom.P
}

func node_is_h2(node *html.Node) bool {
	return node.DataAtom == atom.H2
}

func node_is_ul(node *html.Node) bool {
	return node.DataAtom == atom.Ul
}

func node_is_li(node *html.Node) bool {
	return node.DataAtom == atom.Li
}

func key_in_attr(n *html.Node, k string) (string, error) {
	for i := 0; i < len(n.Attr); i++ {
		if n.Attr[i].Key == k {
			return n.Attr[i].Val, nil
		}
	}
	return "", fmt.Errorf("Key not found: %s", k)
}
