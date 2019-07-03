package main

import (
	"io"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/labstack/echo"
)

type searchResult struct {
	Entries       []entry
	Similar       []string
	BestGuess     string
	KnowledgeBase *knowledgeBase `json:",omitempty"`
}

type entry struct {
	Link        string
	Title       string
	Description string
}

type knowledgeBase struct {
	Title       string
	Subtitle    string
	Description string `json:",omitempty"`
	Modules     map[string][]string
}

func search(imgURL string, c echo.Context) (*searchResult, error) {
	const searchBase = `https://www.google.com/searchbyimage?hl=en-US&image_url=`

	searchFor := searchBase + url.QueryEscape(imgURL)
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	req, err := http.NewRequest(http.MethodGet, searchFor, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", `Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:61.0) Gecko/20100101 Firefox/61.0`)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	ob, _ := os.Create("test.html")
	defer ob.Close()

	return searchParse(io.TeeReader(resp.Body, ob))
}

func searchParse(r io.Reader) (*searchResult, error) {
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return nil, err
	}

	res := &searchResult{}

	doc.Find(`div.rc`).Each(func(u int, s *goquery.Selection) {
		var entry entry

		as := s.Find(`a`)
		if val, exist := as.Attr("href"); exist && val != "" {
			entry.Link = val
		}
		entry.Title = as.Find(`h3`).Text()
		entry.Description = s.Find(`div.s span`).Text()

		res.Entries = append(res.Entries, entry)
	})

	res.Entries = res.Entries[:len(res.Entries)]

	doc.Find(`div #iur a img`).Each(func(u int, s *goquery.Selection) {
		if val, exist := s.Attr("title"); exist && val != "" {
			res.Similar = append(res.Similar, val)
		}
	})

	res.BestGuess = doc.Find(`a.fKDtNb`).Text()

	kb := doc.Find(`div.ifM9O`)
	if kb != nil && kb.Length() > 0 {
		res.KnowledgeBase = &knowledgeBase{
			Title:       kb.Find(`div[data-attrid="title"]`).Text(),
			Subtitle:    kb.Find(`div[data-attrid="subtitle"]`).Text(),
			Description: kb.Find(`div.mod div.kno-rdesc span`).Text(),
			Modules:     make(map[string][]string),
		}

		kb.Find(`div.mod`).Each(func(u int, s *goquery.Selection) {
			name := s.Find(`span.w8qArf a.fl`).Text()
			s.Find(`span.kno-fv`).Each(func(u int, s *goquery.Selection) {
				res.KnowledgeBase.Modules[name] = append(res.KnowledgeBase.Modules[name], s.Text())
			})
		})
	}

	return res, nil
}
