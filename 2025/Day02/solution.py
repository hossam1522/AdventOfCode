with open("input.txt", "r") as f:
    input = f.read()

ranges = input.split(",")
sum1 = 0
sum2 = 0

for value in ranges:
    start, end = value.split("-")
    for number in range(int(start), int(end) + 1):
        number = str(number)

        # First part
        if number.count(number[: len(number) // 2]) == 2 and len(number) % 2 == 0:
            sum1 += int(number)

        # Second part
        for i in range(len(number) // 2, 0, -1):
            if len(number) % i == 0:
                if number.count(number[:i]) == len(number) / i:
                    sum2 += int(number)
                    break

print("First part:", sum1)
print("Second part", sum2)
