CREATE TABLE IF NOT EXISTS vendor_likes_dislikes(
    id integer,
    likes integer,
    dislikes integer,

    CONSTRAINT fk_vendor
        FOREIGN KEY(id) 
        REFERENCES users(id)
        ON DELETE CASCADE
)