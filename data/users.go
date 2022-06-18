package data

import (
	"math/rand"

	"github.com/samber/lo"
)

type User struct {
	Id        string             `json:"id"`
	Portfolio map[string]float32 `json:"portfolio"`
}

var Users []User

func init() {
	Users = []User{}
}

func makeUser(id string) User {
	pfSize := rand.Intn(30) + 4
	portfolio := make(map[string]float32)
	symbol := ""
	for i := 0; i < pfSize; i++ {

		for found := true; found; {
			symbol = Currencies[rand.Intn(len(Currencies))]
			_, found = portfolio[symbol]
		}
		portfolio[symbol] = rand.Float32() * 1000000
	}

	user := User{Id: id, Portfolio: portfolio}
	return user
}

func GetUser(id string) User {
	user, ok := lo.Find(Users, func(u User) bool {
		return u.Id == id
	})
	if !ok {
		user = makeUser(id)
		Users = append(Users, user)
	}
	return user
}
