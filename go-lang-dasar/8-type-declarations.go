package main

import "fmt"

func main() {
	type NoKTP string

	var noKtpPaldi NoKTP = "123321123"
	fmt.Println(noKtpPaldi)
}
