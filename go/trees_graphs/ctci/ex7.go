package ctci

import (
	"errors"
	"sort"
)

// 4.7 Build Order: You are given a list of projects and a list of dependencies
// (which is a list of pairs of project, where the second project is dependent
// on the first project). All of a project's dependencies must be built before
// the project is. Find a build order that will allow the projects to be built.
// If there is no valid build order, return an error.

type Dependency [2]string

// Build order traverses each project's dependencies, checking for cycles. If a
// cycle is found, an error is returned. If the dependency chain is free of
// cycles it is added to the build order in reverse, starting with the project
// with no dependencies. Once all keys have been traversed, the build order is
// returned.
//
// Requires O(n) space to represent the graph in a map and to maintain the slice
// containing the build order and the maps which track which projects have
// already been built. We check each vertex (project) and every edge (dependecy)
// of that vertex, so the time complexity is O(V + E).
//
// Assumptions:
//   * If a cycle exists, we don't need to describe the cycle in the error.
//   * Build order should be deterministic.
func BuildOrder(projects []string, deps []Dependency) ([]string, error) {
	buildOrder := make([]string, 0, len(projects))
	graph := buildAdjacencyList(projects, deps)
	built := make(map[string]bool)
	// A set to check for circular dependencies while building each project.
	// To return the chain of projects containing a cycle in an error message,
	// use a stack.
	inCurrentDependencyChain := make(map[string]bool)

	var generateBuildOrder func(project string) error
	generateBuildOrder = func(project string) error {
		// Check for cycles.
		if inCurrentDependencyChain[project] {
			return ErrNoBuildOrder
		}
		inCurrentDependencyChain[project] = true

		// Skip projects that are already built.
		if !built[project] {
			for _, dep := range graph[project] {
				// Recurse to the bottom of the dependency chains.
				err := generateBuildOrder(dep)
				if err != nil {
					return err
				}
			}

			buildOrder = append(buildOrder, project)
			built[project] = true
		}

		inCurrentDependencyChain[project] = false
		return nil
	}

	// Sort the project list to ensure deterministic build order.
	sort.Strings(projects)
	for _, proj := range projects {
		err := generateBuildOrder(proj)
		if err != nil {
			return nil, err
		}
	}

	return buildOrder, nil
}

func buildAdjacencyList(projects []string, deps []Dependency) AdjacencyList {
	graph := make(AdjacencyList)

	// Add each project's immediate dependencies to the graph.
	for _, dep := range deps {
		proj := dep[1]
		graph[proj] = append(graph[proj], dep[0])
	}

	return graph
}

var ErrNoBuildOrder = errors.New("no valid build order")
