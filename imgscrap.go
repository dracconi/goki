package main

import (
	"net/http"
	"os"
	//"strings"
	"io"
)

func fetchImage(u string, p string) string{
	resp, err := http.Get(u)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	//pa := strings.Split(u,"/")

	if _, err := os.Stat(p); os.IsNotExist(err) {
		// path/to/whatever does not exist


		file, err := os.Create(p)
		if err != nil {
			panic(err)
		}

		_, err = io.Copy(file, resp.Body)
		if err != nil {
			return "file copy"
		}
		file.Close()
	}


	return ""
}