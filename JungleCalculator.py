import json
import sys
from os import path

import distance_calc
import pythag
import radical_simplifier

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


def main():
    while True:
        print(d["str13"])
        print(d["func1"])
        print(d["func2"])
        print(d["func3"])
        print(d["quit"])
        inpt = input().lower()
        if inpt == "1":
            print("\n\n")
            distance_calc.distance_calc(d)
        elif inpt == "2":
            print("\n\n")
            pythag.pythag(d)
        elif inpt == "3":
            print("\n\n")
            radical_simplifier.simplify_radical(d)
        elif inpt == "q":
            print(d["quit"] + d["str10"])
            sys.exit()
        else:
            print(d["str14"])
            main()


main()
