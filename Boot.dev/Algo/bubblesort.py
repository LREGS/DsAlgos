def bubbleSort(nums):
    for i in range(len(nums)):
        for j in range(len(nums) -1 -i):
            if nums[j] > nums[j+1]:
                nums[j], nums[j+1] = nums[j+1], nums[j]
            else:
                continue
    return nums



nums = [1,3,4,43,6,7,1,9,3,8,2,6]
bubbleSort(nums)
print(nums)