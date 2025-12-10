import math


def get_parsed_input(file_input):
    nums = []
    operations = []
    raw_nums = []
    raw_operations = []
    for f in file_input:
        if f[0] == "*" or f[0] == "+":
            operations = f.split()
            raw_operations = f
            break
        nums.append(list(map(lambda x: int(x), f.split())))
        raw_nums.append(f)
    return nums, operations, raw_nums, raw_operations


def task1(nums, operations) -> int:
    result = 0
    for j in range(len(operations)):
        if operations[j] == "+":
            result += sum(nums[i][j] for i in range(len(nums)))
        else:
            result += math.prod(nums[i][j] for i in range(len(nums)))
    return result


def task2(nums, operations) -> int:
    j = 0
    result = 0
    while j < len(operations):
        operator = operations[j]
        count = 0 if operator == "+" else 1
        while j < len(operations) and (
            j == len(operations) - 1 or operations[j + 1] == " "
        ):
            num = 0
            for i in range(len(nums)):
                num = 10 * num + int(nums[i][j]) if nums[i][j] != " " else num
            j += 1
            count = count + num if operator == "+" else count * num
        result += count
        j += 1
    return result


def solution(file_input):
    print("--- Day 6: Trash Compactor ---")
    nums, operations, raw_nums, raw_operations = get_parsed_input(file_input)
    print("Task 1:", task1(nums, operations))
    print("Task 2:", task2(raw_nums, raw_operations))
