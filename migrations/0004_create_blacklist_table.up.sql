CREATE TABLE blacklist
(
    cidr INET NOT NULL,
    CONSTRAINT blacklist_unique_cidr UNIQUE (cidr)
);
