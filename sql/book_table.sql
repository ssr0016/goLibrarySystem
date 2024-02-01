CREATE TABLE books (
    id SERIAL PRIMARY KEY, title VARCHAR(255) NOT NULL, author_id VARCHAR(255) NOT NULL, category_id VARCHAR(255) NOT NULL, published_at TIMESTAMP
);

CREATE TABLE categories (
    category_id SERIAL PRIMARY KEY, name VARCHAR(255) NOT NULL
);

CREATE TABLE authors (
    author_id SERIAL PRIMARY KEY, name VARCHAR(255) NOT NULL
);