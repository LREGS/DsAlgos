import random

def qs(arr, lo, hi):
    if lo < hi:
    
        pidx = partition(arr, lo, hi)

        qs(arr, lo, pidx - 1)
        qs(arr, pidx +1, hi)

def partitionrand(arr, lo, hi):
    pivot = random.randrange(lo, hi)

    arr[lo], arr[pivot] = arr[pivot], arr[lo]

    return partition(arr, lo, hi)

def partition(arr, lo, hi):
    
    pivot = lo

    idx = lo + 1 

    for i in range(lo + 1, hi + 1):
        if arr[i] <= pivot:
            idx = idx + 1
            #moves the indx right which  is the pointer to where the pivot will need to go
            arr[i], arr[idx] = arr[idx], arr[i]

    arr[pivot], arr[i - 1] = arr[i - 1], arr[pivot]

    pivot = i-1

    return (pivot)

if __name__ == "__main__":
    array = [10, 7, 8, 9, 1, 5]
    qs(array, 0, len(array) - 1)
    print(array)
    
    
