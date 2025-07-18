CREATE TABLE whitelist
(
    cidr INET NOT NULL,
    CONSTRAINT whitelist_unique_cidr UNIQUE (cidr)
);
