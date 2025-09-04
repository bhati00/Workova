CREATE TABLE job_skills (
    id SERIAL PRIMARY KEY,
    job_id INT NOT NULL,
    skill VARCHAR(100) NOT NULL,
    type VARCHAR(20),
    CONSTRAINT fk_job FOREIGN KEY (job_id) REFERENCES jobs (id) ON DELETE CASCADE
);