package publishers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	 
	
		
		

)

const publishersFile = "publishers.json"

type Publisher struct {
	ID      string
	Name    string
	Country string
}

type Publishers struct {
	Publishers []Publisher
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

func GetAllPublishers() Publishers {
	err := CheckFile(publishersFile)
	if err != nil {
		log.Println(err)
	}

	jsonFile, err := os.Open(publishersFile)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened", publishersFile)

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var publishers Publishers
	json.Unmarshal(byteValue, &publishers)

	return publishers
}

func (p Publishers) PrintAllPublishers() {
	for i := 0; i < len(p.Publishers); i++ {
		fmt.Println(p.Publishers[i].ToString())
	}
}

func (p Publisher) ToString() string {
	return "Name: " + p.Name + ", " + "Country: " + p.Country
}

func (p Publishers) GetPublisherByName(name string) Publisher {
	for i := 0; i < len(p.Publishers); i++ {
		if p.Publishers[i].Name == name {
			fmt.Printf("Publisher found: Name: %s\n Country: %s\n", p.Publishers[i].Name, p.Publishers[i].Country)
			return p.Publishers[i]
		}
	}

	return Publisher{}
}

func AddPublisher(publisher Publisher) {
	publishers := GetAllPublishers()
	publishers.Publishers = append(publishers.Publishers, publisher)

	updatedPublishersJSON, err := json.Marshal(publishers)
	if err != nil {
		log.Println("Error marshaling publishers:", err)
		return
	}

	err = ioutil.WriteFile(publishersFile, updatedPublishersJSON, 0644)
	if err != nil {
		log.Println("Error writing publishers file:", err)
		return
	}

	fmt.Println("Publisher added successfully.")
}