import sys

def hello(name = 'World'):
    print("Hello", name)

if __name__ == "__main__":
    if len(sys.argv[1:]) != 0:
        hello(' '.join(sys.argv[1:]))
    else:
        hello()