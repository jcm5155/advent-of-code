from intcode import get_pzl_input, run_intputer

def main():
    pzl_input = get_pzl_input('day_7_input.txt')
    print(f"Part 1 = {max(run_intputer('amplify', pzl_input, None))}")
    print(f"Part 2 = {max(run_intputer('loop', pzl_input, None))}")

if __name__ == "__main__":
    main()
