package main

import (
	"context"
	"log"
)

func step1(ctx context.Context) context.Context {
	child := context.WithValue(ctx, "name", "enbin") // create context with key and value
	return child
}

func step2(ctx context.Context) context.Context {
	child := context.WithValue(ctx, "age", 23)
	return child
}

func step3(ctx context.Context) {
	log.Printf("name: %s\n", ctx.Value("name"))
	log.Printf("age: %d\n", ctx.Value("age"))
}

func main() {
	grandpa := context.TODO()
	father := step1(grandpa)
	grandson := step2(father)
	step3(grandson)
}
