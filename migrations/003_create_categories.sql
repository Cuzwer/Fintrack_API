CREATE TABLE categories (
  id_category   SERIAL PRIMARY KEY,
  id_user       INT NOT NULL REFERENCES users(id_user) ON DELETE CASCADE,
  name_category VARCHAR(100) NOT NULL,
  type_category VARCHAR(20) NOT NULL CHECK (type_category IN ('income', 'expense'))
);
