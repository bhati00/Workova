-- Migration file: XXXXXX_add_timestamps_to_jobs_and_job_skills.up.sql

-- Add timestamp fields to jobs table
ALTER TABLE jobs ADD COLUMN created_at DATETIME DEFAULT CURRENT_TIMESTAMP;

ALTER TABLE jobs ADD COLUMN updated_at DATETIME DEFAULT CURRENT_TIMESTAMP;

ALTER TABLE jobs ADD COLUMN deleted_at DATETIME;

-- Add timestamp fields to job_skills table
ALTER TABLE job_skills ADD COLUMN created_at DATETIME DEFAULT CURRENT_TIMESTAMP;

ALTER TABLE job_skills ADD COLUMN updated_at DATETIME DEFAULT CURRENT_TIMESTAMP;

ALTER TABLE job_skills ADD COLUMN deleted_at DATETIME;