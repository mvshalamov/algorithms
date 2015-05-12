def quicksort(data):
    if len(data) <= 1:
        return data

    op_elem = data[0]
    less, pivot, more = [], [], []
    for el in data:
        if el > op_elem:
            more.append(el)
        elif el < op_elem:
            less.append(el)
        else:
            pivot.append(el)
    data = quicksort(less) + pivot + quicksort(more)
    return data

if __name__ == '__main__':
    data = [1, 5, 3, 7, 2, 9, 6]
    print quicksort(data)
