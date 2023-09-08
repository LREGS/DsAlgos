def bubblesort(array):

    for i in range(len(array)):
        for j in range(len(array) - 1 -i):
            if array[j] > array[j + 1]:
                array[j], array[j + 1] = array[j + 1], array[j]
    return array


array = [9,8,7,6]
sorted_array = bubblesort(array)

print(sorted_array)