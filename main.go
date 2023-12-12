package main

import (
	"golang.org/x/sys/unix"
)

func main() {
	// execute a syscall for opening a file descriptor
	fd, err := unix.Open("./hello", unix.O_RDONLY, 0)
	if err != nil {
		panic(err)
	}
	// read the file descriptor
    buf := make([]byte, 1024) // Adjust the size as needed
	_, err = unix.Read(fd, buf)
	if err != nil {
		panic(err)
	}
	// print the content
	println(string(buf))
	// execute a syscall for closing a file descriptor
	err = unix.Close(fd)
	if err != nil {
		panic(err)
	}


}