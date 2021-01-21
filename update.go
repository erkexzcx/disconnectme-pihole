package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

const (
	entitiesLink = "https://raw.githubusercontent.com/disconnectme/disconnect-tracking-protection/master/entities.json"
	servicesLink = "https://raw.githubusercontent.com/disconnectme/disconnect-tracking-protection/master/services.json"
)

var repodirFlag = flag.String("repodir", ".", "path to repo directory")

func main() {
	flag.Parse()

	// Ensure we are in the right location
	if _, err := os.Stat(*repodirFlag + "/update.go"); os.IsNotExist(err) {
		panic("invalid directory")
	}

	deleteBlocklistFiles()
	generateEntitiesFiles()
	generateServicesFiles()
}

type entitiesFileStruct struct {
	Entities map[string]struct {
		Resources []string `json:"resources"`
	} `json:"entities"`
}

func generateEntitiesFiles() {
	content, err := retrieveContents(entitiesLink)
	if err != nil {
		panic(err)
	}

	var efs entitiesFileStruct
	if err := json.Unmarshal(content, &efs); err != nil {
		panic(err)
	}

	// Use map instead of []string to ensure there are no duplicates
	domainsMap := make(map[string]bool, len(efs.Entities)*5)
	for _, entity := range efs.Entities {
		for _, domain := range entity.Resources {
			domainsMap[strings.TrimSpace(domain)] = true
		}
	}

	// Convert this map to []string and sort it
	domainsList := make([]string, 0, len(domainsMap))
	for k := range domainsMap {
		domainsList = append(domainsList, k)
	}
	sort.Strings(domainsList)

	// Write []string to the file
	f, err := os.Create(*repodirFlag + "/entities.txt")
	if err != nil {
		panic(err)
	}
	w := bufio.NewWriter(f)
	for _, domain := range domainsList {
		fmt.Fprintln(w, domain)
	}
	if err = w.Flush(); err != nil {
		panic(err)
	}
	f.Close()
}

type servicesFileStruct struct {
	Categories map[string][]map[string]map[string]interface{} `json:"categories"`
}

func generateServicesFiles() {
	content, err := retrieveContents(servicesLink)
	if err != nil {
		panic(err)
	}

	var sfs servicesFileStruct
	if err := json.Unmarshal(content, &sfs); err != nil {
		panic(err)
	}

	/*
		{
			"license": "Copyright 2010-2020 Disconnect, Inc. / Licens......",
			"categories": {
				"Advertising": [
					{
						"2leep.com": {
						"http://2leep.com/": [
							"2leep.com"
						]
						}
					},
					{
						"2leep.com": {
						"http://2leep.com/": [
							"2leep.com"
						]
						}
					}
				],
				"Content": [
					{
						"2leep.com": {
						"http://2leep.com/": [
							"2leep.com"
						]
						}
					}
				]
			}
		}
	*/

	categoriesAndDomains := make(map[string]map[string]bool)

	// This is nightmare. Do not look to this code at night...
	for category, v := range sfs.Categories {
		for _, vv := range v {
			for _, vvv := range vv {
				for _, vvvv := range vvv {
					// Parse only if vvvv is []interface{}
					switch vvvv.(type) {
					case []interface{}:
						// Convert interface{} to []interface{}
						yetAnotherList := vvvv.([]interface{})
						for _, domainAsInterface := range yetAnotherList {
							// Convert interface{} to string
							domain := domainAsInterface.(string)

							if _, ok := categoriesAndDomains[category]; !ok {
								categoriesAndDomains[category] = make(map[string]bool)
							}
							categoriesAndDomains[category][domain] = true
						}
					}
				}
			}
		}
	}

	// Open main services.txt for writting
	f, err := os.Create(*repodirFlag + "/services.txt")
	if err != nil {
		panic(err)
	}
	w := bufio.NewWriter(f)

	// Create sorted list of categories
	categories := make([]string, 0, len(categoriesAndDomains))
	for category := range categoriesAndDomains {
		categories = append(categories, category)
	}
	sort.Strings(categories)

	// Same for each list of domains in each category + write to files
	for _, category := range categories {
		anotherMap := categoriesAndDomains[category]

		// Open services_CATEGORY.txt for writting
		tmpf, err := os.Create(*repodirFlag + "/services_" + category + ".txt")
		if err != nil {
			panic(err)
		}
		tmpw := bufio.NewWriter(tmpf)

		// Create []string type of list of sorted strings
		domains := make([]string, 0, len(categoriesAndDomains[category]))
		for domain := range anotherMap {
			domains = append(domains, domain)
		}
		sort.Strings(domains)

		// Write to both files
		for _, domain := range domains {
			fmt.Fprintln(w, domain)    // Write to services.txt file
			fmt.Fprintln(tmpw, domain) // Write to services_CATEGORY.txt file
		}

		// Flush & close services_CATEGORY.txt file
		if err = tmpw.Flush(); err != nil {
			panic(err)
		}
		tmpf.Close()
	}

	// Flush & close services.txt file
	if err = w.Flush(); err != nil {
		panic(err)
	}
	f.Close()
}

func retrieveContents(link string) ([]byte, error) {
	resp, err := http.Get(link)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return body, err
}

func deleteBlocklistFiles() {
	files, err := filepath.Glob(*repodirFlag + "/*.txt")
	if err != nil {
		panic(err)
	}
	for _, f := range files {
		if err := os.Remove(f); err != nil {
			panic(err)
		}
	}
}
