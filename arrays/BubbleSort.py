def bubble_sort(array):
    #Gets the first element
    for i in range(len(array)):
        #Gets the second element
        for j in range(len(array) -1 - i):
            #checks the first and second selected element
            if array[j] > array[j + 1]:
                #if the first is higher it swaps to the right. (Bigger bubbles rise)
                array[j], array[j + 1] = array[j + 1], array[j]
    #returns the sorted array 
    return array


sort = [4,3,2,8,6,3,5]

sorted = bubble_sort(sort)

print(sorted)

