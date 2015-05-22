import math


class BinaryHeap(object):
    def __init__(self):
        self.heap = []

    def heap_size(self):
        return len(self.heap)

    def _heap_size(self):
        return len(self.heap) - 1

    def add(self, elem):
        self.heap.append(elem)
        hs = self._heap_size()
        parent = math.ceil(hs / 2) - 1

        while hs > 0 and self.heap[parent] < self.heap[hs]:
            self.heap[parent], self.heap[hs] = self.heap[hs], self.heap[parent]
            hs = parent
            parent = math.ceil(hs / 2) - 1


    def heapify(self, vertex):
        while True:
            left_child = 2*vertex + 1
            right_child = 2*vertex + 2
            largest_child = vertex

            if left_child < self.heap_size() and self.heap[left_child] > self.heap[largest_child]:
                largest_child = left_child

            if right_child < self.heap_size() and self.heap[right_child] > self.heap[largest_child]:
                largest_child = right_child

            if largest_child == vertex:
                break

            self.heap[vertex], self.heap[largest_child] = self.heap[largest_child], self.heap[vertex]
            vertex = largest_child

    def built_heap(self, data):
        self.heap = data
        num = self.heap_size() / 2
        for i in reversed(range(num)):
            self.heapify(i)

    def get_max(self):
        max_elem = self.heap[0]
        self.heap[0] = self.heap[self._heap_size()]
        self.heap.pop(self._heap_size())
        self.heapify(0)
        return max_elem

    def heap_sort(self, data):
        self.built_heap(data)
        ret = []
        for it in reversed(range(self.heap_size())):
            ret.append(self.get_max())
        return ret


if __name__ == '__main__':
    heap = BinaryHeap()
    heap.built_heap([1, 10, 24, 12, 4, 67, 87, 79])
    
    print(heap.heap)
    print(heap.get_max())
    print(heap.get_max())
    print(heap.get_max())
    print(heap.get_max())
    print(heap.get_max())
    print(heap.get_max())
    print(heap.get_max())
    print(heap.heap)
    data = [1, 2, 5, 6, 2, 1, 2, 4, 5, 7, 1,89, ]
    print data
    print(heap.heap_sort(data))
