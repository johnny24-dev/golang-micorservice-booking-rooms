package util

import (
	"fmt"
)

func BuildQuerySearch(searchField string, searchValue string) string {
	var sqlQuery string
	switch searchField {
	case "name":
		sqlQuery = fmt.Sprintf("name LIKE '%%%s%%'", searchValue)
	case "location":
		sqlQuery = fmt.Sprintf("address_id in (SELECT id FROM addresses WHERE province LIKE '%%%s%%' AND type = 'hotel')", searchValue)
	case "room_type":
		sqlQuery = fmt.Sprintf("id in (select rooms.hotel_id from rooms where type like '%%%s%%')", searchValue)
	case "check_in":
		bookedRoomSql := fmt.Sprintf("SELECT * FROM doandb.booked_rooms WHERE check_in >= '%s'", searchValue)
		countBookedRoomSql := fmt.Sprintf("SELECT COUNT(*) FROM doandb.booked_rooms where check_in >= '%s' AND room_id = rooms.id", searchValue)
		roomSql := fmt.Sprintf("SELECT DISTINCT rooms.id FROM doandb.rooms INNER JOIN (%s) AS u ON u.room_id = rooms.id AND rooms.quantity <= (%s)", bookedRoomSql, countBookedRoomSql)
		sqlQuery = fmt.Sprintf("id IN (SELECT DISTINCT hotels.id FROM doandb.hotels INNER JOIN rooms ON rooms.hotel_id = hotels.id AND rooms.id NOT IN (%s) GROUP BY hotels.id)", roomSql)
	case "check_out":
		bookedRoomSql := fmt.Sprintf("SELECT * FROM doandb.booked_rooms WHERE check_out <= '%s'", searchValue)
		countBookedRoomSql := fmt.Sprintf("SELECT COUNT(*) FROM doandb.booked_rooms where check_out <= '%s' AND room_id = rooms.id", searchValue)
		roomSql := fmt.Sprintf("SELECT DISTINCT rooms.id FROM doandb.rooms INNER JOIN (%s) AS u ON u.room_id = rooms.id AND rooms.quantity <= (%s)", bookedRoomSql, countBookedRoomSql)
		sqlQuery = fmt.Sprintf("id IN (SELECT DISTINCT hotels.id FROM doandb.hotels INNER JOIN rooms ON rooms.hotel_id = hotels.id AND rooms.id NOT IN (%s) GROUP BY hotels.id)", roomSql)
	case "star_level":
		sqlQuery = fmt.Sprintf("star_level = '%s'", searchValue)
	}

	return sqlQuery
}

func CaulculateAverange(list []float32) float32 {
	var sum float32
	for _, val := range list {
		sum += val
	}

	return sum / float32(len(list))
}

/*

SELECT DISTINCT hotels.id FROM doandb.hotels INNER JOIN
rooms ON rooms.hotel_id = hotels.id AND rooms.id NOT IN
(SELECT DISTINCT rooms.id FROM doandb.rooms INNER JOIN
(SELECT * FROM doandb.booked_rooms WHERE check_in >= '2023-04-28 10:00:00 +0000 +0000') AS u ON u.room_id = rooms.id
AND rooms.quantity <= (SELECT COUNT(*) FROM doandb.booked_rooms where check_in >= '2023-04-28 10:00:00 +0000 +0000' AND room_id = rooms.id))

*/
