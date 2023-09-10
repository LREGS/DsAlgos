def binary_search(sorted_array, target):
    #defines left and right side of the array
    left = 0
    right = len(sorted_array) - 1

    while left<=right:
        #gets the mid point of the array whlst the left side still has values 
        mid = (left + right) // 2

        #if you get it first time return the values a
        if sorted_array[mid] == target:
            return mid
        #otherwise, if its smaller include only the left hand side
        elif sorted_array[mid] < target:
            left = mid + 1
        #or if its larger include the right hand side 
        elif sorted_array[mid] > target:
            right = mid -1
            
    return -1

target = 2
sorted_array = [1,2,3,4,5,6,7,8]

index = binary_search(sorted_array, target)

print(index)

#Has a O(logn) run time complexity because you're halving the amount of operations with each cycle of the algorithms'

