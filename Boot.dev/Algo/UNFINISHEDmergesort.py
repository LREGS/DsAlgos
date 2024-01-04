def merge(lh, rh):
    final = []
    
    i = 0
    j = 0

    for a, b in range(len(lh)):
        print(lh[0], rh[0])


def merge_sort(nums):
    
    if nums < 2:
        return nums
    

    m = len(nums) // 2
    lh = nums[0:m]
    rh = nums[m:-1]

    mlh = merge_sort(lh)
    mrh = merge_sort(rh)

    merged = merge(mlh, mrh)


lh = [0,1,2,3,4]
rh = [6,4,3,4,5,6]
merge(lh, rh)