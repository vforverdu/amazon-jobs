package main

import (
	"amazon-jobs/jobs"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

func main() {
	baseURL := getURL()
	findJobs(baseURL)
}

func getURL() string {
	baseURL := &url.URL{
		Scheme: "https",
		Host:   "www.amazon.jobs",
		Path:   "/en/search.json",
	}

	values := &url.Values{}
	values.Set("offset", "0")
	values.Set("result_limit", "100")
	values.Set("sort", "recent")
	// values.Set("loc_query", "Remote")
	values.Set("base_query", "Software Development")

	baseURL.RawQuery = values.Encode()
	return baseURL.String()
}

func findJobs(baseURL string) {
	response, err := http.Get(baseURL)

	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {
		body, err := ioutil.ReadAll(response.Body)

		if err != nil {
			log.Fatal(err)
		}

		jobResponse, err := jobs.NewJobResponse(body)
		jobs := jobResponse.GetJobs()

		jobsJSON, err := json.Marshal(jobs)

		if err != nil {
			log.Fatal(err)
		}

		file, err := os.Create("data/jobs.json")

		if err != nil {
			log.Fatal(err)
		}

		defer file.Close()

		_, err = file.Write(jobsJSON)

		if err != nil {
			log.Fatal(err)
		}
	}
}
