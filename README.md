# leetcode-go

- LeetCode 题解项目，使用 [Go](https://go.dev/) 语言实现。

- 帮助开发者更好地了解和学习数据结构与算法。

## 获取代码

```git
git clone https://github.com/shilin1983/leetcode-go.git
```

## 安装依赖

```bash
go get github.com/stretchr/testify
```

```bash
go mod init && go mod tidy
```

## 运行测试

```shell
sh test.sh
```

## 题解列表

| 序号  |                                                 标题                                                 | 难度  |                                     方案                                      |
| :---: | :--------------------------------------------------------------------------------------------------: | :---: | :---------------------------------------------------------------------------: |
| 0001  |                          [两数之和](https://leetcode.cn/problems/two-sum/)                           | 简单  |                   [Go](src/solutions/problem0001/twoSum.go)                   |
| 0002  |                      [两数相加](https://leetcode.cn/problems/add-two-numbers/)                       | 中等  |               [Go](src/solutions/problem0002/addTwoNumbers.go)                |
| 0003  | [无重复字符的最长子串](https://leetcode.cn/problems/longest-substring-without-repeating-characters/) | 中等  | [Go](src/solutions/problem0003/longestSubstringWithoutRepeatingCharacters.go) |
| 0004  |        [寻找两个正序数组的中位数](https://leetcode.cn/problems/median-of-two-sorted-arrays/)         | 困难  |          [Go](src/solutions/problem0004/medianOfTwoSortedArrays.go)           |
