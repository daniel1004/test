package main

import (
	"container/heap"
	"fmt"
)

// Двоичное дерево
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Вставка элемента
func insertNode(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	if val < root.Val {
		root.Left = insertNode(root.Left, val)
	} else {
		root.Right = insertNode(root.Right, val)
	}
	return root
}

// Печать элементов дерева
func inorderTraversal(root *TreeNode) {
	if root != nil {
		inorderTraversal(root.Left)
		fmt.Print(root.Val, " ")
		inorderTraversal(root.Right)
	}
}

// Удаление элемента
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if key < root.Val {
		root.Left = deleteNode(root.Left, key)
	} else if key > root.Val {
		root.Right = deleteNode(root.Right, key)
	} else {
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		}
		minNode := findMin(root.Right)
		root.Val = minNode.Val
		root.Right = deleteNode(root.Right, minNode.Val)
	}
	return root
}
func findMin(node *TreeNode) *TreeNode {
	current := node
	for current.Left != nil {
		current = current.Left
	}
	return current
}

// Поиск элемента
func searchNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if key < root.Val {
		return searchNode(root.Right, key)
	} else if key > root.Val {
		return searchNode(root.Left, key)
	} else {
		return root
	}
}

// Неориентированный граф на матрице смежности
type Graph struct {
	vertices int
	matrix   [][]int
}

func NewGraph(vertices int) *Graph {
	matrix := make([][]int, vertices)
	for i := range matrix {
		matrix[i] = make([]int, vertices)
	}
	return &Graph{vertices, matrix}
}
func (g *Graph) AddEdge(v, w int) {
	g.matrix[v][w] = 1
	g.matrix[w][v] = 1
}
func (g *Graph) Print() {
	for i := 0; i < g.vertices; i++ {
		for j := 0; j < g.vertices; j++ {
			fmt.Printf("%d ", g.matrix[i][j])
		}
		fmt.Println()
	}
}

// Поиск в ширину
func (graph *Graph) BFS(start int) {
	visited := make([]bool, graph.vertices)
	queue := []int{start}
	visited[start] = true

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		fmt.Printf("%d", node)

		for neighbor := 0; neighbor < graph.vertices; neighbor++ {
			if graph.matrix[node][neighbor] == 1 && !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
	}
	fmt.Println()
}

type Edge struct {
	numNeibo int
	Weight   int
}

// Ориентированный граф на списке смежности
type graph struct {
	vertices map[int][]Edge
}

func Newgraph() *graph {
	return &graph{vertices: make(map[int][]Edge)}
}
func (g *graph) AddEdges(u, v, weight int) {
	g.vertices[u] = append(g.vertices[u], Edge{numNeibo: v, Weight: weight})
}

func (g *graph) Print() {
	for u, edges := range g.vertices {
		for _, edge := range edges {
			fmt.Printf("%d->%d (weight: %d)\n", u, edge.numNeibo, edge.Weight)
		}
	}
}

// Поиск наименьшего пути по взвешенному графу через алгоритм Дейкстры
type Priority []*Item

type Item struct {
	vertex   int
	distance int
	index    int
}

func (p Priority) Len() int {
	return len(p)
}
func (p Priority) Less(i, j int) bool {
	return p[i].distance < p[j].distance
}

func (p Priority) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
	p[i].index = i
	p[j].index = j
}

func (p *Priority) Push(x interface{}) {
	n := len(*p)
	item := x.(*Item)
	item.index = n
	*p = append(*p, item)
}

func (p *Priority) Pop() interface{} {
	old := *p
	n := len(old)
	item := old[n-1]
	*p = old[0 : n-1]
	return item
}

func (g *graph) Dijkstra(start int) map[int]int {
	distance := make(map[int]int)
	for v := range g.vertices {
		distance[v] = int(^uint(0) >> 1)
	}
	distance[start] = 0

	p := &Priority{}
	heap.Push(p, &Item{vertex: start, distance: 0})

	for p.Len() > 0 {
		current := heap.Pop(p).(*Item)

		for _, edge := range g.vertices[current.vertex] {
			newDist := distance[current.vertex] + edge.Weight
			if newDist < distance[edge.numNeibo] {
				distance[edge.numNeibo] = newDist
				heap.Push(p, &Item{vertex: edge.numNeibo, distance: newDist})
			}
		}
	}

	return distance
}
func main() {
	//var root *TreeNode
	//root = insertNode(root, 1)
	//root = insertNode(root, 2)
	//root = insertNode(root, 3)
	//root = insertNode(root, 4)
	//root = insertNode(root, 5)
	//root = insertNode(root, 7)
	//root = insertNode(root, 8)
	//root = insertNode(root, 9)
	//root = deleteNode(root, 3)
	//inorderTraversal(root)
	//root = searchNode(root, 4)
	//graph := NewGraph(5)
	//graph.AddEdge(0, 1)
	//graph.AddEdge(0, 4)
	//graph.AddEdge(1, 2)
	//graph.AddEdge(1, 3)
	//graph.AddEdge(1, 4)
	//graph.AddEdge(2, 3)
	//graph.AddEdge(3, 4)
	//graph.Print()
	//graph.BFS(0)
	g := Newgraph()
	g.AddEdges(1, 2, 2)
	g.AddEdges(1, 3, 4)
	g.AddEdges(2, 4, 4)
	g.AddEdges(3, 4, 3)
	g.AddEdges(4, 3, 8)
	g.Print()
	distance := g.Dijkstra(1)
	for key, dis := range distance {
		fmt.Printf("Расстояние от 1 до %d =  %d\n", key, dis)
	}
}
