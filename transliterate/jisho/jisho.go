package jisho

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"net/url"
	"strings"
)

type (
	JishoConceptMeaning struct {
		Value string   `json:"value"`
		Tags  []string `json:"tags"`
	}

	JishoConcept struct {
		Writing  string                `json:"writing"`
		Reading  string                `json:"reading"`
		Meanings []JishoConceptMeaning `json:"meanings"`
		Tags     []string              `json:"tags"`
	}
)

func main() {
	fmt.Println(SearchJisho("しまう"))
}

// SearchJisho will launch a jisho search request and scrape the response page for JishoConcepts
func SearchJisho(query string) []string {

	// TODO: Kanji results
	// TODO: Inflections

	concepts := make([]JishoConcept, 0)

	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36"),
	)

	c.OnHTML("div.concept_light", func(e *colly.HTMLElement) {
		// create a word + "reading" into a computer friendly format.... <kanji>(<kana reading>)
		writingEl := e.DOM.Find("div.concept_light-readings > div.concept_light-representation > span.text")
		writing := strings.TrimSpace(writingEl.Text())

		reading := ""
		readingEl := e.DOM.Find("div.concept_light-readings > div.concept_light-representation > span.furigana > span")
		readingEl.Each(func(i int, s *goquery.Selection) {
			charReading := s.Text()
			if len(charReading) > 0 {
				reading += string([]rune(writing)[i]) + "(" + s.Text() + ")"
			} else {
				reading += string([]rune(writing)[i])
			}
		})

		// collect entry meanings (note that meaning tags and actual meaning elements are siblings in the same parent node)
		meanings := make([]JishoConceptMeaning, 0)
		meaningsContentEl := e.DOM.Find("div.concept_light-meanings > div.meanings-wrapper")

		meaningsContentEl.Find("div.meaning-wrapper").Each(func(i int, ee *goquery.Selection) {
			if len(ee.Prev().Nodes) > 0 {
				meaningTagsEl := ee.Prev().Get(0)
				shouldReadMeaningTagsEl := true

				// check the class Attr of the sibling. If it's not a meaning-tags element, skip it
				for _, attr := range meaningTagsEl.Attr {
					if attr.Key != "class" {
						continue
					}

					if !strings.Contains(attr.Val, "meaning-tags") {
						shouldReadMeaningTagsEl = false
						break
					} else {
						break
					}
				}

				// sometimes meanings don't have any "meaning tags" at all (we check this above). If that's the
				// case just as an empty string instead...
				meaningTagsElText := ""
				if shouldReadMeaningTagsEl {
					meaningTagsElText = meaningTagsEl.FirstChild.Data // this is a text node...
				}

				// ignore any meaning entry that has "Other forms" or "Notes" as tags. These "special" elements have
				// the same structure as the other meaning elements but don't actually define a meaning.
				if !strings.Contains(meaningTagsElText,
					"Other forms") && !strings.Contains(meaningTagsElText, "Notes") {
					// take into account the possibilty of there not being any meaning tags here as well
					meaningTags := []string{}
					if shouldReadMeaningTagsEl {
						meaningTags = strings.Split(meaningTagsElText, ", ")
					}
					meaning := ee.Find("span.meaning-meaning").Text()
					meanings = append(meanings, JishoConceptMeaning{
						Value: meaning,
						Tags:  meaningTags,
					})
				}
			}
		})

		// collect entry tags
		tags := make([]string, 0)
		tagsEl := e.DOM.Find("div.concept_light-status")
		tagsEl.Find("span.concept_light-tag").Each(func(i int, ee *goquery.Selection) {
			tags = append(tags, strings.TrimSpace(ee.Text()))
		})

		// TODO check for the "more words" link (this means there's another page available)

		scrapedConcept := JishoConcept{
			Writing:  writing,
			Reading:  reading,
			Meanings: meanings,
			Tags:     tags,
		}

		concepts = append(concepts, scrapedConcept)
	})

	jishoUrl := fmt.Sprintf("https://jisho.org/search/%s?page=1", url.QueryEscape(query))
	fmt.Println("scraping jisho page %s", jishoUrl)
	err := c.Visit(jishoUrl)
	if err != nil {
		fmt.Println("error occurred while trying to scrape jisho.org")
	}
	var vals []string

	if len(concepts) > 0 {
		concept := concepts[0]

		for _, v := range concept.Meanings {
			for _, vv := range strings.Split(v.Value, "; ") {
				vals = append(vals, vv)
			}
		}
	}

	return vals
}
