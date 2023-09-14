require 'minitest/autorun'


def sum_arr(arr)
  arr.sum
end


def sum_hash_values(hash)
  hash.values.sum
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


class TestSumHashValues < Minitest::Test
  def tset_sum_hash_values_with_empty_hash
    assert_equal 0, sum_hash_values({})
  end

  def test_sum_hash_values_with_one_element
    assert_equal 1, sum_hash_values({ a: 1 })
  end

  def test_sum_hash_values_with_two_elements
    assert_equal 3, sum_hash_values({ a: 1, b: 2 })
  end

  def test_sum_hash_values_with_three_elements
    assert_equal 6, sum_hash_values({ a: 1, b: 2, c: 3 })
  end

  def test_sum_hash_values_with_negative_elements
    assert_equal 0, sum_hash_values({ a: -1, b: 1 })
  end

  def test_sum_hash_values_with_float_elements
    assert_equal 3.5, sum_hash_values({ a: 1.5, b: 2 })
  end
end
