def linear_search(list, v):
    for i in range(len(list)):
        if list[i] == v:
            print(v)
            break
        else:
            print('no there')

v = 2
l = [2,3,5,6,4]
linear_search(l,v)