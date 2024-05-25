package main

import "fmt"

type authenticationInfo struct {
	username string
	password string
}

func (a authenticationInfo) getBasicAuth() string {
	// return fmt.Sprintf("Authorization: Basic %s:%s", a.username, a.password) // slower because it involves parsing the format string and processing each argument
	return "Authorization: Basic " + a.username + ":" + a.password // straightforward and efficient.
}

// don't touch below this line

func test(authInfo authenticationInfo) {
	fmt.Println(authInfo.getBasicAuth())
	fmt.Println("====================================")
}

func main() {
	test(authenticationInfo{
		username: "Google",
		password: "12345",
	})
	test(authenticationInfo{
		username: "Bing",
		password: "98765",
	})
	test(authenticationInfo{
		username: "DDG",
		password: "76921",
	})
}
