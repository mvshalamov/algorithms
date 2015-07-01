public class BinaryHeapTest {
    public static void main(String[] args) {
        int[] arr = {3, 9, 4,  10, 5, 19, 6, 8, 1, 2, 11};
        BinaryHeap heap = new BinaryHeap(arr);
        
        System.out.println(heap.toString());
    }
}

class BinaryHeap {
    private int[] heapArr;

    BinaryHeap(int[] arr) {
        heapArr = arr;
        heapify();
    }

    private void heapify() {
        boolean sort = true;
        int change;
        int i;
        do {
            sort = true;
            for (i = 0; i < (int) Math.ceil(heapArr.length / 2); i++) {
                if (2 * i + 1 < heapArr.length && heapArr[i] < heapArr[2 * i + 1]) {
                    sort = false;
                    change = heapArr[2 * i + 1];
                    heapArr[2 * i + 1] =  heapArr[i];
                    heapArr[i] = change;
                } 

                if (2 * i + 2 < heapArr.length && heapArr[i] < heapArr[2 * i + 2]) {
                    sort = false;
                    change = heapArr[2 * i + 2];
                    heapArr[2 * i + 2] =  heapArr[i];
                    heapArr[i] = change;
                } 
            }
        } while (!sort);
    }

    public String toString() {
        String str = "";
        for (int i = 0; i < heapArr.length; i++) {
            str = str + heapArr[i];
            if (i != heapArr.length - 1) {
                str = str + " | ";
            }
        }
        return str;
    }
}
