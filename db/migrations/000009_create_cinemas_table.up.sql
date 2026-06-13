CREATE TABLE cinemas
(
    id          SERIAL PRIMARY KEY,
    location_id INTEGER,
    name        VARCHAR(255) NOT NULL,
    logo        VARCHAR(255),
    capacity    INTEGER,
    isAvailable BOOLEAN DEFAULT true,
    CONSTRAINT fk_cinemas_location FOREIGN KEY (location_id) REFERENCES locations (id)
);