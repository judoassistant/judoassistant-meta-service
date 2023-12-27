CREATE TABLE IF NOT EXISTS tournaments (
   id INTEGER PRIMARY KEY,
   name TEXT NOT NULL,
   location TEXT NOT NULL,
   date DATETIME NOT NULL,
   is_deleted INTEGER NOT NULL,
   owner INTEGER NOT NULL,
			url_slug TEXT NOT NULL,
   FOREIGN KEY(owner) REFERENCES users(id)
);

CREATE UNIQUE INDEX idx_tournaments_url_slug ON tournaments (url_slug);
