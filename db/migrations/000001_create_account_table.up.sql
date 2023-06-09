CREATE TABLE IF NOT EXISTS account(
  id serial PRIMARY KEY,
  name VARCHAR (50) UNIQUE NOT NULL,
  
  
  -- Audit
  created_at TIMESTAMP WITH TIME ZONE,
  created_by VARCHAR (50),
  updated_at TIMESTAMP WITH TIME ZONE,
  updated_by VARCHAR (50)
)
