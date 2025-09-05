// handlers/job_handler.go
package job

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/bhati00/workova/backend/dtos"
	"github.com/gin-gonic/gin"
)

// JobHandler handles HTTP requests for job operations
type JobHandler struct {
	jobService JobService
}

// NewJobHandler creates a new job handler instance
func NewJobHandler(jobService JobService) *JobHandler {
	return &JobHandler{
		jobService: jobService,
	}
}

// CreateJob godoc
// @Summary Create a new job
// @Description Creates a new job entry
// @Tags Jobs
// @Accept json
// @Produce json
// @Param job body Job true "Job object"
// @Success 201 {object} dtos.APIResponse
// @Failure 400 {object} dtos.APIResponse
// @Failure 500 {object} dtos.APIResponse
// @Router /jobs [post]
func (h *JobHandler) CreateJob(c *gin.Context) {
	var job Job
	if err := c.ShouldBindJSON(&job); err != nil {
		c.JSON(http.StatusBadRequest, dtos.APIResponse{
			Success: false,
			Error:   "Invalid request body: " + err.Error(),
		})
		return
	}

	if err := h.jobService.CreateJob(&job); err != nil {
		c.JSON(http.StatusInternalServerError, dtos.APIResponse{
			Success: false,
			Error:   "Failed to create job: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, dtos.APIResponse{
		Success: true,
		Message: "Job created successfully",
		Data:    job,
	})
}

// GetJob godoc
// @Summary Get job by ID
// @Description Returns a job by its internal ID
// @Tags Jobs
// @Produce json
// @Param id path int true "Job ID"
// @Success 200 {object} dtos.APIResponse
// @Failure 400 {object} dtos.APIResponse
// @Failure 404 {object} dtos.APIResponse
// @Router /jobs/{id} [get]
func (h *JobHandler) GetJob(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, dtos.APIResponse{
			Success: false,
			Error:   "Invalid job ID",
		})
		return
	}

	job, err := h.jobService.GetJobByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, dtos.APIResponse{
			Success: false,
			Error:   "Job not found",
		})
		return
	}

	c.JSON(http.StatusOK, dtos.APIResponse{
		Success: true,
		Data:    job,
	})
}

// GetJobByJobID godoc
// @Summary Get job by external JobID
// @Description Returns a job by its external JobID
// @Tags Jobs
// @Produce json
// @Param job_id path string true "External Job ID"
// @Success 200 {object} dtos.APIResponse
// @Failure 400 {object} dtos.APIResponse
// @Failure 404 {object} dtos.APIResponse
// @Router /jobs/external/{job_id} [get]
func (h *JobHandler) GetJobByJobID(c *gin.Context) {
	jobID := c.Param("job_id")
	if strings.TrimSpace(jobID) == "" {
		c.JSON(http.StatusBadRequest, dtos.APIResponse{
			Success: false,
			Error:   "Job ID cannot be empty",
		})
		return
	}

	job, err := h.jobService.GetJobByJobID(jobID)
	if err != nil {
		c.JSON(http.StatusNotFound, dtos.APIResponse{
			Success: false,
			Error:   "Job not found",
		})
		return
	}

	c.JSON(http.StatusOK, dtos.APIResponse{
		Success: true,
		Data:    job,
	})
}

// UpdateJob godoc
// @Summary Update a job
// @Description Updates a job by ID
// @Tags Jobs
// @Accept json
// @Produce json
// @Param id path int true "Job ID"
// @Param job body Job true "Job object"
// @Success 200 {object} dtos.APIResponse
// @Failure 400 {object} dtos.APIResponse
// @Failure 500 {object} dtos.APIResponse
// @Router /jobs/{id} [put]
func (h *JobHandler) UpdateJob(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, dtos.APIResponse{
			Success: false,
			Error:   "Invalid job ID",
		})
		return
	}

	var job Job
	if err := c.ShouldBindJSON(&job); err != nil {
		c.JSON(http.StatusBadRequest, dtos.APIResponse{
			Success: false,
			Error:   "Invalid request body: " + err.Error(),
		})
		return
	}

	// Ensure the ID in the URL matches the ID in the request body
	job.ID = uint(id)

	if err := h.jobService.UpdateJob(&job); err != nil {
		c.JSON(http.StatusInternalServerError, dtos.APIResponse{
			Success: false,
			Error:   "Failed to update job: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dtos.APIResponse{
		Success: true,
		Message: "Job updated successfully",
		Data:    job,
	})
}

// DeleteJob godoc
// @Summary Delete a job
// @Description Deletes a job by ID
// @Tags Jobs
// @Produce json
// @Param id path int true "Job ID"
// @Success 200 {object} dtos.APIResponse
// @Failure 400 {object} dtos.APIResponse
// @Failure 500 {object} dtos.APIResponse
// @Router /jobs/{id} [delete]
func (h *JobHandler) DeleteJob(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, dtos.APIResponse{
			Success: false,
			Error:   "Invalid job ID",
		})
		return
	}

	if err := h.jobService.DeleteJob(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, dtos.APIResponse{
			Success: false,
			Error:   "Failed to delete job: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dtos.APIResponse{
		Success: true,
		Message: "Job deleted successfully",
	})
}

// DeactivateJob godoc
// @Summary Deactivate a job
// @Description Marks a job as inactive
// @Tags Jobs
// @Produce json
// @Param id path int true "Job ID"
// @Success 200 {object} dtos.APIResponse
// @Failure 400 {object} dtos.APIResponse
// @Failure 500 {object} dtos.APIResponse
// @Router /jobs/{id}/deactivate [patch]
func (h *JobHandler) DeactivateJob(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, dtos.APIResponse{
			Success: false,
			Error:   "Invalid job ID",
		})
		return
	}

	if err := h.jobService.DeactivateJob(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, dtos.APIResponse{
			Success: false,
			Error:   "Failed to deactivate job: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dtos.APIResponse{
		Success: true,
		Message: "Job deactivated successfully",
	})
}

// BatchCreateJobs godoc
// @Summary Batch create jobs
// @Description Creates multiple jobs at once
// @Tags Jobs
// @Accept json
// @Produce json
// @Param jobs body []Job true "List of jobs"
// @Success 201 {object} dtos.APIResponse
// @Failure 400 {object} dtos.APIResponse
// @Failure 500 {object} dtos.APIResponse
// @Router /jobs/batch [post]
func (h *JobHandler) BatchCreateJobs(c *gin.Context) {
	var jobs []Job
	if err := c.ShouldBindJSON(&jobs); err != nil {
		c.JSON(http.StatusBadRequest, dtos.APIResponse{
			Success: false,
			Error:   "Invalid request body: " + err.Error(),
		})
		return
	}

	if len(jobs) == 0 {
		c.JSON(http.StatusBadRequest, dtos.APIResponse{
			Success: false,
			Error:   "No jobs provided",
		})
		return
	}

	if len(jobs) > MaxBatchSize {
		c.JSON(http.StatusBadRequest, dtos.APIResponse{
			Success: false,
			Error:   "Batch size exceeds maximum allowed limit of " + strconv.Itoa(MaxBatchSize),
		})
		return
	}

	result, err := h.jobService.CreateJobsBatch(jobs)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dtos.APIResponse{
			Success: false,
			Error:   "Batch creation failed: " + err.Error(),
		})
		return
	}

	statusCode := http.StatusCreated
	if result.Failed > 0 {
		statusCode = http.StatusPartialContent
	}

	c.JSON(statusCode, dtos.APIResponse{
		Success: result.Successful > 0,
		Message: "Batch operation completed",
		Data:    result,
	})
}

// BatchDeleteJobs godoc
// @Summary Batch delete jobs
// @Description Deletes multiple jobs by IDs or JobIDs
// @Tags Jobs
// @Accept json
// @Produce json
// @Param request body dtos.BatchDeleteRequest true "Batch delete request"
// @Success 200 {object} dtos.APIResponse
// @Failure 400 {object} dtos.APIResponse
// @Failure 500 {object} dtos.APIResponse
// @Router /jobs/batch [delete]
func (h *JobHandler) BatchDeleteJobs(c *gin.Context) {

	var req dtos.BatchDeleteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dtos.APIResponse{
			Success: false,
			Error:   "Invalid request body: " + err.Error(),
		})
		return
	}

	var result *BatchResult
	var err error

	if len(req.IDs) > 0 && len(req.JobIDs) > 0 {
		c.JSON(http.StatusBadRequest, dtos.APIResponse{
			Success: false,
			Error:   "Cannot specify both IDs and JobIDs in the same request",
		})
		return
	}

	if len(req.IDs) > 0 {
		if len(req.IDs) > MaxBatchSize {
			c.JSON(http.StatusBadRequest, dtos.APIResponse{
				Success: false,
				Error:   "Batch size exceeds maximum allowed limit of " + strconv.Itoa(MaxBatchSize),
			})
			return
		}
		result, err = h.jobService.DeleteJobsBatch(req.IDs)
	} else if len(req.JobIDs) > 0 {
		if len(req.JobIDs) > MaxBatchSize {
			c.JSON(http.StatusBadRequest, dtos.APIResponse{
				Success: false,
				Error:   "Batch size exceeds maximum allowed limit of " + strconv.Itoa(MaxBatchSize),
			})
			return
		}
		result, err = h.jobService.DeleteJobsByJobIDsBatch(req.JobIDs)
	} else {
		c.JSON(http.StatusBadRequest, dtos.APIResponse{
			Success: false,
			Error:   "Either IDs or JobIDs must be provided",
		})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, dtos.APIResponse{
			Success: false,
			Error:   "Batch deletion failed: " + err.Error(),
		})
		return
	}

	statusCode := http.StatusOK
	if result.Failed > 0 {
		statusCode = http.StatusPartialContent
	}

	c.JSON(statusCode, dtos.APIResponse{
		Success: result.Successful > 0,
		Message: "Batch deletion completed",
		Data:    result,
	})
}

// GetAllJobs godoc
// @Summary Get all jobs
// @Description Returns paginated list of jobs
// @Tags Jobs
// @Produce json
// @Param page query int false "Page number"
// @Param page_size query int false "Page size"
// @Success 200 {object} dtos.APIResponse
// @Failure 500 {object} dtos.APIResponse
// @Router /jobs [get]
func (h *JobHandler) GetAllJobs(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	result, err := h.jobService.GetAllJobs(page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dtos.APIResponse{
			Success: false,
			Error:   "Failed to get jobs: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dtos.APIResponse{
		Success: true,
		Data:    result,
	})
}

// GetActiveJobs godoc
// @Summary Get active jobs
// @Description Returns paginated list of active jobs
// @Tags Jobs
// @Produce json
// @Param page query int false "Page number"
// @Param page_size query int false "Page size"
// @Success 200 {object} dtos.APIResponse
// @Failure 500 {object} dtos.APIResponse
// @Router /jobs/active [get]
func (h *JobHandler) GetActiveJobs(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	result, err := h.jobService.GetActiveJobs(page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dtos.APIResponse{
			Success: false,
			Error:   "Failed to get active jobs: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dtos.APIResponse{
		Success: true,
		Data:    result,
	})
}

// SearchJobs godoc
// @Summary Search jobs
// @Description Searches jobs with filters like query, work mode, skills, salary, etc.
// @Tags Jobs
// @Produce json
// @Param query query string false "Search keyword"
// @Param currency query string false "Currency code"
// @Param work_mode query string false "Comma-separated work modes"
// @Param work_type query string false "Comma-separated work types"
// @Param skills query string false "Comma-separated skills"
// @Param source query string false "Comma-separated sources"
// @Param min_salary query int false "Minimum salary"
// @Param max_salary query int false "Maximum salary"
// @Param is_active query bool false "Filter active jobs"
// @Param page query int false "Page number"
// @Param page_size query int false "Page size"
// @Success 200 {object} dtos.APIResponse
// @Failure 500 {object} dtos.APIResponse
// @Router /jobs/search [get]
func (h *JobHandler) SearchJobs(c *gin.Context) {
	// Parse query parameters
	params := &JobSearchParams{
		Query:    c.Query("query"),
		Currency: c.Query("currency"),
		Offset:   0,
		Limit:    20,
	}

	// Parse page and page_size
	if page, err := strconv.Atoi(c.DefaultQuery("page", "1")); err == nil && page > 0 {
		params.Offset = (page - 1) * params.Limit
	}
	if pageSize, err := strconv.Atoi(c.DefaultQuery("page_size", "20")); err == nil && pageSize > 0 && pageSize <= 100 {
		params.Limit = pageSize
		params.Offset = (params.Offset / 20) * pageSize // Adjust offset for new page size
	}

	// Parse array parameters
	if workModes := c.Query("work_mode"); workModes != "" {
		params.WorkMode = strings.Split(workModes, ",")
	}
	if workTypes := c.Query("work_type"); workTypes != "" {
		params.WorkType = strings.Split(workTypes, ",")
	}
	if skills := c.Query("skills"); skills != "" {
		params.Skills = strings.Split(skills, ",")
	}
	if sources := c.Query("source"); sources != "" {
		params.Source = strings.Split(sources, ",")
	}

	// Parse salary range
	if minSalaryStr := c.Query("min_salary"); minSalaryStr != "" {
		if minSalary, err := strconv.Atoi(minSalaryStr); err == nil {
			params.MinSalary = &minSalary
		}
	}
	if maxSalaryStr := c.Query("max_salary"); maxSalaryStr != "" {
		if maxSalary, err := strconv.Atoi(maxSalaryStr); err == nil {
			params.MaxSalary = &maxSalary
		}
	}

	// Parse active status
	if activeStr := c.Query("is_active"); activeStr != "" {
		if active, err := strconv.ParseBool(activeStr); err == nil {
			params.IsActive = &active
		}
	}

	result, err := h.jobService.SearchJobs(params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dtos.APIResponse{
			Success: false,
			Error:   "Search failed: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dtos.APIResponse{
		Success: true,
		Data:    result,
	})
}

// GetJobStats godoc
// @Summary Get job statistics
// @Description Returns statistics about jobs
// @Tags Jobs
// @Produce json
// @Success 200 {object} dtos.APIResponse
// @Failure 500 {object} dtos.APIResponse
// @Router /jobs/stats [get]
func (h *JobHandler) GetJobStats(c *gin.Context) {
	stats, err := h.jobService.GetJobStats()
	if err != nil {
		c.JSON(http.StatusInternalServerError, dtos.APIResponse{
			Success: false,
			Error:   "Failed to get job statistics: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dtos.APIResponse{
		Success: true,
		Data:    stats,
	})
}

// RegisterJobRoutes registers all job-related routes
func (h *JobHandler) RegisterJobRoutes(router *gin.RouterGroup) {
	jobs := router.Group("/jobs")
	{
		// Single job operations
		jobs.POST("", h.CreateJob)
		jobs.GET("/:id", h.GetJob)
		jobs.PUT("/:id", h.UpdateJob)
		jobs.DELETE("/:id", h.DeleteJob)
		jobs.PATCH("/:id/deactivate", h.DeactivateJob)

		// Batch operations
		jobs.POST("/batch", h.BatchCreateJobs)
		jobs.DELETE("/batch", h.BatchDeleteJobs)

		// Query operations
		jobs.GET("", h.GetAllJobs)
		jobs.GET("/active", h.GetActiveJobs)
		jobs.GET("/search", h.SearchJobs)
		jobs.GET("/stats", h.GetJobStats)

		// External job ID operations
		jobs.GET("/external/:job_id", h.GetJobByJobID)
	}
}
