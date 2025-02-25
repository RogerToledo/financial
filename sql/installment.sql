-- financial.installment definition

-- Drop table

-- DROP TABLE financial.installment;

CREATE TABLE financial.installment (
	id uuid NOT NULL,
	description varchar(50) NOT NULL,
	"month" numeric(10, 2) NOT NULL,
	value numeric(10, 2) NOT NULL,
	"number" int4 NOT NULL,
	paid bool NULL,
	purchase_id uuid NOT NULL,
	CONSTRAINT installment_pk PRIMARY KEY (id)
);


-- financial.installment foreign keys

ALTER TABLE financial.installment ADD CONSTRAINT installment_purchase_fk FOREIGN KEY (id) REFERENCES financial.purchase(id) ON DELETE CASCADE ON UPDATE CASCADE;