-- financial.payment_type definition

-- Drop table

-- DROP TABLE financial.payment_type;

CREATE TABLE financial.payment_type (
	id uuid NOT NULL,
	"name" varchar(50) NOT NULL,
	CONSTRAINT payment_type_pk PRIMARY KEY (id),
	CONSTRAINT payment_type_unique UNIQUE (name)
);