#!/bin/bash
psql -U root -d Godb << "EOSQL"
create table user (id int auto_increment PRIMARY KEY, name text, email text, created_at datetime default current_timestamp, updated_at timestamp default current_timestamp on update current_timestamp);
EOSQL