package main

import "fmt"

func main() {
	stuMap := make(map[string]map[string]string) //string -> map[string]string
	stuMap["yves"] = make(map[string]string, 3)
	stuMap["yves"]["name"] = "yves"
	stuMap["yves"]["sex"] = "male"
	stuMap["yves"]["address"] = "sjhbdalsjkhfb"

	stuMap["jyl"] = make(map[string]string, 3)
	stuMap["jyl"]["name"] = "jyl"
	stuMap["jyl"]["sex"] = "male"
	stuMap["jyl"]["address"] = "dfgsdgsdf"

	fmt.Println(stuMap["yves"]["name"])
}
