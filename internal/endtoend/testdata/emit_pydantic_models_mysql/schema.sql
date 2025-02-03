CREATE TABLE authors (
          id   BIGINT AUTO_RANDOM PRIMARY KEY,
          name VARCHAR(255)      NOT NULL,
          bio  text,
          created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
