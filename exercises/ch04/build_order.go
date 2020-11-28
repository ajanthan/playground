package ch04

import (
	"errors"
	"playground/ds"
)

const (
	EMPTY = iota
	PARTIAL
	COMPLETE
)

type Node struct {
	Name     string
	children map[string]*Node
	status   int
}
type Graph struct {
	Nodes map[string]*Node
}

//deps[0][0]<-deps[0][1]
func buildGraph(projects []string, deps [][]string) *Graph {
	graph := &Graph{}
	graph.Nodes = make(map[string]*Node)
	for _, proj := range projects {
		graph.Nodes[proj] = &Node{Name: proj, children: make(map[string]*Node)}
	}
	for _, proj := range deps {
		graph.Nodes[proj[1]].children[proj[0]] = graph.Nodes[proj[0]]
	}
	return graph
}

func FindBuildOrder(projects []string, deps [][]string) (ds.Queue, error) {
	graph := buildGraph(projects, deps)
	buildOrder := &ds.Queue{}
	for _, proj := range graph.Nodes {
		err := findBuildOrderHelper(proj, buildOrder)
		if err != nil {
			return nil, err
		}
	}
	return *buildOrder, nil
}

func findBuildOrderHelper(node *Node, order *ds.Queue) error {
	if node.status == PARTIAL {
		return errors.New("no build order")
	}
	if node.status == EMPTY {
		node.status = PARTIAL
		for _, child := range node.children {
			err := findBuildOrderHelper(child, order)
			if err != nil {
				return err
			}
		}
		order.Add(node.Name)
		node.status = COMPLETE
	}
	return nil
}
