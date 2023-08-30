CREATE TABLE users_with_segments (
    user_id BIGINT,
    seg_id BIGINT REFERENCES segments ON DELETE CASCADE,
    PRIMARY KEY (user_id, seg_id)
);