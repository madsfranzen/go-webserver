CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TABLE IF NOT EXISTS users (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  username TEXT NOT NULL,
  email TEXT UNIQUE NOT NULL
);

-- Insert with UUIDs generated automatically
INSERT INTO users (username, email)
VALUES 
  ('Alice', 'alice@example.com'),
  ('Bob', 'bob@example.com');

