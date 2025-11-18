package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type Article struct {
	Id          int
	Title       string
	Content     string
	Picture     string
	CategoryId  string
	Description string
}

type Category struct {
	Id   int
	Name string
}

func GetArticles(db *sql.DB) ([]Article, error) {
	rows, err := db.Query("SELECT * FROM articles")
	if err != nil {
		return nil, err
	}
	articles := []Article{}

	for rows.Next() {
		article := Article{}
		var category_id sql.NullString
		if err := rows.Scan(&article.Id, &article.Title, &article.Content, &article.Picture, &category_id, &article.Description); err != nil {
			return nil, err
		}
		if category_id.Valid {
			article.CategoryId = category_id.String
		}
		articles = append(articles, article)
	}

	return articles, nil
}

func GetCategories(db *sql.DB) ([]Category, error) {
	rows, err := db.Query("SELECT * FROM categories")
	if err != nil {
		return nil, err
	}
	categories := []Category{}

	for rows.Next() {
		category := Category{}
		if err := rows.Scan(&category.Id, &category.Name); err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	return categories, nil
}

func GetCategoryArticles(db *sql.DB, category string) ([]Article, error) {
	stmtOut, err := db.Prepare("SELECT article_id, article_title, article_content, article_picture, article_description FROM articles INNER JOIN categories ON articles.category_id = categories.category_id WHERE category_name = ?;")
	if err != nil {
		return nil, err
	}
	defer stmtOut.Close()

	rows, err := stmtOut.Query(category)
	if err != nil {
		return nil, err
	}
	articles := []Article{}

	for rows.Next() {
		article := Article{}
		if err := rows.Scan(&article.Id, &article.Title, &article.Content, &article.Picture, &article.Description); err != nil {
			return nil, err
		}
		articles = append(articles, article)
	}

	return articles, nil
}
