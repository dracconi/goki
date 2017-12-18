package main

import (
	"net/http"
	"os"
	"strings"
	"io"
)

func fetchImage(u string, p string) string{
	resp, err := http.Get("https:"+u)
	if err != nil {
		return "get request"
	}

	defer resp.Body.Close()

	pa := strings.Split(u,"/")

	file, err := os.Create(p+"/"+pa[len(pa)-2]+"/"+pa[len(pa)-1])
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return "file copy"
	}
	file.Close()

	return ""
}