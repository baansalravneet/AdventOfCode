from typing import List
from typing import Tuple

def get_input(is_debug):
    if is_debug:
        filepath = f'inputs/day1/example.txt'
    else:
        filepath = f'inputs/day1/input.txt'
    with open(filepath) as f:
        for line in f:
            yield line.rstrip('\n')

def get_parsed_input(is_debug):
    input = []
    for line in get_input(is_debug):
        input.append((line[0], int(line[1:])))
    return input

def task1(input: List[Tuple[chr, int]]) -> int:
    current, count = 50, 0
    for direction, distance in input:
        distance = distance if direction == 'R' else -distance
        current = (current + distance) % 100
        count += 1 if current == 0 else 0
    return count

def task2(input: List[str]) -> int:
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

def solution(is_debug: bool):
    print("--- Day 1: Secret Entrance ---")
    input = get_parsed_input(is_debug)
    print("Task 1:", task1(input))
    print("Task 2:", task2(input))
