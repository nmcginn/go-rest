#!/bin/bash -ex
psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" <<-EOSQL
	-- test schema
	CREATE DATABASE docker;
	\c docker;
	CREATE TABLE public.users (
		user_id       SERIAL,
		user_name     VARCHAR(40),
		date_of_birth DATE,
		email         VARCHAR(40),
		is_cool_dude  BOOLEAN
	);
	INSERT INTO users (user_name, date_of_birth, email, is_cool_dude) VALUES
		('nmcginn', current_date, 'nmcginn@users.noreply.github.com', TRUE),
		('nd', current_date, 'nd@cooldomain.com', FALSE),
		('gs', current_date, 'gs@okaydomain.com', TRUE);
	-- privs
	CREATE USER docker PASSWORD 'docker';
	GRANT ALL PRIVILEGES ON DATABASE docker TO docker;
	GRANT USAGE ON SCHEMA public to docker;
	GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO docker;
EOSQL
