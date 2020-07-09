CREATE TABLE task(
	id serial PRIMARY KEY,
	name VARCHAR(50) NOT NULL,
	description VARCHAR(200),
	id_status INTEGER references status(id)
)