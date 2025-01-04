CREATE TYPE book_status AS ENUM ('available', 'checked_out', 'overdue');


CREATE TABLE books (
          id     BIGSERIAL PRIMARY KEY,
          title  text      NOT NULL,
          status book_status DEFAULT 'available'
);
