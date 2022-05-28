package main

func newDog(string) {

}

func main() {
	m := make(map[string]func(string))
	m["dog"] = newDog
	config := "config"

	m["dog"](config)
}
