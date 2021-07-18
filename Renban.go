package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
)

func Make_URL_by_index(template string, index int) string {
	A := ""
	rep := regexp.MustCompile(`(\*)+`)
	A = rep.ReplaceAllString(template, strconv.Itoa(index))

	return A
}

func download_run(url string, index int, root_path string) {
	res, err := http.Get(url)

	if err != nil {
		log.Printf("The file can not be downloaded by \"%s\"\n", url)
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
	format := filepath.Ext(url)

	out_file, err := os.Create(filepath.Join(root_path, strconv.Itoa(index)+format))

	defer out_file.Close()

	io.Copy(out_file, res.Body)

}
