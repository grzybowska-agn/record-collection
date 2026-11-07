CREATE TABLE records (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    artist TEXT NOT NULL,
    title TEXT NOT NULL,
    original_year INTEGER NOT NULL,
    cover_url TEXT,
    pressing_year INTEGER,
    pressing_label TEXT,
    is_reissue BOOLEAN,
    size_inches INTEGER NOT NULL CHECK (size_inches IN (7, 10, 12)),
    release_type TEXT NOT NULL CHECK (release_type IN ('LP', 'EP', 'Single')),
    disc_count INTEGER NOT NULL CHECK (disc_count > 0),
    pressing_country TEXT
);