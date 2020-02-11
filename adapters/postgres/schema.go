package postgres

type schema struct {
	create string
	drop   string
}

var DefaultSchema = schema{
	create: `

	CREATE TABLE users (
		id serial PRIMARY KEY,
		username varchar (60) NOT NULL,
		email text UNIQUE NOT NULL,
		password varchar (60) NOT NULL,
		credit INTEGER NOT NULL DEFAULT 250
	);
	
	CREATE TABLE advisors (
		id integer REFERENCES users(id) ON DELETE CASCADE,
		first_name varchar (30) NOT NULL,
		last_name varchar (45) NOT NULL,
		description text NOT NULL,
		verified boolean NOT NULL DEFAULT TRUE,
		hourly_fee INTEGER NOT NULL DEFAULT 30,
		PRIMARY KEY (id)
	);
	
	CREATE TABLE roles (
		id serial PRIMARY KEY,
		name varchar (30) NOT NULL
	);
	
	CREATE TABLE user_roles (
		user_id integer NOT NULL REFERENCES users(id) ON DELETE CASCADE,
		role_id integer NOT NULL REFERENCES roles(id) ON DELETE CASCADE
	);
	
	CREATE TABLE appointments (
		id serial PRIMARY KEY,
		user_id integer REFERENCES users(id) ON DELETE CASCADE,
		advisor_id integer REFERENCES advisors(id) ON DELETE CASCADE,
		start_datetime timestamptz NOT NULL,
		end_datetime timestamptz NOT NULL,
		cancelled boolean NOT NULL DEFAULT FALSE
	);
	
	CREATE TABLE schedules (
		advisor_id integer REFERENCES advisors(id) ON DELETE CASCADE,
		day_of_week smallint NOT NULL,
		start_time time with time zone NOT NULL,
		end_time time with time zone NOT NULL,
		CHECK (day_of_week >= 0 AND day_of_week <= 6),
		UNIQUE (advisor_id, day_of_week, start_time, end_time)
	);
	
	CREATE TABLE ratings (
		user_id integer REFERENCES users(id),
		appointment_id integer REFERENCES appointments(id),
		score numeric (3,2) NOT NULL
	);
	`,

	drop: `
	DROP TABLE ratings;
	DROP TABLE schedules;
	DROP TABLE appointments;
	DROP TABLE user_roles;
	DROP TABLE roles;
	DROP TABLE advisors;
	DROP TABLE users;
	`,
}