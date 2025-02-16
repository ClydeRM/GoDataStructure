package Graph

import (
    "fmt"
)

func (g *Graph) AddVertex(id string) {
    if _, exists := g.Vertices[id]; exists {
        return
    }

    g.Vertices[id] = &Vertex{
        Id:    id,
        Edges: []*Edge{},
    }
}

func (g *Graph) AddEdge(from, to string, weight int) {
    // 確保 from 和 to 頂點都存在
    fromVertex, fromExists := g.Vertices[from]
    _, toExists := g.Vertices[to] // 只需要檢查是否存在，不需要存變數

    if !fromExists || !toExists {
        return // 頂點不存在，無法新增邊
    }

    // 檢查是否已經存在這條邊，避免重複
    for _, edge := range fromVertex.Edges {
        if edge.To == to {
            return // 已存在該邊，不重複新增
        }
    }

    // 新增邊
    fromVertex.Edges = append(fromVertex.Edges, &Edge{To: to, Weight: weight})
}

func (g *Graph) RemoveVertex(id string) {
    if _, idExists := g.Vertices[id]; !idExists {
        return
    }

    delete(g.Vertices, id)

    for _, vertex := range g.Vertices {
        var refreshEdges []*Edge

        for _, edge := range vertex.Edges {
            if edge.To != id {
                refreshEdges = append(refreshEdges, edge)
            }
        }

        vertex.Edges = refreshEdges
    }
}

func (g *Graph) RemoveEdge(from, to string) {
    fromVertex, fromExist := g.Vertices[from]

    if !fromExist {
        return
    }

    var refreshEdges []*Edge
    for _, edge := range fromVertex.Edges {
        if edge.To != to {
            refreshEdges = append(refreshEdges, edge)
        }
    }
    fromVertex.Edges = refreshEdges
}

func (g *Graph) GetVertices() map[string]*Vertex {
    return g.Vertices
}

func (g *Graph) GetEdges(id string) []*Edge {
    idVertex, idExist := g.Vertices[id]

    if !idExist {
        return nil
    }

    return idVertex.Edges
}

func (g *Graph) HasVertex(id string) bool {
    _, idExist := g.Vertices[id]
    return idExist
}

func (g *Graph) HasEdge(from, to string) bool {
    fromVertex, fromExist := g.Vertices[from]

    if !fromExist {
        return false
    }

    for _, edge := range fromVertex.Edges {
        if edge.To == to {
            return true
        }
    }

    return false
}

func (g *Graph) PrintGraph() {
    for id, vertex := range g.Vertices {
        fmt.Printf("Vertex: %v\n", id)
        for _, edge := range vertex.Edges {
            fmt.Printf("Edge to %v, weight: %v\n", edge.To, edge.Weight)
        }
    }
}

func (g *Graph) DFS(start string) []string {
    visited := make(map[string]bool)
    return g.DFSHelper(start, &visited)
}

func (g *Graph) DFSHelper(start string, visited *map[string]bool) []string {
    if _, isExist := g.Vertices[start]; !isExist {
        return []string{}
    }

    var result []string
    if *visited == nil {
        *visited = make(map[string]bool)
    }

    // if visited: return []
    if (*visited)[start] {
        return []string{}
    }

    (*visited)[start] = true
    stack := []string{start}

    for len(stack) != 0 {
        vertex := stack[len(stack)-1]
        stack = stack[:len(stack)-1]

        result = append(result, vertex)

        for _, edge := range g.Vertices[vertex].Edges {
            if !(*visited)[edge.To] {
                (*visited)[edge.To] = true
                stack = append(stack, edge.To)
            }
        }

    }
    return result
}

func (g *Graph) BFS(start string) []string {
    visited := make(map[string]bool)
    return g.BFSHelper(start, &visited)
}

func (g *Graph) BFSHelper(start string, visited *map[string]bool) []string {
    if _, isExist := g.Vertices[start]; !isExist {
        return []string{}
    }

    var result []string
    if *visited == nil {
        *visited = make(map[string]bool)
    }

    // if visited: return []
    if (*visited)[start] {
        return []string{}
    }

    (*visited)[start] = true
    queue := []string{start}

    for len(queue) != 0 {
        vertex := queue[0]
        queue = queue[1:]

        result = append(result, vertex)

        for _, edge := range g.Vertices[vertex].Edges {
            if !(*visited)[edge.To] {
                (*visited)[edge.To] = true
                queue = append(queue, edge.To)
            }
        }
    }
    return result
}

func (g *Graph) FindConnectedComponents() [][]string {
    visited := make(map[string]bool)
    var components [][]string

    for id := range g.Vertices {
        if !visited[id] {
            currComponent := g.BFSHelper(id, &visited) // can switch with DFSHelper
            components = append(components, currComponent)
        }
    }

    return components
}

func (g *Graph) IsConnected(v1, v2 string) bool {
    if _, v1Existed := g.Vertices[v1]; v1Existed {
        return false
    }
    if _, v2Existed := g.Vertices[v2]; v2Existed {
        return false
    }

    dfsV1 := g.DFS(v1)
    for _, vertex := range dfsV1 {
        if v2 == vertex {
            return true
        }
    }

    return false
}
