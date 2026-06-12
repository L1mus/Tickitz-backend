CREATE TYPE status_role AS ENUM ('user', 'admin');
CREATE TYPE condition_ticket AS ENUM ('active', 'not_active');
CREATE TYPE condition_paid AS ENUM ('paid', 'not_paid');
CREATE TYPE status_transaction AS ENUM ('completed', 'pending', 'failed');

CREATE TABLE locations (
    id SERIAL PRIMARY KEY,
    city VARCHAR(255) NOT NULL
);

CREATE TABLE users(
    id SERIAL PRIMARY KEY,
    email VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    first_name VARCHAR(255),
    last_name VARCHAR(255),
    phone VARCHAR(255),
    photo VARCHAR(255),
    location_id INT NOT NULL,
    isActive BOOLEAN NOT NULL DEFAULT FALSE,
    role role_user NOT NULL DEFAULT 'user',
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP,

    CONSTRAINT fk_forgot_locations FOREIGN KEY (location_id)
        REFERENCES locations (id) ON DELETE CASCADE
);

CREATE TABLE payment_methods (
    id SERIAL PRIMARY KEY,
    logo VARCHAR(255),
    name VARCHAR(255) NOT NULL
);

CREATE TABLE genres (
    id SERIAL PRIMARY KEY,
    genre VARCHAR(255) NOT NULL
);

CREATE TABLE casts (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE directors (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE movies
(
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    duration INTERVAL,
    poster VARCHAR(255),
    release_date DATE,
    synopsis TEXT,
    category VARCHAR(155),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP
);

CREATE TABLE cinemas (
    id SERIAL PRIMARY KEY,
    location_id INTEGER,
    name VARCHAR(255) NOT NULL,
    logo VARCHAR(255),
    capacity INTEGER,
    isAvailable BOOLEAN DEFAULT true,
    CONSTRAINT fk_cinemas_location FOREIGN KEY (location_id) REFERENCES locations(id)
);

CREATE TABLE seats (
    id SERIAL PRIMARY KEY,
    cinema_id INTEGER,
    seat_number INTEGER NOT NULL,
    row VARCHAR(10) NOT NULL,
    seat_type type_seat DEFAULT 'regular',
    CONSTRAINT fk_seats_cinema FOREIGN KEY (cinema_id) REFERENCES cinemas(id)
);

CREATE TABLE showtimes (
    id SERIAL PRIMARY KEY,
    movie_id INTEGER,
    cinema_id INTEGER,
    date DATE NOT NULL,
    time TIME NOT NULL,
    price INTEGER NOT NULL,
    CONSTRAINT fk_showtimes_movie FOREIGN KEY (movie_id) REFERENCES movies(id),
    CONSTRAINT fk_showtimes_cinema FOREIGN KEY (cinema_id) REFERENCES cinemas(id)
);

CREATE TABLE bookings (
    id SERIAL PRIMARY KEY,
    user_id INTEGER,
    showtime_id INTEGER,
    status_ticket condition_ticket,
    status_paid condition_paid,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    CONSTRAINT fk_bookings_user FOREIGN KEY (user_id) REFERENCES users(id),
    CONSTRAINT fk_bookings_showtime FOREIGN KEY (showtime_id) REFERENCES showtimes(id)
);

CREATE TABLE transactions (
    id SERIAL PRIMARY KEY,
    booking_id INTEGER,
    payment_method_id INTEGER,
    virtual_rek INTEGER,
    total_price INTEGER NOT NULL,
    status status_transaction,
    qr_code VARCHAR(255),
    CONSTRAINT fk_transactions_booking FOREIGN KEY (booking_id) REFERENCES bookings(id),
    CONSTRAINT fk_transactions_payment FOREIGN KEY (payment_method_id) REFERENCES payment_methods(id)
);

CREATE TABLE movie_genres (
    movie_id INTEGER,
    genre_id INTEGER,
    PRIMARY KEY (movie_id, genre_id),
    CONSTRAINT fk_mg_movie FOREIGN KEY (movie_id) REFERENCES movies(id),
    CONSTRAINT fk_mg_genre FOREIGN KEY (genre_id) REFERENCES genres(id)
);

CREATE TABLE movie_casts (
    movie_id INTEGER,
    cast_id INTEGER,
    PRIMARY KEY (movie_id, cast_id),
    CONSTRAINT fk_mc_movie FOREIGN KEY (movie_id) REFERENCES movies(id),
    CONSTRAINT fk_mc_cast FOREIGN KEY (cast_id) REFERENCES casts(id)
);

CREATE TABLE movie_directors (
    movie_id INTEGER,
    director_id INTEGER,
    PRIMARY KEY (movie_id, director_id),
    CONSTRAINT fk_md_movie FOREIGN KEY (movie_id) REFERENCES movies(id),
    CONSTRAINT fk_md_director FOREIGN KEY (director_id) REFERENCES directors(id)
);

CREATE TABLE booking_seats (
    id SERIAL PRIMARY KEY,
    booking_id INTEGER,
    seat_id INTEGER,
    showtime_id INTEGER,
    CONSTRAINT fk_bs_booking FOREIGN KEY (booking_id) REFERENCES bookings(id),
    CONSTRAINT fk_bs_seat FOREIGN KEY (seat_id) REFERENCES seats(id),
    CONSTRAINT fk_bs_showtime FOREIGN KEY (showtime_id) REFERENCES showtimes(id),
    CONSTRAINT unique_booked_seat UNIQUE (seat_id, showtime_id)
);