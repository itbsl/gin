package models

type ArticleTag struct {
	*Model
	ArticleId uint32 `json:"article_id"`
	TagId uint32 `json:"tag_id"`
}