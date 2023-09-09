import timeit
from copy import deepcopy
from typing import Any

import numpy as np
import pandas as pd

NUM_LOOP = 1000
DEFAULT_ARR = [1] * 10000

list_ = deepcopy(DEFAULT_ARR)
arr_ = np.array(deepcopy(DEFAULT_ARR))
tuple_ = tuple(deepcopy(DEFAULT_ARR))


def speed_test(xs: Any) -> pd.DataFrame:
    pd.DataFrame({"value": xs})


def case1_from_list() -> None:
    time_ = timeit.timeit("speed_test(list_)", number=NUM_LOOP, globals=globals())
    print(f"case1_from_list: {time_:.6f} sec")


def case2_from_array() -> None:
    time_ = timeit.timeit("speed_test(arr_)", number=NUM_LOOP, globals=globals())
    print(f"case2_from_array: {time_:.6f} sec")


def case3_from_tuple() -> None:
    time_ = timeit.timeit("speed_test(tuple_)", number=NUM_LOOP, globals=globals())
    print(f"case3_from_tuple: {time_:.6f} sec")


if __name__ == "__main__":
    case1_from_list()
    case2_from_array()
    case3_from_tuple()
