CREATE TABLE IF NOT EXISTS users (
   id INTEGER PRIMARY KEY,
   email TEXT UNIQUE NOT NULL,
   password_hash TEXT NOT NULL,
   is_admin INTEGER NOT NULL
);

