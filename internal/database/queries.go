package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type Article struct {
	Id        int
	Otsikko   string
	Teksti    string
	Kuva      string
	Kategoria string
}

func GetArticles(db *sql.DB) ([]Article, error) {
	rows, err := db.Query("SELECT * FROM uutinen")
	if err != nil {
		return nil, err
	}
	articles := []Article{}

	for rows.Next() {
		article := Article{}
		if err := rows.Scan(&article.Id, &article.Otsikko, &article.Teksti, &article.Kuva, &article.Kategoria); err != nil {
			return nil, err
		}
		articles = append(articles, article)
	}

	return articles, nil
}
