CREATE TABLE IF NOT EXISTS Accident_Participants (
    participant_id INT,
    accident_id INT,
    vehicle_id INT NULL,
    is_responsible BOOLEAN DEFAULT FALSE,
    role VARCHAR(50) NOT NULL CHECK (role IN ('водитель', 'пешеход')),
    PRIMARY KEY (participant_id, accident_id),
    FOREIGN KEY (participant_id) REFERENCES Participants(participant_id),
    FOREIGN KEY (accident_id) REFERENCES Accidents(accident_id),
    FOREIGN KEY (vehicle_id) REFERENCES Vehicles(vehicle_id)
);
