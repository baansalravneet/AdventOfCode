import re


def get_parsed_input(file_input):
    current_block = []
    for line in file_input:
        if line == "":
            if current_block:
                current_block = []
        else:
            current_block.append(line)
    return current_block


def task(input) -> int:
    count = 0
    for line in input:
        x, y, *nums = list(map(int, re.findall(r"\d+", line)))
        if (x // 3) * (y // 3) >= sum(nums):
            count += 1
    return count


def solution(file_input):
    print("--- Day 12: Christmas Tree Farm ---")
    input = get_parsed_input(file_input)
    print("Task:", task(input))
