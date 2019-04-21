create table if not exists users (
  id bigserial primary key,
  email text not null,

  bcrypted_password text not null,
  token text not null,

  created_at timestamp not null
);

create unique index if not exists users_email_index on users (email);
create unique index if not exists users_token_index on users (token);
