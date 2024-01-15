-- ddl create table user
create table if not exists users (
	id serial PRIMARY KEY,
	full_name varchar NOT NULL,
	username varchar UNIQUE NOT NULL,
	password varchar NOT NULL,
	age int NOT NULL
);
