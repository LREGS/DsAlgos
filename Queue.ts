type Node<T> = {
    value: T,
    next?: Node<T>
}
export default class Queue<T> {
    public length: number;
    private head?: Node <T>;
    private tail?: Node<T>;

    constructor() {
        this.head = this.tail = undefined;
        this.length = 0
    }

    enqueue(item: T): void {
        //defines node within method 
        const node = {value: item} as Node<T>;
        //inreases the length value as a node is added
        this.length++;
        //checks if the list has a tail, if it doesn't the ndoe is both
        if (!this.tail){
            this.tail = this.head = node;
            return;
        }
        //if the list has a tail the tail now points to the new added node
        this.tail.next = node;
        //the new node is now assigned as the tail
        this.tail = node;
    }
    deque(): T | undefined {

        //makes sure the node is a head as only heads can leave the queue 
        if (!this.head){
            return undefined;
        }
        //shortens the length value as a value is being removed from the queue 
        this.length--;

        const head = this.head;
        //assigns the head to the 2nd value in the queue because the head is being removed 
        this.head = this.head.next

        head.next = undefined;

        return head.value
    }

    peek(): T | undefined {
        return this.head?.value;

    }
}