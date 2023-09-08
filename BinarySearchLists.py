def binary_search(sorted_list, target):
    #allows to get the middle value of the list, values change as the array is divided
    left = 0
    right = len(sorted_list) - 1

    while left <= right:
        #gets the value to get the middle index of the array 
        m = (left+right) // 2
        
        #if equal return
        if sorted_list[m] == target:
            return m
        #narrows the size of the search by essentialy removing
        #the left side of the list
        elif sorted_list[m] < target:
            left = m + 1
        else:
        #otherwise move the right side of the list 
            right = m - 1
    
    return - 1

l = [1,2,3,4,5,6]
t = 3

print(binary_search(l,t))
