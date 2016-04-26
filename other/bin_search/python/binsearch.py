def binsearch(alist, item):
    first = 0
    last = len(alist) - 1
 
    if alist:
        while first <= last:
            mid = (first + last) // 2
            if alist[mid] == item:
                return mid
            elif item < alist[mid]:
                last = mid - 1
            else:
                first = mid + 1
        return mid

if __name__ == '__main__':
    print binsearch([5, 6, 7, 8], 7)
    print binsearch([], 7)
