"""
>>> import io, sys
>>> sys.stdin = io.StringIO(chr(10).join(['Vasya']))
>>> hello()
Hello, Vasya!
>>> sys.stdin = io.StringIO(chr(10).join(['Petya']))
>>> hello()
Hello, Petya!
>>> sys.stdin = io.StringIO(chr(10).join(['Olya']))
>>> hello()
Hello, Olya!
"""


def hello():
    name = input()
    print(f'Hello, {name}!')


if __name__ == "__main__":
    import doctest
    doctest.testmod(verbose=True)
