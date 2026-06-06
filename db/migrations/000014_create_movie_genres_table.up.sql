CREATE TABLE movie_genres
(
    movie_id INTEGER,
    genre_id INTEGER,
    PRIMARY KEY (movie_id, genre_id),
    CONSTRAINT fk_mg_movie FOREIGN KEY (movie_id) REFERENCES movies (id),
    CONSTRAINT fk_mg_genre FOREIGN KEY (genre_id) REFERENCES genres (id)
);