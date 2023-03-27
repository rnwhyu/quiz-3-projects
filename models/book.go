package models

import (
	db "project-quiz3/database"
	"time"
)

type Book struct {
	ID           int       `json:"book_id"`
	Title        string    `json:"book_title"`
	Description  string    `json:"book_desc"`
	Image_url    string    `json:"img_url"`
	Release_year int       `json:"book_year"`
	Price        string    `json:"book_price"`
	Total_page   int       `json:"book_page"`
	Thickness    string    `json:"book_thick"`
	Created_at   time.Time `json:"created_date"`
	Updated_at   time.Time `json:"updated_date"`
	Category_id  int       `json:"category_id"`
}
type Books []Book

func (b *Books) FindAll() error {
	sql := `SELECT * FROM books ORDER BY book_id`

	rows, err := db.DB.Query(sql)

	if err != nil {
		return err
	}

	defer rows.Close()

	for rows.Next() {
		var book = Book{}
		err = rows.Scan(
			&book.ID,
			&book.Title,
			&book.Description,
			&book.Image_url,
			&book.Release_year,
			&book.Price,
			&book.Total_page,
			&book.Thickness,
			&book.Created_at,
			&book.Updated_at,
			&book.Category_id,
		)

		if err != nil {
			return err
		}

		*b = append(*b, book)
	}

	return nil
}

func (bs *Books) FindByCategoryID(categoryID int) error {
	sql := `SELECT * FROM books WHERE category_id = $1 ORDER BY book_id`
	rows, err := db.DB.Query(sql, categoryID)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var book = Book{}
		err = rows.Scan(
			&book.ID,
			&book.Title,
			&book.Description,
			&book.Image_url,
			&book.Release_year,
			&book.Price,
			&book.Total_page,
			&book.Thickness,
			&book.Created_at,
			&book.Updated_at,
			&book.Category_id,
		)
		if err != nil {
			return err
		}
		*bs = append(*bs, book)
	}
	return nil
}
func (b *Book) Create() error {
	sql := `
		INSERT INTO books (category_id,book_title,book_desc,img_url,book_year,book_price,book_page,book_thick)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8)
		Returning *`
	err := db.DB.
		QueryRow(
			sql,
			b.Category_id,
			b.Title,
			b.Description,
			b.Image_url,
			b.Release_year,
			b.Price,
			b.Total_page,
			b.Thickness).
		Scan(
			&b.ID,
			&b.Title,
			&b.Description,
			&b.Image_url,
			&b.Release_year,
			&b.Price,
			&b.Total_page,
			&b.Thickness,
			&b.Created_at,
			&b.Updated_at,
			&b.Category_id)
	if err != nil {
		return err
	}
	return nil
}

func (b *Book) Update() error {
	sql := `
		UPDATE books
		SET category_id = $2, book_title = $3, book_desc = $4, img_url = $5, book_yeaer = $6, book_price = $7, book_page = $8, book_thick = $9, updated_date = $10
		WHERE book_id = $1
		Returning *`
	err := db.DB.
		QueryRow(
			sql,
			b.ID,
			b.Category_id,
			b.Title,
			b.Description,
			b.Image_url,
			b.Release_year,
			b.Price,
			b.Total_page,
			b.Thickness,
			time.Now()).
		Scan(
			&b.ID,
			&b.Title,
			&b.Description,
			&b.Image_url,
			&b.Release_year,
			&b.Price,
			&b.Total_page,
			&b.Thickness,
			&b.Created_at,
			&b.Updated_at,
			&b.Category_id)
	if err != nil {
		return err
	}
	return nil
}

func (b *Book) Delete() error {
	sql := `DELETE from books
			WHERE book_id = $1`
	_, err := db.DB.Exec(sql, b.ID)
	if err != nil {
		return err
	}
	return nil
}
