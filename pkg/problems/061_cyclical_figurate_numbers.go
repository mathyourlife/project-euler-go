package problems

import (
	"fmt"
	"log"
)

type CyclicalFigurateNumbers struct{}

func (p *CyclicalFigurateNumbers) ID() int {
	return 61
}

func (p *CyclicalFigurateNumbers) Text() string {
	return `Triangle, square, pentagonal, hexagonal, heptagonal, and
octagonal numbers are all figurate (polygonal) numbers and are
generated by the following formulae:

Triangle   P(3,n)=n(n+1)/2   1, 3, 6, 10, 15, ...
Square     P(4,n)=n^2        1, 4, 9, 16, 25, ...
Pentagonal P(5,n)=n(3n-1)/2  1, 5, 12, 22, 35, ...
Hexagonal  P(6,n)=n(2n-1)    1, 6, 15, 28, 45, ...
Heptagonal P(7,n)=n(5n-3)/2  1, 7, 18, 34, 55, ...
Octagonal  P(8,n)=n(3n-2)    1, 8, 21, 40, 65, ...

The ordered set of three 4-digit numbers: 8128, 2882, 8281, has
three interesting properties.

1. The set is cyclic, in that the last two digits of each number is
   the first two digits of the next number (including the last number
   with the first).
2. Each polygonal type: triangle (P(3,127)=8128), square (P(4,91)=8281), and
   pentagonal (P(5,44)=2882), is represented by a different number in the set.
3. This is the only set of 4-digit numbers with this property.

Find the sum of the only ordered set of six cyclic 4-digit numbers
for which each polygonal type: triangle, square, pentagonal,
hexagonal, heptagonal, and octagonal, is represented by a different
number in the set.
`
}

func (p *CyclicalFigurateNumbers) trace(edges []map[uint64][]uint64, path []uint64) [][]uint64 {
	if len(path) == 0 {
		paths := [][]uint64{}
		for k, nextSteps := range edges[0] {
			for _, nextStep := range nextSteps {
				curPath := []uint64{k, nextStep}
				results := p.trace(edges[1:], curPath)
				if results != nil && len(results) > 0 {
					paths = append(paths, results...)
				}
			}
		}
		return paths
	}
	if len(edges) == 0 && path[0] == path[len(path)-1] {
		return [][]uint64{
			path,
		}
	}

	lastStep := path[len(path)-1]
	paths := [][]uint64{}
	for i := 0; i < len(edges); i++ {
		for _, nextStep := range edges[i][lastStep] {
			curPath := []uint64{}
			for _, step := range path {
				curPath = append(curPath, step)
			}
			curPath = append(curPath, nextStep)
			result := p.trace(append(edges[:i], edges[i+1:]...), curPath)
			if result != nil {
				paths = append(paths, result...)
			}
		}
	}
	return paths
}

func junk() {
	sequences := map[int]map[int][]int{
		3: map[int][]int{
			14: []int{65, 91},
		},
		5: map[int][]int{
			91: []int{31},
		},
		4: map[int][]int{
			32: []int{65, 14},
		},
		6: map[int][]int{
			31: []int{32},
		},
	}
	fmt.Println(sequences)

	paths := [][]int{}
	// preload paths
	for pn, trees := range sequences {
		for branch, leaves := range trees {
			for _, leaf := range leaves {
				paths = append(paths, []int{pn, branch, leaf})
			}
		}
		break
	}
	fmt.Println("preloaded paths")
	fmt.Println(paths)

	target := len(sequences) * 3
	for pn, trees := range sequences {
		for branch, leaves := range trees {
			for _, leaf := range leaves {
				fmt.Println("new leaf", pn, branch, leaf)
				for _, path := range paths {
					// path already uses the pn
					used := false
					for i := 0; i < len(path); i += 3 {
						if path[i] == pn {
							used = true
							break
						}
					}
					if used {
						continue
					}

					if branch == path[len(path)-1] {
						newPath := append(path, []int{pn, branch, leaf}...)
						if len(newPath) == target && newPath[1] == newPath[len(newPath)-1] {
							fmt.Println("found", newPath)
							prettyPrint(newPath)
						}
						paths = append(paths, newPath)
					}
				}
				paths = append(paths, []int{pn, branch, leaf})
			}
		}
	}
	fmt.Println(paths)
}

func prettyPrint(path []int) {
	for i := 0; i < len(path); i += 3 {
		fmt.Printf("%2d %2d %2d\n", path[i], path[i+1], path[i+2])
	}
}

func walk(sequences map[int]map[int][]int, visited []int, path []int) {
	fmt.Println("walk", "visited", visited, "path", path)

	for pn := 3; pn <= 5; pn++ {
		skip := false
		for _, visit := range visited {
			if pn == visit {
				skip = true
				break
			}
		}
		if skip {
			continue
		}
		fmt.Println("checking pn", pn)
	}
}

func (p *CyclicalFigurateNumbers) preload() map[int]map[int][]int {
	// sequences := map[int]map[int][]int{
	// 	3: map[int][]int{
	// 		3: []int{5, 3},
	// 	},
	// 	4: map[int][]int{
	// 		1: []int{2, 3},
	// 		3: []int{4, 8},
	// 	},
	// 	5: map[int][]int{
	// 		5: []int{6, 1},
	// 	},
	// }

	sequences := map[int]map[int][]int{}
	for pn := 3; pn <= 8; pn++ {
		n := 0
		seq := map[int][]int{}
		for {
			v := PolynomialSequence(pn, n)
			n++
			if v < 1000 {
				continue
			} else if v >= 10000 {
				break
			}

			if seq[int(v/100)] == nil {
				seq[int(v/100)] = []int{int(v % 100)}
			} else {
				seq[int(v/100)] = append(seq[int(v/100)], int(v%100))
			}
		}
		sequences[pn] = seq
	}
	return sequences
}

func (p *CyclicalFigurateNumbers) Solve() (string, error) {

	sequences := p.preload()
	target := len(sequences) * 3

	var g Graph
	g = NewGraph()

	for pn, tree := range sequences {
		for branch, leaves := range tree {
			var x Vertex
			x = NewVertex(fmt.Sprintf("%d", branch))
			x = g.AddVertex(x)

			for _, leaf := range leaves {
				var y Vertex
				y = NewVertex(fmt.Sprintf("%d", leaf))
				y = g.AddVertex(y)
				var e Edge
				e = g.AddEdge(x, y, EdgeDirectionTo)
				var pns map[int]bool
				pnIface := e.Get("pn")
				if pnIface == nil {
					pns = map[int]bool{}
				} else {
					pns = pnIface.(map[int]bool)
				}
				pns[pn] = true
				e.Set("pn", pns)
			}
		}
	}
	for _, v := range g.GetVerticies() {
		fmt.Println(v, v.GetEdges(EdgeDirectionFrom))
	}

	log.Println("-------------------------------------")
	paths := [][]string{}
	for pn := 3; pn <= 3; pn++ {
		for branch, leaves := range sequences[pn] {
			for _, leaf := range leaves {
				paths = append(paths, []string{
					fmt.Sprintf("%d", pn),
					g.GetVertex(fmt.Sprintf("%d", branch)).GetID(),
					g.GetVertex(fmt.Sprintf("%d", leaf)).GetID(),
				})
			}
		}
	}

	for {
		newPaths := [][]string{}
		for j, _ := range paths {
			v := g.GetVertex(paths[j][len(paths[j])-1])
			edges := v.GetEdges(EdgeDirectionFrom)
			log.Printf("for path %v there are %d edges from vertex %s", paths[j], len(edges), v)
			for k, _ := range edges {
				pns := edges[k].Get("pn").(map[int]bool)
				for pn, _ := range pns {
					pnUsed := false
					for idx := 0; idx < len(paths[j]); idx += 3 {
						if paths[j][idx] == fmt.Sprintf("%d", pn) {
							log.Printf("skipping pn=%d already used in path %s", pn, paths[j])
							pnUsed = true
							break
						}
					}
					if pnUsed {
						continue
					}
					id := edges[k].Y().GetID()
					next := []string{
						fmt.Sprintf("%d", pn), v.GetID(), id,
					}
					newPath := append(paths[j], next...)
					log.Printf("adding next hop %s -> %s", v, id)

					if len(newPath) == target && newPath[len(newPath)-1] == newPath[1] {
						log.Printf("FOUND IT: %s", newPath)
						return "0", nil
					}
					newPaths = append(newPaths, newPath)
				}
			}
		}
		if len(newPaths) == 0 {
			log.Printf("error: no valid paths remain")
			return "0", nil
		}
		log.Printf("%d valid paths remaining", len(newPaths))
		for _, path := range newPaths {
			log.Println(path)
		}
		paths = newPaths
	}

	return "0", nil

	for pn := 3; pn <= 5; pn++ {
		n := 0
		seq := map[int][]int{}
		for {
			v := PolynomialSequence(pn, n)
			n++
			if v < 1000 {
				continue
			} else if v >= 10000 {
				break
			}

			if seq[int(v/100)] == nil {
				seq[int(v/100)] = []int{int(v % 100)}
			} else {
				seq[int(v/100)] = append(seq[int(v/100)], int(v%100))
			}
		}
		sequences[pn] = seq
	}

	for start, nodes := range sequences[3] {
		path := []int{start}
		fmt.Println(path, nodes)
		for _, node := range nodes {
			fmt.Println("path", path, "4", sequences[4][node])
			fmt.Println("path", path, "5", sequences[4][node])
		}
	}

	return "0", nil

	set := 3
	edges := make([]map[uint64][]uint64, set)

	for pn := 3; pn < 3+set; pn++ {
		n := 0
		for {
			n++
			v := PolynomialSequence(pn, n)
			if v < 1000 {
				continue
			} else if v >= 10000 {
				break
			}
			fmt.Println(pn, v)
		}
	}
	return "", nil

	for pn := 3; pn < 3+set; pn++ {
		edges[pn-3] = map[uint64][]uint64{}
		n := 1
		for {
			v := PolynomialSequence(pn, n)
			if v >= 1000 && v < 10000 {
				edges[pn-3][v/100] = append(edges[pn-3][v/100], v%100)
			} else if v >= 10000 {
				break
			}
			n++
		}
	}

	path := []uint64{}
	fmt.Println("final", p.trace(edges, path))

	return fmt.Sprintf("%d", 0), nil
}
