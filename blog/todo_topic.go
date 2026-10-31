package blog

import (
	"time"

	"gorm.io/gorm"
)

// all todo topics of the blog
type TodoTopic struct {
	TopicName  string    `gorm:"primaryKey" json:"topicName"`
	TopicOwner string    `gorm:"primaryKey; default:'root'" json:"topicOwner"`
	UpdatedAt  time.Time `gorm:"autoCreateTime" json:"updatedAt"`
	UpdateName string    `json:"updateName"`
}

// 查詢指定使用者的待辦主題，依時間由新到舊排序
func FindTodoTopicsByOwner(db *gorm.DB, owner string) ([]TodoTopic, error) {
	data := []TodoTopic{}

	err := db.Where("topic_owner = ?", owner).Order("updated_at DESC").Find(&data).Error
	return data, err
}

// 新增待辦主題
func CreateTodoTopic(db *gorm.DB, data *TodoTopic) error {
	return db.Create(data).Error
}
