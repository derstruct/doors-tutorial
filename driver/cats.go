package driver

import (
	"database/sql"
	"log"
	"strings"
)

type Cat struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Desc string `json:"desc"`
}

func newCatsDriver(db *sql.DB) *CatsDriver {
	initQuery := `
		CREATE TABLE IF NOT EXISTS cats (
			id TEXT PRIMARY KEY,
			name TEXT NOT NULL,
			desc TEXT
		);
	`
	if _, err := db.Exec(initQuery); err != nil {
		log.Fatal("Failed to create cats table:", err)
	}
	return &CatsDriver{
		db: db,
	}
}

type CatsDriver struct {
	db *sql.DB
}

func (c *CatsDriver) populate() {
	var count int
	err := c.db.QueryRow("SELECT COUNT(*) FROM cats").Scan(&count)
	if err != nil {
		panic(err)
	}
	if count != 0 {
		return
	}
	sampleCats := []Cat{
		{Id: "electronics", Name: "Electronics", Desc: "Phones, laptops, and gadgets"},
		{Id: "books", Name: "Books", Desc: "Fiction, non-fiction, and educational books"},
		{Id: "clothing", Name: "Clothing", Desc: "Shirts, pants, shoes, and accessories"},
		{Id: "home", Name: "Home & Garden", Desc: "Furniture, decor, and gardening supplies"},
		{Id: "sports", Name: "Sports & Outdoors", Desc: "Equipment for fitness and outdoor activities"},
		{Id: "food", Name: "Food & Beverages", Desc: "Snacks, drinks, and cooking ingredients"},
	}
	for _, cat := range sampleCats {
		c.Create(cat)
	}
}

func (c *CatsDriver) Get(id string) (Cat, bool) {
	var cat Cat
	err := c.db.QueryRow("SELECT id, name, desc FROM cats WHERE id = ?", id).
		Scan(&cat.Id, &cat.Name, &cat.Desc)

	if err != nil {
		if err == sql.ErrNoRows {
			return Cat{}, false
		}
		panic(err)
	}

	return cat, true
}

func (c *CatsDriver) List() []Cat {
	rows, err := c.db.Query("SELECT id, name, desc FROM cats")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var cats []Cat
	for rows.Next() {
		var cat Cat
		err := rows.Scan(&cat.Id, &cat.Name, &cat.Desc)
		if err != nil {
			panic(err)
		}
		cats = append(cats, cat)
	}
	if err = rows.Err(); err != nil {
		panic(err)
	}
	return cats
}

func (c *CatsDriver) Create(cat Cat) bool {
	_, err := c.db.Exec("INSERT INTO cats(id, name, desc) VALUES(?, ?, ?)", cat.Id, cat.Name, cat.Desc)
	if err != nil {
		if strings.Contains(err.Error(), "UNIQUE constraint failed") {
			return false // Id already exists
		}
		panic(err)
	}
	return true
}

func (c *CatsDriver) Remove(id string) bool {
	Items.removeItemsByCat(id)
	result, err := c.db.Exec("DELETE FROM cats WHERE id = ?", id)
	if err != nil {
		panic(err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}
	return rowsAffected > 0
}

func (c *CatsDriver) Update(cat Cat) bool {
	result, err := c.db.Exec("UPDATE cats SET name = ?, desc = ? WHERE id = ?",
		cat.Name, cat.Desc, cat.Id)
	if err != nil {
		panic(err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}
	return rowsAffected > 0
}
