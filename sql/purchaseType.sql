-- financial.purchase_type definition

-- Drop table

-- DROP TABLE financial.purchase_type;

CREATE TABLE financial.purchase_type (
	id uuid NOT NULL,
	"name" varchar(50) NOT NULL,
	create_at date DEFAULT now() NOT NULL,
	update_at date NULL,
	CONSTRAINT purchase_type_pk PRIMARY KEY (id),
	CONSTRAINT purchase_type_unique UNIQUE (name)
);