package main

import (
	"errors"
	"fmt"
)

func main() {

	user, err := login("Cyberworm", "123456")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Hoş geldiniz", user)
	}
}

func login(username, password string) (string, error) {

	for i := 0; i < 3; i++ {
		fmt.Print("Kullanıcı adı: ")
		fmt.Scanln(&username)

		fmt.Print("Parola: ")
		fmt.Scanln(&password)

		if username == "Cyberworm" && password == "123456" {
			return "Giriş başarılı", nil
		} else {
			fmt.Println("Kullanıcı adı veya parola yanlış.")
		}
	}
	return "", errors.New("3 başarısız deneme. Hesabınız geçici bir süreliğine kısıtlandı.")
}
