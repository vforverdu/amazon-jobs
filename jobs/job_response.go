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
	CompanyName             string `json:"company_name"`
	Location                string `json:"normalized_location"`
	City                    string `json:"city"`
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
	description := transformText(jobWrapper.Description)
	basicQualifictions := transformQualifications(jobWrapper.BasicQualifications)
	preferredQualifictions := transformQualifications(jobWrapper.PreferredQualifications)

	return &Job{
		ID:                      jobWrapper.ID,
		Title:                   jobWrapper.Title,
		Category:                jobWrapper.Category,
		Description:             description,
		BasicQualifications:     basicQualifictions,
		PreferredQualifications: preferredQualifictions,
		CompanyName:             jobWrapper.CompanyName,
		Location:                jobWrapper.Location,
		City:                    jobWrapper.City,
		Date:                    jobWrapper.Date,
		Path:                    jobWrapper.Path,
	}
}

func transformQualifications(text string) []string {
	parts := transformText(text)

	var result []string

	for _, part := range parts {
		if strings.HasPrefix(part, "· ") {
			part = strings.Replace(part, "· ", "", -1)
			result = append(result, part)
		}
	}

	return result
}

func transformText(text string) []string {

	parts := strings.Split(text, "<br/>")

	var result []string

	for _, part := range parts {
		if part != "" {
			result = append(result, part)
		}
	}

	return result
}
