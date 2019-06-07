package main

import (
	"bufio"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"os"
	"os/exec"
)

func main() {
    //reader
    reader := bufio.NewReader(os.Stdin)
    var pics []string
    var url = "https://www.google.com/search?q=duck+pictures&rlz=1C5GCEA_enUS852US852&source=lnms&tbm=isch&sa=X&ved=0ahUKEwiDj5yx3NXiAhUJVN8KHTo1CM0Q_AUIECgB&biw=1210&bih=798"

    //ask for animal
	fmt.Println("Enter your favorite animal")
	animal , _ := reader.ReadString('\n')
	fmt.Printf("Incorrect. Ducks are better. Ducks > %s ", animal)

	i := len(animal)

    //make request
    response, err := http.Get(url)
    if err != nil {
    	log.Fatal(err)
	}

    defer response.Body.Close()

	page, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal("Error loading HTTP response body. ", err)
	}

	page.Find("img").Each(func(index int, element *goquery.Selection){
		imgSrc, exists := element.Attr("src")
		if exists {
			pics = append(pics, imgSrc)
		}
	})

	cmd := exec.Command("open", pics[i])
	_ = cmd.Run()
}
