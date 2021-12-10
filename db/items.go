package db

import (
	"github.com/google/uuid"

	"gorm.io/gorm"
)

type Item struct {
	Id    string `json:"id,omitempty" gorm:"primary_key"`
	Name  string `json:"name"`
	Stock int    `json:"stock"`
}

func GetItems(items *[]Item) (err error) {
	if err := DB.Find(items).Error; err != nil {
		return err
	}
	return nil
}

func (item *Item) BeforeCreate(tx *gorm.DB) (err error) {
	item.Id = uuid.New().String()
	return
}

func CreateItem(item *Item) error {
	if err := DB.Create(item).Error; err != nil {
		return err
	}

	return nil
}

func GetItem(item *Item, itemId string) error {
	if err := DB.Where("id = ?", itemId).First(item).Error; err != nil {
		return err
	}
	return nil
}

func UpdateItem(item *Item) (err error) {
	DB.Save(item)
	return nil
}

func DeleteItem(item *Item) error {
	if err := DB.Delete(item).Error; err != nil {
		return err
	}

	return nil
}
