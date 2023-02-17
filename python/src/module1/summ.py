"""
>>> import io, sys
>>> sys.stdin = io.StringIO(chr(10).join(['10 12']))  # input
>>> summ()
22
>>> sys.stdin = io.StringIO(chr(10).join(['1 1']))  # input
>>> summ()
2
>>> sys.stdin = io.StringIO(chr(10).join(['10000 10000']))  # input
>>> summ()
20000
"""


def summ():
    a, b = map(int, input().split(' '))
    print(a + b)


if __name__ == "__main__":
    import doctest
    doctest.testmod(verbose=True)
