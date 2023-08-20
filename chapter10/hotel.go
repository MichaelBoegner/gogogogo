/* taken from Chapter 10 14.1's Practical Application section https://www.practical-go-lessons.com/chap-10-functions#practical-application
Refactor this code and create functions that will :

- Compute the occupancy level

- Compute the occupancy rate

- Print the details of a specific room.

package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    const hotelName = "Gopher Paris Inn"
    const totalRooms = 134
    const firstRoomNumber = 110

    rand.Seed(time.Now().UTC().UnixNano())
    roomsOccupied := rand.Intn(totalRooms)
    roomsAvailable := totalRooms - roomsOccupied

    occupancyRate := (float32(roomsOccupied) / float32(totalRooms)) * 100
    var occupancyLevel string
    if occupancyRate > 70 {
        occupancyLevel = "High"
    } else if occupancyRate > 20 {
        occupancyLevel = "Medium"
    } else {
        occupancyLevel = "Low"
    }

    fmt.Println("Hotel:", hotelName)
    fmt.Println("Number of rooms", totalRooms)
    fmt.Println("Rooms available", roomsAvailable)
    fmt.Println("                  Occupancy Level:", occupancyLevel)
    fmt.Printf("                  Occupancy Rate: %0.2f %%\n", occupancyRate)

    if roomsAvailable > 0 {
        fmt.Println("Rooms:")
        for i := 0; roomsAvailable > i; i++ {
            roomNumber := firstRoomNumber + i
            size := rand.Intn(6) + 1
            nights := rand.Intn(10) + 1
            fmt.Println(roomNumber, ":", size, "people /", nights, " nights ")
        }
    } else {
        fmt.Println("No rooms available for tonight")
    }

}*/

package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    const hotelName = "Gopher Paris Inn"
    const totalRooms = 134
    const firstRoomNumber = 0

	rand.Seed(time.Now().UTC().UnixNano())
	roomsOccupied := rand.Intn(totalRooms)
	fmt.Println("Rooms Occupied", roomsOccupied)
    roomsAvailable := totalRooms - roomsOccupied
	fmt.Println("Rooms Available", roomsAvailable)

	occupancyRate := occupancyRate(totalRooms, roomsOccupied)
	occupancyLevel := occupancyLevel(occupancyRate)
	
    fmt.Println("Hotel:", hotelName)
    fmt.Println("Number of rooms", totalRooms)
    fmt.Println("Rooms available", roomsAvailable)
    fmt.Println("                  Occupancy Level:", occupancyLevel)
    fmt.Printf("                  Occupancy Rate: %0.2f %%\n", occupancyRate)
	roomsOccupiedDetails(roomsAvailable, firstRoomNumber, totalRooms, roomsOccupied)
}


func roomsOccupiedDetails(roomsAvailable int, firstRoomNumber int, totalRooms int, roomsOccupied int) {
	if roomsAvailable > 0 {
        for i := 0; i < roomsOccupied; i++ {
            roomNumber := firstRoomNumber + i
            size := rand.Intn(6) + 1
            nights := rand.Intn(10) + 1
            fmt.Println("Room Number:", roomNumber, "-", size, "people for", nights, "nights ")

			if i == (roomsOccupied - 1) {
				nextRoomAvailable := i + 1
				fmt.Println("Room Numbers Left:", nextRoomAvailable, "through", totalRooms)
			}
        }
    } else {
        fmt.Println("No rooms available for tonight")
    }
}

func occupancyRate(totalRooms int, roomsOccupied int) (occupancyRate float32) {
	occupancyRate = (float32(roomsOccupied) / float32(totalRooms)) * 100
	return
}

func occupancyLevel(occupancyRate float32) (occupancyLevel string) {
    if occupancyRate > 70 {
        occupancyLevel = "High"
    } else if occupancyRate > 20 {
        occupancyLevel = "Medium"
    } else {
        occupancyLevel = "Low"
    }
	
	return
}