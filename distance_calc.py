import math
import re


def distance_calc(d):
    # Get input
    print(d["str2"])
    inpt1 = input()

    # Match input to regex
    coord1 = re.split("(-?\\d*.?\\d*)\\s*,\\s*(-?\\d*.?\\d*)", inpt1)

    # Check if input is in correct format
    if not (len(coord1) == 4):
        print(d["str4"])
        distance_calc()

    # Get x and y values from input
    x1 = float(coord1[1])
    y1 = float(coord1[2])

    # Get second input
    print(d["str3"])
    inpt2 = input()

    # Match input to regex
    coord2 = re.split("(-?\\d*.?\\d*)\\s*,\\s*(-?\\d*.?\\d*)", inpt2)

    # Check if input is in correct format
    if not (len(coord2) == 4):
        print(d["str4"])
        distance_calc()

    # Get x and y values from input
    x2 = float(coord2[1])
    y2 = float(coord2[2])

    # Calculate distance
    a = x2 - x1
    b = y2 - y1

    a = a ** 2
    b = b ** 2

    dist = math.sqrt(a + b)

    sqrtDist = '√' + str(round(dist ** 2))
    root = round(dist ** 2)

    # Create unsimplified square root
    rootCoefficient = 1
    simpleRootInt = root

    # Simplify square root
    for i in range(2, round(math.sqrt(root))):
        if (simpleRootInt / i).is_integer():
            if not i == simpleRootInt and not (simpleRootInt / i) == simpleRootInt:
                if (math.sqrt(simpleRootInt / i)).is_integer():
                    simpleRootInt = i
                    rootCoefficient = rootCoefficient * int(math.sqrt(root / i))

    # Create simplified square root string
    simpleRoot = str(rootCoefficient) + '√' + str(simpleRootInt)

    # Create response
    if math.sqrt(root).is_integer() or simpleRootInt == root:
        response = d["str5"] + str(dist) + d["str6"] + str(sqrtDist)
    else:
        response = d["str5"] + str(dist) + d["str7"] + str(sqrtDist) + d["str8"] + str(simpleRoot)

    print(response)

    # Check if user wants to restart
    while True:
        print("\n" + d["str9"])
        yn = input().lower()
        if yn == d["y"]:
            print(d["str12"])
            return True
        elif yn == d["n"]:
            print(d["n"] + d["str10"])
            return False
        else:
            print(d["y"] + d["str6"] + d["n"] + d["str11"])
            return False
