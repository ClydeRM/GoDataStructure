package Graph

import (
    "container/heap"
    "fmt"
    "math"
    "sort"
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

// FindConnectedComponents Undirected Graph (or two way)
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

// FindStronglyConnectedComponents Kosaraju's Algorithm
func (g *Graph) FindStronglyConnectedComponents() [][]string {
    var stack []string
    visited := make(map[string]bool)

    // Step 1: DFS 正向圖 and 記錄完成順序（stack）
    for id := range g.Vertices {
        if !visited[id] {
            g.dfsFillStack(id, visited, &stack)
        }
    }

    // Step 2: 建立反向圖
    reversedGraph := g.reverseGraph()

    // Step 3: 根據 stack 順序，在反向圖上執行 DFS
    visited = make(map[string]bool)
    var components [][]string

    for len(stack) > 0 {
        node := stack[len(stack)-1]
        stack = stack[:len(stack)-1] // pop

        if !visited[node] {
            var component []string
            reversedGraph.dfsCollectSCC(node, visited, &component)
            components = append(components, component)
        }
    }

    return components
}

func (g *Graph) dfsFillStack(vertex string, visited map[string]bool, stack *[]string) {
    visited[vertex] = true

    for _, edge := range g.Vertices[vertex].Edges {
        if !visited[edge.To] {
            g.dfsFillStack(edge.To, visited, stack)
        }
    }

    *stack = append(*stack, vertex)
}

func (g *Graph) reverseGraph() *Graph {
    reversedGraph := NewGraph()

    for id := range g.Vertices {
        reversedGraph.AddVertex(id)
    }

    for from, vertex := range g.Vertices {
        for _, edge := range vertex.Edges {
            reversedGraph.AddEdge(edge.To, from, edge.Weight) // 反向邊
        }
    }

    return reversedGraph
}

func (g *Graph) dfsCollectSCC(vertex string, visited map[string]bool, component *[]string) {
    visited[vertex] = true
    *component = append(*component, vertex)

    for _, edge := range g.Vertices[vertex].Edges {
        if !visited[edge.To] {
            g.dfsCollectSCC(edge.To, visited, component)
        }
    }
}

// KruskalMST Kruskal's Algorithm
func (g *Graph) KruskalMST() []MSTEdge {
    var mst []MSTEdge
    var edges []MSTEdge

    // collect all edge
    for from, vertex := range g.Vertices {
        for _, edge := range vertex.Edges {
            edges = append(edges, MSTEdge{
                from,
                edge.To,
                edge.Weight,
            })
        }
    }

    // sort edge by edge weight
    sort.Slice(edges, func(i, j int) bool {
        return edges[i].Weight < edges[j].Weight
    })

    // init set
    ds := NewDisjointSet()
    for id := range g.Vertices {
        ds.Parent[id] = id
        ds.Rank[id] = 0
    }

    // loop edge
    for _, edge := range edges {
        from := edge.From
        to := edge.To

        if ds.Find(from) != ds.Find(to) { // check cycle
            mst = append(mst, edge)
            ds.Union(from, to)
        }

        // mst edge equal to vertices-1: break
        if len(mst) == len(g.Vertices)-1 {
            break
        }
    }

    return mst
}

func (ds *DisjointSet) Find(v string) string {
    if ds.Parent[v] != v {
        ds.Parent[v] = ds.Find(ds.Parent[v])
    }
    return ds.Parent[v]
}

func (ds *DisjointSet) Union(v1, v2 string) {
    root1 := ds.Find(v1)
    root2 := ds.Find(v2)

    if root1 != root2 {
        if ds.Rank[root1] > ds.Rank[root2] {
            ds.Parent[root2] = root1
        } else if ds.Rank[root1] < ds.Rank[root2] {
            ds.Parent[root1] = root2
        } else {
            ds.Parent[root2] = root1
            ds.Rank[root1]++
        }
    }
}

func (g *Graph) PrimMST(start string) []*MSTEdge {
    if _, exists := g.Vertices[start]; !exists {
        return nil
    }

    var mst []*MSTEdge
    visited := make(map[string]bool)
    pq := &PriorityQueue{}
    heap.Init(pq)

    visited[start] = true
    // add all edge of start vertex in pq
    for _, edge := range g.Vertices[start].Edges {
        heap.Push(pq, &MSTEdge{start, edge.To, edge.Weight})
    }

    for pq.Len() > 0 {
        // pop smallest weight edge from pq
        minEdge := heap.Pop(pq).(*MSTEdge)

        if visited[minEdge.To] {
            continue
        }

        // add into MST
        mst = append(mst, minEdge)
        visited[minEdge.To] = true

        for _, edge := range g.Vertices[minEdge.To].Edges {
            if !visited[edge.To] {
                heap.Push(pq, &MSTEdge{minEdge.To, edge.To, edge.Weight})
            }
        }
    }

    return mst
}

// Dijkstra Algorithm (Non-Negative weight)
func (g *Graph) DijkstraShortPath(start string) (map[string]int, map[string]string) {
    if _, isExists := g.Vertices[start]; !isExists {
        return nil, nil
    }

    dist := make(map[string]int)
    prev := make(map[string]string)

    // set all of vertex in dist map are infinity
    for id := range g.Vertices {
        dist[id] = math.MaxInt
    }
    dist[start] = 0

    // Priority Queue
    pq := &SPPriorityQueue{}
    heap.Init(pq)
    heap.Push(pq, &ShortPath{To: start, Dist: 0})

    for pq.Len() > 0 {
        // get smallest edge
        curr := heap.Pop(pq).(*ShortPath)
        currentVertex := curr.To
        currentDist := curr.Dist

        // skip if curr > dist map
        if currentDist > dist[currentVertex] {
            continue
        }

        // update neighbor
        for _, edge := range g.Vertices[currentVertex].Edges {
            newDist := currentDist + edge.Weight

            // if newDist is smaller than record
            if newDist < dist[edge.To] {
                dist[edge.To] = newDist
                prev[edge.To] = currentVertex
                heap.Push(pq, &ShortPath{To: edge.To, Dist: newDist})
            }
        }
    }

    return dist, prev
}

// Bellman-Ford Algorithm (Negative cycle check)
func (g *Graph) BellmanFordShortPath(start string) (map[string]int, map[string]string, bool) {
    if _, exists := g.Vertices[start]; !exists {
        return nil, nil, false
    }

    dist := make(map[string]int)
    prev := make(map[string]string)

    for id := range g.Vertices {
        dist[id] = math.MaxInt
    }
    dist[start] = 0

    // V-1 times Relaxation.
    V := len(g.Vertices)
    for i := 0; i < V-1; i++ {
        updated := false // 用於提前結束迴圈
        for _, vertex := range g.Vertices {
            for _, edge := range vertex.Edges {
                if dist[vertex.Id] != math.MaxInt && dist[vertex.Id]+edge.Weight < dist[edge.To] {
                    dist[edge.To] = dist[vertex.Id] + edge.Weight
                    prev[edge.To] = vertex.Id
                    updated = true
                }
            }
        }
        // 若這輪沒有更新任何邊，代表已經收斂，提前結束
        if !updated {
            break
        }
    }

    // Negative cycle check
    for _, u := range g.Vertices {
        for _, edge := range u.Edges {
            if dist[u.Id] != math.MaxInt && dist[u.Id]+edge.Weight < dist[edge.To] {
                return nil, nil, true // 發現負權重循環
            }
        }
    }

    return dist, prev, false
}

// Floyd-Warshall Algorithm (All to other SP)
func (g *Graph) FloydWarshallShortPath() (map[string]map[string]int, bool) {
    // 2d map[v][u] = weight
    dist := make(map[string]map[string]int)

    // init dist
    for id1 := range g.Vertices {
        dist[id1] = make(map[string]int)
        for id2 := range g.Vertices {
            if id1 == id2 {
                dist[id1][id2] = 0 // 自己到自己距離為 0
            } else {
                dist[id1][id2] = math.MaxInt // 預設為無限大
            }
        }
    }

    for id, vertex := range g.Vertices {
        for _, edge := range vertex.Edges {
            dist[id][edge.To] = edge.Weight
        }
    }

    for k := range g.Vertices { // 中繼點
        for i := range g.Vertices { // 起點
            for j := range g.Vertices { // 終點
                // 確保 i → k、k → j 可達，避免數值溢出
                if dist[i][k] != math.MaxInt && dist[k][j] != math.MaxInt {
                    if newDist := dist[i][k] + dist[k][j]; newDist < dist[i][j] {
                        dist[i][j] = newDist
                    }
                }
            }
        }
    }

    // Negative cycle check
    for id := range g.Vertices {
        if dist[id][id] < 0 {
            return nil, true // 發現負權重循環
        }
    }

    return dist, false
}
