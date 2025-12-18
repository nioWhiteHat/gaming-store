CREATE TABLE IF NOT EXISTS keys(
    id          SERIAL PRIMARY KEY,
    game_id     INTEGER,
    store_id    INTEGER,
    key_hash    TEXT,
    price       FLOAT, -- Note: NUMERIC(10,2) is usually better for money than FLOAT
    
    

    created_at  TIMESTAMPTZ NOT NULL DEFAULT now(),

    CONSTRAINT fk_game
        FOREIGN KEY(game_id) 
        REFERENCES games(id)
        ON DELETE CASCADE, 

    CONSTRAINT fk_store
        FOREIGN KEY(store_id) 
        REFERENCES stores(id)
        ON DELETE CASCADE
    
);

