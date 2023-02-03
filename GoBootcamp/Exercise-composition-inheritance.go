package main

import "fmt"

type User struct {
	Id             int
	Name, Location string
}

func (u *User) Greetings() string {
	return fmt.Sprintf("Hi %s from %s", u.Name, u.Location)
}

type Player struct {
	*User
	GameId int
}

func newPlayer(id int, name, location string, gameId int) *Player {
	return &Player{
		User: &User{
			Id:       id,
			Name:     name,
			Location: location,
		},
		GameId: gameId,
	}
}

func main() {
	p := newPlayer(123, "Phước", "CÀm bố đi ờ", 123213)
	fmt.Println(p.User)
	fmt.Println(p.Greetings())
}
