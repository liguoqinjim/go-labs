package main

import (
	"fmt"

	gojsonq "github.com/thedevsaddam/gojsonq/v2"
)

func main() {
	const jsonStr = `{"movies":[{"name":"Pirates of the Caribbean","sequels":[{"name":"The Curse of the Black Pearl","released":2003},{"name":"Dead Men Tell No Tales","released":2017}]}]}`
	jq := gojsonq.New().JSONString(jsonStr)

	name := jq.Find("movies.[0].name")
	fmt.Println("name: ", name) // name:  Pirates of the Caribbean

	fsequel := jq.Reset().Find("movies.[0].sequels.[0].name")
	fmt.Println("first sequel: ", fsequel) // first sequel:  The Curse of the Black Pearl
}
