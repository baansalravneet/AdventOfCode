def get_parsed_input(file_input):
    for line in file_input:
        return line.split(",")


def task1_validator(val):
    l = len(val)
    return l % 2 != 0 or val[: l // 2] != val[l // 2 :]


def task2_validator(val):
    def factors(i):
        return (d for d in range(1, i) if i % d == 0)

    return not any(
        all(
            val[idx : idx + part_length] == val[idx + part_length : idx + 2 * part_length]
            for idx in range(0, len(val) - 2 * part_length + 1, part_length)
        )
        for part_length in factors(len(val))
    )


def task(input, validator):
    count = 0
    for num_range in input:
        lower, upper = num_range.split("-")
        for i in range(int(lower), int(upper) + 1):
            count += i if not validator(str(i)) else 0
    return count


def task1(input) -> int:
    return task(input, task1_validator)


def task2(input) -> int:
    return task(input, task2_validator)


def solution(file_input):
    print("--- Day 2: Gift Shop ---")
    input = get_parsed_input(file_input)
    print("Task 1:", task1(input))
    print("Task 2:", task2(input))
