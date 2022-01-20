#!/usr/bin/python3

# Pretty simple here, just checking for 1-9 and then 0-9 for other 5 digits
regex_integer_in_range = r"^[1-9][0-9]{5}$"	# Do not delete 'r'.
# A bit more comlpicated, get digit ref for one digit, and then to a positive look forward for digit followed by the first ref
regex_alternating_repetitive_digit_pair = r"(\d)(?=\d\1)"	# Do not delete 'r'.


import re
P = input()

print (bool(re.match(regex_integer_in_range, P))
and len(re.findall(regex_alternating_repetitive_digit_pair, P)) < 2)
