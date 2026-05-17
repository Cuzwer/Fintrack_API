CREATE TABLE accounts (
  id_account   SERIAL PRIMARY KEY,
  id_user      INT NOT NULL REFERENCES users(id_user) ON DELETE CASCADE,
  name         VARCHAR(100) NOT NULL,
  type_account VARCHAR(50) NOT NULL,
  balance      DECIMAL(15, 2) DEFAULT 0.00,
  currency     VARCHAR(10) DEFAULT 'THB'
);
