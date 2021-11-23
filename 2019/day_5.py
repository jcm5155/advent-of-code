from intcode import get_pzl_input, run_intputer

def main():
    pzl_input = get_pzl_input('day_5_input.txt')
    print(f"Part 1 = {run_intputer('single', pzl_input, 1)}")
    print(f"Part 2 = {run_intputer('single', pzl_input, 5)}")

if __name__ == "__main__":
    main()