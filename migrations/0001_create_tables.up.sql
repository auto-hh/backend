CREATE TYPE JOB_GRADE AS ENUM ('junior', 'middle', 'senior');

CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT UUIDV7(),
    hh_id VARCHAR(32) NOT NULL UNIQUE CHECK (hh_id != ''),
    first_name VARCHAR(32) NOT NULL CHECK (first_name != ''),
    last_name VARCHAR(32) NOT NULL check (last_name != '')
);

CREATE TABLE profiles (
    id UUID PRIMARY KEY DEFAULT UUIDV7(),
    user_id UUID REFERENCES users (id) ON DELETE CASCADE,
    experience VARCHAR(32) NOT NULL CHECK (experience != ''),
    job_title VARCHAR(32) NOT NULL CHECK (work_title != ''),
    grade JOB_GRADE NOT NULL,
    work_format VARCHAR(32) NOT NULL CHECK (work_format != ''),
    salary INTEGER NOT NULL CHECK (salary BETWEEN 0 AND 2000000),
    city VARCHAR(32) NOT NULL CHECK (city != ''),
    about_me VARCHAR(1024) NOT NULL CHECK (about_me != ''),
    recent_jobs VARCHAR(4096) NOT NULL CHECK (recent_jobs != '')
);
