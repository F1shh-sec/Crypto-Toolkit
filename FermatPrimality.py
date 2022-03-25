import math
import random
from threading import Thread

def main():
    carmichaelCound = 0
    for i in range((10**6), 3, -1):
        if carmichaelCound == 3:
            break
        new_thread = Thread(target=checkCar(i))
        new_thread.start()

def checkCar(i):
    if fermat(i, 4):
        if not checkPrime(i):
            print(f"{i}: Is a Carmichael Number")
# retuens False on composite and true on prime
def fermat(p, s):
    for i in range(1, s):
        a = random.randint(2, p-2)
        if ((a**(p-1)) % p) != 1:
            return False
    return True

# Returns True on Prime and False on Composite
def checkPrime(num):
    stoppingPoint = int(num ** 0.5)
    for i in range(2, stoppingPoint+1):
        if (num % i == 0):
            return False
    return True
if __name__ == "__main__":
    main()



