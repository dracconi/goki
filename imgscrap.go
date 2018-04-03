package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	//"strings"
	"io"
)

func fetchImage(u string, p string, v bool) {
	//pa := strings.Split(u,"/")
	if v == true {
		fmt.Println(":! Check for " + u)
	}
	if _, err := os.Stat(p); os.IsNotExist(err) {
		fmt.Println("<- Fetching " + u)
		resp, err := http.Get(u)
		if err != nil {
			panic(err)
		}

		defer resp.Body.Close()
		// path/to/whatever does not exist
		fmt.Printf("-> Saving...")

		file, err := os.Create(p)
		if err != nil {
			panic(err)
		}

		_, err = io.Copy(file, resp.Body)
		if err != nil {
			panic(err)
		}
		s, _ := file.Stat()
		fmt.Printf(" || Saved " + strconv.FormatInt(s.Size()/1000, 10) + "kB \n")
		file.Close()
	}

}
