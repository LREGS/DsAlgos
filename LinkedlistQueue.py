class Node:
    #classifies what a node which in this instance is a value, and a pointer to the next node,
    #if there is one
    def __init__(self, data):
        self.data = data
        self.next = None


class LQueue:
    def __init__(self):
        #defines the first and last element within the list
        self.front = None
        self.rear = None

    def is_empty(self):
        #helps control when the queue is empty or not 
        return self.front is None
    
    def enqueue(self, item):
    #create a new node to add to the list
        new_node = Node(item)
        if self.rear is None:
            #this is assigning the front and the back to the new node in the event that the 
            #list is empty during this step
            self.front = self.rear = new_node
            return 
        
        #adds a node with no values to the back
        self.rear.next = new_node
        self.read = new_node


    
    def dequeue(self):
        if self.is_empty():
            #stops you trying to remove items from a list that has no items to be removed 
            raise IndexError("Queue is empty")
        #you can only remove items from the front of a queue 
        item = self.front.data
        #assigns the front value to the value behind the front as thats next in queue 
        self.front = self.front.next7
        #makes the list empty if there's no more element
        if self.front is None:
            self.rear = None
        return item
    
    def size(self):
        current = self.front
        count = 0
        while current:
            count += 1 
            current = current.next
        return count 
