CREATE SCHEMA test_schema;

CREATE SEQUENCE test_schema.test_table_user_id_seq;

CREATE TABLE test_schema.test_table (
	user_id int4 NOT NULL DEFAULT nextval('test_schema.test_table_user_id_seq'::regclass),
	user_name text NOT NULL DEFAULT ''::text,
    created_at_utc timestamptz NULL,
	UNIQUE (user_id, user_name),
	PRIMARY KEY (user_id)
)
WITH (
	OIDS=FALSE
) ;

-- Sample orgs
INSERT INTO test_schema.test_table(user_name) VALUES ('Mary');
INSERT INTO test_schema.test_table(user_name) VALUES ('John');
INSERT INTO test_schema.test_table(user_name) VALUES ('Tim');
