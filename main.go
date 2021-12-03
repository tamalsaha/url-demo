package main

import (
	"fmt"
	"net/url"
)

func main() {
	var u url.URL
	u.Scheme = "https" // "postgresql"
	//u.User = url.UserPassword("{{username}}", "{{password}}")
	u.Host = "127.0.0.1:5432"
	u.Path = "/"
	q := url.Values{}
	q.Add("name", "my mg")
	q.Add("namespace", "demo & default")
	u.RawQuery = q.Encode()

	fmt.Println(u.String())

	//
	//us := u.String()
	//
	//i := strings.Index(us, "://")
	//
	//fmt.Println(us[:i+3] + "{{username}}:{{password}}@" + us[i+3:])
}
