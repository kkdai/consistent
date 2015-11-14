package consistent_test

import (
	"fmt"
	"testing"

	. "github.com/kkdai/consistent"
)

func TestBasicOp(t *testing.T) {
	ch := NewConsistentHashing()
	ch.Add("server1")
	ch.Add("server2")
	ch.Add("server3")

	fmt.Println(ch.ListNodes())
	targetObj := []string{"client1", "client2", "client3", "client4", "client5", "client6"}
	for _, v := range targetObj {
		server, err := ch.Get(v)
		if err == nil {
			fmt.Printf("client: %s in server: %s \n", v, server)
		}
	}

	fmt.Println("----")
	ch.Add("server4")
	ch.Add("server5")
	for _, v := range targetObj {
		server, err := ch.Get(v)
		if err == nil {
			fmt.Printf("client: %s in server: %s \n", v, server)
		}
	}

	fmt.Println()

}
