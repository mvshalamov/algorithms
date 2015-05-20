def insertsort(data):
    for j in xrange(len(data) - 1):
        i = j
        res = data[j]
        while i > 0 and data[i-1] > res:
            data[i] = data[i-1]
            i -= 1

        data[i] = res

    return data


print insertsort([4, 1, 3, 2, 7, 9])