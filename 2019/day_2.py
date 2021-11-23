op_str = [1,12,2,3,1,1,2,3,1,3,4,
          3,1,5,0,3,2,10,1,19,1,5,
          19,23,1,23,5,27,1,27,13,
          31,1,31,5,35,1,9,35,39,2,
          13,39,43,1,43,10,47,1,47,
          13,51,2,10,51,55,1,55,5,
          59,1,59,5,63,1,63,13,67,
          1,13,67,71,1,71,10,75,1,
          6,75,79,1,6,79,83,2,10,83,
          87,1,87,5,91,1,5,91,95,2,
          95,10,99,1,9,99,103,1,103,
          13,107,2,10,107,111,2,13,
          111,115,1,6,115,119,1,119,
          10,123,2,9,123,127,2,127,9,
          131,1,131,10,135,1,135,2,139,
          1,10,139,0,99,2,0,14,0]


# Part 1
pt_one_str = op_str.copy()
for i in range(0, len(pt_one_str), 4):
    opcode = pt_one_str[i]
    op_idx_1 = pt_one_str[i+1]
    op_idx_2 = pt_one_str[i+2]
    op_res_idx = pt_one_str[i+3]

    if opcode == 99:
        break
    if opcode == 1:
        pt_one_str[op_res_idx] = pt_one_str[op_idx_1] + pt_one_str[op_idx_2]
    else:
        pt_one_str[pt_one_str[i+3]] = pt_one_str[pt_one_str[i+1]] * pt_one_str[pt_one_str[i+2]]
print('-=Part 1=-')
print(pt_one_str[0])

# Part 2
for noun in range(100):
    for verb in range(100):
        curr_op_str = op_str.copy()
        curr_op_str[1] = noun
        curr_op_str[2] = verb
        pointer = 0
        while curr_op_str[pointer] != 99:
            opcode = curr_op_str[pointer]
            op_idx_1 = curr_op_str[pointer+1]
            op_idx_2 = curr_op_str[pointer+2]
            op_res_idx = curr_op_str[pointer+3]

            if opcode == 1:
                curr_op_str[op_res_idx] = curr_op_str[op_idx_1] + curr_op_str[op_idx_2]
            else:
                curr_op_str[op_res_idx] = curr_op_str[op_idx_1] * curr_op_str[op_idx_2]

            pointer += 4
        
        if curr_op_str[0] == 19690720:
            print('-=Part 2=-')
            print(100 * noun + verb)
            break
