# Big O Time Complexity 

## On
Big o categorizing algorithm of time and memory based on input. To make it generalised for hardware etc it will give a general idea of how the time taken to complete the work will grow with how the inputs grow. 

An easy way to tell a big o complexity is within loops. If the loops are defined to keep going based on the size of your input then the complexity will be On?

Constants aren't impornant, if you have an On algorithm but you have 2 instances of it its not O 2n because its still just growing linarly with the inputs.

let sum = 0

for i in range(len.string):
    if i == "E"
        reutrn "E found"

This will loop over a string and return a string if an E is found. Although this could only take one operation, its still O of N time complexity because the worst case scenario is that the programme checks the entire string and doesn't find a capital E. 

We should almost always in an interview situation look at the worst case 


## O(N^2)
Goes over every character of every character 

## O(n log n)
Quicksort is an example - halfs the amount of every space you search but it hals the amount of space every time 

## Three important parts of Big O
1 - Growth is with respect to input 
2 - Cosntants are dropped 
3 - Worst case is usuallt the way we measure 


# Arrays and Data Structures

An array is basically 0 or more pieces of memory in a row of continuos memory ready to be used 

so a = int[3,2,1]

is an array of 3 values stored continuosly next to each other. Then when you index the array a[1] you're telling the computer where to look in the memory to find the value. If its the first value it will be 4 bits along the memory space.

Accessing memory in an array is O of 1 because we know the length of the array already, and the position in the array we want to be so the computer can go directly to the value within a constant amount of operations that don't grow with input at all. 

## Array:
- Fixed sized 
- They cannot grow 
- There is no push, pop, InsertAt etc 


# Linear Search & Kata Setup 

Linear search can be used on an array of values. 

- The first step is to look at the first value and check if its there 
- If it isn't it will keep walking along the array until it finds the value asked for.
- Has an O of N running time O(N)


# Binary Seach Algorithm

 Running time of LogN. 

 Using a sorted array, the algorithm will look at the middle value to see if its the value we're looking for. Depending on if its larger or smaller than what we're looking for the algorithm will then search the required half of the array only after that, continually halving until we find the value or not. 

 Because we are not lineary jumping between every value then the time constraints are not directly linked to the size of the list. 

 
The worst case scenario is logN because you will need to keep halving until you find out that you don't have the value in the array. 
 
## Big O Trick 
If the input halves at each step its likely O(LogN) or O(NlongN)

## Pseudo Code 

func search (arr, lo, hi):

midpoint = lo + (hi - o) / 2 

value = arr[m]

if value = needle 
return true

else if v > neede 
lo = m + 1

else
high = midpoint

# The two crystal ball problem

Given two crystakl balls at which height do they need to be dropped for them to break.

This is an array of falses which at a given value will switch to trues

If you use binary search and get to the middle and it says true to being broke, then you only have one crystal ball left so you will have to walk linearly through the first half of the array meaning time complexity is O(N) again.

The generalised problem is to jump N^2 and then once it breaks we only need to do a linear search on a smaller section and the runtime is O(n^2)


 ## implementing two crystak balls 

export default function two_crystal_balls(breaks: boolean[]): number {

    const jmpAmount = Math.floor(Math.sqrt(breaks.length));
     
    let i = jmpAmount;

    for(; i < breaks.length; i += jmpAmount){
        if (breaks[i]){
            break
        }
    }

    i -= jmpAmount;

    for (let j = 0; j < jmpAmount && i < breaks.length; ++j, ++i){
        if (breaks[i]) {
            return i;
        }
    }
    return -1 
}   


using the square root of n it is the only way of being non linear because there is only two chances to search so the running time it lower than O(N) but not as fast as O(logN)

# Bubble Sort 

very simple algorithm to sort things 

An array that is sorted is Xi < Xi + 1

Bubble sort starts in the first position and checks if its larger or smaller than the value next to it, and if it is they will swap.

A singular iteration will always get the largest item to the end of the array so next time we do the bubble sort we dont need to consider teh last position.

This means the array in question can smaller with every interation

Bubble sort is O(n2)

export default function bubble_sort(arr: number[]): void {
sort = [4,3,2,8,6,3,5]

    // Gets the first element in the array
    for (let i = 0; i < arr.length; i++){
        //gets the second element in the array
        for (let j = 0; j < arr.length - 1 - i; ++j){
            // compares the first to the second
            if(arr[j] > arr[j+1]) {
                // stores the larger of the value to move right 
                const tmp = arr[j];
                //move the larger value to the right (bubble rises, smaller sink)
                arr[j] = arr[j + 1];
                arr[j + 1] = tmp;
            }
        }
    }
}


# Linked Lists  

Linked lists is a node based data structure 

You have a node that contains a value and a referance to another node, so you can walk between values by using the references in the nodes. However, you can always walk forwards because they don't have referance to the nodes behind them,


However, in a doubly linked list you do have a referance to the node before it so you can go forwards and backwards 

Deletion and insertion can be very fast with a linked list. 
You can insert into any point in the list by changing the referances and you can move everything over to the right or the left within the lists by using the next and previos functions. 

No other nodes within the list needed to be accessed so it operates at constant time O(1). Nothing is based of input and you only need to break 4 links. 

To delet something you would need to assign the previous value to the pointer of the value in front of its pointer to the value behind it, and then have the value it points to be the value of the node you're deleting was pointing too eg:

val: a, aNext = b, prev = null,: val:b, bprevious = a, bnext = c, cnxt = d

b = c.previous 
b.next = c.next 

d will then need to be changed to point previously to b, which is c.previous 

c.prev = c.next = null - this will delete the value from the list

this is because b isn't aware of d, it only has referance to a, and c. The same way c isn't aware of a. Because these values are stored in a continious chain of memory it allows us to structure data in a uniformed way without all our pieces of data needing to know about every pother ppiece of data.  

In a linked list there is no index, so the only way your values are stored and through the referances contained in each node

## Linked List Complexity 

If you delete from head or tail is O(1) and a deletion in the middle can be more costly because you will need to traverse across the route which is O(n) and then run the deletion algo 

# Queue Data Structure 
> An, > bn > Cn > Dn < tail 

In a queue when you add an item to the data structure you always add it at the end of the list which means you need to change the tail to point to e, instead of the d node 

> An, > bn > Cn > Dn > tail > E
> An, > bn > Cn > Dn > E < tail -  this is what is left after the tail points to the new piece of data 

A Queue run time is only 2O(1) becuase you're only ever doing 2 operations which is changing where the tail points and then where the last value - 1 is pointing to the last value 

# Stacks

Similar to queue but you only remove and replace at the head of the stack. 

# Array vs Linked List 

- arrays give you indice values 
- under the hood activities are a little easier 
- Always O(1)

- you have to allocate the memory upfront, if you want to store 1000 items you need to alloate the memory upfront but might not use it all. 
    - with a linked list the memory usage is more optimised beause you only use it as a you assign new nodes, however, the runtime implications of this make it slower to use 

- linked lists will always use less memory but have worse runtime implications because you always have to use linear search 

- if you want to scan a list of use values from the middle use an array, otherwise you could use a linked list 

# Array List

we have an array 0[2,   ]3

this has a length of 1 but a capacity of 3, so we can store more than we currently are as thats how we built the array. 

If we wanted to push a value, we could go to the last piece of data, and see whether it is within the capacity of the array. If it is, we can push the data 

If we wanted to push data into the array but its outside of capacity, the array list will create a longer array, and copy the old array into it and then push the data into the end. 

This means you don't use too much memory at the begining, and can grow dynamically as you need to. Realistically, you want to balance the trade off from using the least amount of memory at the begining hilst also doing the least amount of growing after its been initialised.

## enqueue 
You want to avoid doing this, as because its an array, you would essentially be overwriting the value in space one, so instead the array list will shift every value to the right and then insert the queue value. 
In a queue linked list this operation would be O(1) vs the O(N) complexity in an array list - so its much less performant in an array list and should be avoided.

It's the same with a dequeue 

IMPLEMENTATION TO BE DONE

# Array(ring) buffers 

0[    ]n

In an array buffer we don't use 0 as head and length as the tail like an array list. 
    - we can have an index based head and tail instead so everything within the head and tail are the items we have
    - this gives null space either side of the head and tail which make 
    it easier to add items to the head and tail of the list 
    - this makes pop, push, shift and unshift O(1) operations 

If you go off the end of the list the tail will actually be reassigned towards the front creating the ring 

[null, null, tail, val, val, val, head, null, null]

If the tail meets the head, you need to resize, and realign the head and the tail in an array with greater capacity 

# Numerical testing of a Javascript list a = []

could be an array list because when tested get, and push pop show O(1) run times and shift and unshift seem to be O(n) run times


# Recurssion

Recurssion is a function that calls itself 

You always need to have a base case when dealing with recurssion which will tell the loop to stop 

Recurssion can be broken into three steps:

- 1 a pre-recurse - do something before you recurse,
- 2 recurse - which is calling of the function 
- 3 - do something else after recurssion 

## Maze Solver

[#####E#
 "#     #"
 "#S#####"]

Base case:
- it should be the walls "#" 
- off the map you - you have to reuturn to the map
- if its the end 
- places we've already been 

