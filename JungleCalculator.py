import json
import math
import re
import sys
from os import path

print("Please select a language (english (en), magyar (ma) )")
langInpt = input().lower()
if langInpt == "english" or langInpt == "en":
    print("English selected.")
    dictFile = 'en.json'
elif langInpt == "magyar" or langInpt == "ma":
    print("Magyar válogatott.")
    dictFile = 'hu.json'
else:
    print("Language not recognised, defaulting to english.")
    dictFile = 'en.json'

dictFile = 'langs/' + dictFile


def getDict():
    fileName = path.abspath(path.join(path.dirname(__file__), dictFile))
    return json.load(open(fileName))


d = getDict()

print(d["str1"])


def distance_calc():
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

    # Round distance
    rounded = math.ceil(dist)

    # Create unsimplified square root
    rootCoefficient = 1
    simpleRootInt = 0

    sqrtDist = '√' + str(round(dist**2))
    root = round(dist**2)

    # Simplify square root
    for i in range(2, rounded + 1):
        if (root / i).is_integer():
            if not i == root and not (root / i) == root:
                if (math.sqrt(root / i)).is_integer():
                    simpleRootInt = i
                    rootCoefficient = rootCoefficient * int(math.sqrt(root / i))
        if rootCoefficient * math.sqrt(simpleRootInt) == dist:
            break

    # Create simplified square root string
    simpleRoot = str(rootCoefficient) + '√' + str(simpleRootInt)

    # Create response
    if math.sqrt(root).is_integer() or simpleRootInt == 0:
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
        elif yn == d["n"]:
            print(d["n"] + d["str10"])
            break
        else:
            print(d["y"] + d["str6"] + d["n"] + d["str11"])
            break
        distance_calc()
    main()


def pythag():
    # Get first input
    print(d["str15"])
    inpt = input()

    # Match input to regex
    inpt1 = re.search("\\d*.?\\d*", inpt)

    # Check if input is in correct format
    if not inpt1 or float(inpt) <= 0:
        print(d["str4"])
        pythag()

    # Get length of leg 1 from input
    leg1 = float(inpt)

    # Get second input
    print(d["str16"])
    inpt = input()

    # Match input to regex
    inpt2 = re.search("\\d*.?\\d*", inpt)

    # Check if input is in correct format
    if not inpt2 or float(inpt) <= 0:
        print(d["str4"])
        pythag()

    # Get length of leg 2 from input
    leg2 = float(inpt)

    # Calculate hypotenuse
    leg1 = leg1 ** 2
    leg2 = leg2 ** 2

    hyp = math.sqrt(leg1 + leg2)

    # Create radical version of hypotenuse
    root = round(hyp**2)
    sqrtHyp = '√' + str(root)

    # Round hypotenuse
    rounded = math.ceil(hyp)

    # Simplify radical
    rootCoefficient = 1
    simpleRootInt = 0
    for i in range(2, rounded + 1):
        if (root / i).is_integer():
            if not i == root and not (root / i) == root:
                if (math.sqrt(root / i)).is_integer():
                    simpleRootInt = i
                    rootCoefficient = rootCoefficient * int(math.sqrt(root / i))
        if rootCoefficient * math.sqrt(simpleRootInt) == hyp:
            break

    # Create simplified square root string
    simpleRoot = str(rootCoefficient) + '√' + str(simpleRootInt)

    # Create response
    if math.sqrt(root).is_integer() or simpleRootInt == 0:
        response = d["str17"] + str(hyp) + d["str6"] + str(sqrtHyp)
    else:
        response = d["str17"] + str(hyp) + d["str7"] + str(sqrtHyp) + d["str8"] + str(simpleRoot)

    print(response)

    # Check if user wants to restart
    while True:
        print("\n" + d["str9"])
        yn = input().lower()
        if yn == d["y"]:
            print(d["str12"])
        elif yn == d["n"]:
            print(d["n"] + d["str10"])
            break
        else:
            print(d["y"] + d["str6"] + d["n"] + d["str11"])
            break
        pythag()
    main()


def main():
    print(d["str13"])
    print(d["func1"])
    print(d["func2"])
    print(d["quit"])
    inpt = input().lower()
    if inpt == "1":
        print("\n\n")
        distance_calc()
    elif inpt == "2":
        print("\n\n")
        pythag()
    elif inpt == "3":
        print("\n\n")
    elif inpt == "q":
        print(d["quit"] + d["str10"])
        sys.exit()
    else:
        print(d["str14"])
        main()


main()
