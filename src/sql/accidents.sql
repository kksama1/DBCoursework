CREATE TABLE IF NOT EXISTS Accidents (
    accident_id SERIAL PRIMARY KEY,
    date TIMESTAMP NOT NULL,
    location VARCHAR(255),
    description TEXT
);
