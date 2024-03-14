package hotel

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"os"
	"strconv"
	"test/helpers"
)

type Hotel struct {
	Name        string `json:"name"`
	StarLevel   string `json:"star_level"`
	Rule        string `json:"rule"`
	Description string `json:"description"`
	Address     struct {
		DetailAddress string `json:"detail_address"`
		District      string `json:"district"`
		Province      string `json:"province"`
	} `json:"address"`
	ListImage []struct {
		Url  string `json:"url"`
		Type string `json:"type"`
	} `json:"list_image"`
}

func SeedHotel() {
	listNameHotel := readFile("hotelname.txt")
	//listDistrict := readFile("districtname.txt")
	b := []string{"Thành phố Lào Cai", "Huyện Bát Xát", "Huyện Mường Khương", "Huyện Si Ma Cai", "Huyện Bắc Hà", "Huyện Bảo Thắng", "Huyện Bảo Yên", "Huyện Sa Pa", "Huyện Văn Bàn"}
	for i := 0; i < len(listNameHotel); i++ {
		province := "Lào Cai"
		district := b[helpers.Random(0, 9)]
		number := strconv.Itoa(helpers.Random(0, 2000))
		detailAddrs := number + " " + district + ", " + province
		body := Hotel{
			Name:        listNameHotel[i],
			StarLevel:   strconv.Itoa(helpers.Random(0, 6)),
			Rule:        "",
			Description: "",
			Address: struct {
				DetailAddress string `json:"detail_address"`
				District      string `json:"district"`
				Province      string `json:"province"`
			}{
				DetailAddress: detailAddrs,
				District:      district,
				Province:      province,
			},
			ListImage: []struct {
				Url  string `json:"url"`
				Type string `json:"type"`
			}{
				{Url: "https://chatbizfly.mediacdn.vn/2023/03/28backend_chat/_linkimagedefaultpng1679991208.png", Type: "hotel"},
				{Url: "https://chatbizfly.mediacdn.vn/2023/03/28backend_chat/_linkimagedefaultpng1679991208.png", Type: "hotel"},
				{Url: "https://chatbizfly.mediacdn.vn/2023/03/28backend_chat/_linkimagedefaultpng1679991208.png", Type: "hotel"},
				{Url: "https://chatbizfly.mediacdn.vn/2023/03/28backend_chat/_linkimagedefaultpng1679991208.png", Type: "hotel"},
				{Url: "https://chatbizfly.mediacdn.vn/2023/03/28backend_chat/_linkimagedefaultpng1679991208.png", Type: "hotel"},
			},
		}
		variable := strconv.Itoa(helpers.Random(2, 11))
		url := fmt.Sprintf("http://103.184.113.181:81/customer/%s/hotel", variable)
		bodyJson, _ := json.Marshal(body)
		_, err := helpers.SendRequest("POST", url, fiber.Map{}, bodyJson)
		if err != nil {
			fmt.Println(err)
			break
		}
	}
}

func readFile(fileName string) []string {
	var a []string

	filePath := "./seed/hotel/" + fileName
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		a = append(a, scanner.Text())
	}

	return a
}
