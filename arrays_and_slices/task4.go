/*
Слияние

Даны 2 слайса целых чисел nums1 и nums2. Создайте функцию Join(nums1, nums2 []int) []int,
которя создаст новый слайс емкостью, вмещающей в себя ровно два слайса (ёмкость должна быть равна его длине).
Скопируйте в него сначала значения nums1 затем nums2 и верните его.
*/
package arraysandslices

func Join(nums1, nums2 []int) []int {
	l1 := len(nums1)
	l2 := len(nums2)
	res := make([]int, l1+l2)
	for i := 0; i < l1; i++ {
		res[i] = nums1[i]
	}
	for i := 0; i < l2; i++ {
		res[i+l1] = nums2[i]
	}
	return res
}
