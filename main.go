package main

import (
	"fmt"
	"os"
	"encoding/json"
	"net/http"
	"strings"
	"io/ioutil"
	"strconv"
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
	for _, el := range configuration.Links {
		split := strings.Split(el,"/")
		id := strings.Split(split[len(split)-1],".")[0]
		var endpoint,url,name string

		switch split[2]{
		case "boards.4chan.org":
			endpoint = "https://a.4cdn.org/"+split[3]+"/thread/"+split[len(split)-1]+".json"
			url,name = "https://i.4cdn.org/"+split[3]+"/", "4chan"
		case "8ch.net":
			endpoint = "https://8ch.net/"+split[3]+"/res/"+strings.Split(split[len(split)-1],".")[0]+".json"
			url,name = "https://media.8ch.net/file_store/", "8ch"
		}

		path := configuration.Output+"/"+name+split[3]+"/"+id
		os.MkdirAll(path, os.ModePerm)

		resp, err := http.Get(endpoint)
		if err!=nil{
			panic(err)
		}

		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err!=nil{
			panic(err)
		}

		var posts map[string]interface{}
		json.Unmarshal(bodyBytes,&posts)

		// fmt.Printf("%v", posts)
		// _ = url
		for _, v := range posts["posts"].([]interface {}) {

			// fmt.Printf("%v\n",k)
			// fmt.Printf("%v\n",v)

			// fmt.Printf("%v \n",v.(map[string]interface {})["tim"])
			var tim string
			switch v.(map[string]interface {})["tim"].(type){
			case float64:
				tim = strconv.FormatFloat(v.(map[string]interface {})["tim"].(float64),'f', 0, 64)
			case int:
				tim = strconv.FormatInt(int64(v.(map[string]interface {})["tim"].(int)), 10)
			case string:
				tim = v.(map[string]interface {})["tim"].(string)
			default:
				continue
			}

			fmt.Printf(tim + "\n")
			// fmt.Printf(url+tim+v.(map[string]interface {})["ext"].(string)+" \n "+split[3]+"/"+tim)
			err := fetchImage(url+tim+v.(map[string]interface {})["ext"].(string),configuration.Output+"/"+name+split[3]+"/"+id+"/"+tim+v.(map[string]interface {})["ext"].(string))
			if err!=""{
				panic(err)
			}
		}

	}

}