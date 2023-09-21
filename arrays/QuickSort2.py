"""Assign a pivot within the array and sort the elements so the lowest values are to the left
of the pivot and highest to the right. The pivot should then be moved to the middle of the array
and two recursive branches will re-sort all the values to the right and left of the pivot.

The pivot is sorted and the function recurses in a tree like manner until all the values are sorted. 

Sometimes this can be really slow in the example of a reverse sorted list with your pivot furtherst right.

[9,8,7,6,5,4,3,2,{1}]

 because 9 would move to the right of 1, and then be looking across the whole array 
 to look for values larger than it, but there wont be one so it goes back to one, either is placed 
 as the pivot and 9 is sorted and it will keep going. 
 
 I think this is O(n2)
 
 Although if you got it perfectly in the middle of a sorted array its O(logn) becuase the number 
 of required operations halves with each iterations?? (plz fact check me)"""

def qs (arr, hi, lo):
    if lo < hi:

        pivotIdx = partition(arr, lo, hi)

        qs(arr, lo, pivotIdx -1)
        qs(arr, pivotIdx+1, hi)
        

def partition(arr, hi, lo):

    pivot = arr[hi]

    for i in range(lo, hi):
        if arr[i] <= pivot:
            idx = idx + 1
            arr[i], arr[idx] = arr[idx], arr[i]
    
    idx = idx + 1
    arr[hi]. arr[idx] = arr[idx], arr[hi]

    return idx