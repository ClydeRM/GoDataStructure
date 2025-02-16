package Graph

type Graph struct {
    Vertices map[string]*Vertex
}

func NewGraph() *Graph {
    return &Graph{Vertices: make(map[string]*Vertex)}
}

type DisjointSet struct {
    Parent map[string]string
    Rank   map[string]int
}

func NewDisjointSet() *DisjointSet {
    return &DisjointSet{
        Parent: make(map[string]string),
        Rank:   make(map[string]int),
    }
}

type GraphInterface interface {
    // 基本操作
    AddVertex(id string)
    AddEdge(from, to string, weight int)
    RemoveVertex(id string)
    RemoveEdge(from, to string)
    GetVertices() map[string]*Vertex
    GetEdges(id string) []*Edge
    HasVertex(id string) bool
    HasEdge(from, to string) bool
    PrintGraph()

    // 走訪演算法
    DFS(start string) []string // 深度優先搜尋
    DFSHelper(start string, visited *map[string]bool) []string
    BFS(start string) []string // 廣度優先搜尋
    BFSHelper(start string, visited *map[string]bool) []string

    // Connected Components
    FindConnectedComponents() [][]string // find all connected component
    IsConnected(v1, v2 string) bool      // chack both components are connected
    FindStronglyConnectedComponents() [][]string

    // MST
    KruskalMST() []struct {
        From, To string
        Weight   int
    }
}
