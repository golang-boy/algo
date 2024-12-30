/*
题目：

给定N种物品，其中第i种物品的体积为Vi,价值为Wi, 且有**无数个**。容积为M的背包，
要求选择放入一些物品，使得总体积不超过M时，物品价值总和最大

*/

package main

func main() {

	/*

		    f[i][j] 从前i种物品中选择出总体积为j的物品放入背包，物品总值价值最大

		    f[i][j] = f[i-1][j]  不选第i种时

		    if j>Vi{

			    // f[i][j] = f[i-1][j-k*vi] + k*wi
				//  优化为下列
			    f[i][j] = f[i][j-Vi] + Wi
		    }

	*/

	f := make([]int, m+1)

	for i := 1; i < n; i++ {
		for j := v[i]; j <= m; j++ {
			// 完全背包时，得正序
			f[j] = max(f[j], f[j-v[i]]+w[i])
		}
	}

	ans := 0
	for j := 0; j <= m; j++ {
		ans = max(ans, f[j])
	}
}
