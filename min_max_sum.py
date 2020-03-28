#!/usr/bin/python3
import sys
from typing import List


def miniMaxSum(arr: List[int]) -> str:
    total = 0
    min_value = sys.maxsize
    max_value = 0
    for number in arr:
        max_value = max(number, max_value)
        min_value = min(number, min_value)
        total += number
    return f"{total-max_value} {total-min_value}"


print(miniMaxSum([1, 2, 3, 4]))
