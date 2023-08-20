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
    const firstRoomNumber = 110

    rand.Seed(time.Now().UTC().UnixNano())
    roomsOccupied := rand.Intn(totalRooms)
    roomsAvailable := totalRooms - roomsOccupied

    occupancyRate := (float32(roomsOccupied) / float32(totalRooms)) * 100
	
    fmt.Println("Hotel:", hotelName)
    fmt.Println("Number of rooms", totalRooms)
    fmt.Println("Rooms available", roomsAvailable)
    fmt.Println("                  Occupancy Level:", occupancyLevel(occupancyRate))
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