package consistent_test

import (
	"fmt"
	"testing"

	. "github.com/kkdai/consistent"
)

func TestBasicOp(t *testing.T) {
	ch := NewConsistentHashing()
	ch.Add("t1")
	ch.Add("t2")
	ch.Add("t3")

	fmt.Println(ch.ListNodes())
	targetObj := []string{"t1", "t2", "t3", "s1", "s2", "s3"}
	for _, v := range targetObj {
		server, err := ch.Get(v)
		if err == nil {
			fmt.Println(server)
		}
	}
	ch.Add("t4")
	ch.Add("t5")
	for _, v := range targetObj {
		server, err := ch.Get(v)
		if err == nil {
			fmt.Println(server)
		}
	}

}
