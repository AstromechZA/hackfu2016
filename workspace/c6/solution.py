import os
from PIL import Image

FILES = os.path.join(os.path.dirname(__file__))

def rotate0(x, y):
    return x, y

def rotate90(x, y):
    return y, 800 - (x - 800)

def rotate180(x, y):
    return rotate90(*rotate90(x, y))

def rotate270(x, y):
    return rotate90(*rotate90(*rotate90(x, y)))

transforms = [
    rotate0, rotate90, rotate180, rotate270
]

def main():

    img = Image.open(os.path.join(FILES, 'image.jpg')).copy()
    img2 = img.copy()

    pixels = img.load()
    pixels2 = img2.load()

    cycle_width = 10
    cycle_type_count = 4

    imagedata = img.getdata()
    width, height = img.size
    print width, height
    for y in range(height):  # for every pixel:
        for x in range(width):

            distance = max(abs(800 - x), abs(800 - y))

            # determine which cycle it is in
            c = (distance / cycle_width) % cycle_type_count

            nx, ny = transforms[c](x, y)

            d = imagedata[ny * height + nx]


            pixels2[x, y] = d

    img2.show()

if __name__ == '__main__':
    main()
