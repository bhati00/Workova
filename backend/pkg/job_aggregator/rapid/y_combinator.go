package rapid

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/bhati00/workova/backend/constant"
	"github.com/bhati00/workova/backend/dtos"
	jobaggregator "github.com/bhati00/workova/backend/pkg/job_aggregator"
	"github.com/bhati00/workova/backend/pkg/utils"
)

// YCombinatorJobResponse represents the API response structure
type YCombinatorJobResponse struct {
	ID               string  `json:"id"`
	DatePosted       string  `json:"date_posted"`
	DateCreated      string  `json:"date_created"`
	Title            string  `json:"title"`
	Organization     string  `json:"organization"`
	OrganizationURL  string  `json:"organization_url"`
	DateValidThrough *string `json:"date_validthrough"`
	LocationsRaw     []struct {
		Type    string `json:"@type"`
		Address struct {
			Type            string `json:"@type"`
			AddressLocality string `json:"addressLocality"`
			AddressRegion   string `json:"addressRegion"`
			AddressCountry  string `json:"addressCountry"`
		} `json:"address"`
	} `json:"locations_raw"`
	SalaryRaw struct {
		Type     string `json:"@type"`
		Currency string `json:"currency"`
		Value    struct {
			Type     string  `json:"@type"`
			UnitText string  `json:"unitText"`
			MinValue float64 `json:"minValue"`
			MaxValue float64 `json:"maxValue"`
		} `json:"value"`
	} `json:"salary_raw"`
	EmploymentType   []string  `json:"employment_type"`
	URL              string    `json:"url"`
	SourceType       string    `json:"source_type"`
	Source           string    `json:"source"`
	SourceDomain     string    `json:"source_domain"`
	OrganizationLogo string    `json:"organization_logo"`
	CitiesDerived    []string  `json:"cities_derived"`
	CountiesDerived  []string  `json:"counties_derived"`
	RegionsDerived   []string  `json:"regions_derived"`
	CountriesDerived []string  `json:"countries_derived"`
	LocationsDerived []string  `json:"locations_derived"`
	TimezonesDerived []string  `json:"timezones_derived"`
	LatsDerived      []float64 `json:"lats_derived"`
	LngsDerived      []float64 `json:"lngs_derived"`
	RemoteDerived    bool      `json:"remote_derived"`
	RecruiterName    *string   `json:"recruiter_name"`
	RecruiterTitle   *string   `json:"recruiter_title"`
	RecruiterURL     *string   `json:"recruiter_url"`
}

// YCombinatorAggregator implements the JobAggregator interface
type YCombinatorAggregator struct {
	apiKey string
	client *http.Client
}

// NewYCombinatorAggregator creates a new Y Combinator job aggregator
func NewYCombinatorAggregator() *YCombinatorAggregator {
	return &YCombinatorAggregator{
		apiKey: "c756a0c080mshd28444baa7d08c4p140ec0jsnaaaf94596e06",
		client: &http.Client{Timeout: 30 * time.Second},
	}
}

// FetchJobs fetches jobs from Y Combinator API with pagination
func (y *YCombinatorAggregator) FetchJobs(options jobaggregator.FetchOptions) ([]dtos.JobRequest, error) {
	var allJobs []dtos.JobRequest
	offset := 0
	pageSize := 50 // Typical page size, adjust as needed

	for {
		// Create request with offset parameter
		req, err := http.NewRequest("GET", "https://free-y-combinator-jobs-api.p.rapidapi.com/active-jb-7d", nil)
		if err != nil {
			return nil, fmt.Errorf("failed to create request: %w", err)
		}

		// Add offset as query parameter
		q := req.URL.Query()
		q.Add("offset", fmt.Sprintf("%d", offset))
		req.URL.RawQuery = q.Encode()

		req.Header.Set("X-RapidAPI-Key", y.apiKey)
		req.Header.Set("X-RapidAPI-Host", "free-y-combinator-jobs-api.p.rapidapi.com")

		resp, err := y.client.Do(req)
		if err != nil {
			return nil, fmt.Errorf("failed to make request: %w", err)
		}

		if resp.StatusCode != http.StatusOK {
			resp.Body.Close()
			return nil, fmt.Errorf("API returned status %d", resp.StatusCode)
		}

		var rawJobs []YCombinatorJobResponse
		if err := json.NewDecoder(resp.Body).Decode(&rawJobs); err != nil {
			resp.Body.Close()
			return nil, fmt.Errorf("failed to decode response: %w", err)
		}
		resp.Body.Close()

		// If no more jobs, break the loop
		if len(rawJobs) == 0 {
			break
		}

		shouldStop := false

		// Process jobs from this page
		for _, rawJob := range rawJobs {
			// Check if job is older than datePosted filter (if provided)
			if options.DatePosted != nil && rawJob.DatePosted != "" {
				jobDate, err := time.Parse(time.RFC3339, rawJob.DatePosted)
				if err == nil && jobDate.Before(*options.DatePosted) {
					shouldStop = true
					break
				}
			}

			rawJobInterface := make([]any, 1)
			rawJobInterface[0] = rawJob

			job, err := y.RawJobtoDto(rawJobInterface)
			if err != nil {
				continue // Skip jobs that can't be converted
			}

			allJobs = append(allJobs, job)

			// Check if we've reached maxJobs limit
			if options.MaxJobs > 0 && len(allJobs) >= options.MaxJobs {
				shouldStop = true
				break
			}
		}

		// Stop if we found older jobs or reached limits
		if shouldStop {
			break
		}

		// If we got fewer jobs than expected, we've reached the end
		if len(rawJobs) < pageSize {
			break
		}

		// Move to next page
		offset += len(rawJobs)

		// Safety check to prevent infinite loops
		if offset > 10000 { // Adjust as needed
			break
		}
	}

	return allJobs, nil
}

// RawJobtoDto converts raw job data to JobRequest DTO
func (y *YCombinatorAggregator) RawJobtoDto(rawJob []any) (dtos.JobRequest, error) {
	if len(rawJob) == 0 {
		return dtos.JobRequest{}, fmt.Errorf("empty raw job data")
	}

	// Convert the interface{} back to YCombinatorJobResponse
	jobBytes, err := json.Marshal(rawJob[0])
	if err != nil {
		return dtos.JobRequest{}, fmt.Errorf("failed to marshal raw job: %w", err)
	}

	var ycJob YCombinatorJobResponse
	if err := json.Unmarshal(jobBytes, &ycJob); err != nil {
		return dtos.JobRequest{}, fmt.Errorf("failed to unmarshal to YC job: %w", err)
	}

	// Map employment type
	jobType := constant.JobType(1) // Default to full-time
	if len(ycJob.EmploymentType) > 0 {
		switch strings.ToLower(ycJob.EmploymentType[0]) {
		case "full_time":
			jobType = constant.JobType(1)
		case "part_time":
			jobType = constant.JobType(2)
		case "contract":
			jobType = constant.JobType(3)
		}
	}

	// Extract location info
	var city *string
	countryIso := "US" // Default for YC jobs
	if len(ycJob.LocationsRaw) > 0 {
		if ycJob.LocationsRaw[0].Address.AddressLocality != "" {
			city = &ycJob.LocationsRaw[0].Address.AddressLocality
		}
		if ycJob.LocationsRaw[0].Address.AddressCountry != "" {
			countryIso = ycJob.LocationsRaw[0].Address.AddressCountry
		}
	}

	// Extract salary info
	var salaryMin, salaryMax *int
	var salaryCurrency constant.Currency
	if ycJob.SalaryRaw.Currency != "" {
		currency := constant.Currency("USD") // Default USD

		salaryCurrency = currency

		if ycJob.SalaryRaw.Value.MinValue > 0 {
			min := int(ycJob.SalaryRaw.Value.MinValue)
			salaryMin = &min
		}
		if ycJob.SalaryRaw.Value.MaxValue > 0 {
			max := int(ycJob.SalaryRaw.Value.MaxValue)
			salaryMax = &max
		}
	}

	// Work mode
	workMode := constant.WorkMode(1) // Default onsite
	if ycJob.RemoteDerived {
		workMode = constant.WorkMode(2) // Remote
	}

	// Create slug from title and ID
	slug := strings.ToLower(strings.ReplaceAll(ycJob.Title, " ", "-")) + "-" + ycJob.ID

	return dtos.JobRequest{
		ExternalJobID:   &ycJob.ID,
		Slug:            &slug,
		Title:           ycJob.Title,
		Description:     nil, // Not provided in YC API
		CompanyName:     ycJob.Organization,
		CountryIso:      countryIso,
		City:            city,
		JobType:         jobType,
		SalaryMin:       salaryMin,
		SalaryMax:       salaryMax,
		SalaryCurrency:  salaryCurrency,
		IsRemote:        &ycJob.RemoteDerived,
		Skills:          []string{},   // Not provided in YC API
		Category:        "Technology", // Default for YC jobs
		PostedDate:      utils.ParseToRFC3339(ycJob.DatePosted),
		ApplicationURL:  &ycJob.URL,
		Source:          "Y Combinator",
		Tags:            []string{"startup", "ycombinator"},
		Industry:        utils.String("Technology"),
		Department:      nil, // Not provided in YC API
		VisaSponsorship: nil, // Not provided in YC API
		EducationLevel:  nil, // Not provided in YC API
		ExperienceLevel: nil, // Not provided in YC API
		WorkMode:        workMode,
	}, nil
}
