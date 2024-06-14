package main

import (
	"fmt"
	"tokoku/config"
)

func main() {
	setup := config.InportSetting();
	connection, err := config.ConnectDB(setup)

	if err != nil {
		fmt.Println("Failed connect to database", err.Error());
		return;
	}


	fmt.Println(connection)
}