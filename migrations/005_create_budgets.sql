CREATE TABLE budgets (
  id_budg     SERIAL PRIMARY KEY,
  id_user     INT NOT NULL REFERENCES users(id_user) ON DELETE CASCADE,
  id_category INT NOT NULL REFERENCES categories(id_category),
  amount_budg DECIMAL(15, 2) NOT NULL,
  month_budg  INT NOT NULL CHECK (month_budg BETWEEN 1 AND 12),
  year_budg   INT NOT NULL
);
