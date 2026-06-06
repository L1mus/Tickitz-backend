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