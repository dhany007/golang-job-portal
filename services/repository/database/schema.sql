CREATE DATABASE job-portal;

CREATE TABLE users (
	id uuid NOT NULL PRIMARY KEY,
	email varchar(255) NOT NULL,
	hash_password varchar(255) NOT NULL,
	"role" int NOT NULL DEFAULT 1::int,
	is_active int NOT NULL DEFAULT 1::int,
	created_at timestamp NOT NULL DEFAULT now(),
	created_by varchar(255) NOT NULL DEFAULT 'SYSTEM'::CHARACTER VARYING,
	modified_at timestamp NOT NULL DEFAULT now(),
	modified_by varchar(255) NOT NULL DEFAULT 'SYSTEM'::CHARACTER VARYING,
	deleted_at timestamp DEFAULT NULL,
	deleted_by varchar(255) DEFAULT NULL
);

CREATE TABLE companies (
	id uuid NOT NULL PRIMARY KEY,
	email varchar(255) NOT NULL,
	name varchar(255) NOT NULL,
	description TEXT DEFAULT NULL,
	address TEXT DEFAULT NULL,
	website varchar(255) DEFAULT NULL,
	phone_number varchar(255) DEFAULT NULL,
	telp_number varchar(255) DEFAULT NULL,
	profil_picture_url varchar(255) DEFAULT NULL,
	dresscode_code int DEFAULT 1,
	size_code int DEFAULT 1,
	created_at timestamp NOT NULL DEFAULT now(),
	created_by varchar(255) NOT NULL DEFAULT 'SYSTEM'::CHARACTER VARYING,
	modified_at timestamp NOT NULL DEFAULT now(),
	modified_by varchar(255) NOT NULL DEFAULT 'SYSTEM'::CHARACTER VARYING,
	deleted_at timestamp DEFAULT NULL,
	deleted_by varchar(255) DEFAULT NULL
);

ALTER TABLE companies
ADD CONSTRAINT unique_email UNIQUE (email);

CREATE TABLE company_dresscode_codes (
	id serial PRIMARY KEY,
	value varchar(255) DEFAULT NULL,
	created_at timestamp NOT NULL DEFAULT now(),
	created_by varchar(255) NOT NULL DEFAULT 'SYSTEM'::CHARACTER VARYING,
	modified_at timestamp NOT NULL DEFAULT now(),
	modified_by varchar(255) NOT NULL DEFAULT 'SYSTEM'::CHARACTER VARYING,
	deleted_at timestamp DEFAULT NULL,
	deleted_by varchar(255) DEFAULT NULL
);

INSERT INTO
	company_dresscode_codes (value)
VALUES
	('Casual'),
	('Business casual'),
	('Business professional'),
	('Business formal');

CREATE TABLE company_benefits_codes (
	id serial PRIMARY KEY,
	value varchar(255) DEFAULT NULL,
	created_at timestamp NOT NULL DEFAULT now(),
	created_by varchar(255) NOT NULL DEFAULT 'SYSTEM'::CHARACTER VARYING,
	modified_at timestamp NOT NULL DEFAULT now(),
	modified_by varchar(255) NOT NULL DEFAULT 'SYSTEM'::CHARACTER VARYING,
	deleted_at timestamp DEFAULT NULL,
	deleted_by varchar(255) DEFAULT NULL
);


INSERT INTO
	company_benefits_codes (value)
VALUES
	('Benefits that are required by law'),
	('Medical insurance'),
	('Life insurance'),
	('Retirement plans'),
	('Disability insurance'),
	('Fringe benefits');


CREATE TABLE company_size_codes (
	id serial PRIMARY KEY,
	value varchar(255) DEFAULT NULL,
	created_at timestamp NOT NULL DEFAULT now(),
	created_by varchar(255) NOT NULL DEFAULT 'SYSTEM'::CHARACTER VARYING,
	modified_at timestamp NOT NULL DEFAULT now(),
	modified_by varchar(255) NOT NULL DEFAULT 'SYSTEM'::CHARACTER VARYING,
	deleted_at timestamp DEFAULT NULL,
	deleted_by varchar(255) DEFAULT NULL
);


INSERT INTO
	company_size_codes (value)
VALUES
	('Self-employed'),
	('1-10 employees'),
	('11-50 employees'),
	('51-200 employees'),
	('201-500 employees'),
	('501-1.000 employees'),
	('1.001-5.000 employees'),
	('5.001-10.000 employees'),
	('1.0001+ employees');


CREATE TABLE company_benefits (
	id serial PRIMARY KEY,
	company_id varchar(255) NOT NULL,
	benefit_id int NOT NULL,
	created_at timestamp NOT NULL DEFAULT now(),
	created_by varchar(255) NOT NULL DEFAULT 'SYSTEM'::CHARACTER VARYING,
	modified_at timestamp NOT NULL DEFAULT now(),
	modified_by varchar(255) NOT NULL DEFAULT 'SYSTEM'::CHARACTER VARYING,
	deleted_at timestamp DEFAULT NULL,
	deleted_by varchar(255) DEFAULT NULL
);

CREATE TABLE candidates (
	id uuid NOT NULL PRIMARY KEY,
	email varchar(255) NOT NULL,
	first_name varchar(255) NOT NULL,
	last_name varchar(255) DEFAULT NULL,
	phone_number varchar(255) DEFAULT NULL,
	telp_number varchar(255) DEFAULT NULL,
	address TEXT DEFAULT NULL,
	profil_picture_url varchar(255) DEFAULT NULL,
	created_at timestamp NOT NULL DEFAULT now(),
	created_by varchar(255) NOT NULL DEFAULT 'SYSTEM'::CHARACTER VARYING,
	modified_at timestamp NOT NULL DEFAULT now(),
	modified_by varchar(255) NOT NULL DEFAULT 'SYSTEM'::CHARACTER VARYING,
	deleted_at timestamp DEFAULT NULL,
	deleted_by varchar(255) DEFAULT NULL
);

CREATE INDEX idx_email_users ON users(email);
CREATE INDEX idx_email_company ON companies(email);
CREATE INDEX idx_email_candidate ON candidates(email);

CREATE TABLE company_reviews (
	id serial PRIMARY KEY,
	company_id uuid NOT NULL,
	candidate_id uuid NOT NULL,
	rating int DEFAULT NULL,
	review varchar(255) DEFAULT NULL,
	created_at timestamp NOT NULL DEFAULT now(),
	created_by varchar(255) NOT NULL DEFAULT 'SYSTEM'::CHARACTER VARYING,
	modified_at timestamp NOT NULL DEFAULT now(),
	modified_by varchar(255) NOT NULL DEFAULT 'SYSTEM'::CHARACTER VARYING,
	deleted_at timestamp DEFAULT NULL,
	deleted_by varchar(255) DEFAULT NULL
);

CREATE INDEX idx_company_review ON company_reviews(company_id);

ALTER TABLE company_reviews
ADD CONSTRAINT uq_candidate_company UNIQUE(company_id, candidate_id);

