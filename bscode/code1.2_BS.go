package bscode

/*
	1.有序数列中查找num
	2.有序数列中查找>=num最左侧的位置
	3.局部有序数列中查找最小元素
*/
// BSExist
func BSExist(nums []int, num int) bool {
	if len(nums) == 0 {
		return false
	}
	left, right := 0, len(nums)-1
	for left < right {
		mid := (left + right) / 2
		if nums[mid] > num {
			right = mid - 1
		} else if nums[mid] < num {
			left = mid + 1
		} else {
			return true
		}
	}
	return false
}

// BSNearLeft
func BSNearLeft(nums []int, num int) int {
	if len(nums) == 0 {
		return -1
	}
	position := 0
	left, right := 0, len(nums)-1
	for left < right {
		mid := (left + right) / 2
		if nums[mid] >= num {
			position = mid
			right = mid - 1
		} else {
			left = mid + 1
		}

	}
	return position
}

// BSMin
func BSMin(nums []int) (index, value int) {
	if len(nums) == 0 {
		return -1, -1
	}
	if len(nums) == 1 {
		return 0, nums[0]
	}

	if nums[0] < nums[1] {
		return 0, nums[0]
	}

	if nums[len(nums)-1] < nums[len(nums)-2] {
		return len(nums) - 1, nums[len(nums)-1]
	}
	/*
		nums: 9,8,7,5,10,13,15,18
		1.left=1,nums[left]=8,right=6,nums[right]=15, mid=(1+6)/2=3, nums[mid]=5,nums[mid+1]=10,nums[mid-1]=7

		nums:12,1,5,8,9,20
		1.left=1,nums[left]=1,right=4,nums[right]=9, mid=(1+4)/2=2, nums[mid]=5, nums[mid-1]=1,nums[mid+1]=8
		2.left=1,nums[left]=1,right=1,nums[rigth]=1, mid=(1+1)/2=1, nums[mid]=1, nums[mid-1]=12,nums[mid+1]=5

		nums:12,1,1,5,8,9,20
		1.left=1,nums[left]=1,right=5,nums[right]=9, mid=(1+5)/2=3, nums[mid]=5, nums[mid-1]=1,nums[mid+1]=8
		2.left=1,nums[left]=1,right=2,nums[rigth]=1, mid=(1+2)/2=1, nums[mid]=1, nums[mid-1]=12,nums[mid+1]=1

		nums: 2,1
		1.left=1,nums[left]=1,right=0,nums[right]=2,left>right, return left=1,nums[left]=1


		nums:4,2,3
		1.left=1,nums[left]=2,right=1,nums[right]=2,left>=right,return left=1,nums[left]=2

	*/
	left, right := 1, len(nums)-2
	for left < right {
		mid := (left + right) / 2
		if nums[mid] > nums[mid+1] && nums[mid] < nums[mid-1] {
			left = mid + 1
		} else if nums[mid] < nums[mid+1] && nums[mid] > nums[mid-1] {
			right = mid - 1
		} else {
			return mid, nums[mid]
		}
	}
	//  为啥要用left返回
	return left, nums[left]
}
