package postgres

var insertUser = `
	INSERT INTO users (username, email, password, birthdate) VALUES ($1, $2, $3, $4) RETURNING id;
`

var authenticateViaEmail = `SELECT id,username, email, birthdate, password  FROM users WHERE email=$1;`
