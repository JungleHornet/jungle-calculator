import math
import re


def pythag(d):
    # Get first input
    print(d["str15"])
    inpt = input()

    # Match input to regex
    inpt1 = re.search("\\d*.?\\d*", inpt)

    # Check if input is in correct format
    if not inpt1 or float(inpt) <= 0:
        print(d["str4"])
        pythag(d)

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
        pythag(d)

    # Get length of leg 2 from input
    leg2 = float(inpt)

    # Calculate hypotenuse
    leg1 = leg1 ** 2
    leg2 = leg2 ** 2

    hyp = math.sqrt(leg1 + leg2)

    # Create radical version of hypotenuse
    root = round(hyp ** 2)
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
            return True
        elif yn == d["n"]:
            print(d["n"] + d["str10"])
            return False
        else:
            print(d["y"] + d["str6"] + d["n"] + d["str11"])
            return False
