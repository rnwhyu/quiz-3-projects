package models

import (
	"project-quiz3/database"
	"time"
)

type Category struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Created_at time.Time `json:"created_date"`
	Updated_at time.Time `json:"updated_date"`
}
type Categories []Category

func (c *Categories) FindAllCategory() error {
	sql := `SELECT * FROM categories ORDER BY id`
	rows, err := db.DB.Query(sql)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var category = Category{}
		err = rows.Scan(&category.ID, &category.Name, &category.Created_at, &category.Updated_at)
		if err != nil {
			return err
		}
		*c = append(*c, category)
	}
	return nil
}
func (c *Category) CreateCategory() error {
	sql := `INSERT INTO categories (name)
			VALUES ($1)
			Returning *`
	err := db.DB.QueryRow(sql, c.Name).
		Scan(&c.ID, &c.Name, &c.Created_at, &c.Updated_at)
	if err != nil {
		return err
	}
	return nil
}
func (c *Category) FindCategory() error {
	sql := `SELECT * FROM categories
			WHERE id = $1;`
	err := db.DB.QueryRow(sql, c.ID).
		Scan(&c.ID, &c.Name, &c.Created_at, &c.Updated_at)
	if err != nil {
		return err
	}
	return nil
}
func (c *Category) UpdateCategory() error {
	sql := `
		UPDATE categories
		SET name = $2, updated_at = $3
		WHERE id = $1
		Returning *`
	err := db.DB.QueryRow(sql, c.ID, c.Name, time.Now()).
		Scan(&c.ID, &c.Name, &c.Created_at, &c.Updated_at)
	if err != nil {
		return err
	}
	return nil
}

func (c *Category) DeleteCategory() error {
	sql := `DELETE from categories
					WHERE id = $1`

	_, err := db.DB.Exec(sql, c.ID)
	if err != nil {
		return err
	}

	return nil
}
