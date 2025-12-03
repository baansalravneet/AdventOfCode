def get_parsed_input(file_input):
    input = []
    for line in file_input:
        input.append((line[0], int(line[1:])))
    return input

def task1(input) -> int:
    current, count = 50, 0
    for direction, distance in input:
        distance = distance if direction == 'R' else -distance
        current = (current + distance) % 100
        count += 1 if current == 0 else 0
    return count

def task2(input) -> int:
    current, count = 50, 0
    for direction, distance in input:
        if distance > 100:
            count += distance // 100
            distance %= 100
        distance = distance if direction == 'R' else -distance
        if current != 0 and (current+distance > 100 or current+distance < 0):
            count += 1
        current = (current+distance) % 100
        count += 1 if current == 0 else 0
    return count

def solution(file_input):
    print("--- Day 1: Secret Entrance ---")
    input = get_parsed_input(file_input)
    print("Task 1:", task1(input))
    print("Task 2:", task2(input))
