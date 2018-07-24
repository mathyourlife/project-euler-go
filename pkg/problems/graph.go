package problems

import (
	"fmt"
	"strings"
)

type Vertex interface {
	GetID() string
	Equals(Vertex) bool

	String() string
}

type Edge interface {
	X() Vertex
	Y() Vertex
	Direction() EdgeDirection

	String() string
}

type EdgeDirection int

const (
	EdgeDirectionEither EdgeDirection = 0
	EdgeDirectionFrom   EdgeDirection = 1
	EdgeDirectionTo     EdgeDirection = 1 << 1
	EdgeDirectionBoth   EdgeDirection = EdgeDirectionFrom | EdgeDirectionTo
)

type Graph interface {
	// Adds the vertex x, if it is not there
	AddVertex(x Vertex)
	// GetVerticies Return the list of verticies in the graph
	GetVerticies() []Vertex
	// Adds the edge from the vertex x to the vertex y, if it is not there
	AddEdge(x, y Vertex, d EdgeDirection) Edge
	// Lists all vertices y such that there is an edge connecting x to
	// y in the specified direction.  When comparing edge directions,
	// EdgeDirectionTo, EdgeDirectionFrom, EdgeDirectionBoth
	// need to be exact matches.  Pass EdgeDirectionEither if direction
	// does not matter.
	Neighbors(x Vertex, d EdgeDirection) []Vertex
	// Tests whether there is an edge from the vertex x to the vertex y
	Adjacent(x, y Vertex) Edge
	// // Removes the vertex x, if it is there
	// RemoveVertex(x Vertex)
	// // Removes the edge from the vertex x to the vertex y, if it is there
	// RemoveEdge(x, y Vertex)
	// // Returns the value associated with the vertex x
	// GetVertexValue(x Vertex)
	// // Sets the value associated with the vertex x to v
	// SetVertexValue(x, v Vertex)

	String() string
}

type vertex struct {
	ID   string
	Data map[string]interface{}
}

func (v *vertex) GetID() string {
	return v.ID
}

func (v *vertex) Equals(x Vertex) bool {
	return v.ID == x.GetID()
}

func (v *vertex) String() string {
	return v.ID
}

type edge struct {
	x         Vertex
	y         Vertex
	direction EdgeDirection
}

func (e *edge) X() Vertex {
	return e.x
}

func (e *edge) Y() Vertex {
	return e.y
}

func (e *edge) Direction() EdgeDirection {
	return e.direction
}

func (e *edge) String() string {
	if e.direction == EdgeDirectionTo {
		return fmt.Sprintf("%s->%s", e.x, e.y)
	} else if e.direction == EdgeDirectionFrom {
		return fmt.Sprintf("%s<-%s", e.x, e.y)
	} else if e.direction == EdgeDirectionBoth {
		return fmt.Sprintf("%s<->%s", e.x, e.y)
	} else if e.direction == EdgeDirectionEither {
		return fmt.Sprintf("%s-%s", e.x, e.y)
	}
	return "uknown edge direction"
}

func (e *edge) Equals(cmp Edge) bool {
	return e.x.Equals(cmp.X()) &&
		e.y.Equals(cmp.Y()) &&
		e.direction == cmp.Direction()
}

type graph struct {
	verticies []Vertex
	edges     []Edge
}

func NewGraph() *graph {
	return &graph{
		verticies: []Vertex{},
		edges:     []Edge{},
	}
}

func (g *graph) AddVertex(x Vertex) {
	for _, vertex := range g.verticies {
		if vertex.Equals(x) {
			return
		}
	}
	g.verticies = append(g.verticies, x)
}

func (g *graph) GetVerticies() []Vertex {
	return g.verticies
}

func (g *graph) AddEdge(x, y Vertex, d EdgeDirection) Edge {
	g.AddVertex(x)
	g.AddVertex(y)
	e := &edge{
		x:         x,
		y:         y,
		direction: d,
	}
	for _, edge := range g.edges {
		if e.Equals(edge) {
			return edge
		}
	}
	g.edges = append(g.edges, e)
	return e
}

func (g *graph) Neighbors(x Vertex, d EdgeDirection) []Vertex {
	vs := []Vertex{}

	add := func(z Vertex) {
		for _, v := range vs {
			if v.Equals(z) {
				return
			}
		}
		vs = append(vs, z)
	}

	for _, edge := range g.edges {
		isX := edge.X().Equals(x)
		isY := edge.Y().Equals(x)
		if !isX && !isY {
			continue
		}
		if isX && (d == EdgeDirectionFrom || d == EdgeDirectionBoth) {
			add(edge.Y())
		} else if isY && (d == EdgeDirectionTo || d == EdgeDirectionBoth) {
			add(edge.X())
		}
	}
	return vs
}

func (g *graph) Adjacent(x, y Vertex) Edge {
	for _, edge := range g.edges {
		if (edge.Direction() == EdgeDirectionTo || edge.Direction() == EdgeDirectionBoth) &&
			edge.X().Equals(x) && edge.Y().Equals(y) {
			return edge
		} else if (edge.Direction() == EdgeDirectionFrom || edge.Direction() == EdgeDirectionBoth) &&
			edge.X().Equals(y) && edge.Y().Equals(x) {
			return edge
		}
	}
	return nil
}

func (g *graph) String() string {
	vs := make([]string, 0, len(g.verticies))
	for _, vertex := range g.verticies {
		vs = append(vs, vertex.String())
	}

	es := make([]string, 0, len(g.edges))
	for _, edge := range g.edges {
		es = append(es, edge.String())
	}

	return "verticies: " + strings.Join(vs, ",") + "\n" + "edges: " + strings.Join(es, ",")
}
