public class MergeSortTest {
    public static void main(String[] args) {
        int[] arr = {9, 3, 4, 5, 2, 7, 1, 8, 6};

        System.out.println("input array: ");
        for (int v : arr)
            System.out.print(v + " ");
        System.out.println("");

        System.out.println("result array: ");
        for (int v : (MergeSort.sort(arr)))
            System.out.print(v + " ");
        System.out.println("");
    }
}

class MergeSort {
    public static int[] sort(int[] arr) {
        int center = (int) Math.ceil(arr.length / 2);

        int leftArr[] = new int[center];
        System.arraycopy(arr, 0, leftArr, 0, center);

        int rightArr[] = new int[arr.length - center];
        System.arraycopy(arr, center, rightArr, 0, arr.length - center);

        int change;

        if (leftArr.length > 2) {
            leftArr = sort(leftArr);
        } else if (leftArr.length == 2) {
            if (leftArr[0] > leftArr[1]) {
                change = leftArr[0];
                leftArr[0] = leftArr[1];
                leftArr[1] = change;
            }
        }

        if (rightArr.length > 2) {
            rightArr = sort(rightArr);
        } else if (rightArr.length == 2) {
            if (rightArr[0] > rightArr[1]) {
                change = rightArr[0];
                rightArr[0] = rightArr[1];
                rightArr[1] = change;
            }
        }

        int iter = 0;
        int i = 0;
        int j = 0;
        while (i < rightArr.length || j < leftArr.length) {
            if (i < rightArr.length && j < leftArr.length) {
                if (rightArr[i] > leftArr[j]) {
                    arr[iter++] = leftArr[j++];
                } else {
                    arr[iter++] = rightArr[i++];
                }
            } else if (i < rightArr.length) {
                arr[iter++] = rightArr[i++];
            } else {
                arr[iter++] = leftArr[j++];
            }
        }

        return arr;
    }
}
