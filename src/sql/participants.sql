CREATE TABLE IF NOT EXISTS Participants (
    participant_id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    date_of_birth DATE,
    license_number VARCHAR(255),
    is_driver BOOLEAN DEFAULT FALSE
);
