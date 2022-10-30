package elementInfo

//package main

import (
	"strings"
)

func Director(text string) string {
	var firstInd int
	newText := strings.Split(text, "\n")[0]
	if strings.Contains(newText, "реж.") {
		firstInd = strings.Index(newText, "реж.")
		return newText[firstInd:]
	}
	return ""
}

func Genres(text string) string {
	newText := strings.Split(text, "\n")[1]
	if strings.Contains(newText, "(") {
		firstInd := strings.Index(text, "(")
		lastInd := strings.Index(text, ")")
		genres := text[firstInd:lastInd]
		genres = strings.TrimPrefix(genres, "(")
		return genres
	}
	return ""
}

func MainActors(text string) string {
	firstInd := strings.Index(text, ")") + 1
	actors := text[firstInd:]
	actors = strings.TrimSpace(actors)
	return actors
}
