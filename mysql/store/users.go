package store

import (
	"context"
	"database/sql"
	"go-db-labs/mysql/db"
	"go-db-labs/mysql/models"
	"go-db-labs/mysql/query"
	"log"
)

type UsersStore interface{
	Create(ctx context.Context, user *models.User) error
	GetAll(ctx context.Context) ([]*models.User, error)
	Get(ctx context.Context, id string) (*models.User, error)
	Update(ctx context.Context, user *models.User) error
	Delete(ctx context.Context, id string) error
}

type usersStore struct {
	conn *sql.Conn
}

func NewUsersStore(conn *sql.Conn) UsersStore {
	return &usersStore{conn}
}

func (s *usersStore) Create(ctx context.Context, user *models.User) error {

	q := query.NewInsert("users", []string{
		"id",
		"name",
		"username",
		"email",
		"password",
		"status",
		"created_at",
		"updated_at",
	})

	stmt, err := s.conn.PrepareContext(ctx, q)
	if err != nil {
		return err
	}

	defer db.Close(stmt)

	result, err := stmt.Exec(
		user.ID,
		user.Name,
		user.Username,
		user.Email,
		user.Password,
		user.Status,
		user.CreatedAt,
		user.UpdatedAt,
	)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err == nil {
		_ = rows
		//log.Printf("%d Rows affected\n", rows)
	}

	return err
}

func (s *usersStore) GetAll(ctx context.Context) ([]*models.User, error) {

	q := "select * from users"

	rows, err := s.conn.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}
	defer db.Close(rows)

	var users []*models.User

	for rows.Next() {
		var user models.User

		err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Username,
			&user.Email,
			&user.Password,
			&user.Status,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			return users, err
		}

		users = append(users, &user)
	}

	log.Printf("Get all: %d\n", len(users))

	return users, nil
}

func (s *usersStore) Get(ctx context.Context, id string) (*models.User, error) {

	q := "select * from users where id = ?"

	row := s.conn.QueryRowContext(ctx, q, id)
	if row.Err() != nil {
		return nil, row.Err()
	}

	var user models.User

	err := row.Scan(
		&user.ID,
		&user.Name,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.Status,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	return &user, err
}

func (s *usersStore) Update(ctx context.Context, user *models.User) error {

	q := query.NewUpdate("users", []string{
		"name",
		"username",
		"email",
		"password",
		"status",
		"created_at",
		"updated_at",
	}, "id")

	stmt, err := s.conn.PrepareContext(ctx, q)
	if err != nil {
		return err
	}

	defer db.Close(stmt)

	result, err := stmt.Exec(
		user.Name,
		user.Username,
		user.Email,
		user.Password,
		user.Status,
		user.CreatedAt,
		user.UpdatedAt,
		user.ID,
	)

	rows, err := result.RowsAffected()
	if err == nil {
		log.Printf("%d Rows affected\n", rows)
	}

	return err
}

func (s *usersStore) Delete(ctx context.Context, id string) error {
	q := "delete from users where = ?"
	return s.conn.QueryRowContext(ctx, q, id).Err()
}