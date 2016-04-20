# coding=utf-8

from collections import defaultdict, Counter


def xor(a, b):
    return a ^ b


def nott(a):
    return a ^ 255


def op(a, b):
    return nott(xor(a, b))


with open('bytes', 'r') as f:
    data = f.read()

# now group into pairs
bytepairs = [data[i:i+2] for i in xrange(0, len(data), 2)]

# convert to ints
shorts = [ord(p[0]) * 255 + ord(p[1]) for p in bytepairs]

print bin(op(shorts[0], ord('')))


h = 'NOT SIGNED'
mp = (len(data) - len(h)) / 2
print len(h), mp
print h
print data[mp:mp+len(h)]

