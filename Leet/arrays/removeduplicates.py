def removeDuplicates(nums):
    s = set(nums)
    sorted = [num for num in s]

    for i in range(len(sorted)):
        for j in range (len(sorted) -1 -i):
            if sorted[j] > sorted[j+1]:
                sorted[j], sorted[j+1] = sorted[j+1], sorted[j]

    print(sorted)

array = [1,1,2]
removeDuplicates(array)     

#not accepting answer because I didn't remove elements in place - so although my output is correct it wasn't
#done the way leetcode wanted me to complete the question so will try again