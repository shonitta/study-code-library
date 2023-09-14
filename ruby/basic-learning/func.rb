require 'minitest/autorun'


def sum_arr(arr)
  arr.sum
end


class TestSumArr < Minitest::Test
  def test_sum_arr_with_empty_array
    assert_equal 0, sum_arr([])
  end

  def test_sum_arr_with_one_element
    assert_equal 1, sum_arr([1])
  end

  def test_sum_arr_with_two_elements
    assert_equal 3, sum_arr([1, 2])
  end

  def test_sum_arr_with_three_elements
    assert_equal 6, sum_arr([1, 2, 3])
  end

  def test_sum_arr_with_negative_elements
    assert_equal 0, sum_arr([-1, 1])
  end

  def test_sum_arr_with_float_elements
    assert_equal 3.5, sum_arr([1.5, 2])
  end
end
