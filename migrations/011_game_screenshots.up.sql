CREATE TABLE game_screenshots (
    game_id         INT NOT NULL,
    screenshot_id  INT NOT NULL,    
    
    PRIMARY KEY (screenshot_id,game_id),
    FOREIGN KEY (game_id) REFERENCES games(id) ON DELETE CASCADE,
    FOREIGN KEY (screenshot_id) REFERENCES screenshots(id) ON DELETE CASCADE
);
