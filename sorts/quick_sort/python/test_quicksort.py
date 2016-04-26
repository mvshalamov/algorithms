import unittest
from quicksort import quicksort, partition

class TestStringMethods(unittest.TestCase):
    def test__sort(self):
        data = [1, 3, 2, 8]
        quicksort(data, 0, len(data) - 1)
        self.assertEqual([1,2,3, 8], data)
        
        data1 = [150, 4, 2, 7, 1, 9, 160]
        quicksort(data1, 0, 6)
        self.assertEqual([1, 2, 4, 7, 9, 150, 160], data1)

