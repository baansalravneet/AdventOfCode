def get_parsed_input(file_input):
    input = []
    for f in file_input:
        input.append(list(f))
    return input


def task1(grid) -> int:
    count = 0
    for i in range(1, len(grid)):
        j = 0
        while j < len(grid[i]):
            if grid[i-1][j] == '|' or grid[i-1][j] == 'S':
                if grid[i][j] == '.':
                    grid[i][j] = '|'
                else:
                    count += 1
                    grid[i][j-1] = '|'
                    grid[i][j+1] = '|'
                    j += 1
            j += 1
    return count


def task2(grid) -> int:
    counts = [[0] * len(grid[0]) for _ in range(len(grid))]
    for i in range(0, len(grid)):
        for j in range(len(grid[i])):
            if grid[i-1][j] == 'S':
                counts[i][j] += 1
                break
            if grid[i][j] == '^':
                counts[i][j-1] += counts[i-1][j]
                counts[i][j+1] += counts[i-1][j]
            elif grid[i][j] == '|':
                counts[i][j] += counts[i-1][j]
    return sum(counts[-1])


def solution(file_input):
    print("--- Day 7: Laboratories ---")
    grid = get_parsed_input(file_input)
    print("Task 1:", task1(grid))
    print("Task 2:", task2(grid))
