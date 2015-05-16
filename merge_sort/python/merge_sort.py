def merge_sort(data):
    if len(data) <= 1:
        return data

    midle = len(data) / 2
    left = merge_sort(data[:midle])
    right = merge_sort(data[midle:])
    return merge(left, right)


def merge(left, right):
    res = []

    while left and right:
        if left[0] < right[0]:
            res.append(left.pop(0))
        else:
            res.append(right.pop(0))

    if left:
        res.extend(left)
    if right:
        res.extend(right)

    return res

if __name__ == '__main__':
    data = [1, 3, 5, 7, 8]
    data1 = [2, 4, 6]
    print merge(data, data1)
    d = [8, 3, 6, 1, 90, 12, 3]
    print merge_sort(d)

