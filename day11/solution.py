from functools import cache


def get_parsed_input(file_input):
    graph = {}
    for line in file_input:
        parts = line.split(": ")
        source = parts[0]
        destinations = parts[1].split(" ")
        graph[source] = destinations
    return graph


def solution(file_input):
    print("--- Day 11: Reactor ---")
    graph = get_parsed_input(file_input)

    @cache
    def dfs(current, target):
        if current == target:
            return 1
        if current not in graph:
            return 0
        return sum(dfs(next, target) for next in graph[current])

    print("Task 1:", dfs("you", "out"))
    print(
        "Task 2:",
        dfs("svr", "dac") * dfs("dac", "fft") * dfs("fft", "out")
        + dfs("svr", "fft") * dfs("fft", "dac") * dfs("dac", "out"),
    )
