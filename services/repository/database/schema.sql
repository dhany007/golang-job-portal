CREATE SCHEMA job-portal;

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
