package Graph

type Graph struct {
    Vertices map[string]*Vertex
}

func NewGraph() *Graph {
    return &Graph{Vertices: make(map[string]*Vertex)}
}

type MSTEdge struct {
    From, To string
    Weight   int
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

type PriorityQueue []*MSTEdge

func (pq *PriorityQueue) Len() int           { return len(*pq) }
func (pq *PriorityQueue) Less(i, j int) bool { return (*pq)[i].Weight < (*pq)[j].Weight }
func (pq *PriorityQueue) Swap(i, j int)      { (*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i] }

func (pq *PriorityQueue) Push(x interface{}) {
    *pq = append(*pq, x.(*MSTEdge))
}

func (pq *PriorityQueue) Pop() interface{} {
    old := *pq
    n := len(old)
    item := old[n-1]
    *pq = old[:n-1]
    return item
}

type ShortPath struct {
    Node  string
    Dist  int
    Index int
}

type SPPriorityQueue []*ShortPath

func (sppq *SPPriorityQueue) Len() int {
    return len(*sppq)
}

func (sppq *SPPriorityQueue) Less(i, j int) bool {
    return (*sppq)[i].Dist < (*sppq)[j].Dist
}

func (sppq *SPPriorityQueue) Swap(i, j int) {
    (*sppq)[i], (*sppq)[j] = (*sppq)[j], (*sppq)[i]
    (*sppq)[i].Index = i
    (*sppq)[j].Index = j
}

func (sppq *SPPriorityQueue) Push(x interface{}) {
    item := x.(*ShortPath)
    item.Index = len(*sppq)
    *sppq = append(*sppq, item)
}

func (sppq *SPPriorityQueue) Pop() interface{} {
    old := *sppq
    n := len(*sppq)
    item := old[n-1]
    item.Index = -1
    old[n-1] = nil
    *sppq = old[:n-1]

    return item
}

type Interface interface {
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
    KruskalMST() []MSTEdge
    PrimMST(start string) []*MSTEdge

    // ShortestPath
    DijkstraShortPath(start string) map[string]int
}
