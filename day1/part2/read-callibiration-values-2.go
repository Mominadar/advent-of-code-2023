package main

import "fmt"
import "os"
import "bufio"
import "strconv"
import "regexp"

var letters_map = map[string]int{
	"zero":  0,
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
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

func is_numeric(word string) bool {
	return regexp.MustCompile(`\d`).MatchString(word)
}

func is_number_in_letters(line string, index int) int {
	for i := 1; i < 6 && index+i <= len(line); i++ {
		substring := line[index : index+i]
		value, check_existence := letters_map[substring]
		if check_existence {
			return value
		}
	}
	return -1
}

func get_sum_of_first_and_last_number_from_line(line string) int {
	// get all numeric characters
	numbers := []string{}
	for index, char := range line {
		character := string(char)
		if is_numeric(character) {
			numbers = append(numbers, character)
		} else if is_number_in_letters(line, index) != -1 {
			numbers = append(numbers, convert_number_to_string(is_number_in_letters(line, index)))
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
