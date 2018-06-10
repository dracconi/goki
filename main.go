package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

func main() {

	type Configuration struct {
		Links  []string
		Output string
	}

	type Chan struct {
		Name     string
		Endpoint string
		Url      string
	}
	verbose := flag.Bool("v", false, "verbose")
	mconf := flag.Bool("make-conf", false, "make config")
	flag.Parse()
	if len(flag.Args()) >= 1 {
		if *mconf != true {
			file, _ := os.Open(flag.Args()[0])
			decoder := json.NewDecoder(file)
			configuration := Configuration{}
			err := decoder.Decode(&configuration)
			if err != nil {
				fmt.Println("error:", err)
			}
			for _, el := range configuration.Links {
				split := strings.Split(el, "/")
				id := strings.Split(split[len(split)-1], ".")[0]
				var c Chan

				switch split[2] {
				case "boards.4chan.org":
					c.Endpoint = "https://a.4cdn.org/" + split[3] + "/thread/" + split[len(split)-1] + ".json"
					c.Url, c.Name = "https://i.4cdn.org/"+split[3]+"/", "4chan"
				case "8ch.net":
					c.Endpoint = "https://8ch.net/" + split[3] + "/res/" + strings.Split(split[len(split)-1], ".")[0] + ".json"
					c.Url, c.Name = "https://media.8ch.net/file_store/", "8ch"
				}

				path := configuration.Output + "/" + c.Name + "/" + split[3] + "/" + id
				os.MkdirAll(path, os.ModePerm)

				resp, err := http.Get(c.Endpoint)
				if err != nil {
					panic(err)
				}

				bodyBytes, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					panic(err)
				}

				var posts map[string]interface{}
				json.Unmarshal(bodyBytes, &posts)

				// fmt.Printf("%v", posts)
				// _ = url
				for _, v := range posts["posts"].([]interface{}) {

					// fmt.Printf("%v\n",k)
					// fmt.Printf("%v\n",v)

					// fmt.Printf("%v \n",v.(map[string]interface {})["tim"])
					var tim string
					switch v.(map[string]interface{})["tim"].(type) {
					case float64:
						tim = strconv.FormatFloat(v.(map[string]interface{})["tim"].(float64), 'f', 0, 64)
					case int:
						tim = strconv.FormatInt(int64(v.(map[string]interface{})["tim"].(int)), 10)
					case string:
						tim = v.(map[string]interface{})["tim"].(string)
					default:
						continue
					}

					// fmt.Printf("Downloading " + tim + " \n")
					// fmt.Printf(url+tim+v.(map[string]interface {})["ext"].(string)+" \n "+split[3]+"/"+tim)
					// fmt.Printf("%d / %d", i+1, len(posts["posts"].([]interface{})))
					fetchImage(c.Url+tim+v.(map[string]interface{})["ext"].(string), configuration.Output+"/"+c.Name+"/"+split[3]+"/"+id+"/"+tim+v.(map[string]interface{})["ext"].(string), *verbose)
				}

			}
		} else if *mconf == true {
			if _, err := os.Stat(flag.Args()[0]); os.IsNotExist(err) {
				configuration := Configuration{Output: "out"}
				reader := bufio.NewReader(os.Stdin)
				fmt.Print("URLs (space sep.): ")
				links, _ := reader.ReadString('\n')
				configuration.Links = strings.Split(links, " ")
				for i := range configuration.Links {
					configuration.Links[i] = strings.TrimSuffix(configuration.Links[i], "\r\n")
					_, err := url.Parse(configuration.Links[i])
					if err != nil {
						panic(err)
					}
				}
				fmt.Print("Output directory: ")
				fmt.Scanln(&configuration.Output)
				f, _ := os.Create(flag.Args()[0])
				encoder := json.NewEncoder(f)
				err := encoder.Encode(configuration)
				if err != nil {
					panic(err)
				}
				f.Close()
			} else {
				fmt.Printf("Remove the file " + flag.Args()[0] + " manually")
			}
		}
	} else {
		fmt.Printf("Specify config JSON as last argument")
	}

}

func Add(a int, b int) int {
	return a + b
}
