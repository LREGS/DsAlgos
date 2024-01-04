def insertionSort(nums):
    for i in range(1, len(nums)):
        j = i

        while j > 0 and nums[j-1] > nums[j]:
            nums[j-1], nums[j] = nums[j], nums[j-1]
            j -= 1

    return nums

nums = [9,6,4,3,2]
sorted = insertionSort(nums)
print(sorted)

