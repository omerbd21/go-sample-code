package main

import "fmt"

func main() {
	x := []string{"James", "Bond", "Shaken, not strirred"}
	y := []string{"Miss", "Moneypenny", "Helloooooooo, James"}
	xy := [][]string{x, y}
	for i := range xy {
		for j := 0; j < len(x); j++ {
			fmt.Println(xy[i][j])
		}
	}
}
