def removeDuplicates(nums):
    # s = set(nums)   
    # sorted = [num for num in s]

    # for i in range(len(sorted)):
    #     for j in range (len(sorted) -1 -i):
    #         if sorted[j] > sorted[j+1]:
    #             sorted[j], sorted[j+1] = sorted[j+1], sorted[j]        duplicates = []7
    un = []
    for num in nums:
        if num not in un:
            un.append(num)

    x = slice(len(un))
    print(nums, un)


array = [0,0,1,1,1,2,2,3,3,4]
n = removeDuplicates(array)     
print(n)  

#not accepting answer because I didn't remove elements in place - so although my output is correct it wasn't
#done the way leetcode wanted me to complete the question so will try again