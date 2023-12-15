import re
import math
import sys

print("Welcome to JungleHornet's 2D distance calculator")


def main():
    # Get input
    print("Please input point 1 (x,y)")
    inpt1 = input()

    # Match input to regex
    coord1 = re.split("(\\d+)\\s*,\\s*(\\d+)", inpt1)

    # Check if input is in correct format
    if not (len(coord1) == 4):
        print("Error: Input not in correct format")
        return

    # Get x and y values from input
    x1 = int(coord1[1])
    y1 = int(coord1[2])

    # Get second input
    print("Please input point 2 (x,y)")
    inpt2 = input()

    # Match input to regex
    coord2 = re.split("(\\d+)\\s*,\\s*(\\d+)", inpt2)

    # Check if input is in correct format
    if not (len(coord2) == 4):
        print("Error: Input not in correct format")
        return

    # Get x and y values from input
    x2 = int(coord2[1])
    y2 = int(coord2[2])

    # Calculate distance
    a = x2 - x1
    b = y2 - y1

    a = a ** 2
    b = b ** 2

    dist = math.sqrt(a + b)

    # Round distance
    rounded = round(dist)
    if rounded < dist:
        rounded = rounded + 1

    # Create unsimplified square root
    rootCoefficient = 1
    simpleRootInt = 0
    exp = dist ** 2

    sqrtDist = '√' + str(round(exp))
    root = round(exp)

    # Simplify square root
    for i in range(2, rounded + 1):
        if (root / i).is_integer():
            if not i == root and not (root / i) == root:
                if (math.sqrt(root / i)).is_integer():
                    simpleRootInt = i
                    rootCoefficient = rootCoefficient * int(math.sqrt(root / i))
                    print(str(rootCoefficient) + '√' + str(simpleRootInt) + "\n")
        if rootCoefficient * math.sqrt(simpleRootInt) == dist:
            break

    # Create simplified square root string
    simpleRoot = str(rootCoefficient) + '√' + str(simpleRootInt)

    # Create response
    if math.sqrt(root).is_integer() or simpleRootInt == 0:
        response = "The distance is " + str(dist) + " or " + str(sqrtDist)
    else:
        response = "The distance is " + str(dist) + ", " + str(sqrtDist) + ", or " + str(simpleRoot)

    print(response)

    # Check if user wants to restart
while True:
    main()
    print("\n Restart? (y/n)")
    yn = input().lower()
    if yn == "y":
        print("Restarting Program")
    elif yn == "n":
        print("n recieved, ending program.")
        break
    else:
        print("y or n not recieved, ending program.")
        break
sys.exit()
