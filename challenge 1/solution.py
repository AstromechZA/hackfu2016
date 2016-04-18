import os
import re


FILES = os.path.join(os.path.dirname(__file__), '..', 'hackfu2016', 'container', 'Challenge 1')

NUMBER_PAIRS = (
    (85, 8),
    (124, 11),
    (1984, 8),
    (3, 5),
    (901, 1),
    (3, 13),
    (8546, 12),
    (5, 2),
    (3, 4),
    (85, 10),
    (3437, 7),
)


def get_line_n(file_path, n):
    with open(file_path, 'r') as f:
        lines = f.readlines()
        return lines[n]

def get_char_n(file_path, n):
    with open(file_path, 'r') as f:
        content = f.read()
        content = re.sub('\W', '', content)
        return content[n]


word_regex = re.compile('[^\W]+')


def get_words(file_path):
    with open(file_path, 'r') as f:
        all = f.read()
        return word_regex.findall(all)


def main():

    books_path = os.path.join(FILES, 'Books')
    books = os.listdir(books_path)
    books = sorted([b for b in books if os.path.isfile(os.path.join(books_path, b))])
    book_paths = [os.path.join(books_path, b) for b in books]

    words_per_book = {}
    for b in book_paths:
        words_per_book[b] = get_words(b)

    for k, v in words_per_book.iteritems():
        try:
            words = [(v[i + 1], j) for i, j in NUMBER_PAIRS]
            print words
        except IndexError:
            pass




if __name__ == '__main__':
    main()