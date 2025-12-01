with open("input.txt", "r") as f:
    rotations = f.readlines()
current = 50
pass1 = 0
pass2 = 0

for rotation in rotations:
    direction = rotation[0]
    distance = int(rotation[1:])

    if distance > 100:
        pass2 += int(distance / 100)
        distance %= 100

    if direction == "L":
        if distance > current and distance < 100 and current != 0:
            pass2 += 1
        current = (current - distance) % 100
    elif direction == "R":
        if current + distance > 100 and distance < 100 and current != 0:
            pass2 += 1
        current = (current + distance) % 100

    if current == 0:
        pass1 += 1
        pass2 += 1

print("Password 1:", pass1)
print("Password 2:", pass2)
