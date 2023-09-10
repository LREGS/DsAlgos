def isAnagram(s, t):
    shared_letters = []
    first = [letter for letter in s]
    second = [letter for letter in t]

    for i in range(len(first)):
        if first[i] in second:
            print(first[i])
            shared_letters.append(first[i])
    print(shared_letters)
    if shared_letters < second:
        return False
    else:
        return True
            
    

one = ["anagram"]
two = ["margana"]

print(isAnagram(one, two))

#not right - just try and sort the strings in alphabetical order and compare strings... using built in function mayb