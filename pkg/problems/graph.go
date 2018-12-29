package problems

import (
	"fmt"
	"strings"
)

type Vertex interface {
	GetID() string
	Equals(Vertex) bool

	AddEdge(e Edge)
	GetEdges(d EdgeDirection) []Edge

	String() string
}

type Edge interface {
	ID() string
	X() Vertex
	Y() Vertex
	Direction() EdgeDirection

	Get(prop string) interface{}
	Set(prop string, val interface{})

	String() string
}

type EdgeDirection int

const (
	EdgeDirectionEither EdgeDirection = 0
	EdgeDirectionFrom   EdgeDirection = 1
	EdgeDirectionTo     EdgeDirection = 1 << 1
	EdgeDirectionBoth   EdgeDirection = EdgeDirectionFrom | EdgeDirectionTo
)

func (d EdgeDirection) String() string {
	if d == EdgeDirectionTo {
		return "->"
	} else if d == EdgeDirectionFrom {
		return "<-"
	} else if d == EdgeDirectionBoth {
		return "<->"
	} else if d == EdgeDirectionEither {
		return "-"
	}
	panic("unknown direction")
}

type Graph interface {
	// Get the vertex x by ID
	GetVertex(id string) Vertex
	// Adds the vertex x, if it is not there
	AddVertex(x Vertex) Vertex
	// GetVerticies Return the list of verticies in the graph
	GetVerticies() map[string]Vertex
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
	ID    string
	Data  map[string]interface{}
	edges map[string]Edge
}

func NewVertex(id string) Vertex {
	return &vertex{
		ID:    id,
		Data:  map[string]interface{}{},
		edges: map[string]Edge{},
	}
}

func (v *vertex) GetID() string {
	return v.ID
}

func (v *vertex) Equals(x Vertex) bool {
	return v.ID == x.GetID()
}

func (v *vertex) GetEdges(d EdgeDirection) []Edge {
	var edges []Edge
	for _, e := range v.edges {
		if (d == EdgeDirectionTo || d == EdgeDirectionBoth || d == EdgeDirectionEither) &&
			e.Y().Equals(v) {
			edges = append(edges, e)
			continue
		}
		if (d == EdgeDirectionFrom || d == EdgeDirectionBoth || d == EdgeDirectionEither) &&
			e.X().Equals(v) {
			edges = append(edges, e)
			continue
		}
	}
	return edges
}

func (v *vertex) AddEdge(e Edge) {
	v.edges[e.ID()] = e
}

func (v *vertex) String() string {
	return v.ID
}

func newEdge() *edge {
	return &edge{
		data: map[string]interface{}{},
	}
}

type edge struct {
	x         Vertex
	y         Vertex
	direction EdgeDirection
	data      map[string]interface{}
}

func (e *edge) ID() string {
	return e.String()
}

func (e *edge) X() Vertex {
	return e.x
}

func (e *edge) Y() Vertex {
	return e.y
}

func (e *edge) Get(prop string) interface{} {
	return e.data[prop]
}

func (e *edge) Set(prop string, val interface{}) {
	e.data[prop] = val
}

func (e *edge) Direction() EdgeDirection {
	return e.direction
}

func (e *edge) String() string {
	if e.direction == EdgeDirectionTo {
		return fmt.Sprintf("%s%s%s", e.x, e.Direction().String(), e.y)
	} else if e.direction == EdgeDirectionFrom {
		return fmt.Sprintf("%s%s%s", e.x, e.Direction().String(), e.y)
	} else if e.direction == EdgeDirectionBoth {
		return fmt.Sprintf("%s%s%s", e.x, e.Direction().String(), e.y)
	} else if e.direction == EdgeDirectionEither {
		return fmt.Sprintf("%s%s%s", e.x, e.Direction().String(), e.y)
	}
	return "uknown edge direction"
}

func (e *edge) Equals(cmp Edge) bool {
	return e.x.Equals(cmp.X()) &&
		e.y.Equals(cmp.Y()) &&
		e.direction == cmp.Direction()
}

type graph struct {
	verticies map[string]Vertex
	edges     map[string]Edge
}

func NewGraph() *graph {
	return &graph{
		verticies: map[string]Vertex{},
		edges:     map[string]Edge{},
	}
}

func (g *graph) GetVertex(id string) Vertex {
	return g.verticies[id]
}

func (g *graph) AddVertex(x Vertex) Vertex {
	for _, vertex := range g.verticies {
		if vertex.Equals(x) {
			return vertex
		}
	}
	g.verticies[x.GetID()] = x
	return x
}

func (g *graph) GetVerticies() map[string]Vertex {
	return g.verticies
}

func (g *graph) AddEdge(x, y Vertex, d EdgeDirection) Edge {
	x = g.AddVertex(x)
	y = g.AddVertex(y)
	e := newEdge()
	e.x = x
	e.y = y
	e.direction = d
	for _, edge := range g.edges {
		if e.Equals(edge) {
			return edge
		}
	}

	x.AddEdge(e)
	y.AddEdge(e)
	g.edges[e.ID()] = e
	return e
}

func (g *graph) Neighbors(x Vertex, d EdgeDirection) []Vertex {
	var vs []Vertex

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
