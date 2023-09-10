def binary_search(sorted_list, target):

    left  = 0 
    right = len(sorted_list) - 1

    while left <= right:
        middle = (left + right) // 2

        if sorted_list[middle] == target:
            return middle
        elif sorted_list[middle] > target:
            right = middle - 1
        elif sorted_list[middle] < target:
            left = middle + 1 

    return -1  

sorted_list = [1,2,3,4,5,6,7,8,9]
target = 5  

print(binary_search(sorted_list, target))