package store

const (
	createSchemaSQL = `
		CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
		CREATE TABLE IF NOT EXISTS users (
			city_lat DECIMAL(8,6),
			city_lng DECIMAL(9,6),
			email TEXT,
			id uuid NOT NULL DEFAULT uuid_generate_v4() PRIMARY KEY,
			name TEXT,
			password TEXT,
			phone TEXT,
			photo_url TEXT,
			username TEXT
		);
	`
)

const (
	createUserSQL = `
		INSERT into users (
			email,
			name,
			password,
			phone,
			username
		) VALUES (
			:email,
			:name,
			:password,
			:phone,
			:username
		)
	`
)

const (
	getUserByIdSQL = `
		SELECT * FROM user WHERE id=:id
	`
)