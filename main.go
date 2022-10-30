package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/gocolly/colly"
	elementInfo "github.com/menyasosali/go-scraper-kinopoisk/pkg"
	"log"
	"os"
)

type Movie struct {
	Title      string `json:"title"`
	Year       string `json:"year"`
	Director   string `json:"director,omitempty"`
	Genres     string `json:"genres"`
	MainActors string `json:"main_actors,omitempty"`
}

func main() {
	firstDate := flag.String("firstDate", "2013", "Movies from this year")
	lastDate := flag.String("lastDate", "2017", "Movies before this year")
	crawl(*firstDate, *lastDate)
}

func crawl(firstDate string, lastDate string) {
	fmt.Println("Scraper started working")
	movies := []Movie{}
	c := colly.NewCollector(
		colly.AllowedDomains("kinopoisk.ru", "www.kinopoisk.ru"),
	)

	infoCollector := c.Clone()

	c.OnHTML(".element", func(e *colly.HTMLElement) {
		filmPageUrl := e.ChildAttr("p.pic > a", "href")
		filmPageUrl = e.Request.AbsoluteURL(filmPageUrl)
		infoCollector.Visit(filmPageUrl)
		tmpMovie := Movie{}
		tmpMovie.Title = e.ChildText("p.name > a")
		tmpMovie.Year = e.ChildText("p.name > span")

		{
			text := e.ChildText("div.info > span.gray")
			tmpMovie.Director = elementInfo.Director(text)
			tmpMovie.Genres = elementInfo.Genres(text)
			tmpMovie.MainActors = elementInfo.MainActors(text)
		}
		movies = append(movies, tmpMovie)
	})

	c.OnHTML("li.arr > a", func(e *colly.HTMLElement) {
		nextPage := e.Request.AbsoluteURL(e.Attr("href"))
		c.Visit(nextPage)
	})

	startUrl := fmt.Sprintf("https://www.kinopoisk.ru/s/type/film/list/1/m_act[from_year]/%s/m_act[to_year]/%s/", firstDate, lastDate)
	c.Visit(startUrl)
	writeJSON(movies)
	fmt.Println("The scraper has finished")
}

func writeJSON(data []Movie) {
	f, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Fatal(err)
		return
	}
	_ = os.WriteFile("movies.json", f, 0644)
}
