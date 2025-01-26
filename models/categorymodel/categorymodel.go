package categorymodel

import (
	"produk_go/config"
	"produk_go/entities"
)

func GetAll() []entities.Category {
	rows, err := config.DB.Query(`SELECT * FROM categories`)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	//variabel untuk menampung seluruh data categori
	var categories []entities.Category

	for rows.Next() {
		var category entities.Category
		if err := rows.Scan(&category.Id, &category.Name, &category.CreatedAt, &category.UpdatedAt); err != err{
			panic(err)
		}

		categories = append(categories, category)
	}

	return categories
}

func Create(category entities.Category)  {
	// config.DB.Exec(``)
}