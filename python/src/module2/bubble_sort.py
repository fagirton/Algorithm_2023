"""
>>> import io, sys
>>> sys.stdin = io.StringIO(chr(10).join(['4','4 3 2 1']))
>>> bubble_sort()
3 4 2 1
3 2 4 1
3 2 1 4
2 3 1 4
2 1 3 4
1 2 3 4
"""


def bubble_sort():
    n = input()
    inp_string = input()
    str_lst = inp_string.split(" ")
    res = []
    for item in str_lst:
        res.append(int(item))

    n = len(res)
    num_swap = 0
    for i in range(0, n - 1):
        for j in range(0, n - 1 - i):
            if res[j] > res[j + 1]:
                res[j], res[j + 1] = res[j + 1], res[j]
                num_swap += 1
                print(" ".join(map(str, res)))
    if num_swap == 0:
        print("0")


if __name__ == "__main__":
    import doctest
    doctest.testmod(verbose=True)
