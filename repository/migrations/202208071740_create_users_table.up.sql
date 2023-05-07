CREATE TABLE IF NOT EXISTS users (
   id INTEGER PRIMARY KEY,
   email TEXT UNIQUE NOT NULL,
   first_name TEXT NOT NULL,
   last_name TEXT NOT NULL,
   password_hash TEXT NOT NULL,
   is_admin INTEGER NOT NULL
);

