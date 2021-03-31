package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func gerarSenha() {
	password := "yivLTC12"
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	fmt.Println(string(bytes))
}
