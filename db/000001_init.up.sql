CREATE TABLE tasks
(
    id SERIAL PRIMARY KEY,
    status VARCHAR(12) NOT NULL DEFAULT 'opened',
    title VARCHAR(30),
    description VARCHAR(120)
);