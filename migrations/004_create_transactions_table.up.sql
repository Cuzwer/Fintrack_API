CREATE TABLE transactions (
  id_trans         SERIAL PRIMARY KEY,
  id_account       INT NOT NULL REFERENCES accounts(id_account) ON DELETE CASCADE,
  id_category      INT REFERENCES categories(id_category),
  amount_trans     DECIMAL(15, 2) NOT NULL,
  type_trans       VARCHAR(20) NOT NULL CHECK (type_trans IN ('income', 'expense')),
  descrip_trans    VARCHAR(255),
  transaction_date DATE NOT NULL DEFAULT CURRENT_DATE,
  created_at       TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

