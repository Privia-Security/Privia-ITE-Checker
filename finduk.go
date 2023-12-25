package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var cmd = "go run klasor.go -u %s"
var cmd2 = "go run klasor.go -u URL | grep DIRECTORY -A10 | grep -v \"STATUS CODE\" | grep -v -i url | awk '{print $3}' | tr '\\n' ' '"
var cmd3 = "go run uzanti.go -url %s -dir %s -ext %s"

var url = os.Args[1]
var ex = os.Args[2]

func main() {
	var exts string

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

