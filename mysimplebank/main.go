package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/webvillain/mysimplebank/config"
)

func main() {
	// remove pre existing file if any
	os.Remove("./Users.db")
	db, err := config.ConnectDb()
	if err != nil {
		log.Fatal("Not Connected To Database !")
	}
	fmt.Println("Successfully Connected To Database.")
	defer db.Close()
	http.HandleFunc("/mybank", Userhandler)
	// log.Fatal(http.ListenAndServe(":8080", nil))
	// fmt.Println("Your Application Is Started On LocalMahine At Port :8080/mybank")
	UserChoices()

}

func Userhandler(w http.ResponseWriter, r *http.Request) {

}

func UserChoices() {
	var opts int
	fmt.Println("Choose From Below Operations  : ")
	fmt.Println(" Type `1` And Press Enter To Get-All Users . ")
	fmt.Println(" Type `2` And Press Enter To Get Single User . ")
	fmt.Println(" Type `3` And Press Enter To Create New User .")
	fmt.Println(" Type `4` And Press Enter To Update User .")
	fmt.Println(" Type `5` And Press Enter To Delete User .")

	fmt.Scanf("%d", &opts)

	switch opts {
	case 1:
		fmt.Println("hi , you opt for GET ALL USERS INFO ")

	case 2:
		fmt.Println("hi , you opt for EGT SINGEL user INFO ")

	case 3:
		fmt.Println("hi , you opt for CREATE NEW user ")

	case 4:
		fmt.Println("hi , you opt for UPDATE EXISTING user ")

	case 5:
		fmt.Println("hi , you opt for DELETE user  ")
	}
}
