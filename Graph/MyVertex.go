package Graph

type Edge struct {
	To     string
	Weight int
}

type Vertex struct {
	Id    string
	Edges []*Edge
}
