def main():
    # x ^ e mod(n)
    n = 3763
    e = 11
    elmsToCrack = [2912, 2929, 3368, 153, 3222, 3335, 153, 1222]
    CodeBook = {}

    print("CHAR | ASCII | RSA")
    # Generates a dictionary mapping a RSA encoded character to its matching Capital letter
    for i in range(65, 91):
        elm = (i**e)%n
        CodeBook[elm] = chr(i)
        print(f'{chr(i)} | {i} | {elm}')

    print("DECODED TEXT: ", end="")
    # Iterates over the encoded text and prints out the matching plaintext char
    for elm in elmsToCrack:
        print(CodeBook[elm], end="")

main()