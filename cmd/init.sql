create database film_collection_db;
create table if not exists actors(
	gender text NOT NULL,
	actor_name text NOT NULL,
	birth_date TIMESTAMP NOT NULL
);
	
create table if not exists films(
	actors_cast text[],
	names text UNIQUE,
	overview text NULL,
	release_date TIMESTAMP NOT NULL,
	rate integer NULL
);