def get_parsed_input(file_input):
    return [tuple(map(lambda x: int(x), f.split(","))) for f in file_input]


def closest(coordinates):
    dists = []
    for i in range(len(coordinates)):
        x1, y1, z1 = coordinates[i]
        for j in range(i + 1, len(coordinates)):
            x2, y2, z2 = coordinates[j]
            dist = (x1 - x2) ** 2 + (y1 - y2) ** 2 + (z1 - z2) ** 2
            dists.append((dist, i, j))
    dists.sort()
    return [(i, j) for _, i, j in dists]


def task1(coordinates) -> int:
    uf = UnionFind(len(coordinates))
    count = 1000
    for x, y in closest(coordinates):
        uf.union(x, y)
        count -= 1
        if count == 0:
            break
    return uf.max_3_product()


def task2(coordinates) -> int:
    uf = UnionFind(len(coordinates))
    for x, y in closest(coordinates):
        if uf.checkedUnion(x, y):
            return coordinates[x][0] * coordinates[y][0]


def solution(file_input):
    print("--- Day 8: Playground ---")
    input = get_parsed_input(file_input)
    print("Task 1:", task1(input))
    print("Task 2:", task2(input))


class UnionFind:
    def __init__(self, n: int):
        self.parent = list(range(n))
        self.size = [1] * n
        self.n = n

    def find(self, x: int) -> int:
        while self.parent[x] != x:
            self.parent[x] = self.parent[self.parent[x]]
            x = self.parent[x]
        return x

    def union(self, x: int, y: int):
        px = self.find(x)
        py = self.find(y)
        if px == py:
            return
        if px < py:
            self.parent[py] = px
            self.size[px] += self.size[py]
        else:
            self.parent[px] = py
            self.size[py] += self.size[px]

    def checkedUnion(self, x: int, y: int) -> bool:
        self.union(x, y)
        return self.size[self.find(x)] == self.n or self.size[self.find(y)] == self.n

    def max_3_product(self):
        sorted_sizes = sorted(self.size, reverse=True)
        return sorted_sizes[0] * sorted_sizes[1] * sorted_sizes[2]
