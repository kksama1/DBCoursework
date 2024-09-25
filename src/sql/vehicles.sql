CREATE TABLE IF NOT EXISTS Vehicles (
    vehicle_id SERIAL PRIMARY KEY,
    license_plate VARCHAR(20) NOT NULL,
    model VARCHAR(100),
    year INT,
    owner_id INT,
    FOREIGN KEY (owner_id) REFERENCES Participants(participant_id)
);
