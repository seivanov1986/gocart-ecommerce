package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"regexp"
)

const (
	dirBasePath = "/Users/seivanov/shop/gocart-ecommerce/schemes/public/static"
	webBasePath = "https://shofy-nuxt.vercel.app"
)

func main() {
	content, err := os.ReadFile("file.html")
	if err != nil {
		fmt.Println(err)
		return
	}

	//getImg(string(content))
	getHref(string(content))
}

func getHref(content string) {
	reg, _ := regexp.Compile(`link(.*)href="([a-zA-Z0-9-\/\._]+)"`)
	for _, match := range reg.FindAllStringSubmatch(content, -1) {
		iurl := match[2]
		parsedUrl, err := url.Parse(iurl)
		if err != nil {
			fmt.Println(err)
			continue
		}

		if parsedUrl.Host != "" {
			continue
		}

		fmt.Println(iurl)
		path := path.Dir(iurl)

		fmt.Println(dirBasePath + path)

		err = os.MkdirAll(dirBasePath+path, 0777)
		if err != nil {
			fmt.Println(err)
			continue
		}

		DownloadToFile(webBasePath+iurl, dirBasePath+iurl)

		fmt.Println(webBasePath+iurl, dirBasePath+iurl)
	}
}

func getImg(content string) {
	reg, _ := regexp.Compile(`src="([a-zA-Z0-9-\/\.]+)"`)
	for _, match := range reg.FindAllStringSubmatch(content, -1) {
		iurl := match[1]
		parsedUrl, err := url.Parse(iurl)
		if err != nil {
			fmt.Println(err)
			continue
		}

		if parsedUrl.Host != "" {
			continue
		}

		fmt.Println(iurl)
		path := path.Dir(iurl)

		fmt.Println(dirBasePath + path)

		err = os.MkdirAll(dirBasePath+path, 0777)
		if err != nil {
			fmt.Println(err)
			continue
		}

		DownloadToFile(webBasePath+iurl, dirBasePath+iurl)

		fmt.Println(webBasePath+iurl, dirBasePath+iurl)
	}
}

func HTTPDownload(uri string) ([]byte, error) {
	fmt.Printf("HTTPDownload From: %s.\n", uri)
	res, err := http.Get(uri)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	d, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ReadFile: Size of download: %d\n", len(d))
	return d, err
}

func WriteFile(dst string, d []byte) error {
	fmt.Printf("WriteFile: Size of download: %d\n", len(d))
	err := ioutil.WriteFile(dst, d, 0666)
	if err != nil {
		log.Fatal("1", err)
	}
	return err
}

func DownloadToFile(uri string, dst string) {
	fmt.Printf("DownloadToFile From: %s.\n", uri)
	if d, err := HTTPDownload(uri); err == nil {
		fmt.Printf("downloaded %s.\n", uri)
		if WriteFile(dst, d) == nil {
			fmt.Printf("saved %s as %s\n", uri, dst)
		}
	}
}
