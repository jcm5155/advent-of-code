package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"text/template"
)

const (
	inputUrl     = "https://adventofcode.com/%v/day/%v/input"
	downloadPath = "./%v/inputs/day%v.input"
	solutionPath = "./%v/day%v.go"
)

func main() {
	year := flag.String("y", "2021", "year")
	day := flag.String("d", "1", "day")
	flag.Parse()
	fullUrl := fmt.Sprintf(inputUrl, *year, *day)
	savepath := fmt.Sprintf(downloadPath, *year, *day)

	// check if .input already exists
	_, err := os.Open(savepath)
	if err == nil {
		log.Fatalf("%v already exists!", savepath)
	}

	savefile, err := os.Create(savepath)
	if err != nil {
		log.Fatal(err)
	}

	// send request
	var client http.Client
	req, _ := http.NewRequest(http.MethodGet, fullUrl, nil)
	req.AddCookie(&http.Cookie{Name: "session", Value: os.Getenv("AOC_SESSION")})
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	// dump to .input file
	io.Copy(savefile, res.Body)
	fmt.Printf("created %v!\n", savepath)

	// create .go file
	sfn := fmt.Sprintf(solutionPath, *year, *day)
	_, err = os.Open(sfn)
	if err != nil {
		solutionFile, _ := os.Create(sfn)
		tmpl, _ := template.ParseFiles("common/templates/day.tpl")
		_ = tmpl.Execute(solutionFile, map[string]interface{}{
			"Day":  day,
			"Year": year,
		})
		fmt.Printf("created %v!\n", sfn)
	}
}
