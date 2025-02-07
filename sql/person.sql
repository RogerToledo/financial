-- financial.person definition

-- Drop table

-- DROP TABLE financial.person;

CREATE TABLE financial.person (
	id serial4 NOT NULL,
	"name" varchar(50) NOT NULL,
	CONSTRAINT person_pk PRIMARY KEY (id),
	CONSTRAINT person_unique UNIQUE (name)
);