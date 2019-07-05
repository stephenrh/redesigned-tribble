package main

import "selfscale/users/app"

func main() {
	app.NewService().Run(":8080")
}
