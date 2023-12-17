package main

import "fmt"
import "os"
import "bufio"
import "strconv"
import "strings"

type ball_ struct {
	color string
	value int
}

type balls struct {
	red   int
	blue  int
	green int
}

type id_with_balls struct {
	id        int
	balls_map balls
}

func check_error(err error) {
	if err != nil {
		panic(err)
	}
}

func convert_string_to_number(letter string) int {
	number, err := strconv.Atoi(letter)
	check_error(err)
	return number
}

func convert_number_to_string(number int) string {
	letter := strconv.Itoa(number)
	return letter
}

func get_id_of_game(line string) int {
	id := convert_string_to_number(line[strings.Index(line, "Game")+5:])
	return id
}
func convert_to_ball(word string) ball_ {
	result := strings.Split(word, " ")
	ball := ball_{color: result[1], value: convert_string_to_number(strings.TrimSpace(result[0]))}
	return ball
}

func get_balls_of_game(line string) balls {
	combos := strings.Split(line, ";")
	balls_map := balls{red: 0, blue: 0, green: 0}

	for _, combo := range combos {
		config := strings.Split(combo, ",")
		for _, conf := range config {
			ball := convert_to_ball(conf[1:])
			if ball.color == "red" && balls_map.red < ball.value {
				balls_map.red = ball.value
			} else if ball.color == "blue" && balls_map.blue < ball.value {
				balls_map.blue = ball.value
			} else if ball.color == "green" && balls_map.green < ball.value {
				balls_map.green = ball.value
			}
		}
	}

	return balls_map
}

func get_sum_of_all_game_ids(ids []int) int {
	sum := 0
	for _, id := range ids {
		sum += id
	}
	return sum
}

func get_power_set(balls balls) int {
	return balls.red * balls.blue * balls.green
}

func main() {
	readFile, err := os.Open("input.txt")
	check_error(err)
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	power_set_list := []int{}

	for fileScanner.Scan() {
		line := fileScanner.Text()
		balls := get_balls_of_game(line[strings.Index(line, ":")+1:])
		power := get_power_set(balls)
		power_set_list = append(power_set_list, power)
	}
	sum := get_sum_of_all_game_ids(power_set_list)
	fmt.Println("Sum: ", sum)
	readFile.Close()

}
