package room

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"test/helpers"
	"test/util"
)

type Room struct {
	Name        string    `json:"name"`
	Type        string    `json:"type"`
	Price       float32   `json:"price"`
	Quantity    int       `json:"quantity"`
	Description string    `json:"description"`
	Class       string    `json:"class"`
	ListAmenity []Amenity `json:"list_amenity"`
}

type Amenity struct {
	Name string `json:"name"`
}

//Note:
func SeedRoom() {
	a := map[string][]Amenity{}
	mapPrice := map[string][]float32{}

	singleAmenity := []Amenity{{Name: "Bed"}, {Name: "Tivi"}, {Name: "Fridge"}, {Name: "Wardrobe"}}
	doubleAmenity := []Amenity{{Name: "Double Bed"}, {Name: "Tivi"}, {Name: "Fridge"}, {Name: "Wardrobe"}}
	vipAmenity := []Amenity{{Name: "King Bed"}, {Name: "Tivi"}, {Name: "Fridge"}, {Name: "Wardrobe"}, {Name: "Bathtub"}, {Name: "Massage Chair"}}

	listRoomName := util.ReadFile("roomname.txt")
	typeR := []string{"Single", "Double", "VIP"}

	a[typeR[0]] = singleAmenity
	a[typeR[1]] = doubleAmenity
	a[typeR[2]] = vipAmenity

	singlePrice := []float32{100000, 200000, 300000}
	doublePrice := []float32{400000, 500000, 600000}
	vipPrice := []float32{700000, 800000, 900000}

	mapPrice[typeR[0]] = singlePrice
	mapPrice[typeR[1]] = doublePrice
	mapPrice[typeR[2]] = vipPrice
	count := 0
	for j := 6; j < 120; j++ {
		typeN := typeR[0]
		priceN := mapPrice[typeN][helpers.Random(0, 3)]
		body := Room{
			Name:        listRoomName[count] + " Room",
			Type:        typeN,
			Price:       priceN,
			Quantity:    helpers.Random(1, 10),
			Description: "",
			ListAmenity: a[typeN],
		}
		hotelID := strconv.Itoa(j)
		url := fmt.Sprintf("http://103.184.113.181:82/hotel/%s/room", hotelID)
		bodyJson, _ := json.Marshal(body)
		_, err := helpers.SendRequest("POST", url, fiber.Map{}, bodyJson)
		if err != nil {
			fmt.Println(err)
			break
		} else {
			count++
		}

		typeN1 := typeR[1]
		priceN1 := mapPrice[typeN1][helpers.Random(0, 3)]
		body1 := Room{
			Name:        listRoomName[count] + " Room",
			Type:        typeN1,
			Price:       priceN1,
			Quantity:    helpers.Random(1, 10),
			Description: "",
			ListAmenity: a[typeN1],
		}
		url1 := fmt.Sprintf("http://103.184.113.181:82/hotel/%s/room", hotelID)
		bodyJson1, _ := json.Marshal(body1)
		_, err1 := helpers.SendRequest("POST", url1, fiber.Map{}, bodyJson1)
		if err1 != nil {
			fmt.Println(err1)
			break
		} else {
			count++
		}

		typeN2 := typeR[2]
		priceN2 := mapPrice[typeN][helpers.Random(0, 3)]
		body2 := Room{
			Name:        listRoomName[count] + " Room",
			Type:        typeN2,
			Price:       priceN2,
			Quantity:    helpers.Random(1, 10),
			Description: "",
			ListAmenity: a[typeN2],
		}
		url2 := fmt.Sprintf("http://103.184.113.181:82/hotel/%s/room", hotelID)
		bodyJson2, _ := json.Marshal(body2)
		_, err2 := helpers.SendRequest("POST", url2, fiber.Map{}, bodyJson2)
		if err2 != nil {
			fmt.Println(err2)
			break
		} else {
			count++
		}
		fmt.Println(j)
	}

}

//Class: Single -> Amenity: Bed, Tivi, Fridge, Wardrobe
//Class: Double -> Double Bed, Tivi, Fridge, Wardrobe
//Class VIP -> King Bed, Tivi, Fridge, Wardrobe, Bathtub, Massage Chair
