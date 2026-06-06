CREATE TABLE seats (
                       cinema_id INTEGER,
                       seat_number INTEGER NOT NULL,
                       row VARCHAR(10) NOT NULL,
                       CONSTRAINT fk_seats_cinema FOREIGN KEY (cinema_id) REFERENCES cinemas(id)
);