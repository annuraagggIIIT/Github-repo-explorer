package main

import(
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Repository  struct {
	Name string `json:"name"`
	Description string `json:"description"`
	stars int `json:"stargazers_count"`
	forks int `json:"forks_count"`
}

func getRepositories(username string) ([]Repository, error) {
	// END: ed8c6549bwf9
	url := fmt.Sprintf("https://api.github.com/users/%s/repos", username)
	response , err := http.Get(url)
	if err!=nil {
		fmt.Println("Error while fetching the data")
		return nil , err
	}
	defer response.Body.Close()

	Body, err := ioutil.ReadAll(response.Body)
	if err!=nil	{
		fmt.Println("Error while reading the data")
		return nil ,err
	}
	
	var repos []Repository
	err= json.Unmarshal(Body, &repos)
	if err!=nil	{
		fmt.Println("Error while unmarshalling the data")
		return nil ,err
	}	

	return repos , nil
}



func main () {
	fmt.Println("Enter the Username")
	var username string
	fmt.Scanln(&username)
	repos ,err :=getRepositories(username)
	if err!=nil	{
		fmt.Println("Error while fetching the data")
		return
	}

	fmt.Println("The Repositories for %s are :",username)
	for _ , repo := range repos	{
		fmt.Println(repo.Name)
		fmt.Println(repo.Description)
		fmt.Println(repo.stars)
		fmt.Println(repo.forks)
	}

}