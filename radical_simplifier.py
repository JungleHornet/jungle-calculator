import math


def simplify_radical(d):
    print(d["str18"])
    try:
        root = int(input())
    except ValueError:
        print(d["str4"])
        return True
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
    if rootCoefficient == 1:
        simpleRoot = '√' + str(root)
    elif simpleRootInt == root:
        simpleRoot = '√' + str(root)
    else:
        simpleRoot = str(rootCoefficient) + '√' + str(simpleRootInt)

    print(simpleRoot)

    #  Ask if user wants to simplify another square root
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
