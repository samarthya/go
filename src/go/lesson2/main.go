package main

import (
	"fmt"
	lesson1 "go/lesson1"
	lesson3 "go/lesson3"

	"time"

	"github.com/inancgumus/myhttp"
)

type number struct {
	f float32
}

type nr number        // new distinct type
type nrAlias = number // alias type

func main() {
	fmt.Printf(" Welcome to Lesson 2!\n")
	fmt.Printf(" Version: %s\n", lesson1.HelloVersion())
	fmt.Printf(" %T : %f", lesson3.NewVersion(), lesson3.NewVersion())
	mh := myhttp.New(time.Second)
	response, _ := mh.Get("https://jsonip.com/")
	fmt.Println("\nHTTP status code: ", response.StatusCode)

}
