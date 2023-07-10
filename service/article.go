package service

import (
	"ArticleSV/database"
	"ArticleSV/models"
	"ArticleSV/utils"
	"database/sql"
	"fmt"
	"github.com/gosimple/slug"
	"log"
	"time"
)

func AddArticle(article models.Article) (int64, error) {
	db := database.DB

	var datetime = time.Now()
	dt := datetime.Format(time.RFC3339)

	randomString, err := utils.GenerateRandomString(5)
	if err != nil {
		return 0, fmt.Errorf("Error generating random string:", err)
	}
	resultSlug := slug.Make(article.Title) + "-" + randomString
	result, err := db.Exec("INSERT INTO articles (id, title, slug, content, category, created_date, updated_date, status) VALUES (?, ?, ?, ?, ?, ?, ?, ?)", article.ID, article.Title, resultSlug, article.Content, article.Category, dt, dt, article.Status)
	if err != nil {
		return 0, fmt.Errorf("AddArticle: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("AddCycleTime: %v", err)
	}
	return id, nil
}

//func GetArticles(limit, offset int, status string) ([]models.Article, error) {
//	offset = (offset - 1) * limit
//	var Articles []models.Article
//	//if status == "" {
//	//} else {
//	//	string := []string{status}
//	//}
//	log.Println("Status.GETARTICLES => ", status)
//	db := database.DB // HOW TO GET CONNECTION/CALL DB FROM DATABASE/DATABASE.go ????
//	//query := fmt.Sprintf("SELECT * FROM articles LIMIT %s OFFSET %s", limit, offset)
//	rows, err := db.Query("SELECT * FROM articles LIMIT (?) OFFSET (?) WHERE status  (?)", limit, offset, status)
//	if err != nil {
//		return nil, fmt.Errorf("Failed to load articles: %v", err)
//	}
//	defer rows.Close()
//	for rows.Next() {
//		var Article models.Article
//		if err := rows.Scan(&Article.ID, &Article.Title, &Article.Slug, &Article.Content, &Article.Category, &Article.CreatedDate, &Article.UpdatedDate, &Article.Status); err != nil {
//			return nil, fmt.Errorf("failed scan article: %v", err)
//		}
//		Articles = append(Articles, Article)
//	}
//	if err := rows.Err(); err != nil {
//		return nil, fmt.Errorf("Gagal ketika load article: %v", err)
//	}
//	return Articles, nil
//}

func GetArticles(limit, offset int, status string) ([]models.Article, error) {
	offset = (offset - 1) * limit
	var Articles []models.Article

	db := database.DB // Menggunakan koneksi DB dari package database

	query := "SELECT * FROM articles"
	if status == "" {
		query += fmt.Sprintf(" WHERE status IN (%v, %v, %v)", "'Publish'", "'Draft'", "'Thrash'")
	} else {
		query += fmt.Sprintf(" WHERE status IN ('%v')", status)
	}
	query += " LIMIT ? OFFSET ?"
	log.Println(query)
	rows, err := db.Query(query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("Failed to load articles: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var Article models.Article
		if err := rows.Scan(&Article.ID, &Article.Title, &Article.Slug, &Article.Content, &Article.Category, &Article.CreatedDate, &Article.UpdatedDate, &Article.Status); err != nil {
			return nil, fmt.Errorf("Failed to scan article: %v", err)
		}
		Articles = append(Articles, Article)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("Failed when loading articles: %v", err)
	}

	return Articles, nil
}

func GetTotalArticles(status string) (int, error) {
	db := database.DB

	query := "SELECT COUNT(*) FROM articles"
	if status == "" {
		query += fmt.Sprintf(" WHERE status IN (%v, %v, %v)", "'Publish'", "'Draft'", "'Thrash'")
	} else {
		query += fmt.Sprintf(" WHERE status IN ('%v')", status)
	}
	var total int
	err := db.QueryRow(query).Scan(&total)
	if err != nil {
		return 0, fmt.Errorf("failed to get total count of articles: %v", err)
	}

	return total, nil
}

func GetArticleByIdOrSlug(ID string) (models.Article, error) {
	var Article models.Article
	db := database.DB // HOW TO GET CONNECTION/CALL DB FROM DATABASE/DATABASE.go ????
	row := db.QueryRow("SELECT * FROM articles WHERE id = ? OR slug = ?", ID, ID)
	if err := row.Scan(&Article.ID, &Article.Title, &Article.Slug, &Article.Content, &Article.Category, &Article.CreatedDate, &Article.UpdatedDate, &Article.Status); err != nil {
		if err == sql.ErrNoRows {
			return Article, fmt.Errorf("Artikel %d: tidak ditemukan", ID)
		}
		return Article, fmt.Errorf("Artikel %q: %v", ID, err)
	}
	return Article, nil
}

func UpdateArticle(ID string, article models.Article) (int64, error) {
	log.Println("UpdateArticle")
	log.Println(ID)
	db := database.DB

	var datetime = time.Now()
	dt := datetime.Format(time.RFC3339)

	randomString, err := utils.GenerateRandomString(5)
	if err != nil {
		return 0, fmt.Errorf("Error generating random string:", err)
	}
	resultSlug := slug.Make(article.Title) + "-" + randomString
	result, err := db.Exec("UPDATE articles SET title = ?, slug = ?, content = ?, category = ?, updated_date = ?, status = ? WHERE id = ?", article.Title, resultSlug, article.Content, article.Category, dt, article.Status, ID)
	if err != nil {
		return 0, fmt.Errorf("UpdateArticle: %v", err)
	}
	res, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("UpdateArticle: %v", err)
	}
	return res, nil
}

func DeleteArticle(ID string) (int64, error) {
	db := database.DB
	_, err := GetArticleByIdOrSlug(ID)
	if err != nil {
		return 0, fmt.Errorf(err.Error())
	}
	row, err := db.Exec("UPDATE articles SET status = 'Thrash' WHERE id = ?", ID)
	if err != nil {
		return 0, fmt.Errorf(err.Error())
	}
	if err != nil {
		return 0, fmt.Errorf(err.Error())
	}

	// Get the number of affected rows
	rowsAffected, err := row.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf(err.Error())
	}
	return rowsAffected, nil
}
