package main

import (

"fmt"
"net/http"
"io/ioutil"
"io"
"os"

)

func main(){
test, error := http.Get("http://www1.search-results.com/web?q=asw");
if error != nil{
	fmt.Print("Error Gan")
}
defer test.Body.Close()

data, err := os.Create("trash.html")
if err != nil {
	fmt.Print("no data")
}

_, err = io.Copy(data, test.Body)
if err != nil {
	fmt.Print("Invalid Data")
}

bukafile, err := os.Open("trash.html")
if err != nil {
	fmt.Print("")
}

read, err := ioutil.ReadAll(bukafile)
if err != nil {
	fmt.Print("")
}

fmt.Print(string(read))

os.Remove("trash.html")

}
