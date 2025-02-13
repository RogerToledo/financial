-- financial.person definition

-- Drop table

-- DROP TABLE financial.person;

CREATE TABLE financial.person (
	id uuid NOT NULL,
	"name" varchar(50) NOT NULL,
	create_at date DEFAULT now() NOT NULL,
	update_at date NULL,
	CONSTRAINT person_pk PRIMARY KEY (id),
	CONSTRAINT person_unique UNIQUE (name)
);