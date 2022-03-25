
def main():
    squareMult(1234567, 23456789,  3333337)

def squareMult(base, exponent, mod):
    binaryExp = format(exponent, 'b')
    print(f'In Binary:{binaryExp}')

    square = 1
    mult = 0
    print(f'Length of item: {len(binaryExp)}')
    for i in range(len(binaryExp)):
        # print(f'The {i}th Element is: {binaryExp[i]}')
        print(f'Round: {i+1}')
        square = (square ** 2) % mod
        print(f'square: {square}')
        if binaryExp[i] == '1':
            mult = (square * base) % mod
            square = mult
            print(f'mult: {mult}')
        else:
            print("mult: -")
        print('~~~~~~~~~~~~~~~~~')
    print(f'Final Answer: {square}')
main()
