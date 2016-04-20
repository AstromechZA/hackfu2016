import os
from PIL import Image

FILES = os.path.join(os.path.dirname(__file__))


def main():
    img = Image.open(os.path.join(FILES, 'image.bmp')).copy()

    pixels = img.load()

    imagedata = img.getdata()
    width, height = img.size
    for y in range(height):  # for every pixel:
        for x in range(width):
            d = imagedata[y * height + x]
            pixels[x, y] = (1 - (d & 1)) * 255

    img.save(os.path.join(FILES, 'image-revealed.bmp'))
    img.show()


if __name__ == '__main__':
    main()
