DO $$ BEGIN
    CREATE TYPE role_user AS ENUM ('user', 'admin');
EXCEPTION WHEN duplicate_object THEN NULL;
END $$;

CREATE TABLE IF NOT EXISTS locations(
    id SERIAL PRIMARY KEY,
    city VARCHAR(255)
);