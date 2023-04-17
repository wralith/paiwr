create table
  if not exists users (
    id uuid primary key,
    username varchar unique,
    email varchar unique,
    hashed_password bytea,
    bio varchar,
    created_at timestamptz,
    updated_at timestamptz
  );
