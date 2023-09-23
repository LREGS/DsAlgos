#not there yet - HINT - track the majority element against a count?!

def majorityElement(arr):
    map = {}
    for num in arr:
        if map[num]:
            map[num] += 1
        elif not map[num]:
            map[num] = 1
    print(map)

arr = [1,1,2,1,2]

majorityElement(arr)

