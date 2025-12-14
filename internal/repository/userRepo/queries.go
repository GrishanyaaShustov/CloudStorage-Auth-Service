package userRepo

const (
	// Create user and return generated UUID
	queryCreateUser = `
		INSERT INTO users (email, password_hash)
		VALUES ($1, $2)
		RETURNING id
	`

	// Get password hash by email (for login)
	queryGetPasswordHashByEmail = `
		SELECT password_hash
		FROM users
		WHERE email = $1
	`

	// Get full user info by ID
	queryGetUserByID = `
		SELECT id, email, registered_at
		FROM users
		WHERE id = $1
	`
)
