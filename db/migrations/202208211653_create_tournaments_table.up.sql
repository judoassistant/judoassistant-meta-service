CREATE TABLE IF NOT EXISTS tournaments (
   id INTEGER PRIMARY KEY,
   name TEXT NOT NULL,
   location TEXT NOT NULL,
   date DATETIME NOT NULL,
   is_deleted INTEGER NOT NULL,
   owner INTEGER NOT NULL,
   FOREIGN KEY(owner) REFERENCES users(id)
);

