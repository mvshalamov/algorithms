import unittest
from quicksort import quicksort

class TestStringMethods(unittest.TestCase):
    def test_success_sort(self):
        self.assertEqual([1,2,4], quicksort([4, 2, 1], 0, 2))
        self.assertEqual([1,2,4,7, 9, 150, 160], quicksort([150, 4, 2, 7, 1, 9, 160], 0, 6))
