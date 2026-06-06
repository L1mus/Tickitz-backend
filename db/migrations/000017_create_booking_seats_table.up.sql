CREATE TABLE booking_seats
(
    id          SERIAL PRIMARY KEY,
    booking_id  INTEGER,
    seat_id     INTEGER,
    showtime_id INTEGER,
    CONSTRAINT fk_bs_booking FOREIGN KEY (booking_id) REFERENCES bookings (id),
    CONSTRAINT fk_bs_seat FOREIGN KEY (seat_id) REFERENCES seats (id),
    CONSTRAINT fk_bs_showtime FOREIGN KEY (showtime_id) REFERENCES showtimes (id)
);