import argparse
import importlib
import traceback

def get_input(day, is_debug):
    if is_debug:
        filepath = f'inputs/day{day}/example.txt'
    else:
        filepath = f'inputs/day{day}/input.txt'
    with open(filepath) as f:
        for line in f:
            yield line.rstrip('\n')

def main():
    parser = argparse.ArgumentParser(description='Advent of Code')
    parser.add_argument('day', type=int, help='The day to run')
    parser.add_argument('--debug', '-d', action='store_true', help='Use example input')
    args = parser.parse_args()

    day = args.day
    is_debug = args.debug

    if is_debug:
        print("Using example input.")

    try:
        module = importlib.import_module(f'day{day}.solution')
    except ModuleNotFoundError:
        print(f"No solution module for day {day}")
        return

    func = getattr(module, 'solution', getattr(module, 'main', None))
    if func is None:
        print(f"No callable 'solution' or 'main' in day{day}.solution")
        return

    try:
        func(get_input(day, is_debug))
    except Exception as e:
        print("Could not call the solution function", e)
        print(traceback.format_exc())

if __name__ == '__main__':
    main()
