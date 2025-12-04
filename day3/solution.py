from functools import cache


def get_parsed_input(file_input):
    banks = []
    for l in file_input:
        banks.append(list(map(lambda x: int(x), list(l))))
    return banks


def task(input, calculator) -> int:
    result = 0
    for bank in input:
        result += calculator(bank)
    return result


def task1(input) -> int:
    def calculator(bank):
        joltage = bank[0] * 10 + bank[1]
        max_so_far = max(bank[0], bank[1])
        for i in range(2, len(bank)):
            joltage = max(joltage, max_so_far * 10 + bank[i])
            max_so_far = max(max_so_far, bank[i])
        return joltage

    return task(input, calculator)


def task2(input) -> int:
    def calculator(bank):
        @cache
        def helper(idx, remaining) -> int:
            if remaining <= 0 or idx >= len(bank):
                return 0
            if len(bank) - idx < remaining:
                return 0
            return max(
                bank[idx] * (10 ** (remaining - 1)) + helper(idx + 1, remaining - 1),
                helper(idx + 1, remaining),
            )

        return helper(0, 12)

    return task(input, calculator)


def solution(file_input):
    print("--- Day 3: Lobby ---")
    input = get_parsed_input(file_input)
    print("Task 1:", task1(input))
    print("Task 2:", task2(input))
