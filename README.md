Consistent: Consistent Hashing implement in Golang
==============

[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/kkdai/consistent/master/LICENSE)  [![GoDoc](https://godoc.org/github.com/kkdai/consistent?status.svg)](https://godoc.org/github.com/kkdai/consistent)  [![Build Status](https://travis-ci.org/kkdai/consistent.svg?branch=master)](https://travis-ci.org/kkdai/consistent)

![](http://blog.codinglabs.org/uploads/pictures/consistent-hashing/6.png)

What is this "Consistent Hashing"
=============

Consistent hashing is a special kind of hashing such that when a hash table is resized and consistent hashing is used, only K/n keys need to be remapped on average, where K is the number of keys, and n is the number of slots. In contrast, in most traditional hash tables, a change in the number of array slots causes nearly all keys to be remapped. (cited from [Wiki](https://en.wikipedia.org/wiki/Consistent_hashing))

Explanation for how it work
=============

Consistent Hashing has 4 features:

- Balance
- Monotonicity
- Spread
- Load

In this application you will see:

#### Balance

We use virtual node to make replica for load balance. (not all node pick the same server).

#### Monotonicity

Use the hash key function, it will make sure it always pick the same server (in the same quantity of server)


Here is more [detail](http://blog.csdn.net/cywosp/article/details/23397179) in Chinese
 

Installation and Usage
=============


Install
---------------

    go get github.com/kkdai/consistent


Usage
---------------

Following is sample code:


```go

package main

import (
	"fmt"
    "github.com/kkdai/consistent"
)

func main() {
	ch := NewConsistentHashing()
	ch.Add("server1")
	ch.Add("server2")
	ch.Add("server3")
//[server1 server2 server3]

	fmt.Println(ch.ListNodes())
	targetObj := []string{"client1", "client2", "client3", "client4", "client5", "client6"}
	for _, v := range targetObj {
		server, err := ch.Get(v)
		if err == nil {
			fmt.Printf("client: %s in server: %s \n", v, server)
		}
	}
//client: client1 in server: server1
//client: client2 in server: server3
//client: client3 in server: server1
//client: client4 in server: server3
//client: client5 in server: server1
//client: client6 in server: server3

	ch.Add("server4")
	ch.Add("server5")
	for _, v := range targetObj {
		server, err := ch.Get(v)
		if err == nil {
			fmt.Printf("client: %s in server: %s \n", v, server)
		}
	}

//client: client1 in server: server4
//client: client2 in server: server3
//client: client3 in server: server4
//client: client4 in server: server3
//client: client5 in server: server4
//client: client6 in server: server3

}
```

Inspired By
=============

- [每天进步一点点——五分钟理解一致性哈希算法(consistent hashing)](http://blog.csdn.net/cywosp/article/details/23397179)
- [https://github.com/stathat/consistent](https://github.com/stathat/consistent)
- [Wiki: Consistent Hashing](https://en.wikipedia.org/wiki/Consistent_hashing)


Project52
---------------

It is one of my [project 52](https://github.com/kkdai/project52).


License
---------------

This package is licensed under MIT license. See LICENSE for details.
