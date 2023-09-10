def bubblesort(array):

    #shortens the searchable list with every iteration as sorted values arrive in their correct position 
    for i in range(len(array)):
        #gets the index for the values to be sorted
        for j in range(len(array) - 1 -i):
            #big bubbles rise to the right
            if array[j] > array[j + 1]:
                array[j], array[j + 1] = array[j + 1], array[j]
    return array


array = [9,8,7,6]
sorted_array = bubblesort(array)

print(sorted_array)