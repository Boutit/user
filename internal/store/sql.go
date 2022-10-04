package store

const (
	createUserSQL = `
		INSERT into user (
			phone,
			email,
			birthday,
			gender,
			city,
			city_lat,
			city_lng,
			username,
			name
		) VALUES (
			:phone,
			:email,
			:birthday,
			:gender,
			:city,
			:city_lat,
			:city_lng,
			:username,
			:name
		)
	`
)