package ch04

import "playground/ds/tree"

func FindPathsWithWightV0(root *tree.BTNode, targetSum int) int {
	if root == nil {
		return 0
	}
	pathsFromRoot := countPathFromRoot(root, targetSum, 0)
	pathCountsLeft := FindPathsWithWightV0(root.Left, targetSum)
	pathCountsRight := FindPathsWithWightV0(root.Right, targetSum)
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
func FindPathsWithWightV1(root *tree.BTNode, targetSum int) int {
	if root == nil {
		return 0
	}
	return FindPathsWithWightV1Helper(root, targetSum, make(map[int]int), 0)
}

func FindPathsWithWightV1Helper(root *tree.BTNode, targetSum int, sumMap map[int]int, pathSum int) int {
	pathCount := 0
	if root == nil {
		return 0
	}
	pathSum += root.Data.(int)
	delta := pathSum - targetSum
	pathCount += sumMap[delta]

	if pathSum == targetSum {
		pathCount++
	}
	sumMap[pathSum] += 1
	pathCount += FindPathsWithWightV1Helper(root.Left, targetSum, sumMap, pathSum)
	pathCount += FindPathsWithWightV1Helper(root.Right, targetSum, sumMap, pathSum)
	sumMap[pathSum] -= 1
	return pathCount
}
