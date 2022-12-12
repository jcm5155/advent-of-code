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
	inputUrl         = "https://adventofcode.com/%v/day/%v/input"
	templateFilePath = "common/templates/day.%v.tmpl"

	// go
	downloadPath = "./%v/inputs/day%v.input"
	solutionPath = "./%v/day%v.go"

	// rust
	rustSolutionPath = "./aoc%v/src/bin/%v.rs"
	rustDownloadPath = "./aoc%v/src/bin/inputs/%v.input"
)

func main() {
	year := flag.String("y", "2022", "year")
	day := flag.String("d", "1", "day")
	language := flag.String("l", "rust", "language")
	flag.Parse()

	var downloadLoc, solutionLoc, templateFile string
	switch *language {
	case "go":
		downloadLoc = downloadPath
		solutionLoc = solutionPath
		templateFile = fmt.Sprintf(templateFilePath, "go")
	case "rust", "rs":
		downloadLoc = rustDownloadPath
		solutionLoc = rustSolutionPath
		templateFile = fmt.Sprintf(templateFilePath, "rs")
	default:
		panic("language not found")
	}
	fullUrl := fmt.Sprintf(inputUrl, *year, *day)
	savepath := fmt.Sprintf(downloadLoc, *year, *day)

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

	// create solution file
	sfn := fmt.Sprintf(solutionLoc, *year, *day)
	_, err = os.Open(sfn)
	if err != nil {
		solutionFile, _ := os.Create(sfn)
		tmpl, _ := template.ParseFiles(templateFile)
		_ = tmpl.Execute(solutionFile, map[string]interface{}{
			"Day":  day,
			"Year": year,
		})
		fmt.Printf("created %v!\n", sfn)
	}
}
