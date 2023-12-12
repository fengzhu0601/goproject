package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type Movie struct {
	Title         string  `xml:"title"`
	OriginalTitle string  `xml:"originaltitle"`
	Plot          string  `xml:"plot"`
	Genre         string  `xml:"genre"`
	Director      string  `xml:"director"`
	Year          string  `xml:"year"`
	Runtime       string  `xml:"runtime"`
	Actors        []Actor `xml:"actor"`
}

type Actor struct {
	Name  string `xml:"name"`
	Image string `xml:"image"`
}

func main() {
	movie := Movie{
		Title:         "Movie Title",
		OriginalTitle: "Original Title",
		Plot:          "This is the plot",
		Genre:         "Action",
		Director:      "Director Name",
		Year:          "2023",
		Runtime:       "120 mins",
		Actors: []Actor{
			{
				Name:  "Actor 1",
				Image: "image1.jpg",
			},
			{
				Name:  "Actor 2",
				Image: "image2.jpg",
			},
		},
	}

	// 编码为XML
	xmlData, err := xml.MarshalIndent(movie, "", "    ")
	if err != nil {
		fmt.Println("编码XML时出错:", err)
		return
	}
	fmt.Println("编码后的XML数据:")
	os.Stdout.Write(xmlData)
	fmt.Println()

	// 解码XML
	var decodedMovie Movie
	if err := xml.Unmarshal(xmlData, &decodedMovie); err != nil {
		fmt.Println("解码XML时出错:", err)
		return
	}
	fmt.Println("从XML解码得到的Movie实例:")
	fmt.Printf("%+v\n", decodedMovie)
}
