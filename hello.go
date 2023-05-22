package sayprinter

// #cgo CFLAGS: -Ihello -Igoodbye
// #cgo LDFLAGS: -static -L. -lstdc++ -lgoodbye
// #include "hello.h"
// #include "goodbye.h"
import "C"

func SayHello() {
	C.say_hello()
}

func SayWorld() {
	C.say_world()
}

func SayGoodbye() {
	C.say_goodbye()
}
