CREATE TABLE showtimes
(
    id        SERIAL PRIMARY KEY,
    movie_id  INTEGER,
    cinema_id INTEGER,
    date      DATE    NOT NULL,
    time      TIME    NOT NULL,
    price     INTEGER NOT NULL,
    CONSTRAINT fk_showtimes_movie FOREIGN KEY (movie_id) REFERENCES movies (id),
    CONSTRAINT fk_showtimes_cinema FOREIGN KEY (cinema_id) REFERENCES cinemas (id)
);