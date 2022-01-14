#!/bin/python3

import math
import os
import random
import re
import sys

#
# Complete the 'arrayManipulation' function below.
#
# The function is expected to return a LONG_INTEGER.
# The function accepts following parameters:
#  1. INTEGER n
#  2. 2D_INTEGER_ARRAY queries
#

def arrayManipulation(n, queries):
    tmp_list = [0] * n
    # For each query just going to step up or down on the list accordingly
    for query in queries:
        # Simply add at the initial entry no matter what
        tmp_list[query[0]-1] += query[2]
        # If it doesn't continue to end of array then subtract at end index
        if query[1] != len(tmp_list):
            tmp_list[query[1]] -= query[2]
    # Going to create a running sum and return values, replacing as necessary
    ret = 0
    run_sum = 0
    for value in tmp_list:
        run_sum += value
        if run_sum > ret:
            ret = run_sum
    return ret

if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    first_multiple_input = input().rstrip().split()

    n = int(first_multiple_input[0])

    m = int(first_multiple_input[1])

    queries = []

    for _ in range(m):
        queries.append(list(map(int, input().rstrip().split())))

    result = arrayManipulation(n, queries)

    fptr.write(str(result) + '\n')

    fptr.close()
