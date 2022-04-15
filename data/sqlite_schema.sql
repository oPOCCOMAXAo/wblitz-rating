CREATE TABLE IF NOT EXISTS players (
    id INTEGER PRIMARY KEY NOT NULL DEFAULT 0,
    nickname TEXT NOT NULL DEFAULT "",
    battles INTEGER NOT NULL DEFAULT 0,
    damage INTEGER NOT NULL DEFAULT 0,
    wins INTEGER NOT NULL DEFAULT 0
);

CREATE TABLE IF NOT EXISTS rating (
    number INTEGER PRIMARY KEY NOT NULL DEFAULT 0,
    player_id INTEGER NOT NULL DEFAULT 0,
    score INTEGER NOT NULL DEFAULT 0,
    mmr REAL NOT NULL DEFAULT 0
);
