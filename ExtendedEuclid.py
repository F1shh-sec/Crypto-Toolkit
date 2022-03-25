import math
def main():
    #Grabs the First Number
    firstnum = int(input("Please Enter first number: "))
    #Grabs the second number
    secondnum = int(input("Please Enter second number: "))
    # Grabs the GCD, Err, and Bezout Coefficent as return val
    CalcGCD, err, t = xgcd(firstnum, secondnum, 1, 1)
    #Prints out GCD Of the two numbers
    print(f'The GCD of {firstnum} and {secondnum} is {CalcGCD}')
    # If we have a GCD of one, Prints the Mult-Inverse
    if (CalcGCD == 1):
        print(f'The Mult Inverse is: {t%secondnum}')
def xgcd(a, b, s1 = 1, s2 = 0, t1 = 0, t2 = 1):
    # Break Condition, B=0
    if (b == 0):
        #returns the absolute value of A, 1, and 0
        return abs(a), 1, 0
    # Q = a/b drops the decimal place.
    q = math.floor(a / b)
    # R is set to A - the floored value, mult B
    r = a - q * b
    # S3 is set to the previous S1 minus the floored int, times s2
    s3 = s1 - q * s2
    # S3 gets set to the previous T1 minus the floored int times the previous T2
    t3 = t1 - q * t2
    # if r==0, then b will be the gcd and s2, t2 the Bezout coefficients
    # Recurses down the loop untill we hit r=0.
    return (abs(b), s2, t2) if (r == 0) else xgcd(b, r, s2, s3, t2, t3)

if (__name__ == "__main__"):
    main()