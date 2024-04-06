CREATE TABLE links (
  id SERIAL PRIMARY KEY,
  short_url VARCHAR(255) NOT NULL,
  original_url VARCHAR(255) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  access_count INT DEFAULT 0
);

CREATE INDEX links_short_url_index ON links (short_url);