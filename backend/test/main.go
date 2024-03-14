package main

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/nekizz/final-project/backend/hotel/model"
	"gorm.io/gorm"
	"test/helpers"
)

func main() {
	var listImages []*model.Image

	listImages = append(listImages, &model.Image{
		Model: gorm.Model{
			ID: 554,
		},
		Url:  "123",
		Type: "hotel",
	})

	listImages = append(listImages, &model.Image{
		Model: gorm.Model{
			ID: 555,
		},
		Url:  "123",
		Type: "hotel",
	})

	listImages = append(listImages, &model.Image{
		Model: gorm.Model{
			ID: 556,
		},
		Url:  "123",
		Type: "hotel",
	})

	listImages = append(listImages, &model.Image{
		Model: gorm.Model{
			ID: 557,
		},
		Url:  "123",
		Type: "hotel",
	})

	listImages = append(listImages, &model.Image{
		Model: gorm.Model{
			ID: 558,
		},
		Url:  "123",
		Type: "hotel",
	})

	hotel := map[string]interface{}{
		"id":         120,
		"name":       "Leminatse Anis Motel",
		"star_level": 4,
		"address": map[string]interface{}{
			"id":             126,
			"district":       "Quận Ba Đình",
			"province":       "Hà Nội",
			"detail_address": "167 Quận Ba Đình, Hà Nội",
		},
		"list_image": listImages,
	}
	customerJson, _ := json.Marshal(hotel)
	url := "http://localhost:81/hotel/120"

	data, err := helpers.SendRequest("PUT", url, fiber.Map{}, customerJson)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(data))
}
