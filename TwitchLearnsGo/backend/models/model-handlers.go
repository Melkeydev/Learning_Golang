package models

import (
	"log"
  "time"
  "context"
)

func (m *DBModel) GetUser(id int) (*User, error) {
  ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select id, username, password from users where id = $1`

	row := m.DB.QueryRowContext(ctx, query, id)

	var user User 

	err := row.Scan(
		&user.ID,
		&user.Username,
    &user.Password,
	)

	if err != nil {
		return nil, err
	}

  return &user, nil
}

func (m *DBModel) RegisterUser(user User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `insert into users(username, password) VALUES ($1, $2)`

	_, err := m.DB.ExecContext(ctx, query,
		user.Username,
		user.Password,
	)

	if err != nil {
		log.Print(err)
		return err
	}

	return nil
}


func (m *DBModel) InsertJob(job Job) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `insert into jobs(title, company, link, description, total_compensation) VALUES ($1, $2, $3, $4, $5)`

	_, err := m.DB.ExecContext(ctx, query,
		job.Title,
		job.Company,
		job.Link,
		job.Description,
		job.TotalCompensation,
	)

	if err != nil {
		log.Print(err)
		return err
	}

	return nil
}

func (m *DBModel) GetJob(id int) (*Job, error) {
  ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select id, title, company, link, description, total_compensation from jobs where id = $1`

	row := m.DB.QueryRowContext(ctx, query, id)

	var job Job 

	err := row.Scan(
		&job.ID,
		&job.Title,
    &job.Company,
    &job.Link,
    &job.Description,
    &job.TotalCompensation,
	)

	if err != nil {
		return nil, err
	}

  return &job, nil
}
























