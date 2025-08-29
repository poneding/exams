/*
 * @lc app=leetcode.cn id=70 lang=java
 *
 * [70] 爬楼梯
 *
 * https://leetcode.cn/problems/climbing-stairs/description/
 *
 * algorithms
 * Easy (55.11%)
 * Likes:    3760
 * Dislikes: 0
 * Total Accepted:    1.8M
 * Total Submissions: 3.2M
 * Testcase Example:  '2'
 *
 * 假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
 * 
 * 每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：n = 2
 * 输出：2
 * 解释：有两种方法可以爬到楼顶。
 * 1. 1 阶 + 1 阶
 * 2. 2 阶
 * 
 * 示例 2：
 * 
 * 
 * 输入：n = 3
 * 输出：3
 * 解释：有三种方法可以爬到楼顶。
 * 1. 1 阶 + 1 阶 + 1 阶
 * 2. 1 阶 + 2 阶
 * 3. 2 阶 + 1 阶
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= n <= 45
 * 
 * 
 */

// @lc code=start
class Solution {
    // 爬 N 阶楼梯：第一步 1 或者 2，后面就是 爬 N-1 或 N-2 阶楼梯的方法
    // fn(n) = fn(n-1) + fn(n-2)
    // 斐波那契
    public int climbStairs(int n) {
        // 1 阶 1 种方法
        if (n == 1) {
            return 1;
        }
        // 2 阶 2 种方法
        if (n == 2) {
            return 2;
        }

        int a = 1;
        int b = 2;
        int c = 0;
        for (int i = 3; i <= n; i++) {
            c = a + b;
            a = b;
            b = c;
        }
        return b;
    }
}
// @lc code=end
