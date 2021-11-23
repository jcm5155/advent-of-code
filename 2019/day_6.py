import networkx as nx

orbit_data = [i.split(')') for i in open('day_6_input.txt').read().splitlines()]

G = nx.DiGraph(orbit_data)

# part 1
print(len(nx.transitive_closure(G).edges()))

# part 2
print(nx.has_path(G, "YOU", "SAN"))