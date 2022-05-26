package main

import "fmt"

func calculate(amount int) int {

	welcomeBonus := 5
	oneBonus := 5
	bonus := (amount / oneBonus) + welcomeBonus
	return bonus
}
func main() {
	fmt.Println(calculate(666))

}
