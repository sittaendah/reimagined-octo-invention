package user

import (
	"database/sql"

	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct {
	DB *sql.DB
}

func (r *UserRepository) CreateUser(user User) (int, error) {
	hash, err := hashPassword(user.Password)
	if err != nil {
		return 0, err
	}

	var id int
	err = r.DB.QueryRow(`INSERT INTO users(username, role, password) VALUES($1, $2, $3) RETURNING id`, user.Username, user.Role, hash).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *UserRepository) GetUser(id int) (User, error) {
	var user User
	row := r.DB.QueryRow(`SELECT id, username, role FROM users WHERE id = $1`, id)
	err := row.Scan(&user.ID, &user.Username, &user.Role)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func (r *UserRepository) GetUserByUsername(username string) (User, error) {
	var user User
	row := r.DB.QueryRow(`SELECT id, username, password, role FROM users WHERE username = $1`, username)
	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Role)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func (r *UserRepository) GetAllUsers() ([]User, error) {
	rows, err := r.DB.Query(`SELECT id, username, role FROM users`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Username, &user.Role); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (r *UserRepository) UpdateUser(user User) error {
	_, err := r.DB.Exec(`UPDATE users SET username = $1, role = $2 WHERE id = $3`, user.Username, user.Role, user.ID)
	return err
}

func (r *UserRepository) DeleteUser(id int) error {
	_, err := r.DB.Exec(`DELETE FROM users WHERE id = $1`, id)
	return err
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
