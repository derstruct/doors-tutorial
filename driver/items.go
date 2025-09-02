package driver

import (
	"database/sql"
	"log"
)

type Item struct {
	Id     int    `json:"id"`
	Cat    string `json:"cat"`
	Name   string `json:"name"`
	Desc   string `json:"desc"`
	Rating int    `json:"rating"`
}

func newItemsDriver(db *sql.DB) *ItemsDriver {
	initQuery := `
		CREATE TABLE IF NOT EXISTS items (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			cat TEXT NOT NULL,
			name TEXT NOT NULL,
			desc TEXT,
			rating INTEGER
		);
	`
	if _, err := db.Exec(initQuery); err != nil {
		log.Fatal("Failed to create items table:", err) // Fixed
	}
	return &ItemsDriver{
		db: db,
	}
}

type ItemsDriver struct {
	db *sql.DB
}

const onPage = 6

func (d *ItemsDriver) CountPages(catId string) int {
	var total int
	err := d.db.QueryRow("SELECT COUNT(*) FROM items WHERE cat = ?", catId).Scan(&total)
	if err != nil {
		panic(err)
	}
	pages := total / onPage
	if total%onPage > 0 {
		pages += 1
	}
	return pages
}

func (d *ItemsDriver) List(catId string, page int) []Item {
	offset := page * onPage
	rows, err := d.db.Query("SELECT id, cat, name, desc, rating FROM items WHERE cat = ? LIMIT ? OFFSET ?",
		catId, onPage, offset)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var items []Item
	for rows.Next() {
		var item Item
		err := rows.Scan(&item.Id, &item.Cat, &item.Name, &item.Desc, &item.Rating)
		if err != nil {
			panic(err)
		}
		items = append(items, item)
	}
	if err = rows.Err(); err != nil {
		panic(err)
	}
	return items
}

func (d *ItemsDriver) Get(id int) (Item, bool) {
	var item Item
	err := d.db.QueryRow("SELECT id, cat, name, desc, rating FROM items WHERE id = ?", id).
		Scan(&item.Id, &item.Cat, &item.Name, &item.Desc, &item.Rating)
	if err != nil {
		if err == sql.ErrNoRows {
			return Item{}, false
		}
		panic(err)
	}
	return item, true
}

func (d *ItemsDriver) Create(item Item) {
	_, err := d.db.Exec("INSERT INTO items(cat, name, desc, rating) VALUES(?, ?, ?, ?)",
		item.Cat, item.Name, item.Desc, item.Rating)
	if err != nil {
		panic(err)
	}
}

func (d *ItemsDriver) Remove(id int) bool {
	result, err := d.db.Exec("DELETE FROM items WHERE id = ?", id)
	if err != nil {
		panic(err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}
	return rowsAffected > 0
}

func (d *ItemsDriver) Update(item Item) bool {
	result, err := d.db.Exec("UPDATE items SET cat = ?, name = ?, desc = ?, rating = ? WHERE id = ?",
		item.Cat, item.Name, item.Desc, item.Rating, item.Id)
	if err != nil {
		panic(err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}
	return rowsAffected > 0
}

func (d *ItemsDriver) removeItemsByCat(catId string) {
	_, err := d.db.Exec("DELETE FROM items WHERE cat = ?", catId)
	if err != nil {
		panic(err)
	}
}
