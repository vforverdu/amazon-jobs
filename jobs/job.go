package jobs

// Job struct
type Job struct {
	ID                      string   `json:"id"`
	Title                   string   `json:"title"`
	Category                string   `json:"category"`
	Description             string   `json:"description"`
	BasicQualifications     []string `json:"basicQualifications"`
	PreferredQualifications []string `json:"preferredQualifications"`
	CompanyName             string   `json:"companyName"`
	Location                string   `json:"location"`
	City                    string   `json:"city"`
	Date                    string   `json:"date"`
	URL                     string   `json:"url"`
}
