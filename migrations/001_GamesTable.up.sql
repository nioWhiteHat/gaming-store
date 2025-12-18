CREATE TABLE games (
    id                SERIAL PRIMARY KEY,  -- internal auto-increment ID
    external_id       INT NOT NULL,        -- keeps the RAWG external ID
    slug              TEXT NOT NULL,
    name              TEXT NOT NULL,
    description       TEXT,
    released          DATE,
    main_image  TEXT,
    image             TEXT,
    region_id         INTEGER,
    platform_id          INTEGER,
    
    -- Ratings split into 4 integer columns
    added              INT DEFAULT 0,
    rating_recommended INT DEFAULT 0,
    rating_meh         INT DEFAULT 0,
    rating_exceptional INT DEFAULT 0,
    rating_skip        INT DEFAULT 0
    
   
    
    
    
);
