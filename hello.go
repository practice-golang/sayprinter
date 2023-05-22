package sayprinter

// #cgo CFLAGS: -Ihello -Igoodbye
// #cgo LDFLAGS: -static -L. -lstdc++ -lgoodbye
// #include "hello.h"
// #include "goodbye.h"
// #include "hello/hello.c"
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
