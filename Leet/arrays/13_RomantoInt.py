

def dank (s: str):
    roman = {
    'I': 1,
    'V': 5,
    'X': 10,
    'L': 50,
    'C': 100,
    'D': 500,
    'M': 1000
}

    total = 0
    for i in range(len(s) - 1):
        if roman[s[i]] < roman[s[i + 1]]:
            value = roman[s[i + 1]] - roman[s[i]]
            total += value
        elif roman[s[i]] >= roman[s[i + 1]]:
            value = roman[s[i]]
            total += value
    print(total)


dank('III')
