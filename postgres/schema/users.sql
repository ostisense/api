create table if not exists users (
  id bigserial primary key,
  email text not null,
  bcrypted_password text not null,
  token text not null
);
