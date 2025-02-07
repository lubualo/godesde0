package files

import (
	"bufio"
	"fmt"
	"os"

	"github.com/lubualo/godesde0/excercises"
)

const fileName string = "./files/txt/table.txt"

func SaveTable()  {
	var text string = excercises.Excercise02()
	file, err := os.Create(fileName)
	if (err != nil) {
		fmt.Println("Error while creating the file: " + err.Error())
		return
	}
	defer file.Close()
	fmt.Fprint(file, text)	
}

func AppendTableToFile()  {
	var text string = excercises.Excercise02()
	if (!append(fileName, text)) {
		fmt.Println("Error while concatenating new table")
	}
}

func append(filen string, text string) bool  {
	file, err := os.OpenFile(filen, os.O_WRONLY | os.O_APPEND, 0644)
	if (err != nil) {
		fmt.Println("Error while appending new table: " + err.Error())
		return false
	}
	defer file.Close()
	_, err = file.WriteString(text)
	if (err != nil) {
		fmt.Println("Error while appending new table: " + err.Error())
		return false
	}
	return true
}

// Loads the file line by line
func ReadFile() {
	file, err := os.Open(fileName)
	if (err != nil) {
		fmt.Println("Error while reading the file: " + err.Error())
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for (scanner.Scan()) {
		record := scanner.Text()
		fmt.Println("> " + record)
	}
}

// Loads the entire file at once, more memory consuming
func ReadFileV2() {
	file, err := os.ReadFile(fileName)
	if (err != nil) {
		fmt.Println("Error while reading the file: " + err.Error())
		return
	}
	fmt.Println(string(file))
}
