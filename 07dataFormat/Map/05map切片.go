package main

import "fmt"

func main() {
	var users []map[string]string //map切片
	users = make([]map[string]string, 2)

	if users[0] == nil {
		users[0] = make(map[string]string)
		users[0]["name"] = "yves"
		users[0]["age"] = "22"
		users[0]["sex"] = "male"
	}
	if users[1] == nil {
		users[1] = make(map[string]string)
		users[1]["name"] = "jack"
		users[1]["age"] = "43"
		users[1]["sex"] = "female"
	}

	//追加slice
	newUser := map[string]string{"name": "ella", "age": "99", "sex": "female"}
	users = append(users, newUser)

	fmt.Println(users)
}
