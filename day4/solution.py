def get_parsed_input(file_input):
    grid = []
    for l in file_input:
        grid.append(list(l))
    return grid


def task(input):
    n, m = len(input), len(input[0])
    marks = [[False] * m for _ in range(n)]

    def neighbours(i, j):
        for x in range(-1, 2):
            for y in range(-1, 2):
                yield (i + x, j + y)

    def valid(i, j):
        count = 0
        for nx, ny in neighbours(i, j):
            if nx < 0 or ny < 0 or nx >= n or ny >= m or (nx == i and ny == j):
                continue
            count += 1 if input[nx][ny] == "@" else 0
        return count < 4

    for i in range(n):
        for j in range(m):
            if input[i][j] == ".":
                continue
            if valid(i, j):
                marks[i][j] = True
    return sum(sum(m) for m in marks), marks


def task1(input) -> int:
    return task(input)[0]


def task2(input) -> int:
    count = 0
    while True:
        removed, marks = task(input)
        for i in range(len(input)):
            for j in range(len(input[0])):
                if marks[i][j]:
                    input[i][j] = "."
        count += removed
        if not removed:
            break
    return count


def solution(file_input):
    print("--- Day 4: Printing Department ---")
    input = get_parsed_input(file_input)
    print("Task 1:", task1(input))
    print("Task 2:", task2(input))
