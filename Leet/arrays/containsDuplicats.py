class Solution(object):
    def containsDuplicate(self, nums):
        
        check_duplicate = set()
        for num in nums:
            if num in check_duplicate:
                return True
        else:
            check_duplicate.add(num)

        return False