package main

import "fmt"

type User struct {
	Id             int
	Name, Location string
}

func (u *User) Greetings() string {
	return fmt.Sprintf("Hi %s from %s used Id is %d", u.Name, u.Location, u.Id)
}

type Player struct {
	User
	GameId int
}

func Greetings() {

}

func main() {
	p := &Player{
		User{
			Id:       42,
			Name:     "Matt",
			Location: "LA",
		},
		90404,
	}

	fmt.Println(p.Greetings())
}
