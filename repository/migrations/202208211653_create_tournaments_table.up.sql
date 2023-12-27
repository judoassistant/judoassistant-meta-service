CREATE TABLE IF NOT EXISTS tournaments (
   id INTEGER PRIMARY KEY,
   name TEXT NOT NULL,
   location TEXT NOT NULL,
   date DATETIME NOT NULL,
   is_deleted INTEGER NOT NULL,
   owner INTEGER NOT NULL,
			short_name TEXT NOT NULL,
   FOREIGN KEY(owner) REFERENCES users(id)
);

CREATE UNIQUE INDEX idx_tournaments_short_name ON tournaments (short_name);
