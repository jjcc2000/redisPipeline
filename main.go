package main

import (
	"bufio"
	"context"
	"fmt"
	"os"

	
	"github.com/redis/go-redis/v9"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
		Protocol: 2,
	})

	ctx := context.Background()

	testPing(client)

	SetAndGetValues(client, ctx)

}

func testPing(client *redis.Client) error {

	ctx := context.Background()
	fmt.Println(client.Ping(ctx))

	info, err := client.ClientInfo(ctx).Result()
	if err != nil {
		return fmt.Errorf("method ClientInfo failed: %w", err)
	}

	fmt.Printf("%#v\n", info)

	return nil
}

func SetAndGetValues(client *redis.Client, ctx context.Context) {
	sc := bufio.NewScanner(os.Stdin)

	fmt.Println("Set the title: ")
	sc.Scan()
	title := sc.Text()

	fmt.Println("Set the value: ")
	sc.Scan()
	value := sc.Text()

	fmt.Println(title, value)
	err := client.Set(ctx, title, value, 0).Err()

	if err != nil {
		fmt.Println("Error here!")
		panic(err)
	}

	val, err := client.Get(ctx, title).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Result", val)
}
