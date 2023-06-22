package main

import (
	"github.com/jordinislic/repo/RepoExcV"
	"github.com/jordinislic/utilities/utilities/ClientJordi"
)

func main() {
	client := ClientJordi.New("http://openexchangerates.org/api")
	repo := RepoExcV.New("localhost", 5432, "jordinislic", "Jordi05Nislic", "jordinislic")

	param := []string{}
	param = append(param, "/latest.json")
	param = append(param, "app_id")
	param = append(param, "9c42286105114b6fb1a9ca36ebd7b63d")

	client = client.SetUrl(client.CreateNewUrlWithParams(param))
	resp := client.Get("")
	value := client.GetBodyResp(resp)
	val := RepoExcV.GetValue(value)
	repo.AddToDB(val)
}
