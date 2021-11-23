from itertools import permutations
from collections import defaultdict


class Intputer:
    def __init__(self, pzl_input):
        self.mem = pzl_input.copy()
        self.ptr = 0
        self.rel_base = 0

class Amp(Intputer):
    """Used in amp_circuit(), loop_circuit()"""
    def __init__(self, phase, pzl_input):
        Intputer.__init__(self, pzl_input)
        self.name = 'Amp'
        self.phase = phase
        self.phased = False

    def __repr__(self):
        return f"Amp: {self.phase}"

class Comp(Intputer):
    """Used in non-looping circuits"""
    def __init__(self, pzl_input):
        Intputer.__init__(self, pzl_input)
        self.name = 'Comp'


def get_pzl_input(file_name):
    """Sets puzzle input"""
    with open(file_name) as f:
        pzl_input = defaultdict(int)
        pzl_input = {pos: int(x) for pos, x in enumerate(f.readline().split(','))}
    return pzl_input


def run_intputer(mode, pzl_input, val):
    if mode == 'single':
        return(single_circuit(val, pzl_input))
    elif mode == 'amplify':
        return amp_circuit(permutations(range(5), 5), pzl_input)
    elif mode == 'loop':
        return loop_circuit(permutations(range(5, 10), 5), pzl_input)
    else: 
        return 'ERROR––mode not found.'


def set_params(cpu, param_range):
    param = {}
    for i in param_range:
        mode = cpu.mem[cpu.ptr]//int('100'.ljust(i+2, '0')) % 10
        # Position mode
        if mode == 0:
            param[i] = cpu.mem[cpu.ptr + i]
        # Immediate mode
        elif mode == 1:
            param[i] = cpu.ptr + i
        # Relative mode
        elif mode == 2:
            param[i] = cpu.mem[cpu.ptr + i] + cpu.rel_base
    return param


def intcode_cycle(cpu, val):
    param, output, skip = {}, None, (4,4,2,2,3,3,4,4,2)
    while cpu.mem[cpu.ptr] != 99:
        # Set opcode
        opcode = cpu.mem[cpu.ptr] % 100
        # Set parameters
        param = set_params(cpu, range(1,skip[opcode-1]))
        # Add
        if opcode == 1:
            cpu.mem[param[3]] = cpu.mem[param[1]] + cpu.mem[param[2]]

        # Multiply
        elif opcode == 2:
            cpu.mem[param[3]] = cpu.mem[param[1]] * cpu.mem[param[2]]
        # Input
        elif opcode == 3:
            if cpu.name == 'Amp':
                if not cpu.phased:
                    cpu.mem[param[1]] = cpu.phase
                    cpu.phased = True
                else:
                    cpu.mem[param[1]] = val
            else:
                cpu.mem[param[1]] = val
        # Output
        elif opcode == 4:
            if cpu.name == 'Amp':
                output = cpu.mem[param[1]]
                cpu.ptr += skip[opcode-1]
                return output
            else:
                output = cpu.mem[param[1]]
                print(output)
        # Jump if T/F
        elif opcode == 5 and cpu.mem[param[1]] or opcode == 6 and not cpu.mem[param[1]]:
            cpu.ptr = cpu.mem[param[2]] - 3
        # 1/0 if less
        elif opcode == 7:
            cpu.mem[param[3]] = 1 if cpu.mem[param[1]] < cpu.mem[param[2]] else 0
        # 1/0 if equal
        elif opcode == 8:
            cpu.mem[param[3]] = 1 if cpu.mem[param[1]] == cpu.mem[param[2]] else 0
        # Change relative base
        elif opcode == 9:
            cpu.rel_base += cpu.mem[param[1]]
        cpu.ptr += skip[opcode-1]
    return output


def single_circuit(val, pzl_input):
    return intcode_cycle(Comp(pzl_input), val)


def build_amp_sequence(seq, pzl_input):
    return [Amp(i, pzl_input) for i in seq]


def amp_circuit(phase_combos, pzl_input):
    thrust_vals = {}
    for seq in phase_combos:
        amp_seq = build_amp_sequence(seq, pzl_input)
        output = {-1: 0}
        for i in range(len(seq)):
            output[i] = intcode_cycle(amp_seq[i], output[i-1])
        thrust_vals[output[len(seq)-1]] = seq
    return thrust_vals


def loop_circuit(phase_combos, pzl_input):
    thrust_vals = {}
    for seq in phase_combos:
        cyc_idx, output = 0, 0
        amp_seq = build_amp_sequence(seq, pzl_input)
        while output != None:
            curr_amp_val = output
            curr_amp = amp_seq[cyc_idx]
            output = intcode_cycle(curr_amp, output)
            cyc_idx = (cyc_idx + 1) % 5
        thrust_vals[curr_amp_val] = seq
    return thrust_vals
