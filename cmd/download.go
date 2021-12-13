package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

const (
	puzzleInputUrl         = "https://adventofcode.com/%v/day/%v/input"
	puzzleDownloadLocation = "./%v/inputs/day%v.input"
)

func main() {
	year := flag.String("y", "2021", "year")
	day := flag.String("d", "1", "day")
	flag.Parse()
	fullUrl := fmt.Sprintf(puzzleInputUrl, *year, *day)
	savepath := fmt.Sprintf(puzzleDownloadLocation, *year, *day)

	// send request
	var client http.Client
	req, _ := http.NewRequest(http.MethodGet, fullUrl, nil)
	req.AddCookie(&http.Cookie{Name: "session", Value: os.Getenv("AOC_SESSION")})
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	// dump to file
	savefile, err := os.Create(savepath)
	if err != nil {
		log.Fatal(err)
	}
	io.Copy(savefile, res.Body)
}
