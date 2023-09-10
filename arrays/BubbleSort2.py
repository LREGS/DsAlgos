#checks the value to the left and moves it right if its larger
#big bubbles rise small bubbles fall 

def bubble_sort(arr):
    #sets the boundary and ensures that the largest, and thus correctly sorted values 
    #are no longer checked on the next loops 
    for v in range(len(arr)):
        print(arr)
        #checks the array and swaps the values if the one on the left is largest.
        #bubble always rises to the top 
        #-1 -v shrinks the array so that the loop doesn't bother checking the already sorted 
        #largest values 
        for j in range(len(array) -1 - v):
            print(arr)
            if arr[j] > arr[j + 1]:
                #swaps the larges values to the right
                arr[j], arr[j + 1] = arr[j + 1], arr[j]
    return arr

array = [9,8,7,6]
sorted__array = bubble_sort(array)
print(sorted__array)