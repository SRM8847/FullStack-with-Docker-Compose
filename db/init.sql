
-- This file is going to be execute after the db is built and created using the given username and password
create table if not exists notes(
id serial primary key,
name text not null,
heading text,
content text,
created_at Timestamp with time zone default CURRENT_TIMESTAMP
);
