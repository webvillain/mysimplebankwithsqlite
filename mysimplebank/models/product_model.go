// so fincally all of my 5 functions are ready to interact with db and my applicaation

package models

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/webvillain/mysimplebank/entities"
)

type Database struct {
	Db *sql.DB
}

//   ProductModel = UserModel
//  Enitities.Product = eEtities.User

// find all users in our table
func (database Database) FindAll() ([]entities.User, error) {
	var user entities.User
	var users []entities.User
	FindAllUsersQuery := `SELECT * FROM USERS;`
	rows, err := database.Db.Query(FindAllUsersQuery)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		rows.Scan(&user.Id, &user.Name, &user.Email)
	}
	users = append(users, user)
	return users, nil
}

// find single user in our DATABASE

// i think i need to work more on this because we are quering a single row means that i want a single user info
// so if a want to pass a parameter in this query then it must be   db.prepare not db.query
func (database Database) FindOne() (entities.User, error) {

	// 1st attampt
	// FindOneQuery := `SELECT * FROM USERS WHERE ID = ?`
	// var user entities.User
	// row, err := database.Db.Query(FindOneQuery)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// for row.Next() {
	// 	row.Scan(&user.Id, &user.Name, &user.Email)
	// }
	// return user, nil

	// 2nd attampt
	// FindOneQuery := `SELECT * FROM USERS WHERE ID = ?`
	// var id int
	// fmt.Println("Enter Id Of User Of Which You Want To Get Data : ")
	// fmt.Scanf("%d", id)
	// var user entities.User
	// stmt, err := database.Db.Prepare(FindOneQuery)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// res, err = stmt.Exec()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// re

	// for row.Next() {
	// 	row.Scan(&user.Id, &user.Name, &user.Email)
	// }
	// return user, nil
	var id int
	fmt.Println("Enter User Id To Fatch Data : ")
	fmt.Scanf("%d", id)
	FindOneQuery := `SELECT * FROM USERS WHERE Id = ?`
	var user entities.User
	row, err := database.Db.Query(FindOneQuery, id)
	if err != nil {
		log.Fatal(err)
	}
	for row.Next() {
		row.Scan(&user.Id, &user.Name, &user.Email)
	}
	return user, nil

}

func (database Database) CreateNewUser() {
	var name, email string
	CreateNewUserQuery := `INSERT INTO USERS(Name, Email)VALES(?,?);`
	stmt, err := database.Db.Prepare(CreateNewUserQuery)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("To Create A New User In Table Press Enter And Fill The Details As Per Message Displayed On Your Screen ........")
	fmt.Println("Please Enter` Name` and Then Press Enter : ")
	fmt.Scanf("%s\n", &name)
	fmt.Println("Please Enter `Email` and Then Press Enter: ")
	fmt.Scanf("%s\n", &email)
	res, err := stmt.Exec(name, email)
	if err != nil {
		log.Print(err)
	}
	fmt.Println("..............oohhh....yeeeeaaahhhhh............----\nNew User Created Successfull")
	n, _ := res.RowsAffected()
	fmt.Printf("No. Of Rows Affected: %d", n)
}

func (database Database) UpdateUser() {
	//1st attampt
	var opts int
	var name, email string
	fmt.Println("Enter User Id Which You Want To Delete From Table : ")
	fmt.Scanf("%d", &opts)
	fmt.Println("Enter Updated User Name : ")
	fmt.Scanf("%g", &name)
	fmt.Println("Enter Update User Email : ")
	fmt.Scanf("%g", &email)
	updateuserquery := `UPDATE table
	SET Name = ?,Email = ?,
	WHERE Id = ?;
	`
	stmt, err := database.Db.Prepare(updateuserquery)
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec(name, email, opts)
	if err != nil {
		log.Fatal(err)
	}

	n, _ := res.RowsAffected()
	fmt.Println("No. Of Rows Affected:", n)
	fmt.Println("User Is Successfully Updated From Table `USERS`")

}

func (database Database) DeleteUser() {
	// 1st attampt
	// var opts int
	// fmt.Println("Enter User Id Which You Want To Delete From Table : ")
	// fmt.Scanf("%d", &opts)
	// deleteuserquery := `DELETE FROM table
	// WHERE Id = ?;`
	// rows, err := database.Db.Query(deleteuserquery, opts)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer rows.Close()
	// fmt.Println("Credentials Are Deleted From Table")

	//2nd attampt
	var opts int
	fmt.Println("Enter User Id Which You Want To Delete From Table : ")
	fmt.Scanf("%d", &opts)
	deleteuserquery := `DELETE FROM table
	WHERE Id = ?;`
	stmt, err := database.Db.Prepare(deleteuserquery)
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec(opts)
	if err != nil {
		log.Fatal(err)
	}

	n, _ := res.RowsAffected()
	fmt.Println("No. Of Rows Affected:", n)
	fmt.Println("User Is Successfully Deleted From Table `USERS`")

}
