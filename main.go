package main

import (
	"fmt"
	"os"
	"encoding/json"
	"net/http"
	"golang.org/x/net/html"
	"github.com/yhat/scrape"
	"golang.org/x/net/html/atom"
	"strings"
)

func main() {

	type Configuration struct {
		Links []string
		Output string
	}

	file, _ := os.Open("conf.json")
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Print("goki\noutput folder:"+configuration.Output)

	for _, el := range configuration.Links {
		os.MkdirAll(configuration.Output+"/"+strings.Split(el,"/")[2], os.ModePerm)

		resp, err := http.Get(el)
		if err != nil {
			panic(err)
		}
		root, err := html.Parse(resp.Body)
		if err != nil {
			panic(err)
		}
		resp.Body.Close()

		posts := scrape.FindAll(root, scrape.ByClass("postContainer"))

		for _, foo := range posts {
			postText, err := scrape.Find(foo, scrape.ByClass("fileText"))
			if err==true{
				postUrl, err := scrape.Find(postText,scrape.ByTag(atom.A))

				if err==true{
					err := fetchImage(scrape.Attr(postUrl,"href"),configuration.Output)
					if err!="" {
						panic(err)
					}
				}
			}
		}


	}

}
