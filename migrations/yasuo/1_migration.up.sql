CREATE TABLE IF NOT EXISTS `users` (
  user_id TEXT NOT NULL,
  email TEXT NOT NULL,
  name TEXT NOT NULL,
  gender TEXT,
  avatar TEXT,
  birthday Timestamp with time zone,
  street TEXT NOT NULL,
  city TEXT NOT NULL,
  state TEXT NOT NULL,
  country TEXT NOT NULL,
  created_at TIMESTAMP with time zone NOT NULL,
  updated_at TIMESTAMP with time zone NOT NULL,
  deleted_at TIMESTAMP with time zone;
  CONSTRAINT users_pk PRIMARY KEY (user_id)
)