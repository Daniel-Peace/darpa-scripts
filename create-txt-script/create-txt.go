package main

import (
	"fmt"
	"os"
	"os/exec"
)

// variables
var (
	dir_path string
	prefix   string
	err      error
)

func main() {
	// prompting user
	clear_terminal()
	fmt.Println("-------------------------------------------------------------------")
	fmt.Println("Enter the directory path that contains all of the labels:")
	fmt.Println("-------------------------------------------------------------------")

	// looping until valid path is enterer
	for {
		fmt.Print("-> ")
		_, err = fmt.Scan(&dir_path)
		if err != nil {
			fmt.Println("-------------------------------------------------------------------")
			fmt.Print("system: \x1b[31mERROR - ", err)
			fmt.Println("\x1b[0m")
			fmt.Println("-------------------------------------------------------------------")
			return
		}
		fmt.Println("-------------------------------------------------------------------")

		if path_exists() {
			fmt.Print("system: \x1b[32mSuccessfully found directory\x1b[0m\n")
			fmt.Println("-------------------------------------------------------------------")
			break
		} else {
			fmt.Print("system: \x1b[31mERROR - ", err)
			fmt.Println("\x1b[0m")
			fmt.Println("-------------------------------------------------------------------")
		}
	}

	fmt.Println("Enter the prefix you would like to add to the \nbeginning of each file name. Just enter \"[-]\" \nif you wish to add nothing:")
	fmt.Println("-------------------------------------------------------------------")
	fmt.Print("-> ")
	_, err = fmt.Scan(&prefix)
	if err != nil {
		fmt.Println("-------------------------------------------------------------------")
		fmt.Print("system: \x1b[31mERROR - ", err)
		fmt.Println("\x1b[0m")
		fmt.Println("-------------------------------------------------------------------")
		return
	}

	if prefix == "[-]" {
		prefix = ""
	}
	fmt.Println("-------------------------------------------------------------------")

	// creating array of files from directory
	fmt.Println("system: \x1b[34mOpening directory...\x1b[0m")
	fmt.Println("-------------------------------------------------------------------")
	files, err := os.ReadDir(dir_path)
	if err != nil {
		fmt.Print("system: \x1b[31mERROR - ", err)
		fmt.Println("\x1b[0m")
		fmt.Println("-------------------------------------------------------------------")
		return
	} else {
		fmt.Println("system: \x1b[32mSuccessfully opened directory\x1b[0m")
		fmt.Println("-------------------------------------------------------------------")
	}

	// create test.txt
	fmt.Println("system: \x1b[34mCreating test.txt...\x1b[0m")
	fmt.Println("-------------------------------------------------------------------")
	test_file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	} else {
		fmt.Println("system: \x1b[32mSuccessfully created test.txt\x1b[0m")
		fmt.Println("-------------------------------------------------------------------")
	}
	defer test_file.Close()

	// create train.txt
	fmt.Println("system: \x1b[34mCreating train.txt...\x1b[0m")
	fmt.Println("-------------------------------------------------------------------")
	train_file, err := os.Create("train.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	} else {
		fmt.Println("system: \x1b[32mSuccessfully created train.txt\x1b[0m")
		fmt.Println("-------------------------------------------------------------------")
	}
	defer train_file.Close()

	// looping over files in directory
	file_counter := 1
	for _, file := range files {
		// ensuring file is a text file
		if file.Name()[len(file.Name())-3:] != "png" {
			continue
		}

		// printing file name
		fmt.Printf("\tcurrent file: \x1b[35m%s\x1b[0m\n", file.Name())
		data := prefix + file.Name() + "\n"

		// checking if the file needs to be added to the test.txt
		if file_counter%5 == 0 {
			fmt.Println("\t\x1b[34madding file name to test.txt...\x1b[0m")
			_, err = test_file.WriteString(data)
			if err != nil {
				fmt.Println("/tERROR - ", err)
				fmt.Println("-------------------------------------------------------------------")
				return
			}
			fmt.Println("\t\x1b[32msuccessfully added file name to test.txt...\x1b[0m")
		} else {
			fmt.Println("\t\x1b[34madding file name to train.txt...\x1b[0m")
			_, err = train_file.WriteString(data)
			if err != nil {
				fmt.Println("/tERROR - ", err)
				fmt.Println("-------------------------------------------------------------------")
				return
			}
			fmt.Println("\t\x1b[32msuccessfully added file name to train.txt...\x1b[0m")
		}
		fmt.Println("-------------------------------------------------------------------")

		file_counter++
	}
}

func clear_terminal() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func path_exists() bool {
	_, err = os.Stat(dir_path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}
