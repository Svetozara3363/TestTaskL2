package main

import (
	"flag"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/opesun/goquery"
)

/*
Реализовать утилиту wget с возможностью скачивать сайты целиком.
*/

// Парсим URL сайта, чтобы использлвать как имя для нового файла
func fileNameParse(site string) string {
	urls := strings.Split(site, "/")
	return urls[2] + ".html"
}

func download(site string) {
	// Отправляем get запрос
	resp, err := http.Get(site)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	fileName := fileNameParse(site)

	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	// Записываем в файл тело ответа
	_, err = io.Copy(file, resp.Body)
}

func parseResources(site string) {
	x, _ := goquery.ParseUrl(site)
	for _, url := range x.Find("").Attrs("href") {
		var str []string
		switch {
		case strings.Contains(url, ".png"):
			str = strings.Split(url, "/")
			downloadResources(str[len(str)-1], url)
		case strings.Contains(url, ".jpg"):
			str = strings.Split(url, "/")
			downloadResources(str[len(str)-1], url)
		case strings.Contains(url, ".css"):
			str = strings.Split(url, "/")
			downloadResources(str[len(str)-1], url)
		}
	}
}

func downloadResources(filepath string, url string) error {
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

func main() {
	site := flag.String("s", "https://www.youtube.com/", "site")

	flag.Parse()

	// Проверяем, что аргументов передан сайт
	if ok, err := regexp.MatchString("^(http|https)://", *site); ok == true && err == nil {
		download(*site)
		parseResources(*site)
	} else {
		log.Fatal("invalid url")
	}

}
