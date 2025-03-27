CREATE TABLE jobs (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP,
    title VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    job_type VARCHAR(50), -- full-time, part-time, etc.
    location VARCHAR(255) NOT NULL,
    status VARCHAR(50) NOT NULL, -- active, inactive, etc.
    category_id UUID NOT NULL,
    CONSTRAINT fk_category FOREIGN KEY (category_id) REFERENCES job_categories(id) ON DELETE SET NULL
);
