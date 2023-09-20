def qs(arr, lo, hi):
    if  lo < hi:
     
        #gets pivot index in middle of arr
        pivotIdx = partition(arr, lo, hi)

        #this runs to branches to the right(higher) and left (lower) of the pivot creating 
        #more branches until eventually almost every value has been the pivot and sorted
        #or was one of the remaining values that got sorted left or right of the final pivot
        qs(arr, lo, pivotIdx - 1)
        qs(arr, pivotIdx +1, hi)

def partition(arr, lo, hi):
    
    #make the pivot the last element in the array
    pivot = arr[hi]

    #idex before the array to find the first time a number is less than the pivot and make sure its 
    #in the first element in the array at the begining
    idx = lo -1

    for i in range(lo, hi):
        if arr[i] <= pivot:
            #this moves the index initially to the first element in the aray and will swap the first 
            #element it finds that is smaller than the pivot
            idx = idx +1 
            #swaps the elements
            arr[i], arr[idx] = arr[idx], arr[i]
    
    idx = idx +1
    arr[hi], arr[idx] = arr[idx], arr[hi]
    

    return idx



arr = [9,4,5,3,2,7,9,3,2,5,66,5,34,4,43,2,5,1]
size = len(arr)
ss = qs(arr, 0, (size - 1))
print(arr)