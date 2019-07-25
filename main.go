package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type IssuesData struct {
	HtmlUrl         string `json:"html_url"`
	OpenIssuesCount int    `json:"open_issues_count"`
}


func main() {
	userName := os.Args[1]
	userURL := "https://api.github.com/users/" + userName + "/repos"

	gitResp := getGitResp(userURL)
	filteredData := filterGitResp(gitResp)
	printData(filteredData)
}

func printData(data []IssuesData) {
	for _, d := range data {
		fmt.Printf("%d Issues in %s\n", d.OpenIssuesCount, d.HtmlUrl)
	}
}

func filterGitResp(g []IssuesData) []IssuesData {
	nid := []IssuesData{}
	for _, id := range g {
		if id.OpenIssuesCount > 0 {
			nid = append(nid, id)
		}
	}
	return nid
}

func getGitResp(userURL string) []IssuesData {
	resp, err := http.Get(userURL)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	var data []IssuesData
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	return data
}
