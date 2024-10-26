package puppy

import (
	bigdog "github.com/KamilWlodarczyk2003/go_test_rep_2"
)

func main() {

}

func Bark() string {
	return "Woof!"
}

func Barks() string {
	return "Woof! Woof! Woof!"
}

func BigBark() string {
	return bigdog.WhenGrownUp(Barks())
}
