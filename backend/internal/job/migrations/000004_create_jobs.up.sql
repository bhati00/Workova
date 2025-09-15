CREATE TABLE jobs (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    slug VARCHAR(150),
    external_job_id VARCHAR(255),

    title VARCHAR(255) NOT NULL,
    description TEXT,
    summary TEXT,

    company_name VARCHAR(255) NOT NULL,
    company_size VARCHAR(50),
    company_logo_url TEXT,
    company_website TEXT,
    company_industry VARCHAR(100),

    job_type INTEGER NOT NULL DEFAULT 1,
    work_mode INTEGER NOT NULL DEFAULT 1,
    experience_level INTEGER,

    is_remote BOOLEAN DEFAULT 0,

    salary_min INTEGER,
    salary_max INTEGER,
    salary_currency VARCHAR(10) DEFAULT 'USD',
    salary_period VARCHAR(20),
    salary_is_estimate BOOLEAN DEFAULT 0,
    equity_offered BOOLEAN,

    years_experience_min INTEGER,
    years_experience_max INTEGER,

    benefits TEXT,
    health_insurance BOOLEAN,
    paid_time_off BOOLEAN,
    flexible_schedule BOOLEAN,

    application_url TEXT,
    application_email VARCHAR(255),
    contact_person VARCHAR(255),
    how_to_apply TEXT,

    education_level VARCHAR(100),
    education_required VARCHAR(255),

    contract_duration INTEGER,
    start_date DATETIME,
    is_urgent BOOLEAN DEFAULT 0,
    travel_required VARCHAR(50),
    remote_location_restriction VARCHAR(255),

    interview_process VARCHAR(50),
    hiring_process_steps INTEGER,
    expected_hire_date DATETIME,

    visa_sponsorship BOOLEAN,
    security_clearance VARCHAR(50),
    background_check_required BOOLEAN,

    source VARCHAR(100) NOT NULL,
    source_url TEXT,
    posted_date DATETIME,
    expiry_date DATETIME,
    application_deadline DATETIME,

    keywords TEXT,
    industry VARCHAR(100),
    department VARCHAR(100),
    tags TEXT,

    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    deleted_at DATETIME,

    UNIQUE(source)
);
