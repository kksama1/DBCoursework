CREATE TABLE IF NOT EXISTS Participants (
    participant_id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    age INT,
    contact_info VARCHAR(255),
    is_driver BOOLEAN DEFAULT FALSE
);
