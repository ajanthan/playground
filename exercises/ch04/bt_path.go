package ch04

import "playground/ds/tree"

func FindPathsWithWight(root *tree.BTNode, targetSum int) int {
	if root == nil {
		return 0
	}
	pathsFromRoot := countPathFromRoot(root, targetSum, 0)
	pathCountsLeft := FindPathsWithWight(root.Left, targetSum)
	pathCountsRight := FindPathsWithWight(root.Right, targetSum)
	return pathsFromRoot + pathCountsLeft + pathCountsRight

}

func countPathFromRoot(root *tree.BTNode, targetSum, pathSum int) int {
	pathCounts := 0
	if root == nil {
		return pathCounts
	}
	pathSum += root.Data.(int)
	if pathSum == targetSum {
		pathCounts++
	}
	pathCountsLeft := countPathFromRoot(root.Left, targetSum, pathSum)
	pathCountsRight := countPathFromRoot(root.Right, targetSum, pathSum)
	return pathCounts + pathCountsLeft + pathCountsRight
}
