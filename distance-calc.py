import re
import math
import sys    

print("Welcome to JungleHornet's 2D distance calculator")
print("Please input point 1 (x,y)")
inpt1 = input()
coord1 = re.split("[(]\s*(\d+)\s*[,]\s*(\d+)\s*[)]", inpt1)
x1 = coord1[1]
y1 = coord1[2]

x1 = int(x1)
y1 = int(y1)

print("Please input point 2 (x,y)")
inpt2 = input()
coord2 = re.split("[(]\s*(\d+)\s*[,]\s*(\d+)\s*[)]", inpt2)
x2 = coord2[1]
y2 = coord2[2]

x2 = int(x2)
y2 = int(y2)

a = x2 - x1
b = y2 - y1

a = a**2
b = b**2


dist = math.sqrt(a + b)

rounded = round(dist)
if rounded < dist:
    rounded = rounded + 1

root = 0
simpleRoot = 0
rootCoefficient = 1
sqrtDist = 0

exp = dist**2

sqrtDist = '√' + str(round(exp))
root = exp

for i in range(2, rounded**2 + 1):
    if (root / i).is_integer():
        if (math.sqrt(root / i)).is_integer():
            simpleRoot = i
            rootCoefficient = rootCoefficient * int(math.sqrt(root / i))
    for x in range(2, simpleRoot**2 + 1):
        if (root / i).is_integer():
            i = 2
            break

simpleRoot = str(rootCoefficient) + '√' + str(simpleRoot)

if math.sqrt(root).is_integer():
    response = "The distance is " + str(dist) + " or " + str(sqrtDist)
else:
    response = "The distance is " + str(dist) + ", " + str(sqrtDist) + ", or " + str(simpleRoot)

print(response)

sys.exit()