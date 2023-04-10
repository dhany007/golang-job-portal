CREATE TABLE public.users (
	id uuid NOT NULL,
	email varchar(255) NOT NULL,
	hash_password varchar(255) NOT NULL,
	"role" int4 NOT NULL DEFAULT 1,
	is_active int4 NOT NULL DEFAULT 1,
	created_at timestamp NOT NULL DEFAULT now(),
	created_by varchar(255) NOT NULL DEFAULT 'SYSTEM'::character varying,
	modified_at timestamp NOT NULL DEFAULT now(),
	modified_by varchar(255) NOT NULL DEFAULT 'SYSTEM'::character varying,
	deleted_at timestamp NULL,
	deleted_by varchar(255) NULL DEFAULT NULL::character varying,
	CONSTRAINT users_pkey PRIMARY KEY (id)
);
CREATE INDEX idx_email_users ON public.users USING btree (email);

CREATE TABLE public.companies (
	id uuid NOT NULL,
	email varchar(255) NOT NULL,
	"name" varchar(255) NOT NULL,
	description text NULL,
	address text NULL,
	website varchar(255) NULL DEFAULT NULL::character varying,
	phone_number varchar(255) NULL DEFAULT NULL::character varying,
	telp_number varchar(255) NULL DEFAULT NULL::character varying,
	profil_picture_url varchar(255) NULL DEFAULT NULL::character varying,
	dresscode_code int4 NULL DEFAULT 1,
	size_code int4 NULL DEFAULT 1,
	created_at timestamp NOT NULL DEFAULT now(),
	created_by varchar(255) NOT NULL DEFAULT 'SYSTEM'::character varying,
	modified_at timestamp NOT NULL DEFAULT now(),
	modified_by varchar(255) NOT NULL DEFAULT 'SYSTEM'::character varying,
	deleted_at timestamp NULL,
	deleted_by varchar(255) NULL DEFAULT NULL::character varying,
	CONSTRAINT companies_pkey PRIMARY KEY (id),
	CONSTRAINT unique_email UNIQUE (email)
);
CREATE INDEX idx_email_company ON public.companies USING btree (email);

CREATE TABLE public.candidates (
	id uuid NOT NULL,
	email varchar(255) NOT NULL,
	first_name varchar(255) NOT NULL,
	last_name varchar(255) NULL DEFAULT NULL::character varying,
	phone_number varchar(255) NULL DEFAULT NULL::character varying,
	telp_number varchar(255) NULL DEFAULT NULL::character varying,
	address text NULL,
	profil_picture_url varchar(255) NULL DEFAULT NULL::character varying,
	created_at timestamp NOT NULL DEFAULT now(),
	created_by varchar(255) NOT NULL DEFAULT 'SYSTEM'::character varying,
	modified_at timestamp NOT NULL DEFAULT now(),
	modified_by varchar(255) NOT NULL DEFAULT 'SYSTEM'::character varying,
	deleted_at timestamp NULL,
	deleted_by varchar(255) NULL DEFAULT NULL::character varying,
	CONSTRAINT candidates_pkey PRIMARY KEY (id)
);
CREATE INDEX idx_email_candidate ON public.candidates USING btree (email);

CREATE TABLE public.candidate_experiences (
	id serial4 NOT NULL,
	candidate_id varchar(255) NOT NULL,
	company_name text NOT NULL,
	title text NOT NULL,
	description text NULL,
	date_start timestamp NULL,
	date_end timestamp NULL,
	created_at timestamp NOT NULL DEFAULT now(),
	created_by varchar(255) NOT NULL DEFAULT 'SYSTEM'::character varying,
	modified_at timestamp NOT NULL DEFAULT now(),
	modified_by varchar(255) NOT NULL DEFAULT 'SYSTEM'::character varying,
	deleted_at timestamp NULL,
	deleted_by varchar(255) NULL DEFAULT NULL::character varying,
	CONSTRAINT candidate_experiences_pkey PRIMARY KEY (id)
);

CREATE TABLE public.company_benefits (
	id serial4 NOT NULL,
	company_id varchar(255) NOT NULL,
	benefit_id int4 NOT NULL,
	created_at timestamp NOT NULL DEFAULT now(),
	created_by varchar(255) NOT NULL DEFAULT 'SYSTEM'::character varying,
	modified_at timestamp NOT NULL DEFAULT now(),
	modified_by varchar(255) NOT NULL DEFAULT 'SYSTEM'::character varying,
	deleted_at timestamp NULL,
	deleted_by varchar(255) NULL DEFAULT NULL::character varying,
	CONSTRAINT company_benefits_pkey PRIMARY KEY (id)
);

CREATE TABLE public.company_benefits_codes (
	id serial4 NOT NULL,
	value varchar(255) NULL DEFAULT NULL::character varying,
	created_at timestamp NOT NULL DEFAULT now(),
	created_by varchar(255) NOT NULL DEFAULT 'SYSTEM'::character varying,
	modified_at timestamp NOT NULL DEFAULT now(),
	modified_by varchar(255) NOT NULL DEFAULT 'SYSTEM'::character varying,
	deleted_at timestamp NULL,
	deleted_by varchar(255) NULL DEFAULT NULL::character varying,
	CONSTRAINT company_benefits_codes_pkey PRIMARY KEY (id)
);

CREATE TABLE public.company_dresscode_codes (
	id serial4 NOT NULL,
	value varchar(255) NULL DEFAULT NULL::character varying,
	created_at timestamp NOT NULL DEFAULT now(),
	created_by varchar(255) NOT NULL DEFAULT 'SYSTEM'::character varying,
	modified_at timestamp NOT NULL DEFAULT now(),
	modified_by varchar(255) NOT NULL DEFAULT 'SYSTEM'::character varying,
	deleted_at timestamp NULL,
	deleted_by varchar(255) NULL DEFAULT NULL::character varying,
	CONSTRAINT company_dresscode_codes_pkey PRIMARY KEY (id)
);

CREATE TABLE public.company_reviews (
	id serial4 NOT NULL,
	company_id uuid NOT NULL,
	candidate_id uuid NOT NULL,
	rating int4 NULL,
	review varchar(255) NULL DEFAULT NULL::character varying,
	created_at timestamp NOT NULL DEFAULT now(),
	created_by varchar(255) NOT NULL DEFAULT 'SYSTEM'::character varying,
	modified_at timestamp NOT NULL DEFAULT now(),
	modified_by varchar(255) NOT NULL DEFAULT 'SYSTEM'::character varying,
	deleted_at timestamp NULL,
	deleted_by varchar(255) NULL DEFAULT NULL::character varying,
	CONSTRAINT company_reviews_pkey PRIMARY KEY (id),
	CONSTRAINT uq_candidate_company UNIQUE (company_id, candidate_id)
);
CREATE INDEX idx_company_review ON public.company_reviews USING btree (company_id);

CREATE TABLE public.company_size_codes (
	id serial4 NOT NULL,
	value varchar(255) NULL DEFAULT NULL::character varying,
	created_at timestamp NOT NULL DEFAULT now(),
	created_by varchar(255) NOT NULL DEFAULT 'SYSTEM'::character varying,
	modified_at timestamp NOT NULL DEFAULT now(),
	modified_by varchar(255) NOT NULL DEFAULT 'SYSTEM'::character varying,
	deleted_at timestamp NULL,
	deleted_by varchar(255) NULL DEFAULT NULL::character varying,
	CONSTRAINT company_size_codes_pkey PRIMARY KEY (id)
);

CREATE TABLE public.job_specializations (
	id serial4 NOT NULL,
	value varchar(255) NOT NULL,
	created_at timestamp NOT NULL DEFAULT now(),
	created_by varchar(255) NOT NULL DEFAULT 'SYSTEM'::character varying,
	modified_at timestamp NOT NULL DEFAULT now(),
	modified_by varchar(255) NOT NULL DEFAULT 'SYSTEM'::character varying,
	deleted_at timestamp NULL,
	deleted_by varchar(255) NULL DEFAULT NULL::character varying,
	CONSTRAINT job_specializations_pkey PRIMARY KEY (id)
);