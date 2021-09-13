package cache

import "redisTesting/models"

type PostCache interface{
	Set(key string, value *models.Post)
	Get(key string) *models.Post
}
