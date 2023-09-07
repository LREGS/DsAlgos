def binary_search(sorted_list, target):

     left = 0
     right = len(sorted_list) - 1 #-1 stops us going over the array end

     while left <= right:
        #updates where to search as left and mid are adjusted which
        #controls the halving of the array 
        mid = (left + right) //2

        if sorted_list[mid] == target:
            return sorted_list[mid]
        elif sorted_list[mid] < target:
            #if target is right hadn its plus one(bigger number = +)
            left = mid + 1
        else:
            #if the target is left half its -1
            right = mid - 1 
            
