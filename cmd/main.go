package main

import "gitlab.com/ricardo134/link-service/boot"

func main() {
	boot.LoadEnv()
	boot.LoadDb()
	boot.LoadServices()

	boot.ServeHTTP()
}
