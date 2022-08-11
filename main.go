package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"strconv"
)

func main() {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		// ... handle error if file missing
		log.Fatal(err)
	}
    myString := string(content) 
	windows := strings.Replace(myString, "\n", "\r\n", -1)
	my_arr := strings.Split(strings.Replace(windows, "\r\r\n", "\n", -1), "\n")

	// default variable filed
	position_x := 0
	position_y := 0
	direction_degree := 90
	direction := "N"

	//check & assign variable size 
	size, err := strconv.Atoi(my_arr[0])
    if err != nil {
        // ... handle error if first input not integer
        log.Fatal(err)
    }

	// loop check all input
	for i := 0; i < len(my_arr); i++ {

		//change degree if found input right , left  
		if my_arr[i] == "R" {
			direction_degree += 90
			if direction_degree == 360{
				direction_degree = 0
			}
		} else if my_arr[i] == "L" {
			if direction_degree == 0{
				direction_degree = 360
			}
			direction_degree -= 90
		}

		//assign now direction and check moving
		switch direction_degree {
		case 90:
			direction = "N"
			if my_arr[i] == "F"{
				position_y += 1
				//check if moving reaches the edge
				if position_y > size{
					position_y = size
				}
			}
		case 180:
			direction = "E"
			if my_arr[i] == "F"{
				position_x += 1
				//check if moving reaches the edge
				if position_x > size{
					position_x = size
				}
			}
		case 270:
			direction = "S"
			if my_arr[i] == "F"{
				position_y -= 1
				//check if moving to negative position
				if position_y < 0{
					position_y = 0
				}
			}
		case 0:
			direction = "w"
			if my_arr[i] == "F"{
				position_x -= 1
				//check if moving to negative position
				if position_x < 0{
					position_x = 0
				}
			}
		default:
			
		}
		// print result
		fmt.Printf("Input = %s , Output %s:%d,%d \n", my_arr[i],direction, position_x, position_y)
	}
}