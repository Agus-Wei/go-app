CREATE EXTENSION "uuid-ossp";

CREATE TABLE items (
    id uuid DEFAULT uuid_generate_v4(),
    name VARCHAR NOT NULL,
	stock int NOT NULL,

    PRIMARY KEY (id)
);