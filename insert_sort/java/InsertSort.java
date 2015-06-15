import java.util.ArrayList;
import java.util.Arrays;

public class InsertSort {
    public static void main(String[] args) {
        ArrayList<Integer> list = new ArrayList<Integer>();
        list.addAll(Arrays.asList(5, 1, 9, 2, 7, 3, 8, 10));

        System.out.println(sort(list).toString());
    }

    public static ArrayList<Integer> sort(ArrayList<Integer> list) {
        Integer item;
        int j;
        
        for (int i = 1; i < list.size(); i++) {
            item = list.get(i);
            j = i - 1;

            while (j >= 0 && list.get(j) > item) {
                list.set(j + 1, list.get(j));
                j = j - 1;
            }
            list.set(j + 1, item);
        }

        return list;
    }
}
