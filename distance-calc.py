import re
import math
import sys
import json

print("Please select a language (english (en), magyar (ma) )")
langInpt = input().lower()
if langInpt == "english" or langInpt == "en":
    print("English selected.")
    dictionaryFile = "en.json"
elif langInpt == "magyar" or langInpt == "ma":
    print("Magyar válogatott.")
    dictionaryFile = "hu.json"
else:
    print("Language not recognised, defaulting to english.")
    dictionaryFile = "en.json"

with open("langs/" + dictionaryFile) as json_file:
    d = json.load(json_file)

print(d["str1"])




def main():
    # Get input
    print(d["str2"])
    inpt1 = input()

    # Match input to regex
    coord1 = re.split("(-?\\d*.?\\d*)\\s*,\\s*(-?\\d*.?\\d*)", inpt1)

    # Check if input is in correct format
    if not (len(coord1) == 4):
        print(d["str4"])
        return

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
        return

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
    main()
    print("\n" + d["str9"])
    yn = input().lower()
    if yn == d["y"]:
        print(d["str12"])
    elif yn == d["n"]:
        print(d["n"] + d["str11"])
        break
    else:
        print(d["y"] + d["str6"] + d["n"] + d["str12"])
        break
sys.exit()
