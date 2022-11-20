package store

const (
	createUserSQL = `
		INSERT into user (
			username,
			name,
			city_lat
		) VALUES (
			:username,
			:name,
			:citylat
		)
	`
)

const (
	getUserByIdSQL = `
		SELECT * FROM user WHERE id=:id
	`
)