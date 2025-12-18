CREATE TABLE IF NOT EXISTS vendor_keys (
    vendor_id   INTEGER,
    key_id      INTEGER,
  
    CONSTRAINT fk_vendor
        FOREIGN KEY(vendor_id) 
        REFERENCES users(id)
        ON DELETE CASCADE, 

    CONSTRAINT fk_key
        FOREIGN KEY(key_id) 
        REFERENCES keys(id)
        ON DELETE CASCADE 
);