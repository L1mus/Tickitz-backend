CREATE TABLE movie_casts
(
    movie_id INTEGER,
    cast_id  INTEGER,
    PRIMARY KEY (movie_id, cast_id),
    CONSTRAINT fk_mc_movie FOREIGN KEY (movie_id) REFERENCES movies (id),
    CONSTRAINT fk_mc_cast FOREIGN KEY (cast_id) REFERENCES casts (id)
);