package repositories

import (
	"database/sql"

	"GoCRUDApplicationMySQL/app/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) GetAllUsers() ([]models.User, error) {
	query := "SELECT * FROM users"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []models.User{}
	for rows.Next() {
		user := models.User{}
		err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserRepository) CreateUser(user *models.User) error {
	query := "INSERT INTO users (first_name, last_name, email) VALUES (?, ?, ?)"
	result, err := r.db.Exec(query, user.FirstName, user.LastName, user.Email)
	if err != nil {
		return err
	}

	userID, err := result.LastInsertId()
	if err != nil {
		return err
	}

	user.ID = int(userID)
	return nil
}

// Implement other database operations like GetUser, UpdateUser, and DeleteUser
