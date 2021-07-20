package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

func Make_URL_by_index(template string, index int) string {
	A := ""
	rep := regexp.MustCompile(`(\*)+`)
	A = rep.ReplaceAllString(template, strconv.Itoa(index))

	return A
}

func download_run(url string, index int, root_path,format string) {
	res, err := http.Get(url)

	if err != nil {
		log.Printf("The file can not be downloaded by \"%s\"\n", url)
		return
	}
	defer res.Body.Close()

	/*
		buf := new(bytes.Buffer)
		io.Copy(buf, res.Body)

		_, format, err := image.DecodeConfig(buf)
		format += "."
		if err != nil {
			log.Println(err)
			format = ""
		}
	*/

	/*
		rep := regexp.MustCompile(`\.(.+?)$`)
		format := rep.FindString(url)
		fmt.Println(format)
	*/
	//format := filepath.Ext(url)

	out_file, err := os.Create(filepath.Join(root_path, strconv.Itoa(index)+format))

	defer out_file.Close()

	io.Copy(out_file, res.Body)
}

func download_list_by_HTML(url string) []string {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatalln(err)
	}
	A := []string{}
	doc.Find("img").Each(func(_ int, s *goquery.Selection) {
		s_url, _ := s.Attr("src")
		A = append(A, s_url)
	})

	return A
}
