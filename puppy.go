package puppy

import (
	"fmt"

	"github.com/LilechkaYa/dog"
)

func Bark() string {
	return "Woof!"
}

func Barks() string {
	return "Woof! Woof! Woooof!!"
}

func BigBark() string {
	return dog.WhenGrownUp(Bark())
}

func BigBarks() string {
	return dog.WhenGrownUp(Barks())
}

func from110() {
	fmt.Println("I am from v1.1.0")
}

func from120() {
	fmt.Println("I am from v1.2.0")
}
