package main

import (
	"image"
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
	rep := regexp.MustCompile(`\**`)
	A = rep.ReplaceAllString(template, "")

	return A
}

func download_run(url string, index int, root_path string) {
	res, err := http.Get(url)

	defer res.Body.Close()

	if err != nil {
		log.Printf("The file can not be downloaded by \"%s\"\n", url)
	}
	_, format, err := image.DecodeConfig(res.Body)
	format += "."
	if err != nil {
		format = ""
	}
	out_file, err := os.Create(filepath.Join(root_path, strconv.Itoa(index)+format))

	defer out_file.Close()

	io.Copy(out_file, res.Body)

}
