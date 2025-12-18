CREATE TABLE IF NOT EXISTS vendor_reviews(
    vendor_id integer,
    client_id integer,
    review text,
    rating integer,

    CONSTRAINT fk_vendorid
        FOREIGN KEY(vendor_id) 
        REFERENCES users(id)
        ON DELETE CASCADE,
    
    CONSTRAINT fk_clientid
        FOREIGN KEY(client_id) 
        REFERENCES users(id)
        ON DELETE CASCADE
)