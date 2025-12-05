def get_parsed_input(file_input):
    ranges, ingredients = [], []
    process_ranges = True
    for l in file_input:
        if l == "":
            process_ranges = False
            continue
        if process_ranges:
            start, end = l.split("-")
            ranges.append([int(start), int(end)])
        else:
            ingredients.append(int(l))
    return ranges, ingredients


def task1(ranges, ingredients) -> int:
    count = 0
    for i in ingredients:
        count += 1 if any(r[0] <= i <= r[1] for r in ranges) else 0
    return count


def task2(ranges) -> int:
    sorted_ranges = sorted(ranges, key=lambda x: x[0])
    merged_ranges = [sorted_ranges[0]]
    for i in range(1, len(sorted_ranges)):
        if sorted_ranges[i][0] <= merged_ranges[-1][1]:
            merged_ranges[-1][1] = max(sorted_ranges[i][1], merged_ranges[-1][1])
        else:
            merged_ranges.append(sorted_ranges[i])
    return sum(r[1] - r[0] + 1 for r in merged_ranges)


def solution(file_input):
    print("--- Day 5: Cafeteria ---")
    ranges, ingredients = get_parsed_input(file_input)
    print("Task 1:", task1(ranges, ingredients))
    print("Task 2:", task2(ranges))
