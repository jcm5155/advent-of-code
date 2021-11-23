import numpy as np

with open('day_8_input.txt', 'r') as f:
    pzl_input = np.array(list(f.read().strip())).reshape((-1, 6, 25))

fewest_zeros = min(pzl_input, key=lambda layer:np.sum(layer == '0'))
print(np.sum(fewest_zeros == '1') * np.sum(fewest_zeros == '2')) # Part 1

decoded = pzl_input[0]
for layer in pzl_input:
    decoded = np.where(decoded != '2', decoded, layer)

decoded = np.where(decoded == '1', '#', ' ') # To make it easier to see
for row in decoded: # Part 2
    print(*row, sep='')
