package main

import (
	"fmt"
	"os"

	"github.com/trilgon/pas_brute_forcer/bruteforcer"
)

func main() {
	var md5_password string
	passwordCh := make(chan string)

	fmt.Print("Введите md5 пароля: ")
	fmt.Fscan(os.Stdin, &md5_password)

	go bruteforcer.Force(33, 57, md5_password, passwordCh)

	go bruteforcer.Force(57, 80, md5_password, passwordCh)

	go bruteforcer.Force(80, 103, md5_password, passwordCh)

	go bruteforcer.Force(103, 126, md5_password, passwordCh)

	fmt.Println(<-passwordCh)
}
