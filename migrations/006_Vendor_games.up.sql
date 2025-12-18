CREATE TABLE IF NOT EXISTS vendor_games(
    vendor_id       INTEGER,
    game_id         INTEGER,

    CONSTRAINT fk_vendor
        FOREIGN KEY(vendor_id) 
        REFERENCES users(id)
        ON DELETE CASCADE, 
    
    CONSTRAINT fk_game
        FOREIGN KEY(game_id) 
        REFERENCES games(id)
        ON DELETE CASCADE
);