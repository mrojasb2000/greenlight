package data

import "database/sql"

// Define a BookModel struct type which wraps a sql.DB connection pool.
type BookModel struct {
	DB *sql.DB
}

func (m BookModel) Insert(book *Book) error {
	return nil
}

func (m BookModel) Get(id int64) (*Book, error) {
	return nil, nil
}

func (m BookModel) Update(book *Book) error {
	return nil
}

func (m BookModel) Delete(id int64) error {
	return nil
}

type Book struct {
}
