-- schema.sql

-- Drop existing tables if they exist
DROP TABLE IF EXISTS emoji_prefixes;
DROP TABLE IF EXISTS emojis;

-- Create emojis table
CREATE TABLE IF NOT EXISTS emojis (
    id INTEGER PRIMARY KEY AUTOINCREMENT, 
    name TEXT NOT NULL UNIQUE COLLATE NOCASE, 
    symbol TEXT NOT NULL COLLATE NOCASE
);

-- Create emoji_prefixes table
CREATE TABLE IF NOT EXISTS emoji_prefixes (
    prefix TEXT NOT NULL, 
    emoji_name TEXT NOT NULL, 
    FOREIGN KEY (emoji_name) REFERENCES emojis(name) ON DELETE CASCADE
);
