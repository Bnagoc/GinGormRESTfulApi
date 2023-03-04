package models

import "github.com/bnagos/GinGormRESTfulApi/storage"

func GetAllArticles(a *[]Article) error {
	return storage.DB.Find(a).Error
}

func AddNewArticles(a *Article) error {
	return storage.DB.Create(a).Error
}

func GetArticleById(a *Article, id string) error {
	return storage.DB.Where("id = ?", id).First(a).Error
}

func DeleteArticleById(a *Article, id string) error {
	return storage.DB.Where("id = ?", id).Delete(a).Error
}

func UpdateArticleById(a *Article, id string) error {
	return storage.DB.Where("id = ?", id).Update(a).Error
}
