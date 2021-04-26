package ctci

import "strings"

// 4.1 Route Between Nodes: Given a directed graph, design an algorithm to find
// out whether there is a route between two nodes.
//
// A classic breadth-first search problem. Depth-first search is not appropriate
// because it may traverse many nodes in the wrong direction before moving on
// to a more suitable path.
//
// All solutions assume that the graph is implemented using an adjacency list.

// Node represents a node in a directed graph.
type Node struct {
	name     string
	children []*Node
}

func NewNode(name string) *Node {
	return &Node{
		name:     name,
		children: make([]*Node, 0),
	}
}

// FindRoute implements a breadth-first search (BFS) to find a route (if
// present) in O(k^d) time, where k is the maximal number of neighbouring nodes
// an d is the length of the shortest path between the two nodes. It takes
// pointers to the start and end nodes and returns a slice of pointers to the
// nodes comprising the shortest path between them. If no path is possible, nil
// is returned. If the start and the end node are the same, a slice containing
// only that node is returned.
func FindRoute(s, t *Node) []*Node {
	if s == nil || t == nil {
		return nil
	}

	if s == t {
		return []*Node{s}
	}

	worklist := []*Node{s}
	seen := make(map[*Node][]*Node)

	for len(worklist) > 0 {
		nodes := worklist
		worklist = nil

		for _, current := range nodes {
			routeToCurrent := append(seen[current], current)

			for _, child := range current.children {
				if child == t {
					return append(routeToCurrent, t)
				}

				if _, ok := seen[child]; !ok {
					seen[child] = routeToCurrent
					worklist = append(worklist, child)
				}
			}

		}
	}

	return nil
}

// RouteToString returns a string representation of a route between two nodes.
func RouteToString(route []*Node) string {
	sb := strings.Builder{}
	for i, n := range route {
		if i > 0 {
			sb.WriteString(" -> ")
		}
		sb.WriteString(n.name)
	}

	return sb.String()
}

// BidirecNode is a directed node which maintains knowledge of its parents.
type BidirecNode struct {
	name     string
	children []*BidirecNode
	parents  []*BidirecNode
}

func NewBirdirecNode(name string) *BidirecNode {
	return &BidirecNode{
		name:     name,
		children: make([]*BidirecNode, 0),
		parents:  make([]*BidirecNode, 0),
	}
}

// FindRouteWithParents implements a bidirectional breadth-first search (BFS) to
// find a route (if present) in O(k^(d/2)) time, where k is the maximal number
// of neighboring nodes and d is the length of the shortest path between the two
// nodes. It takes pointers to the start and the end nodes and returns a slice
// of pointers to the nodes comprising the shortest path between the input
// nodes. If no path is possible, nil is returned. If the start and the end node
// are the same, a slice containing only that node is returned.
//
// For a directed graph, bidirectional search is only possible if the
// "downstream" nodes have knowledge of their parents, i.e. of the nodes with
// edges leading to them. Clarify this assumption in the interview.
func FindRouteBidirectional(s, t *BidirecNode) []*BidirecNode {
	if s == nil || t == nil {
		return nil
	}

	if s == t {
		return []*BidirecNode{s}
	}

	// sSeen and tSeen record the nodes visited by the BFSs started from nodes
	// s and t. Values are the paths taken to arrive at the keys.
	sSeen := make(map[*BidirecNode][]*BidirecNode)
	tSeen := make(map[*BidirecNode][]*BidirecNode)

	// sQueue and tQueue contain the nodes whose neighbours must still be
	// traversed.
	sQueue := []*BidirecNode{s}
	tQueue := []*BidirecNode{t}

	// Process the queues until a route is found or one of the searches
	// completes, at which point we know a route is impossible.
	for len(sQueue) > 0 && len(tQueue) > 0 {
		sNodes, tNodes := sQueue, tQueue
		sQueue, tQueue = nil, nil

		for _, sCur := range sNodes {
			routeFromSToCurrent := append(sSeen[sCur], sCur)

			// Check whether the current node has been seen by the search from t to
			// see if we've found a route.
			if tToCurrent, ok := tSeen[sCur]; ok {
				reverse(tToCurrent)
				return append(routeFromSToCurrent, tToCurrent...)
			}

			for _, n := range sCur.children {
				if _, ok := sSeen[n]; !ok {
					sSeen[n] = routeFromSToCurrent
					sQueue = append(sQueue, n)
				}
			}
		}

		for _, tCur := range tNodes {
			routeFromTToCurrent := append(tSeen[tCur], tCur)

			if sToCurrent, ok := sSeen[tCur]; ok {
				reverse(routeFromTToCurrent)
				return append(sToCurrent, routeFromTToCurrent...)
			}

			for _, n := range tCur.parents {
				if _, ok := tSeen[n]; !ok {
					tSeen[n] = routeFromTToCurrent
					tQueue = append(tQueue, n)
				}
			}
		}
	}

	return nil
}

func reverse(nodes []*BidirecNode) {
	for i, j := 0, len(nodes)-1; i < len(nodes)/2; i, j = i+1, j-1 {
		nodes[i], nodes[j] = nodes[j], nodes[i]
	}
}

// BidirecRouteToString returns a string representation of a route between two nodes.
func BidirecRouteToString(route []*BidirecNode) string {
	sb := strings.Builder{}
	for i, n := range route {
		if i > 0 {
			sb.WriteString(" -> ")
		}
		sb.WriteString(n.name)
	}

	return sb.String()
}
