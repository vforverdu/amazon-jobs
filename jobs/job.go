package jobs

// Job struct
type Job struct {
	ID                      string   `json:"id"`
	Company                 string   `json:"company"`
	Title                   string   `json:"title"`
	Category                string   `json:"category"`
	Description             string   `json:"description"`
	BasicQualifications     []string `json:"basicQualifications"`
	PreferredQualifications []string `json:"preferredQualifications"`
	Location                string   `json:"location"`
	Date                    string   `json:"date"`
	URL                     string   `json:"url"`
}
