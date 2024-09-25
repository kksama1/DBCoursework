CREATE TABLE IF NOT EXISTS Drivers (
    driver_id INT PRIMARY KEY,
    license_number VARCHAR(50) NOT NULL,
    experience_years INT,
    FOREIGN KEY (driver_id) REFERENCES Participants(participant_id)
);
