package consistent

import (
	"errors"
	"hash"
	"hash/fnv"
	"sort"
	"strconv"
)

type SortedKeys []uint32

func (x SortedKeys) Len() int           { return len(x) }
func (x SortedKeys) Less(i, j int) bool { return x[i] < x[j] }
func (x SortedKeys) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

type ConsistentHashing struct {
	NumOfVirtualNode int

	hashSortedKeys SortedKeys
	h              hash.Hash32

	circleRing map[uint32]string
	dataSet    map[string]bool
}

func NewConsistentHashing() *ConsistentHashing {
	newCH := &ConsistentHashing{h: fnv.New32(), NumOfVirtualNode: 20}
	newCH.circleRing = make(map[uint32]string)
	newCH.dataSet = make(map[string]bool)
	return newCH
}

func (c *ConsistentHashing) getVirtualNodeKey(index int, obj string) uint32 {
	newObjStr := strconv.Itoa(index) + "-" + obj
	return c.hasKey(newObjStr)
}

// Add a node into this consistent hashing ring
func (c *ConsistentHashing) Add(node string) {
	if _, find := c.dataSet[node]; find {
		return
	}

	c.dataSet[node] = true
	key := c.hasKey(node)
	c.circleRing[key] = node

	//Add virtual node for "balance"
	for i := 0; i < c.NumOfVirtualNode; i++ {
		vk := c.getVirtualNodeKey(i, node)
		c.circleRing[vk] = node
	}

	c.updateSortHashKeys()
}

// Remove a node from this consistent hashing ring
func (c *ConsistentHashing) Remove(node string) {
	if _, find := c.dataSet[node]; !find {
		return //not in our dataset
	}

	delete(c.dataSet, node)
	key := c.hasKey(node)
	delete(c.circleRing, key)

	//Delete virtual node
	for i := 0; i < c.NumOfVirtualNode; i++ {
		vk := c.getVirtualNodeKey(i, node)
		delete(c.circleRing, vk)
	}

	c.updateSortHashKeys()
}

func (c *ConsistentHashing) searchNearRingIndex(obj string) int {
	targetKey := c.hasKey(obj)

	targetIndex := sort.Search(len(c.hashSortedKeys), func(i int) bool { return c.hashSortedKeys[i] >= targetKey })

	//fmt.Println("key=", targetKey, " index=", targetIndex)
	if targetIndex >= len(c.hashSortedKeys) {
		targetIndex = 0
	}

	return targetIndex
}

func (c *ConsistentHashing) updateSortHashKeys() {
	c.hashSortedKeys = nil

	for node, _ := range c.dataSet {
		key := c.hasKey(node)
		c.hashSortedKeys = append(c.hashSortedKeys, key)
	}
	sort.Sort(c.hashSortedKeys)
}

// Get a nearest object name from input object in consistent hashing ring
func (c *ConsistentHashing) Get(obj string) (string, error) {
	if len(c.dataSet) == 0 {
		return "", errors.New("Empty struct")
	}

	nearObj, _ := c.circleRing[c.hashSortedKeys[c.searchNearRingIndex(obj)]]
	//fmt.Println("Get:", nearObj, " size circle=", len(c.circleRing), " ring=", c.circleRing)
	return nearObj, nil
}

// List the whole nodes in consistent hashing ring
func (c *ConsistentHashing) ListNodes() []string {
	var retList []string
	for k, _ := range c.dataSet {
		retList = append(retList, k)
	}
	return retList
}

func (c *ConsistentHashing) hasKey(obj string) uint32 {
	data := []byte(obj)
	c.h.Reset()
	c.h.Write(data)
	return c.h.Sum32()
}
