def array_diff(a, b):
    result = []
    for i in range(len(a)):
        if a[i] not in b:
            result.append(a[i])
    for i in range(len(b)):
        if b[i] not in a:
            result.append(b[i])
    return result
