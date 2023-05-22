package sayprinter

// #include "hello.h"
import "C"

func SayHello() {
	C.say_hello()
}

func SayWorld() {
	C.say_world()
}
