CREATE TABLE bucket_configs
(
    type      VARCHAR(255) NOT NULL,
    try_count INTEGER      NOT NULL CHECK (try_count > 0)
);
