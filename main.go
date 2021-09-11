package main

func main() {
	app := App {

	}

	app.Initialize("main.db")

	app.Run(":8080")
}