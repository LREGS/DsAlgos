
def LargestPrefix(words: list) -> str:
    for word in words:
        for letter in word:
            






list1 = ['flow', 'flower', 'flight']
list2 = ['car', 'racers', 'dog']

'''Find the longest common prefix if there is one, else return an empty string 

I think you just want to loop through each letter and iterate a counter after each letter is the same. 
If no letters match theres no prefix, and when there is a prefix it will be the
same length for everyone

dont count, just track the letters that are the same, if there are any the string will be populated, otherwise, it wont be. In the event all 
the words start with the same letter but dont contain any prefixes we should only count strs with lengths greater than 1 to be considered '''

