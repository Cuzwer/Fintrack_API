CREATE TABLE IF NOT EXISTS users (
  id_user       SERIAL PRIMARY KEY,
  email_user    VARCHAR(100) NOT NULL UNIQUE,
  password_hash VARCHAR(255) NOT NULL,
  created_at    TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);



CREATE INDEX idx_users_email ON users(email_user);
