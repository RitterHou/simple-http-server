package main

import (
	"flag"
	"fmt"
	"time"
)

type Flag struct {
	root   string
	help   bool
	upload bool
	ip     string
	port   uint16
}

var customFlag Flag

func init() {
	root := flag.String("root", getRoot(), "Root directory")
	help := flag.Bool("help", false, "Print help message")
	upload := flag.Bool("upload", false, "Enable upload files")
	ip := flag.String("ip", "0.0.0.0", "IP address to bind [default: 0.0.0.0]")
	port := flag.Uint("port", 9257, "Port number [default: 9257]")

	customFlag = Flag{
		root:   *root,
		help:   *help,
		upload: *upload,
		ip:     *ip,
		port:   uint16(*port),
	}

	prompt := ""
	prompt += fmt.Sprintf("   Root: %s\n", customFlag.root)
	prompt += fmt.Sprintf(" Upload: %t\n", customFlag.upload)
	prompt += fmt.Sprintf("Address: http://%s:%d/\n", customFlag.ip, customFlag.port)
	if customFlag.help {
		prompt += "Help Message\n"
	}

	prompt += fmt.Sprintf("======== [%s] ========", time.Now().Format("2006-01-02 15:04:05"))
	fmt.Println(prompt)
}

func main() {

}
