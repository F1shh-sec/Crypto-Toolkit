def main():
    # Grabs Input from terminal
    firstNumString = input("Enter In the first number: ")
    secondNumString = input("Enter In the Second number: ")

    #Converts Strings into Hex
    firstNum = int(firstNumString, 16)
    secondNum = int(secondNumString, 16)

    #Array Of all Values we want to XOR
    result = []

    # Russian Peasant
    while secondNum != 0:
        if (secondNum % 2 != 0):
            #Appends to array
            result.append(firstNum)
            firstNum = firstNum * 2
            secondNum = secondNum // 2

        if (secondNum % 2 == 0):
            firstNum = firstNum * 2
            secondNum = secondNum // 2

    # Passes Result array to XOR Helper Function function
    n = len(result)
    FinalValue = xorOfArray(result, n)

    #Returns the Final value, Prints In HEX
    print(f'Result = {hex(FinalValue)}')

# XOR Helper Function From: https://www.geeksforgeeks.org/find-xor-of-all-elements-in-an-array/
def xorOfArray(arr, n):
    # Resultant variable
    xor_arr = 0

    # Iterating through every element in
    # the array
    for i in range(n):
        # Find XOR with the result
        xor_arr = xor_arr ^ arr[i]

    # Return the XOR
    return xor_arr

main()