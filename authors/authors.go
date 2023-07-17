package authors

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

const authorsFile = "authors.json"

type Author struct {
	ID      string
	Name    string
	Country string
	PenName string
}

type Authors struct {
	Authors []Author
}

func CheckFile(filename string) error {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		_, err := os.Create(filename)
		if err != nil {
			return err
		}
	}
	return nil
}

func GetAllAuthors() Authors {
	err := CheckFile(authorsFile)
	if err != nil {
		log.Println(err)
	}

	jsonFile, err := os.Open(authorsFile)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened", authorsFile)

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var authors Authors
	json.Unmarshal(byteValue, &authors)

	return authors
}

func (b Authors) PrintAllAuthors() {
	for i := 0; i < len(b.Authors); i++ {
		fmt.Println(b.Authors[i].ToString())
	}
}

func (b Author) ToString() string {
	return "Name: " + b.Name + ", " + "Pen Name: " + b.PenName + ", " + "Country: " + b.Country
}

func (a Authors) GetAuthorByName(name string) Author {
	for i := 0; i < len(a.Authors); i++ {
		if a.Authors[i].Name == name {
			fmt.Printf("Author found: Name: %s\n Country: %s\n Pen Name: %s\n", a.Authors[i].Name, a.Authors[i].Country, a.Authors[i].PenName)
			return a.Authors[i]
		}
	}

	return Author{}
}
