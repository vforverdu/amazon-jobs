package jobs

import (
	"encoding/json"
	"strings"
)

// JobResponse is the response obtained from the Amazon Jobs endpoint
type JobResponse struct {
	Error string       `json:"error"`
	Hits  int          `json:"hits"`
	Jobs  []JobWrapper `json:"jobs"`
}

// JobWrapper representes the job json structure received in a JobResponse
type JobWrapper struct {
	ID                      string `json:"id_icims"`
	Title                   string `json:"title"`
	Category                string `json:"job_category"`
	Description             string `json:"description"`
	BasicQualifications     string `json:"basic_qualifications"`
	PreferredQualifications string `json:"preferred_qualifications"`
	Location                string `json:"normalized_location"`
	Date                    string `json:"posted_date"`
	Path                    string `json:"job_path"`
}

// NewJobResponse crates a JobResponse from a response body
func NewJobResponse(body []byte) (*JobResponse, error) {
	var jobResponse JobResponse
	err := json.Unmarshal(body, &jobResponse)
	return &jobResponse, err
}

// GetJobs returns an slice of jobs
func (jobResponse *JobResponse) GetJobs() []*Job {
	var jobs []*Job

	for _, jobWrapper := range jobResponse.Jobs {
		job := jobWrapper.ToJob()
		jobs = append(jobs, job)
	}

	return jobs
}

// ToJob converts a JobWrapper into a Job
func (jobWrapper *JobWrapper) ToJob() *Job {
	job := &Job{
		ID:                      jobWrapper.ID,
		Company:                 "AMAZON",
		Title:                   transformTitle(jobWrapper.Title),
		Category:                transformText(jobWrapper.Category),
		Description:             transformText(jobWrapper.Description),
		BasicQualifications:     transformQualifications(jobWrapper.BasicQualifications),
		PreferredQualifications: transformQualifications(jobWrapper.PreferredQualifications),
		Location:                transformText(jobWrapper.Location),
		Date:                    transformText(jobWrapper.Date),
		URL:                     "www.amazon.jobs" + jobWrapper.Path,
	}

	return job
}

func transformTitle(title string) string {
	cleanTitle := strings.Split(title, "-")[0]
	return transformText(cleanTitle)
}

func transformQualifications(qualifications string) []string {
	parts := strings.Split(qualifications, "<br/>")

	var result []string

	for _, part := range parts {
		part = strings.Replace(part, "&", "and", -1)
		part = strings.Replace(part, "· ", "", -1)

		if part != "" {
			result = append(result, part)
		}
	}

	return result
}

func transformText(text string) string {
	text = strings.Replace(text, "<br/>", "\n", -1)
	text = strings.Replace(text, "&", "and", -1)
	text = strings.Trim(text, "\n")
	text = strings.TrimSpace(text)
	return text
}
