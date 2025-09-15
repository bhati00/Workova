CREATE TABLE job_locations (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    job_id INTEGER NOT NULL,
    country_id INTEGER NOT NULL,
    city VARCHAR(100),

    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    deleted_at DATETIME,

    CONSTRAINT fk_job
        FOREIGN KEY (job_id) REFERENCES jobs(id)
        ON DELETE CASCADE,

    CONSTRAINT fk_country
        FOREIGN KEY (country_id) REFERENCES countries(id)
        ON DELETE CASCADE
);

-- indexes
CREATE INDEX idx_job_locations_job_id ON job_locations(job_id);
CREATE INDEX idx_job_locations_country_id ON job_locations(country_id);
