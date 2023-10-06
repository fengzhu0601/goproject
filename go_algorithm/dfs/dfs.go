package main

import (
	"fmt"
)

type Graph struct {
	Nodes []*Node
}

type Node struct {
	Value     int
	Neighbors []*Node
	Visited   bool
}

func NewGraph() *Graph {
	return &Graph{
		Nodes: make([]*Node, 0),
	}
}

func (g *Graph) AddNode(value int) *Node {
	node := &Node{
		Value:     value,
		Neighbors: make([]*Node, 0),
		Visited:   false,
	}
	g.Nodes = append(g.Nodes, node)
	return node
}

func (n *Node) AddNeighbor(neighbor *Node) {
	n.Neighbors = append(n.Neighbors, neighbor)
}

func DFS(node *Node) {
	if node == nil {
		return
	}

	fmt.Println(node.Value) // 输出节点值
	node.Visited = true

	for _, neighbor := range node.Neighbors {
		if !neighbor.Visited {
			DFS(neighbor)
		}
	}
}

func main() {
	graph := NewGraph()

	// 添加节点
	node1 := graph.AddNode(1)
	node2 := graph.AddNode(2)
	node3 := graph.AddNode(3)
	node4 := graph.AddNode(4)
	node5 := graph.AddNode(5)

	// 添加边
	node1.AddNeighbor(node2)
	node2.AddNeighbor(node1)
	node1.AddNeighbor(node3)
	node3.AddNeighbor(node1)
	node4.AddNeighbor(node5)
	node5.AddNeighbor(node4)

	for _, node := range graph.Nodes {
		if !node.Visited {
			DFS(node)
		}
	}
}
