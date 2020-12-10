package ch05

//m->n; n[j:i]=m
//100000 110 j
//  110
//j=3 i=1
func Insert(m, n int, j, i int) int {
	mask1 := 1 << j       //10000
	mask1 = mask1 - 1     //01111
	mask1 = ^mask1        //10000
	mask2 := 1 << i       //00010
	mask2 = mask2 - 1     //00001
	mask := mask1 | mask2 //11110001
	n1 := n & mask        //
	m = m << i            // 1100
	return n1 | m
}
