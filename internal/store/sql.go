package store

const (
	createSchemaSQL = `
		CREATE TABLE IF NOT EXISTS users (
			city_lat decimal(8,6),
			city_lng decimal(9,6),
			email text,
			id uuid,
			name text,
			password text,
			phone text,
			photo_url text,
			username text
		);
	`
)

const (
	createUserSQL = `
		INSERT into users (
			city_lat,
			city_lng,
			email,
			name,
			password,
			phone,
			photo_url,
			username
		) VALUES (
			:city_lat,
			:city_lng,
			:email,
			:name,
			:password,
			:phone,
			:photo_url,
			:username
		)
	`
)

const (
	getUserByIdSQL = `
		SELECT * FROM user WHERE id=:id
	`
)