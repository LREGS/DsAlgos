function qs(arr: number[], lo: number, hi:number): void{
//calls parttition, does the pivot and does the recurssive step 
    if (lo >= hi){
        return
    }

    const pivotIdx = partition(arr, lo, hi);

    qs(arr, lo, pivotIdx -1);
    qs(arr, lo, pivotIdx + 1, hi); 
}

function partition(arr: number[], lo:number, hi:number): number {
//gets the pivot and moves the items to their correct side of the pivot

    const pivot = arr[hi];

    
    let idx = lo -1;
    
    //walks from low to hi but not the pivot, and everything is being compared to the pivot 
    for (let i = lo; i < hi; ++i){
        //if you find an element which is smaller than or equal to the pivot 
        if (arr[i] <= pivot){
            //incremenets the index to move along the array 
            idx++;
            const tmp = arr[i];
            arr[i] = arr[idx];
            arr[idx] = tmp;
        } 
    }
    idx++;
    arr[hi] = arr[idx];
    arr[idx] = pivot;

    return idx
}

export default function quick_sort(arr: number[]): void{
    qs(arr, 0, arr.length - 1);
    
}