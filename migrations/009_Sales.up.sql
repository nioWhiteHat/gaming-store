CREATE TABLE IF NOT EXISTS orders (
    order_id SERIAL PRIMARY KEY,
    
    
    client_id bigint NOT NULL,
    game_id bigint NOT NULL,
    vendor_id bigint NOT NULL,
    
    price_at_sale NUMERIC(10, 2) NOT NULL,
    
    created_at timestamptz NOT NULL DEFAULT NOW(),

    game_key text NOT NULL UNIQUE,

  
    CONSTRAINT fk_client
        FOREIGN KEY(client_id) 
        REFERENCES users(id)
        ON DELETE SET NULL, 

    CONSTRAINT fk_vendor
        FOREIGN KEY(vendor_id) 
        REFERENCES users(id) 
        ON DELETE RESTRICT, 

    CONSTRAINT fk_game
        FOREIGN KEY(game_id) 
        REFERENCES games(id) 
        ON DELETE RESTRICT 
);