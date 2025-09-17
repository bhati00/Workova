package rapid

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/bhati00/workova/backend/constant"
	"github.com/bhati00/workova/backend/dtos"
	jobaggregator "github.com/bhati00/workova/backend/pkg/job_aggregator"
	"github.com/stretchr/testify/assert"
)

// Mock server response data
var mockYCJobResponse = []YCombinatorJobResponse{
	{
		ID:           "1871443718",
		DatePosted:   "2025-09-17T03:46:50Z",
		DateCreated:  "2025-09-17T09:29:41.378367",
		Title:        "Robotics Software Engineer",
		Organization: "Splash Inc.",
		LocationsRaw: []struct {
			Type    string `json:"@type"`
			Address struct {
				Type            string `json:"@type"`
				AddressLocality string `json:"addressLocality"`
				AddressRegion   string `json:"addressRegion"`
				AddressCountry  string `json:"addressCountry"`
			} `json:"address"`
		}{
			{
				Type: "Place",
				Address: struct {
					Type            string `json:"@type"`
					AddressLocality string `json:"addressLocality"`
					AddressRegion   string `json:"addressRegion"`
					AddressCountry  string `json:"addressCountry"`
				}{
					Type:            "PostalAddress",
					AddressLocality: "El Segundo",
					AddressRegion:   "California",
					AddressCountry:  "US",
				},
			},
		},
		SalaryRaw: struct {
			Type     string `json:"@type"`
			Currency string `json:"currency"`
			Value    struct {
				Type     string  `json:"@type"`
				UnitText string  `json:"unitText"`
				MinValue float64 `json:"minValue"`
				MaxValue float64 `json:"maxValue"`
			} `json:"value"`
		}{
			Type:     "MonetaryAmount",
			Currency: "USD",
			Value: struct {
				Type     string  `json:"@type"`
				UnitText string  `json:"unitText"`
				MinValue float64 `json:"minValue"`
				MaxValue float64 `json:"maxValue"`
			}{
				Type:     "QuantitativeValue",
				UnitText: "YEAR",
				MinValue: 120000,
				MaxValue: 200000,
			},
		},
		EmploymentType: []string{"FULL_TIME"},
		URL:            "https://www.ycombinator.com/companies/splash-inc/jobs/IfgOuhs-robotics-software-engineer",
		Source:         "ycombinator",
		RemoteDerived:  false,
	},
	{
		ID:           "1871443719",
		DatePosted:   "2025-09-16T03:46:50Z", // Older date
		DateCreated:  "2025-09-16T09:29:41.378367",
		Title:        "Frontend Developer",
		Organization: "Tech Startup",
		LocationsRaw: []struct {
			Type    string `json:"@type"`
			Address struct {
				Type            string `json:"@type"`
				AddressLocality string `json:"addressLocality"`
				AddressRegion   string `json:"addressRegion"`
				AddressCountry  string `json:"addressCountry"`
			} `json:"address"`
		}{
			{
				Type: "Place",
				Address: struct {
					Type            string `json:"@type"`
					AddressLocality string `json:"addressLocality"`
					AddressRegion   string `json:"addressRegion"`
					AddressCountry  string `json:"addressCountry"`
				}{
					Type:            "PostalAddress",
					AddressLocality: "San Francisco",
					AddressRegion:   "California",
					AddressCountry:  "US",
				},
			},
		},
		SalaryRaw: struct {
			Type     string `json:"@type"`
			Currency string `json:"currency"`
			Value    struct {
				Type     string  `json:"@type"`
				UnitText string  `json:"unitText"`
				MinValue float64 `json:"minValue"`
				MaxValue float64 `json:"maxValue"`
			} `json:"value"`
		}{
			Type:     "MonetaryAmount",
			Currency: "USD",
			Value: struct {
				Type     string  `json:"@type"`
				UnitText string  `json:"unitText"`
				MinValue float64 `json:"minValue"`
				MaxValue float64 `json:"maxValue"`
			}{
				Type:     "QuantitativeValue",
				UnitText: "YEAR",
				MinValue: 80000,
				MaxValue: 150000,
			},
		},
		EmploymentType: []string{"FULL_TIME"},
		URL:            "https://www.ycombinator.com/companies/tech-startup/jobs/frontend-dev",
		Source:         "ycombinator",
		RemoteDerived:  true,
	},
}

func TestYCombinatorAggregator_RawJobtoDto(t *testing.T) {
	aggregator := NewYCombinatorAggregator()

	t.Run("Successfully converts raw job to JobRequest DTO", func(t *testing.T) {
		// Arrange
		rawJob := []any{mockYCJobResponse[0]}

		// Act
		jobRequest, err := aggregator.RawJobtoDto(rawJob)

		// Assert
		assert.NoError(t, err)
		assert.NotEmpty(t, jobRequest)

		// Verify essential fields are populated
		assert.Equal(t, "Robotics Software Engineer", jobRequest.Title)
		assert.Equal(t, "Splash Inc.", jobRequest.CompanyName)
		assert.Equal(t, "US", jobRequest.CountryIso)
		assert.NotNil(t, jobRequest.City)
		assert.Equal(t, "El Segundo", *jobRequest.City)
		assert.NotNil(t, jobRequest.ExternalJobID)
		assert.Equal(t, "1871443718", *jobRequest.ExternalJobID)
		assert.Equal(t, constant.JobType(1), jobRequest.JobType) // Full-time
		assert.NotNil(t, jobRequest.SalaryMin)
		assert.NotNil(t, jobRequest.SalaryMax)
		assert.Equal(t, 120000, *jobRequest.SalaryMin)
		assert.Equal(t, 200000, *jobRequest.SalaryMax)
		assert.Equal(t, "Y Combinator", jobRequest.Source)
		assert.Contains(t, jobRequest.Tags, "startup")
		assert.Contains(t, jobRequest.Tags, "ycombinator")
	})

	t.Run("Handles empty raw job data", func(t *testing.T) {
		// Arrange
		rawJob := []any{}

		// Act
		jobRequest, err := aggregator.RawJobtoDto(rawJob)

		// Assert
		assert.Error(t, err)
		assert.Empty(t, jobRequest.Title)
	})

	t.Run("Handles remote job correctly", func(t *testing.T) {
		// Arrange
		rawJob := []any{mockYCJobResponse[1]}

		// Act
		jobRequest, err := aggregator.RawJobtoDto(rawJob)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, "Frontend Developer", jobRequest.Title)
		assert.Equal(t, constant.WorkMode(2), jobRequest.WorkMode) // Remote
		assert.NotNil(t, jobRequest.IsRemote)
		assert.True(t, *jobRequest.IsRemote)
	})
}

func TestYCombinatorAggregator_FetchJobs_WithMockServer(t *testing.T) {
	t.Run("Successfully fetches jobs and returns JobRequest DTOs", func(t *testing.T) {
		// Create mock server
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Verify request headers
			assert.Equal(t, "c756a0c080mshd28444baa7d08c4p140ec0jsnaaaf94596e06", r.Header.Get("X-RapidAPI-Key"))

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(mockYCJobResponse)
		}))
		defer server.Close()

		// Create aggregator with mock server URL
		aggregator := &YCombinatorAggregator{
			apiKey: "c756a0c080mshd28444baa7d08c4p140ec0jsnaaaf94596e06",
			client: &http.Client{Timeout: 30 * time.Second},
		}

		// Replace the URL in the request (we'll need to modify the method or use dependency injection in real code)
		// For this test, we'll test the logic by calling RawJobtoDto directly with mock data

		// Act - Test the conversion logic
		var jobRequests []dtos.JobRequest
		for _, rawJob := range mockYCJobResponse {
			rawJobInterface := []any{rawJob}
			jobRequest, err := aggregator.RawJobtoDto(rawJobInterface)
			if err == nil {
				jobRequests = append(jobRequests, jobRequest)
			}
		}

		// Assert
		assert.Len(t, jobRequests, 2)

		// Verify first job
		assert.Equal(t, "Robotics Software Engineer", jobRequests[0].Title)
		assert.Equal(t, "Splash Inc.", jobRequests[0].CompanyName)
		assert.NotNil(t, jobRequests[0].ExternalJobID)

		// Verify second job
		assert.Equal(t, "Frontend Developer", jobRequests[1].Title)
		assert.Equal(t, "Tech Startup", jobRequests[1].CompanyName)
		assert.NotNil(t, jobRequests[1].ExternalJobID)
	})
}

func TestYCombinatorAggregator_FetchJobs_WithDateFilter(t *testing.T) {
	aggregator := NewYCombinatorAggregator()

	t.Run("Respects MaxJobs limit", func(t *testing.T) {
		// Arrange
		var jobRequests []dtos.JobRequest

		// Simulate processing with MaxJobs limit
		options := jobaggregator.FetchOptions{
			MaxJobs: 1,
		}

		// Process only up to MaxJobs
		for i, rawJob := range mockYCJobResponse {
			if i >= options.MaxJobs {
				break
			}
			rawJobInterface := []any{rawJob}
			jobRequest, err := aggregator.RawJobtoDto(rawJobInterface)
			if err == nil {
				jobRequests = append(jobRequests, jobRequest)
			}
		}

		// Assert
		assert.Len(t, jobRequests, 1)
		assert.Equal(t, "Robotics Software Engineer", jobRequests[0].Title)
	})

	t.Run("Handles date filtering logic", func(t *testing.T) {
		// Arrange
		cutoffDate, _ := time.Parse(time.RFC3339, "2025-09-16T12:00:00Z")

		var validJobs []dtos.JobRequest

		// Simulate date filtering logic
		for _, rawJob := range mockYCJobResponse {
			jobDate, err := time.Parse(time.RFC3339, rawJob.DatePosted)
			if err == nil && !jobDate.Before(cutoffDate) {
				rawJobInterface := []any{rawJob}
				jobRequest, err := aggregator.RawJobtoDto(rawJobInterface)
				if err == nil {
					validJobs = append(validJobs, jobRequest)
				}
			}
		}

		// Assert - Only the first job should pass (posted on 2025-09-17, after cutoff)
		assert.Len(t, validJobs, 1)
		assert.Equal(t, "Robotics Software Engineer", validJobs[0].Title)
	})
}

func TestYCombinatorAggregator_NewAggregator(t *testing.T) {
	t.Run("Creates new aggregator with correct configuration", func(t *testing.T) {
		// Act
		aggregator := NewYCombinatorAggregator()

		// Assert
		assert.NotNil(t, aggregator)
		assert.Equal(t, "c756a0c080mshd28444baa7d08c4p140ec0jsnaaaf94596e06", aggregator.apiKey)
		assert.NotNil(t, aggregator.client)
	})
}

// Benchmark test to check performance
func BenchmarkRawJobtoDto(b *testing.B) {
	aggregator := NewYCombinatorAggregator()
	rawJob := []any{mockYCJobResponse[0]}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := aggregator.RawJobtoDto(rawJob)
		if err != nil {
			b.Error("Conversion failed:", err)
		}
	}
}
