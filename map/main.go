package main

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"black": "#000000",
		"white": "#ffffff",
	}

	printColorMap((colors))

}

func printColorMap(c map[string]string) {
	for color, hex := range c {
		println("hex code for", color, "is", hex)
	}
}
