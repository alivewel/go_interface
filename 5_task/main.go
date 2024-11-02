package main

func main() {
	var ptr *struct{}
	var iface interface{}
	iface = ptr
	if iface == nil {
		println("It's nil!")
	}
} 
