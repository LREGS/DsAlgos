export default function bs_list(haystack: number[], needle: number): boolean {

    let lo = 0;
    let high = haystack.length;

     do{

        const m = Math.floor(lo + (high - lo)/2);
        const v = haystack[m];

        if (v === needle){
            return true;
        } else if (v > needle) {
            high = m;
        }else{
            lo = m + 1; 
        }
      } while (lo < high )
     
      return false;
}