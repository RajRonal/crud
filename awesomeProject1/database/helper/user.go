package helper

import (
	"awesomeProject1/database"
	"awesomeProject1/models"
	"database/sql"
)

func CreateUser(name, email string) (string, error) {

	SQL := `INSERT INTO users(name, email) VALUES ($1, $2) RETURNING id;`
	var userID string
	err := database.Db.Get(&userID, SQL, name, email)
	if err != nil {
		return "", err
	}
	return userID, nil
}
func GetUser(userID string) (*models.User, error) {

	SQL := `SELECT id, name, email, created_at, archived_at FROM users WHERE archived_at IS NULL AND id = $1`
	var user models.User
	err := database.Db.Get(&user, SQL, userID)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &user, nil
}

func UpdateUser(name, email, id string) (string, error) {

	SQL := `UPDATE users SET name=$1, email=$2 WHERE id=$3 RETURNING id;`
	var userID string
	err := database.Db.Get(&userID, SQL, name, email, id)
	if err != nil {
		return "", err
	}
	return userID, nil
}

func DeleteUser(uid string) sql.Result {

	SQL := `DELETE FROM users WHERE id=$1 RETURNING id;`

	err, _ := database.Db.Exec(SQL, uid)

	return err
}
