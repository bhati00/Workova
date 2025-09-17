package jobaggregator

// import (
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// 	"strings"
// 	"time"
// 	"github.com/bhati00/workova/backend/pkg/utils"
// 	"github.com/bhati00/workova/backend/internal/job/model"
// 	"github.com/bhati00/workova/backend/dtos"
// )

// type ArbeitNowAggregator struct {
// 	Data         ArbeitNowResponse
// 	FetchOptions FetchOptions
// }
// type ArbeitNowResponse struct {
// 	ArbeitNowJob []ArbeitNowJob
// 	Meta         ArbeitNowMeta
// }

// // ArbeitNowMeta represents pagination metadata
// type ArbeitNowMeta struct {
// 	CurrentPage int `json:"current_page"`
// 	TotalPages  int `json:"total_pages"`
// 	TotalJobs   int `json:"total_jobs"`
// }

// type ArbeitNowJob struct {
// 	Slug        string   `json:"slug"`
// 	CompanyName string   `json:"company_name"`
// 	Title       string   `json:"title"`
// 	Description string   `json:"description"`
// 	Remote      bool     `json:"remote"`
// 	URL         string   `json:"url"`
// 	Tags        []string `json:"tags"`
// 	JobTypes    []string `json:"job_types"`
// 	Location    string   `json:"location"`
// 	CreatedAt   int64    `json:"created_at"` // Unix timestamp
// }

// func NewArbeitNowAggregator() *ArbeitNowAggregator {
// 	return &ArbeitNowAggregator{Data: ArbeitNowResponse{}, FetchOptions: FetchOptions{Pages: 1}}
// }

// // FetchJobs calls the ArbeitNow API and returns job data with loop until old data
// func (a *ArbeitNowAggregator) FetchJobs(fetchOptions FetchOptions) ([]dtos.JobRequest, error) {
// 	var allJobs []dtos.JobRequest
// 	currentPage := 1

// 	// Create HTTP client with timeout
// 	client := &http.Client{
// 		Timeout: 30 * time.Second,
// 	}

// 	// Set cutoff date - if DatePosted is provided, use it; otherwise use 30 days ago
// 	var cutoffDate time.Time
// 	if fetchOptions.DatePosted != nil {
// 		cutoffDate = *fetchOptions.DatePosted
// 	} else {
// 		cutoffDate = time.Now().AddDate(0, 0, -30) // 30 days ago
// 	}

// 	fmt.Printf("Fetching jobs newer than: %s\n", cutoffDate.Format("2006-01-02 15:04:05"))

// 	for {
// 		// Check if we've exceeded max pages
// 		if fetchOptions.Pages > 0 && currentPage > fetchOptions.Pages {
// 			break
// 		}

// 		// Check if we've exceeded max jobs
// 		if fetchOptions.MaxJobs > 0 && len(allJobs) >= fetchOptions.MaxJobs {
// 			break
// 		}

// 		// Fetch jobs from current page
// 		pageJobs, shouldContinue, err := a.fetchJobsFromPage(client, currentPage, cutoffDate, fetchOptions.VisaSponsorship)
// 		if err != nil {
// 			return nil, fmt.Errorf("error fetching page %d: %w", currentPage, err)
// 		}

// 		// Add fetched jobs to our collection
// 		allJobs = append(allJobs, pageJobs...)

// 		// If we found old jobs or no jobs, stop fetching
// 		if !shouldContinue {
// 			fmt.Printf("Stopping fetch: found old jobs or no more jobs on page %d\n", currentPage)
// 			break
// 		}

// 		fmt.Printf("Fetched %d jobs from page %d\n", len(pageJobs), currentPage)
// 		currentPage++

// 		// Add small delay to be respectful to the API
// 		time.Sleep(500 * time.Millisecond)
// 	}

// 	// Limit to MaxJobs if specified
// 	if fetchOptions.MaxJobs > 0 && len(allJobs) > fetchOptions.MaxJobs {
// 		allJobs = allJobs[:fetchOptions.MaxJobs]
// 	}

// 	fmt.Printf("Total jobs fetched: %d\n", len(allJobs))
// 	return allJobs, nil
// }

// // fetchJobsFromPage fetches jobs from a specific page and returns whether to continue fetching
// func (a *ArbeitNowAggregator) fetchJobsFromPage(client *http.Client, page int, cutoffDate time.Time, visaSponsorship bool) ([]dtos.JobRequest, bool, error) {
// 	// Build API URL
// 	baseURL := "https://www.arbeitnow.com/api/job-board-api"
// 	url := fmt.Sprintf("%s?page=%d&visa_sponsorship=%t", baseURL, page, visaSponsorship)

// 	// Create HTTP GET request
// 	req, err := http.NewRequest("GET", url, nil)
// 	if err != nil {
// 		return nil, false, fmt.Errorf("failed to create request: %w", err)
// 	}

// 	// Set headers
// 	req.Header.Set("Accept", "application/json")
// 	req.Header.Set("User-Agent", "Workova-JobAggregator/1.0")

// 	// Make the API call
// 	fmt.Printf("Calling ArbeitNow API: %s\n", url)
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		return nil, false, fmt.Errorf("API request failed: %w", err)
// 	}
// 	defer resp.Body.Close()

// 	// Check HTTP status code
// 	if resp.StatusCode != http.StatusOK {
// 		return nil, false, fmt.Errorf("API returned status code: %d", resp.StatusCode)
// 	}

// 	// Parse JSON response
// 	var apiResponse ArbeitNowResponse
// 	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
// 		return nil, false, fmt.Errorf("failed to decode JSON response: %w", err)
// 	}

// 	// If no jobs on this page, stop fetching
// 	if len(apiResponse.ArbeitNowJob) == 0 {
// 		return nil, false, nil
// 	}

// 	// Transform jobs and check for old jobs
// 	var jobs []model.Job
// 	foundOldJob := false

// 	for _, rawJob := range apiResponse.ArbeitNowJob {
// 		// Convert raw job to our internal job structure
// 		transformedJob, err := a.TransformJob(rawJob)
// 		if err != nil {
// 			fmt.Printf("Error transforming job %s: %v\n", rawJob.Slug, err)
// 			continue
// 		}

// 		// Check if job is older than cutoff date
// 		if transformedJob.CreatedAt.Before(cutoffDate) {
// 			fmt.Printf("Found old job: %s (created: %s)\n", transformedJob.JobTitle, transformedJob.CreatedAt.Format("2006-01-02"))
// 			foundOldJob = true
// 			break // Stop processing this page
// 		}

// 		jobs = append(jobs, transformedJob)
// 	}

// 	// Return jobs and whether to continue fetching more pages
// 	shouldContinue := !foundOldJob && len(apiResponse.ArbeitNowJob) > 0
// 	return jobs, shouldContinue, nil
// }

// // TransformJob converts raw ArbeitNow job data to our internal Job structure
// func (a *ArbeitNowAggregator) TransformJob(rawJob interface{}) (dtos.JobRequest, error) {
// 	// Type assertion to convert interface{} to ArbeitNowJob
// 	arbeitJob, _ := rawJob.(ArbeitNowJob)
// 	job := dtos.JobRequest{
// 		Title: arbeitJob.Title,
// 		Slug:  &arbeitJob.Slug,
// 		VisaSponsorship: &a.FetchOptions.VisaSponsorship,
// 		IsRemote : &arbeitJob.Remote,
// 		ApplicationURL: &arbeitJob.URL,
// 		CompanyName: arbeitJob.CompanyName,
// 		PostedDate: utils.UnixToTime(arbeitJob.CreatedAt),
// 	}
// }

// // normalizeJobType converts ArbeitNow job types to our internal job types
// func normalizeJobType(jobType string) string {
// 	jobType = strings.ToLower(strings.TrimSpace(jobType))

// 	switch {
// 	case strings.Contains(jobType, "full"):
// 		return "full-time"
// 	case strings.Contains(jobType, "part"):
// 		return "part-time"
// 	case strings.Contains(jobType, "contract"):
// 		return "contract"
// 	case strings.Contains(jobType, "freelance"):
// 		return "freelance"
// 	case strings.Contains(jobType, "intern"):
// 		return "internship"
// 	default:
// 		return "full-time" // default
// 	}
// }
