

def MajorityElement(nums: list) -> int:
    numbers = dict()
    for num in nums:
        numbers[num] = 1
    
    # for num in nums:
    #     for key in numbers:
    #         if num == key:
    #             numbers[key] += 1
    
    top_letter = max(numbers, key=numbers.get)

    print(numbers)



lst = [1,2,3,3,4,5,6,7,4,5,6,5]
MajorityElement(lst)

"""No understanding of what majority element means. 

Majority element means to appear most within the array. So a majority element 
within an array length of ten would need to appear six times. 

Even if everything appeared once and something twice the repeated element isn't 
a majority element because of this. 

With this rule in mind you can sort the list and look in the middle, if there is a 
majority element it will always be in the middle because it must take up more 
than half of the list"""