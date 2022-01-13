#!/usr/local/bin/python3

# Read in the first line
first_line = input()
k = int(first_line.split(' ')[0])
m = int(first_line.split(' ')[1])

# Setting up list for all ragged entries
full_entries = []

# Go through each entry, trim just incase, and get only values that are declared
for entry in range(k):
    tmp_line = str.strip(input())
    tmp_list = list(map(int, tmp_line.split(' ')))
    full_entries.append(tmp_list[1:tmp_list[0]+1])

"""
Recursive function to take all entries with the mod value to be used. First
just evaluates the split and whether this is the end of the recursion. If it is
the final run, then just do the simple max of the math using running_count.
Otherwise recursive call with -1 entries, sending same mode, and updated running
count.
"""
def eval_entries(entries, mod, running_count=0):
    entry=entries[0]
    if len(entries)==1:
        return max([
            (running_count+x**2)%mod for x in entry
        ])
    return max([
        eval_entries(entries[1:], mod, running_count+x**2) for x in entry
    ])

print(eval_entries(full_entries,m))
