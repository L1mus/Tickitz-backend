CREATE TABLE bookings
(
    id            SERIAL PRIMARY KEY,
    user_id       INTEGER,
    showtime_id   INTEGER,
    status_ticket condition_ticket,
    status_paid   condition_paid,
    quantity      INTEGER NOT NULL,
    created_at    TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at    TIMESTAMP,
    CONSTRAINT fk_bookings_user FOREIGN KEY (user_id) REFERENCES users (id),
    CONSTRAINT fk_bookings_showtime FOREIGN KEY (showtime_id) REFERENCES showtimes (id)
);