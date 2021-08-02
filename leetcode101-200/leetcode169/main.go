package main

// 常规方法，时间复杂度O(n), 空间复杂度O(n)
func majorityElement1(nums []int) int {
	// 异常情况处理
	if len(nums) == 1 {
		return nums[0]
	}

	// 基本情况处理
	var max, maxCount int
	m := make(map[int]int)

	for _, num := range nums {
		v, ok := m[num]
		if !ok {
			m[num] = 1
		} else {
			if v >= maxCount {
				max = num
				maxCount = v + 1
			}
			m[num] = v + 1
		}
	}

	return max
}

// 多数投票算法
// 博耶-摩尔多数投票算法（英语：Boyer–Moore majority vote algorithm）,中文常作多数投票算法、摩尔投票算法等，是一种用来寻找一组元素中占多数元素的常数空间级时间复杂度算法。这一算法由罗伯特·S·博耶和J·斯特罗瑟·摩尔在1981年发表[1]，也是处理数据流的一种典型算法。
// 这一算法应用的问题原型是在集合中寻找可能存在的多数元素，这一元素在输入的序列重复出现并占到了序列元素的一半以上；在第一遍遍历之后应该再进行一个遍历以统计第一次算法遍历的结果出现次数，确定其是否为众数；如果一个序列中没有占到多数的元素，那么第一次的结果就可能是无效的随机元素。对于数据流而言，则不太可能在亚线性空间复杂度的情况下中就寻找到出现频率最高的元素；而对于序列，其元素的重复次数也有可能很低。

// 算法可以用伪代码如下表示：

// 初始化元素m并给计数器i赋初值i = 0
//   对于输入队列中每一个元素x：
//     若i = 0, 那么 m = x and i = 1
//     否则若m = x, 那么 i = i + 1
//     否则 i = i − 1
//   返回 m
// 即便输入序列没有多数元素，这一算法也会返回一个序列元素。然而如果能够进行第二轮遍历，检验返回元素的出现次数，就能判断返回元素是否为多数元素。因此算法需要两次遍历，亚线性空间算法无法通过一次遍历就得出输入中是否存在多数元素。

// 摩尔投票法，时间复杂度O(n), 空间复杂度O(1)
func majorityElement2(nums []int) int {
	var max, maxCount int
	for _, num := range nums {
		if maxCount == 0 {
			max = num
			maxCount = 1
		} else if max == num {
			maxCount += 1
		} else {
			maxCount -= 1
		}
	}

	return max
}
