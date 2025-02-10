-- financial.purchase definition

-- Drop table

-- DROP TABLE financial.purchase;

CREATE TABLE financial.purchase (
	id uuid NOT NULL,
	description varchar(150) NULL,
	amount numeric(10, 2) NOT NULL,
	"date" date NOT NULL,
	installment_number int4 DEFAULT 0 NULL,
	installment numeric(10, 2) DEFAULT 0 NULL,
	place varchar(60) NULL,
	id_payment_type uuid NOT NULL,
	id_purchase_type uuid NOT NULL,
	id_credit_card uuid NOT NULL,
	id_person uuid NOT NULL,
	CONSTRAINT purchase_pk PRIMARY KEY (id)
);


-- financial.purchase foreign keys

ALTER TABLE financial.purchase ADD CONSTRAINT purchase_credit_card_fk FOREIGN KEY (id_credit_card) REFERENCES financial.credit_card(id) ON DELETE CASCADE ON UPDATE CASCADE;
ALTER TABLE financial.purchase ADD CONSTRAINT purchase_payment_type_fk FOREIGN KEY (id_payment_type) REFERENCES financial.payment_type(id) ON DELETE CASCADE ON UPDATE CASCADE;
ALTER TABLE financial.purchase ADD CONSTRAINT purchase_person_fk FOREIGN KEY (id_person) REFERENCES financial.person(id) ON DELETE CASCADE ON UPDATE CASCADE;
ALTER TABLE financial.purchase ADD CONSTRAINT purchase_purchase_type_fk FOREIGN KEY (id_purchase_type) REFERENCES financial.purchase_type(id) ON DELETE CASCADE ON UPDATE CASCADE;