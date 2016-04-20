
STRINGS = (
    'OIOOb',
    '_IOI0',
    'OI_If',
    'IOOO6',
    '_OOO8',
    'I_II9',
    'II__3',
    '_IOO1',
    'O__I2',
    'OIOI9',
    'OOOOa',
    '_III8',
    'IOI_f',
    'OO_Of',
    'OIIIe',
    '_OIO0',
    'OOIOb',
    '_IO_1',
    '_IIO4',
    'IO_Oe',
    'OIO_d',
    '__IOc',
    'IIOIe',
    'O_OO4',
    'IIIId',
    'IOIO9',
    'IOOI8',
    'IO_Ic',
    'IO_I0',
    'IIIO8',
    'OO_I5',
    'OOII8'
)

converted = []
for s in STRINGS:
    bins = s[:4]
    bins = bins.replace('O', '0')
    bins = bins.replace('I', '1')
    end = int(s[-1], 16)
    converted.append(bins + format(end, "04b"))


def generate_underscore_replaced(s):

    if '_' not in s:
        return [s]

    i = s.find('_')

    return generate_underscore_replaced(s[:i] + '0' + s[i+1:]) + generate_underscore_replaced(s[:i] + '1' + s[i+1:])


for c in converted:
    print c
    print generate_underscore_replaced(c)
    for d in generate_underscore_replaced(c):
        print chr(int(d, 2))
