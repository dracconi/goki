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

// type Post struct{
// 	No            int    `json:"no"`
// 	Sub           string `json:"sub,omitempty"`
// 	Com           string `json:"com"`
// 	Name          string `json:"name"`
// 	Capcode       string `json:"capcode,omitempty"`
// 	Time          int    `json:"time"`
// 	OmittedPosts  int    `json:"omitted_posts,omitempty"`
// 	OmittedImages int    `json:"omitted_images,omitempty"`
// 	Sticky        int    `json:"sticky"`
// 	Locked        int    `json:"locked"`
// 	Cyclical      string `json:"cyclical"`
// 	Bumplocked    string `json:"bumplocked"`
// 	LastModified  int    `json:"last_modified"`
// 	ID            string `json:"id"`
// 	TnH           int    `json:"tn_h,omitempty"`
// 	TnW           int    `json:"tn_w,omitempty"`
// 	H             int    `json:"h,omitempty"`
// 	W             int    `json:"w,omitempty"`
// 	Fsize         int    `json:"fsize,omitempty"`
// 	Filename      string `json:"filename,omitempty"`
// 	Ext           string `json:"ext,omitempty"`
// 	Tim           string `json:"tim,omitempty"`
// 	Fpath         int    `json:"fpath,omitempty"`
// 	Spoiler       int    `json:"spoiler,omitempty"`
// 	Md5           string `json:"md5,omitempty"`
// 	Resto         int    `json:"resto"`
// 	Embed         string `json:"embed,omitempty"`
// 	EmbedThumb    string `json:"embed_thumb,omitempty"`
// 	Email         string `json:"email,omitempty"`
// 	ExtraFiles    []struct {
// 		TnH      int    `json:"tn_h"`
// 		TnW      int    `json:"tn_w"`
// 		H        int    `json:"h"`
// 		W        int    `json:"w"`
// 		Fsize    int    `json:"fsize"`
// 		Filename string `json:"filename"`
// 		Ext      string `json:"ext"`
// 		Tim      string `json:"tim"`
// 		Fpath    int    `json:"fpath"`
// 		Spoiler  int    `json:"spoiler"`
// 		Md5      string `json:"md5"`
// 	} `json:"extra_files,omitempty"`
// }

// type Posts struct {
// 	Posts []Post
// }

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

		path := configuration.Output+"/"+name+split[3]+"/"+id
		os.MkdirAll(path, os.ModePerm)

		var endpoint,url,name string

		switch split[2]{
		case "boards.4chan.org":
			endpoint = "https://a.4cdn.org/"+split[3]+"/thread/"+split[len(split)-1]+".json"
			url,name = "https://i.4cdn.org/"+split[3]+"/", "4chan"
		case "8ch.net":
			endpoint = "https://8ch.net/"+split[3]+"/res/"+strings.Split(split[len(split)-1],".")[0]+".json"
			url,name = "https://media.8ch.net/file_store/", "8ch"
		}

		resp, err := http.Get(endpoint)
		if err!=nil{
			panic(err)
		}

		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err!=nil{
			panic(err)
		}

		var posts map[string]interface{}
		//httpdec := json.NewDecoder(resp.Body)
		//httpdec.Decode(&posts)
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