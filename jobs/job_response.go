package jobs

import (
	"encoding/json"
	"log"
	"strings"
	"time"
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
	return &Job{
		ID:          jobWrapper.ID,
		Company:     "AMAZON",
		Title:       jobWrapper.Title,
		Category:    "SOFTWARE",
		Description: transformDescription(jobWrapper.Description, jobWrapper.BasicQualifications, jobWrapper.PreferredQualifications),
		Location:    jobWrapper.Location,
		Date:        transformDate(jobWrapper.Date),
		URL:         "www.amazon.jobs" + jobWrapper.Path,
	}
}

func transformDescription(description, basicQualifications, preferredQualifications string) string {
	parts := []string{
		description,
		basicQualifications,
		preferredQualifications,
	}

	return strings.Join(parts, "<br/>")
}

func transformDate(dateString string) string {
	date, err := time.Parse("January 2, 2006", dateString)

	if err != nil {
		log.Fatal("Error parsing date")
	}

	return date.Format("2006-01-02")
}
