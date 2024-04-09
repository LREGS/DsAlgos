fibs = []

def fib(nums):
    numz = nums[0] + nums[1]
    for num in nums:
        fibs.append(fib(numz))
    print(fibs)

fib([1,2,3,4,5])
