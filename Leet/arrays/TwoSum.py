def twosum(arr, target):
    for i in range(len(arr)):
        for j in range(i + 1, len(arr)):
            if arr[i] + arr[j] == target:
                return i,j

"""
Pretty bad solution because its just brute force and worst case 
scenario is O(n^2)

A better approach is the two pointers approach

2184ms
Beats 37.71%of users with Python

14.41MB
Beats 12.09%of users with Python
"""

#Two Pointers Approach
class Solution(object):
    def twoSum(self, nums, target):
        # Create a dictionary to store the index of each element
        num_dict = {}
        
        for i, num in enumerate(nums):
            complement = target - num
            # Check if the complement exists in the dictionary
            if complement in num_dict:
                return [num_dict[complement], i]
            # Store the current number and its index in the dictionary
            num_dict[num] = i

        # If no solution is found, return an empty list or handle it as needed
        return []