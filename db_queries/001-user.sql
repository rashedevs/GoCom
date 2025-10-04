CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    is_shop_owner BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP WITH ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH ZONE DEFAULT CURRENT_TIMESTAMP
);

/*
CLASS : 62

Basic Data Types
--------------------

SERIAL - 1,2 3,4 etc (32 bits) approx 429 crore
BIGSERIAL - 1,2,3 (64bits) huge

SMALLINT - 16 bits
INT - 32 bits
BIGINT - 64 bits

REAL - 10.33 (less precision) - 32 bits
DOUBLEPRECISION - (high precision for money like things) - 64 bits

CHAR(n) - Fixed character size - first_name CHAR(6) - a -> "a....." -> memory waste

*/