package core

import (
	"fmt"
)

func PrintRoutes() {
	fmt.Println()
	fmt.Println("---------------------------Mcache v1.0.0--------------------------------")
	fmt.Println("-------------------Light weight | High performace-----------------------")
	fmt.Println()
	fmt.Println("Available routes:")
	fmt.Println("\n------------------------------------------------------------------------")
	fmt.Println("| Method | Route                           | Description                 |")
	fmt.Println("|--------|---------------------------------|-----------------------------|")
	fmt.Println("| GET    | /keys                           | Retrieve all keys of cache  |")
	fmt.Println("| GET    | /get?key={key}                  | Retrieve value for given key|")
	fmt.Println("| POST   | /set?key={key}&value={value}    | Add new key-value to cache  |")
	fmt.Println("| DELETE | /delete?key={key}               | Delete key from cache       |")
	fmt.Println("--------------------------------------------------------------------------")
}
