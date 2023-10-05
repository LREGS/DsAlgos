
    # I think you need to identify the lowest number, and then slice the list
    #find the value with the largest difference (profit) between the lowest 
    #number and itself
def maxProfit(prices: list[int]) -> int:
    
    lowest_val = min(prices)
    print(f'lowest value is: {lowest_val}')

    lowest_idx = prices.index(lowest_val)

    new_prices = prices[lowest_idx:]
    print(new_prices)
    
    most_profit = 0

    for i in range(len(new_prices)):
        if (new_prices[i] - new_prices[0]) > most_profit:
            most_profit = new_prices[i] - new_prices[0]
        else:
            continue
        
    print(most_profit)


price = [4,3,5,3,1,2,3]
maxProfit(price)