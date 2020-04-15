package main

var biggestSize int32
var group, size map[int32]int32

func MaxCircle(queries [][]int32) (result []int32) {
	initialize()
	for _, q := range queries {
		registerIfNew(q[0])
		registerIfNew(q[1])
		makeFriend(q[0], q[1])
		result = append(result, biggestSize)
	}
	return result
}

func initialize() {
	biggestSize = 0
	group = make(map[int32]int32)
	size = make(map[int32]int32)
}

func registerIfNew(person int32) {
	if group[person] == 0 {
		group[person] = person
		size[person] = 1
	}
}

func isFriend(person1, person2 int32) bool {
	return root(person1) == root(person2)
}

func root(person int32) int32 {
	for group[person] != person {
		person = group[person]
	}
	return person
}

func makeFriend(person1, person2 int32) {
	if isFriend(person1, person2) {
		return
	}
	root1 := root(person1)
	root2 := root(person2)
	if size[root1] < size[root2] {
		union(root1, root2)
	} else {
		union(root2, root1)
	}
}

func union(rootA, rootB int32) {
	group[rootA] = rootB
	size[rootB] += size[rootA]
	if size[rootB] > biggestSize {
		biggestSize = size[rootB]
	}
}
