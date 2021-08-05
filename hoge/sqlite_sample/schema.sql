CREATE TABLE IF NOT EXISTS tasks (
    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    name varchar(255) not null,
    detail text,
    created_at TIMESTAMP DEFAULT (DATETIME('now','localtime')),
    updated_at TIMESTAMP DEFAULT (DATETIME('now','localtime')),
    deleted_at TIMESTAMP
);
