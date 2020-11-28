package ch04

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"playground/ds"
	"testing"
)

func TestFindBuildOrder(t *testing.T) {
	projects := []string{"a", "b", "c", "d", "e", "f"}
	deps := [][]string{{"a", "d"}, {"f", "b"}, {"b", "d"}, {"f", "a"}, {"d", "c"}}
	buildOrder, err := FindBuildOrder(projects, deps)
	assert.NoError(t, err)

	graph := buildGraph(projects, deps)
	assert.NoError(t, build(buildOrder, graph))
}

func build(projects ds.Queue, graph *Graph) error {
	artifactory := make(map[string]bool)
	for !projects.IsEmpty() {
		proj, _ := projects.Remove()
		deps := graph.Nodes[proj].children
		for dep := range deps {
			if !artifactory[dep] {
				return errors.New("build failure")
			}
		}
		artifactory[proj] = true
	}
	return nil
}

func TestBuild(t *testing.T) {
	projects := []string{"a", "b", "c", "d", "e", "f"}
	deps := [][]string{{"a", "d"}, {"f", "b"}, {"b", "d"}, {"f", "a"}, {"d", "c"}}
	graph := buildGraph(projects, deps)

	buildQueue0 := ds.Queue{}
	buildQueue0.Add("f")
	buildQueue0.Add("e")
	buildQueue0.Add("a")
	buildQueue0.Add("b")
	buildQueue0.Add("d")
	buildQueue0.Add("c")

	assert.NoError(t, build(buildQueue0, graph))

	buildQueue1 := ds.Queue{}
	buildQueue1.Add("f")
	buildQueue1.Add("a")
	buildQueue1.Add("b")
	buildQueue1.Add("d")
	buildQueue1.Add("c")
	buildQueue1.Add("e")

	assert.NoError(t, build(buildQueue1, graph))

	buildQueue2 := ds.Queue{}
	buildQueue2.Add("e")
	buildQueue2.Add("a")
	buildQueue2.Add("b")
	buildQueue2.Add("f")
	buildQueue2.Add("d")
	buildQueue2.Add("c")

	assert.Error(t, build(buildQueue2, graph))

}
