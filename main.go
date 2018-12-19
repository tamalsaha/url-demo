package main

import (
	"fmt"
	"log"
	"net/url"
	"strings"
)

func main() {
	u, err := url.Parse("")
	if err != nil {
		log.Fatalln(err)
	}
	u.Scheme = "postgresql"
	//u.User = url.UserPassword("{{username}}", "{{password}}")
	u.Host = "127.0.0.1:5432"
	u.Path = "/"
	u.RawQuery = "sslmode=false"
	fmt.Println(u.String())

	us := u.String()

	i := strings.Index(us, "://")

	fmt.Println(us[:i+3]+ "{{username}}:{{password}}@" + us[i+3:])
}
