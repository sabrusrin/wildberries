package leetcode_twosum

func TwoSum(nums []int, target int) []int	{
	m := make(map[int]int)
	var i, num, num2 int
	pair := []int{}
	ok := false
	for i, num = range nums	{
		m[num] = i
	}
	for i, num = range nums	{
		num2 = target - num
		if _, ok = m[num2]; ok	{
			if i != m[num2]	{
				pair = append(pair, i, m[num2])
				goto end
			}
		}
	}
	end:
	return pair
}
