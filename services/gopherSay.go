package services

/***
 *
 * This file mainly generat gopher withe ASCII code art
 *
 * @author: Boris
 * @version: 2022-03-21
 *
 */

import (
	"embed"
	"fmt"
	"log"
)

//go:embed gophers
var embedGopherFiles embed.FS

//This function mainly display a gopher ASCII code art
func GopherSay(sentence string) {
	nbChar := len(sentence)

	line := " "
	for i := 0; i <= nbChar; i++ {
		line += "-"
	}

	fmt.Println(line)
	fmt.Println("< " + sentence + " >")
	fmt.Println(line)
	fmt.Println("         \\")

	fileData, err := embedGopherFiles.ReadFile("gophers/gopher.txt")
	if err != nil {
		log.Fatal("Error during read gopher ascii file", err)
	}
	fmt.Println(string(fileData))
}
