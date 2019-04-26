package main

import (

"fmt"
"net/http"
"github.com/PuerkitoBio/goquery"
"github.com/mgutz/ansi"
"time"

)

func main(){
title := ansi.Color("Dork-Search By Fajar Firdaus", "green+b:white")
b1 := ansi.Color("\n[=======================]\n", "blue")
b2 := ansi.Color("Coder : Fajar Firdaus\n", "green")
b3 := ansi.Color("Fb : Fajar Firdaus / Ace.of.spades729\n", "green")
b4 := ansi.Color("IG : fajar_firdaus_7\n", "green")
b5 := ansi.Color("YT : iTech7732\n", "green")
b6 := ansi.Color("[=======================]\n", "blue")
tanda1 := ansi.Color("[", "green")
tanda2 := ansi.Color("]", "green")
input2 := ansi.Color("[ Input ] > ", "white+b:blue")
result := ansi.Color("[ Result ]", "green+b:blue")

fmt.Print(title)
fmt.Print(b1)
fmt.Print(b2)
fmt.Print(b3)
fmt.Print(b4)
fmt.Print(b5)
fmt.Print(b6)

var waktu = time.Now()

var in string
fmt.Print(tanda1, waktu.Hour(), waktu.Minute(), waktu.Second(), tanda2, " ", input2)
fmt.Scan(&in)
test, error := http.Get("http://www1.search-results.com/web?q="+ in);
if error != nil{
	fmt.Print("Error Gan")
}
defer test.Body.Close()

dom, err := goquery.NewDocumentFromReader(test.Body)
if err != nil{
	fmt.Print("[!] Error")
}

dom.Find(".algo-title").Each(func(index int, element *goquery.Selection){
	hasil, tampil := element.Attr("href")
	if tampil {
		fmt.Print(result)
	fmt.Print(hasil + "\n")
}
})
}
