CREATE TABLE movie_directors
(
    movie_id    INTEGER,
    director_id INTEGER,
    PRIMARY KEY (movie_id, director_id),
    CONSTRAINT fk_md_movie FOREIGN KEY (movie_id) REFERENCES movies (id),
    CONSTRAINT fk_md_director FOREIGN KEY (director_id) REFERENCES directors (id)
);