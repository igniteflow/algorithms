import re


def compress(text):
    if len(text) == 1:
        return '1' + text

    i = 0
    previous = None
    compressed = ''
    for idx, char in enumerate(text):
        if idx == 0:
            # first iteration
            previous = char
            i += 1
            continue

        if char == previous:
            i += 1
        else:
            compressed += unicode(i) + previous + ','
            i = 1

        previous = char

        if idx == len(text) - 1:
            # last iteration so record and omit trailing comma
            compressed += unicode(i) + char

    return compressed


def decompress(compressed):
    text = ''
    for part in compressed.split(','):
        num, char, _ = re.split('(\D)+', part, flags=re.IGNORECASE)
        for i in range(int(num)):
            text += char

    return text


if __name__ == '__main__':
    text = u'AAAAAAAAAAAAAAAAAAAAAAAAAAAAAABCBCAADEF'
    compressed = compress(text)
    print('Before: {} After: {}'.format(len(text), len(compressed)))
    print('Text: {}\nCompressed: {}'.format(text, compressed))
    decompressed_text = decompress(compressed)
    print(decompressed_text)
    assert decompressed_text == text
