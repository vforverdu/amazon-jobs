package main

import (
	"amazon-jobs/jobs"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	baseURL := "https://www.amazon.jobs/en/search.json?offset=0&result_limit=100&sort=relevant&loc_query=Remote&base_query=Software%20Development"
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

		file, err := os.Create("result.json")

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
