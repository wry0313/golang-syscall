package main

import (
	"syscall"
)

func main() {
	// execute a syscall for opening a file descriptor
	fd, err := syscall.Open("./hello", syscall.O_RDONLY, 0)
	if err != nil {
		panic(err)
	}
	// read the file descriptor
    buf := make([]byte, 1024) // Adjust the size as needed
	_, err = syscall.Read(fd, buf)
	if err != nil {
		panic(err)
	}
	// print the content
	println(string(buf))
	// execute a syscall for closing a file descriptor
	err = syscall.Close(fd)
	if err != nil {
		panic(err)
	}

}