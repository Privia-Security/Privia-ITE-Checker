package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var cmd = "go run folder.go -u %s"
var cmd2 = "go run folder.go -u URL | grep DIRECTORY -A10 | grep -v \"STATUS CODE\" | grep -v -i url | awk '{print $3}' | tr '\\n' ' '"
var cmd3 = "go run extension.go -url %s -dir %s -ext %s"



func main() {
	var exts string

	/*if len(os.Args) < 2 {
		fmt.Println("Usage: <full URL> <extensions>\n");
		fmt.Println("Example:\n");
		fmt.Println("finduk http://192.168.112.136/ txt,png\n");
		return
	}*/

	var url = os.Args[1]

	if url== "-help"{
		fmt.Println("Usage: <full URL> <extensions>\n");
		fmt.Println("Example:");
		fmt.Println("\tfinduk http://192.168.112.136/ txt,png\n");
		return
	}

	var ex = os.Args[2]

	out, err := exec.Command("sh", "-c", fmt.Sprintf(cmd, url)).Output()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(out))

	cmd23 := strings.Replace(cmd2, "URL", url, -1)
	out2, err2 := exec.Command("sh", "-c", cmd23).Output()
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	parameters := strings.TrimSpace(string(out2))
	parameters = strings.Replace(parameters, " ", ",", -1)

	tmp_exts := strings.Split(ex, ",")

	exts = strings.Join(tmp_exts, ",")

	out3, err3 := exec.Command("sh", "-c", fmt.Sprintf(cmd3, url, parameters, exts)).Output()
	if err3 != nil {
		fmt.Println(err3)
		return
	}
	fmt.Println(string(out3))

}

