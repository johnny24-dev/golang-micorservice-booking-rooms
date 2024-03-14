package comment

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/xuri/excelize/v2"
	"strconv"
	"test/helpers"
)

func Excelconn() (*excelize.File, error) {
	f, err := excelize.OpenFile("./seed/comment/comment.xlsx")
	if err != nil {
		return &excelize.File{}, err
	}
	defer func() {
		if err := f.Close(); err != nil {
			return
		}
	}()
	return f, nil
}

type Comment struct {
	Text string
	Rate float64
}

func SeedComment() {
	mapping := map[int]Comment{}

	conn, err := Excelconn()
	if err != nil {
		fmt.Println(err)
		return
	}

	for i := 1; i <= 130; i++ {
		stt, _ := conn.GetCellValue("Sheet1", "A"+strconv.Itoa(i))
		comment, _ := conn.GetCellValue("Sheet1", "B"+strconv.Itoa(i))
		rate, _ := conn.GetCellValue("Sheet1", "C"+strconv.Itoa(i))

		if s, err := strconv.ParseFloat(rate, 32); err == nil {
			st, _ := strconv.Atoi(stt)
			mapping[st] = Comment{
				Text: comment,
				Rate: s,
			}
		}
	}

	fmt.Println(mapping)

	//var wg sync.WaitGroup

	for i := 6; i < 125; i++ {
		for j := 0; j < 50; j++ {
			obj := mapping[helpers.Random(1, 129)]
			text := obj.Text
			rate := obj.Rate
			url := "http://103.184.113.181:81/hotel/add_comment"
			body := map[string]interface{}{
				"text":        text,
				"type":        "Text",
				"rate":        rate,
				"hotel_id":    i,
				"customer_id": helpers.Random(2, 11),
			}
			bodyJson, _ := json.Marshal(body)
			_, err2 := helpers.SendRequest("POST", url, fiber.Map{}, bodyJson)
			if err2 != nil {
				fmt.Println(err2)
				break
			}
		}
		fmt.Println("Done hotel", i)
	}

	fmt.Println("OK")
}
