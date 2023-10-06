package main

import (
	"fmt"
	"math"
)

// 定义一个节点
type Node struct {
	State     string
	G         float64            // 从起始节点到当前节点的实际代价
	H         float64            // 当前节点到目标节点的预估代价（启发函数）
	F         float64            // 估计总代价 F = G + H
	Neighbors map[string]float64 // 相邻节点及其对应的实际代价
	Parent    *Node              // 父节点
}

// 计算启发函数
func (n *Node) CalculateH(target string) {
	n.H = float64(math.Abs(float64(len(n.State)) - float64(len(target))))
}

// 计算实际代价
func (n *Node) CalculateG(parent *Node) {
	n.G = parent.G + 1
}

// 计算总代价
func (n *Node) CalculateF() {
	n.F = n.G + n.H
}

// 获取相邻节点
func (n *Node) GetNeighbors(states []string) map[string]float64 {
	neighbors := make(map[string]float64)
	for _, state := range states {
		if len(n.State) == len(state) {
			continue
		}
		neighbors[state] = 1
	}
	return neighbors
}

// A*算法实现
func AStar(states []string, start, target string) []string {
	// 初始化起始节点
	startNode := Node{State: start, H: 0, Neighbors: make(map[string]float64)}
	startNode.CalculateH(target)
	startNode.CalculateF()

	// 创建优先队列
	pq := make(priorityQueue, 0)
	pq.Push(&startNode)

	// 访问过的节点
	visited := make(map[string]bool)

	for pq.Len() > 0 {
		// 获取优先队列中最小的节点
		currentNode := pq.Pop().(*Node)
		visited[currentNode.State] = true

		// 到达目标节点
		if currentNode.State == target {
			var path []string
			for ; currentNode != nil; currentNode = currentNode.Parent {
				path = append(path, currentNode.State)
			}
			return reverse(path)
		}

		// 获取相邻节点
		neighbors := currentNode.GetNeighbors(states)
		for state, _ := range neighbors {
			// 如果没有访问过
			if !visited[state] {
				// 计算实际代价
				neighborNode := Node{State: state, G: 0, H: 0, Neighbors: make(map[string]float64), Parent: currentNode}
				neighborNode.CalculateG(currentNode)
				neighborNode.CalculateH(target)
				neighborNode.CalculateF()

				// 将邻居节点加入优先队列
				pq.Push(&neighborNode)
			}
		}
	}

	return []string{}
}

// 反转路径
func reverse(path []string) []string {
	reversedPath := make([]string, 0)
	for i := len(path) - 1; i >= 0; i-- {
		reversedPath = append(reversedPath, path[i])
	}
	return reversedPath
}

type priorityQueue []*Node

func (pq priorityQueue) Len() int { return len(pq) }

func (pq priorityQueue) Less(i, j int) bool {
	return pq[i].F < pq[j].F
}

func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *priorityQueue) Push(x interface{}) {
	node := x.(*Node)
	*pq = append(*pq, node)
}

func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	node := old[n-1]
	*pq = old[0 : n-1]
	return node
}

func main() {
	states := []string{"hello", "world", "apple", "banana", "orange"}
	start := "hello"
	target := "world"

	path := AStar(states, start, target)
	if len(path) > 0 {
		fmt.Printf("Path from %s to %s: %v\n", start, target, path)
	} else {
		fmt.Printf("No path found from %s to %s\n", start, target)
	}
}
