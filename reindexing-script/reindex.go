package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func main() {
	// declaring vars
	var directory string
	var pairs [][]int

	// clearing terminal
	clear_terminal()

	// getting user parameters
	directory, pairs = gets_user_input()

	// confirming with user
	confirm_parameters(directory, pairs)

	// replacing labels
	replace_labels(directory, pairs)
}

// clears the terminal
func clear_terminal() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// gets user paramters
func gets_user_input() (string, [][]int) {
	// declaring vars
	var dir string
	var pairs [][]int
	var list_of_pair_strings []string

	// prompting user
	fmt.Println("--------------------------------------------------------------------------------------")
	fmt.Println("label: Enter the file path of the folder containing the files to be relabeled")
	fmt.Print("-> ")

	// creating scanner
	scanner := bufio.NewScanner(os.Stdin)

	// getting directory path from user
	if scanner.Scan() {
		dir = scanner.Text()
	}

	// prompting user
	fmt.Println("--------------------------------------------------------------------------------------")
	fmt.Println("system: Enter your desired replacemnt pairs. ")
	fmt.Println("\tEx. \"14,1 15,2 16,3\" creates the following rules: ")
	fmt.Println("\t\t14 -> 1")
	fmt.Println("\t\t15 -> 2")
	fmt.Println("\t\t16 -> 3")

	for {
		fmt.Println("--------------------------------------------------------------------------------------")
		fmt.Print("-> ")

		// scanning and parsing user input
		if scanner.Scan() {
			if scanner.Text() == "" {
				fmt.Println("No replacement pair provided")
				continue
			}
			list_of_pair_strings = strings.Split(scanner.Text(), " ")
			break
		}
	}

	for _, current_pair_string := range list_of_pair_strings {
		current_pair := strings.Split(current_pair_string, ",")

		num1, err := strconv.Atoi(current_pair[0])
		if err != nil {
			fmt.Println("label: Error converting string -", err)
		}

		num2, err := strconv.Atoi(current_pair[1])
		if err != nil {
			fmt.Println("label: Error converting string -", err)
		}
		pairs = append(pairs, []int{num1, num2})
	}

	return dir, pairs
}

func confirm_parameters(directory string, pairs [][]int) {
	// printing what the user has entered to confirm
	fmt.Println("--------------------------------------------------------------------------------------")
	fmt.Println("The path you entered was:")
	fmt.Printf("\t%s\n", directory)
	fmt.Println("--------------------------------------------------------------------------------------")
	fmt.Println("The replacement pairs you entered were:")
	for _, pair := range pairs {
		fmt.Printf("\t%d -> %d\n", pair[0], pair[1])
	}

	// looping until user has entered a valid option
	for {
		// prompting user
		fmt.Println("--------------------------------------------------------------------------------------")
		fmt.Println("Enter 'p' to proceed or 'c' to cancel:")
		fmt.Print("-> ")

		// creating scanner
		scanner := bufio.NewScanner(os.Stdin)

		// getting user input
		var option string
		if scanner.Scan() {
			option = scanner.Text()
		}

		// responding to selected option
		if option == "c" {
			fmt.Println("--------------------------------------------------------------------------------------")
			fmt.Println("system: Exiting...")
			os.Exit(0)
		} else if option != "p" {
			fmt.Println("Invalid option")
		} else {
			break
		}
	}

	fmt.Println("--------------------------------------------------------------------------------------")
}

func replace_labels(directory string, pairs [][]int) {
	// Open the directory
	files, err := os.ReadDir(directory)
	if err != nil {
		fmt.Println("systen: Error opening directory")
		return
	}

	// looping through files
	for _, file := range files {
		// skipping class file
		if file.Name() == "classes.txt" {
			fmt.Println("system: Skipping class file")
			continue
		}

		// creating filepath from file name and directory path
		filePath := filepath.Join(directory, file.Name())

		// checking if the file is a txt
		if file.Name()[len(file.Name())-3:] != "txt" {
			continue
		}

		// opening file
		file_descriptor, err := os.OpenFile(string(filePath), os.O_RDWR, 0644)
		if err != nil {
			fmt.Println("system: Error opening file")
			continue
		}
		fmt.Println("--------------------------------------------------------------------------------------")
		fmt.Printf("system: Replacing labels in file \"%s\"\n", file.Name())

		// creating scanner to read file
		file_scanner := bufio.NewScanner(file_descriptor)

		// buffer to hold contents to be written to the file
		var content_buffer strings.Builder

		// looping over file to find label indexes and swap replace as needed
		for file_scanner.Scan() {
			// getting contents of the current line in the file
			current_line := file_scanner.Text()

			// getting the first element of the line and converting it to an integer
			label_number, err := strconv.Atoi(strings.Split(current_line, " ")[0])
			if err != nil {
				fmt.Println("system: Error converting string -", err)
			}

			not_replaced := true
			for _, pair := range pairs {
				// checking if label needs to be replaced
				if label_number == pair[0] {
					// replacing label
					modified_line := strings.Replace(current_line, strconv.Itoa(label_number), strconv.Itoa(pair[1]), 1)

					fmt.Printf("\tReplacing label %d with label %d\n", label_number, pair[1])

					// write to content_buffer
					content_buffer.WriteString(modified_line + "\n")

					// setting replaced flag
					not_replaced = false

					break
				}
			}

			if not_replaced {
				// write to content_buffer
				content_buffer.WriteString(current_line + "\n")
			}
		}
		// writing to file
		file_descriptor.Truncate(0) // Clear the file
		file_descriptor.Seek(0, 0)  // Reset the file pointer to the beginning
		bytes_written, err := file_descriptor.WriteString(content_buffer.String())
		if err != nil {
			fmt.Println("system: Error writing to file")
			continue
		}
		fmt.Printf("system: wrote %d bytes to file \"%s\"\n", bytes_written, file.Name())

		file_descriptor.Close()
		time.Sleep(100 * time.Millisecond)
	}
}

func error_shut_down() {
	fmt.Println("systen: Shutting down...")
	os.Exit(1)
}
