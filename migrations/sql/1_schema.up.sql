-- Users table
CREATE TABLE users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    email TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL
);

-- Static Sites table
CREATE TABLE static_sites (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    owner_id INTEGER NOT NULL,
    hostname TEXT UNIQUE NOT NULL,
    file_path TEXT,
    FOREIGN KEY (owner_id) REFERENCES users(id)
);

-- Uploads table
CREATE TABLE uploads (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    site_id INTEGER NOT NULL,
    file_name TEXT NOT NULL,
    upload_path TEXT NOT NULL,
    upload_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (site_id) REFERENCES static_sites(id)
);
