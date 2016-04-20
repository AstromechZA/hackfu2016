# Challenge 1

```
85,8
124,11
1984,8
3,5
901,1
3,13
8546,12
5,2
3,4
85,10
3437,7
```

- Looks to be some kind of indexes into the list and content of the books. There are 50 books.
- Each book ranges from 32K to 3.1M.
- There are 11 lines each with a pair of numbers
- We can presume that each line maps to a character? although none are duplicated.
- first number ranges from 3 to 8546
- second number ranges from 1 to 13
- files were last modified on 3 dates (Dec 14 2015, Jan 6, Feb 15)
- each date group has 16 files, 16 files, 18 files (== 52)
- we assume the first number is a line number, since they seem to be on the same scale as the number of lines per book.
- from idea 1 we assume the line number only counts line with text?

```
-rw-r--r--   1 benmeier  staff   847K Dec 14 11:24 Dracula
-rw-r--r--   1 benmeier  staff   431K Dec 14 11:24 Frankenstein
-rw-r--r--   1 benmeier  staff   758K Dec 14 11:24 A Tale of Two Cities
-rw-r--r--   1 benmeier  staff   160K Dec 14 11:25 Alice in Wonderland
-rw-r--r--   1 benmeier  staff   443K Dec 14 11:26 The Picture of Dorian Gray
-rw-r--r--   1 benmeier  staff   136K Dec 14 11:27 Metamorphosis
-rw-r--r--   1 benmeier  staff   156K Dec 14 11:28 Dr Jekyll and Mr Hyde
-rw-r--r--   1 benmeier  staff   653K Dec 14 11:28 Wuthering Heights
-rw-r--r--   1 benmeier  staff   194K Dec 14 11:28 The Time Machine
-rw-r--r--   1 benmeier  staff   588K Dec 14 11:29 Gulliver's Travels
-rw-r--r--   1 benmeier  staff   375K Dec 14 11:29 Treasure Island
-rw-r--r--   1 benmeier  staff   287K Dec 14 11:29 Beowulf
-rw-r--r--   1 benmeier  staff   222K Dec 14 11:30 The Wonderful Wizard of Oz
-rw-r--r--   1 benmeier  staff   688K Dec 14 11:30 Pride and Prejudice
-rw-r--r--   1 benmeier  staff   484K Dec 14 11:30 Paradise Lost
-rw-r--r--   1 benmeier  staff   393K Dec 14 11:31 Beyond Good and Evil
-rw-r--r--   1 benmeier  staff   580K Jan  6 09:20 Huckleberry Finn
-rw-r--r--   1 benmeier  staff    32K Jan  6 09:21 The Awakening
-rw-r--r--   1 benmeier  staff   3.1M Jan  6 09:21 War and Peace
-rw-r--r--   1 benmeier  staff   568K Jan  6 09:21 The Adventures of Sherlock Holmes
-rw-r--r--   1 benmeier  staff   403K Jan  6 09:21 Tom Sawyer
-rw-r--r--   1 benmeier  staff   136K Jan  6 09:23 The Importance of Being Earnest
-rw-r--r--   1 benmeier  staff   174K Jan  6 09:23 A Christmas Carol
-rw-r--r--   1 benmeier  staff   1.1M Jan  6 09:24 The Iliad
-rw-r--r--   1 benmeier  staff   2.5M Jan  6 09:24 The Count of Monte Cristo
-rw-r--r--   1 benmeier  staff   3.1M Jan  6 09:25 Les Miserables
-rw-r--r--   1 benmeier  staff   1.2M Jan  6 09:28 Leviathan
-rw-r--r--   1 benmeier  staff   286K Jan  6 09:29 The Jungle Book
-rw-r--r--   1 benmeier  staff    87K Jan  6 09:29 The Legend of Sleepy Hollow
-rw-r--r--   1 benmeier  staff   231K Jan  6 09:30 Siddhartha
-rw-r--r--   1 benmeier  staff   177K Jan  6 09:32 Through the Looking Glass
-rw-r--r--   1 benmeier  staff   224K Jan  6 09:33 Heart of Darkness
-rw-r--r--   1 benmeier  staff   1.9M Feb 15 15:49 Anna Karenina
-rw-r--r--   1 benmeier  staff   381K Feb 15 15:50 Around the World in 80 Days
-rw-r--r--   1 benmeier  staff   256K Feb 15 15:51 A Study in Scarlet
-rw-r--r--   1 benmeier  staff   1.9M Feb 15 15:51 David Copperfield
-rw-r--r--   1 benmeier  staff   380K Feb 15 15:52 Dubliners
-rw-r--r--   1 benmeier  staff   881K Feb 15 15:53 Emma
-rw-r--r--   1 benmeier  staff   1.0M Feb 15 15:53 Jane Eyre
-rw-r--r--   1 benmeier  staff   239K Feb 15 15:54 Life of Frederick Douglass
-rw-r--r--   1 benmeier  staff   1.0M Feb 15 15:55 Little Women
-rw-r--r--   1 benmeier  staff   896K Feb 15 15:56 Oliver Twist
-rw-r--r--   1 benmeier  staff   272K Feb 15 15:56 Peter Pan
-rw-r--r--   1 benmeier  staff   203K Feb 15 15:57 Pygmalion
-rw-r--r--   1 benmeier  staff   630K Feb 15 15:58 Robinson Crusoe
-rw-r--r--   1 benmeier  staff   677K Feb 15 15:58 Sense and Sensibility
-rw-r--r--   1 benmeier  staff   331K Feb 15 15:59 The Hound of the Baskervilles
-rw-r--r--   1 benmeier  staff   809K Feb 15 16:00 The Jungle
-rw-r--r--   1 benmeier  staff   379K Feb 15 16:01 Three Men in a Boat
-rw-r--r--   1 benmeier  staff   1.5M Feb 15 16:02 Ulysses
```

pairs in orders:

```
3,4
3,5
3,13
5,2
85,8
85,10
124,11
901,1
1984,8
3437,7
8546,12
```

```
901,1
5,2
3,4
3,5
3437,7
85,8
1984,8
85,10
124,11
8546,12
3,13
```

## first idea

1st character of Nth line of the M'th book, books 0 indexed

(N,M)       1 indexed lines
85,8        - X                  - X
124,11      - "                  - (space)
1984,8      - (space)            - '
3,5         - T                  - a
901,1       - (space)            - "
3,13        - T                  - a
8546,12     - E                  - w
5,2         - r                  - w
3,4         - (space)            - T
85,10       - J                  - (space)
3437,7      - a                  - a

Result: neither work, not too suprised there, we assume the key will be an english word.

## second idea

hurrr durrr "break the book cipher" ---> https://en.wikipedia.org/wiki/Book_cipher

first number could be word number
second number could be letter of that word

hard to do, since there is no schema of what to do about weird punctuation, and where to start, can't just split on whitespace.
Also there are some long word requirements.

## Tried a variety of things

- first tried (line, character) with skipping blank lines and skipping spaces
- tested one off errors
- nothing gave reasonable passphrases

- need to test words since its quite possibly a very long passphrase
