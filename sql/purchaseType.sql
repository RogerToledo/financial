-- financial.purchase_type definition

-- Drop table

-- DROP TABLE financial.purchase_type;

CREATE TABLE financial.purchase_type (
	id serial4 NOT NULL,
	"name" varchar(50) NOT NULL,
	CONSTRAINT purchase_type_pk PRIMARY KEY (id),
	CONSTRAINT purchase_type_unique UNIQUE (name)
);