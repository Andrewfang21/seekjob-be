CREATE TABLE jobs (
  id VARCHAR(50) PRIMARY KEY NOT NULL UNIQUE,
  url VARCHAR(255) NOT NULL,
  title TEXT NOT NULL,
  company VARCHAR(255) NOT NULL,
  description TEXT NOT NULL,
  category VARCHAR(255) NOT NULL,
  country VARCHAR(100) NOT NULL,
  type VARCHAR(100) NOT NULL,
  time INTEGER NOT NULL,
  source VARCHAR(20) NOT NULL
);