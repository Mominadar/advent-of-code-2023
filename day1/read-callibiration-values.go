package main

import "fmt"
import "os"
import "bufio"
import "strconv"
import "regexp"

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

func is_numeric(word string) bool {
	return regexp.MustCompile(`\d`).MatchString(word)
}

func get_sum_of_first_and_last_number_from_line(line string) int {
	// get all numeric characters
	numbers := []string{}
	for _, char := range line {
		character := string(char)
		if is_numeric(character) {
			numbers = append(numbers, character)
		}
	}

	// get first and last digits
	first_number := numbers[0]
	second_number := numbers[len(numbers)-1]

	// convert to two digit number
	number := convert_string_to_number(first_number + second_number)

	return number
}

func main() {
	readFile, err := os.Open("input.txt")
	check_error(err)
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	sum := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		sum += get_sum_of_first_and_last_number_from_line(line)
	}

	fmt.Println("Sum: ", sum)
	readFile.Close()

}
