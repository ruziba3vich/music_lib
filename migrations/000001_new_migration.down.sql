CREATE TABLE IF NOT EXISTS songs (
    id UUID PRIMARY KEY,
    artists VARCHAR[] NOT NULL,
    group VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    lyrics TEXT,
    is_deleted BOOLEAN DEFAULT FALSE,
    release_date TIMESTAMP,
    created_at TIMESTAMP DEFAULT NOW()
);
