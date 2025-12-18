CREATE TABLE IF NOT EXISTS screenshots(
    id  SERIAL PRIMARY KEY,
    screenshot_url text,
    width INT,
    height INT
)