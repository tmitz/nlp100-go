package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"regexp"

	"github.com/tmitz/nlp100-go/ch03/020/british"
)

const baseURL = "https://en.wikipedia.org/w/api.php"

type Output struct {
	Query struct {
		Pages struct {
			Number struct {
				Imageinfo []struct {
					URL                 string `json:"url"`
					Descriptionurl      string `json:"descriptionurl"`
					Descriptionshorturl string `json:"descriptionshorturl"`
				} `json:"imageinfo"`
			} `json:"23473560"`
		} `json:"pages"`
	} `json:"query"`
}

func main() {
	file := os.Args[1:]
	body := british.Parse(file[0])
	lines := regexp.MustCompile(`\n[\|}]`).Split(body, -1)

	for _, line := range lines {
		re := regexp.MustCompile(`(?s)^(.*?)\s=\s(.*)`)
		m := re.FindStringSubmatch(line)
		if len(m) > 0 && m[1] == "国旗画像" {
			q := urlQuery("File:" + m[2])
			url := baseURL + "?" + q
			res, err := http.Get(url)
			if err != nil {
				panic(err)
			}
			defer res.Body.Close()

			data, err := ioutil.ReadAll(res.Body)
			if err != nil {
				panic(err)
			}

			var out Output
			err = json.Unmarshal(data, &out)
			if err != nil {
				panic(err)
			}
			fmt.Println(out.Query.Pages.Number.Imageinfo[0].URL)
		}
	}
}

func urlQuery(title string) string {
	v := url.Values{}
	v.Add("action", "query")
	v.Add("titles", title)
	v.Add("prop", "imageinfo")
	v.Add("format", "json")
	v.Add("iiprop", "url")
	return v.Encode()
}
