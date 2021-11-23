pzl_input = range(372304,847060)

def get_valid_passwords(rnge):
    valid_passes = []
    for num in rnge:
        str_num = str(num)
        adj = False
        for i in range(0, len(str_num)-1):
            if str_num[i] == str_num[i+1]:
                adj = True
        if adj and list(str_num) == sorted(str_num):
            valid_passes.append(str_num)
    return valid_passes


def get_valid_passwords_for_real_tho(lst):
    valid_count = 0
    for str_num in lst:
        valid = False
        for char in str_num:
            if str_num.count(char) == 2:
                valid = True
        if valid:
            valid_count += 1
    return valid_count


# Part 1
part_1 = get_valid_passwords(pzl_input)
print(f"part 1 answer = {len(part_1)}")

# Part 2
part_2 = get_valid_passwords_for_real_tho(part_1)
print(f"part 2 answer = {part_2}")

