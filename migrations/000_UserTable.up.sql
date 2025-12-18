CREATE TABLE IF NOT EXISTS users(
    id SERIAL PRIMARY KEY,
    type TEXT,
    username TEXT,
    password TEXT,
    email TEXT,
    created_at TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT NOW(),
    image TEXT,
    bio text
);