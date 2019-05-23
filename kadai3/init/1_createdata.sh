#!/bin/bash
set -e
psql -U root -d Godb <<-EOSQL
  create table users
  (
      id serial,
      name text,
      email text, 
      created_at timestamp not null default current_timestamp,
      updated_at timestamp not null default current_timestamp,
      PRIMARY KEY (id)
  );
  create function set_update_time() returns opaque as '
    begin
      new.updated_at := ''now'';
      return new;
    end;
  ' language 'plpgsql';
  create trigger update_tri before update on users for each row
    execute procedure set_update_time();
EOSQL