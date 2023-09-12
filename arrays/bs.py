def BinarySearch(sorted_array, target):
    """uses divide and conque to look through a sorted array until it finds its 
        target with an O(n) run time"""
    
    l = 0
    r = len(sorted_array) - 1

    while l <= r:
        m = (l + r) // 2
        if sorted_array[m] == target:
            return m
        elif sorted_array[m] > target:
            r = m - 1
        else:
            l = m + 1
    return -1


target = 7

sorted_arr = [2,3,4,5,6,7]


print(BinarySearch(sorted_arr, target))
