-- financial.credit_card definition

-- Drop table

-- DROP TABLE financial.credit_card;

CREATE TABLE financial.credit_card (
	id uuid NOT NULL,
	"owner" varchar(50) NOT NULL,
	CONSTRAINT credit_card_pk PRIMARY KEY (id),
	CONSTRAINT credit_card_unique UNIQUE (owner)
);