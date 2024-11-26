CREATE TYPE operator AS ENUM (
    '=',
    '>',
    '<',
    '>=',
    '<='
);

CREATE TABLE operations (
          a int,
          b int,
          operation operator
);
