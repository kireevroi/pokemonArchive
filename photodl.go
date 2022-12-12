package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	if _, err := os.Stat("pokemon"); errors.Is(err, os.ErrNotExist) {
		if err := os.Mkdir("pokemon", os.ModePerm); err != nil {
			log.Fatal(err)
		}
	}
	for i := 1; i <= 905; i++ {
		s := fmt.Sprintf("%03d", i)
		fileName := s + ".png"
		URL := "https://assets.pokemon.com/assets/cms2/img/pokedex/full/" + fileName
		err := downloadFile(URL, "pokemon/"+fileName)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func downloadFile(URL, fileName string) error {
	//Get the response bytes from the url
	response, err := http.Get(URL)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return errors.New("Received non 200 response code")
	}
	//Create a empty file
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	//Write the bytes to the fiel
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	return nil
}