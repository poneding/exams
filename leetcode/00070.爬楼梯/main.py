#
# @lc app=leetcode.cn id=70 lang=python
#
# [70] 爬楼梯
#
# https://leetcode.cn/problems/climbing-stairs/description/
#
# algorithms
# Easy (55.11%)
# Likes:    3760
# Dislikes: 0
# Total Accepted:    1.8M
# Total Submissions: 3.2M
# Testcase Example:  '2'
#
# 假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
#
# 每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
#
#
#
# 示例 1：
#
#
# 输入：n = 2
# 输出：2
# 解释：有两种方法可以爬到楼顶。
# 1. 1 阶 + 1 阶
# 2. 2 阶
#
# 示例 2：
#
#
# 输入：n = 3
# 输出：3
# 解释：有三种方法可以爬到楼顶。
# 1. 1 阶 + 1 阶 + 1 阶
# 2. 1 阶 + 2 阶
# 3. 2 阶 + 1 阶
#
#
#
#
# 提示：
#
#
# 1 <= n <= 45
#
#
#

# @lc code=start
class Solution(object):
    def climbStairs(self, n):
        """
        :type n: int
        :rtype: int
        """
        a = 1
        b = 1
        c = 0
        if n == 1:
            return 1
        for i in range(2, n + 1):
            c = a + b
            a = b
            b = c
        return c


# @lc code=end
