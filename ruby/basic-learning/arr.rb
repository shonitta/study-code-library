arr = [1, 2, 3, 4, 5]

arr.push(6)  # arr << 6
p "Add 6 to the end of the array"
p arr

arr.shift  # arr.delete_at(0), update arr
p "Remove the first element of the array (destructive)"
p arr

arr.slice(1..-1)  # arr[1..-1], not destructive
p "Remove the first element of the array (non-destructive)"
p arr, arr.slice(1..-1)

last_element = arr[-1]  # arr.last
p "Get the last element of the array"
p last_element

sub_arr = arr[1..3]  # 2 ~ 4
p "Get a sub-array from the array"
p sub_arr

len_arr = arr.length  # arr.size
p "Get the length of the array"
p len_arr

sum_arr = arr.sum
p "Get the sum of the array"
p sum_arr

avg_arr = arr.sum / arr.size.to_f
p "Get the average of the array"
p avg_arr

arr.sort!.reverse!
p "Sort the array in descending order"
p arr

contains_1 = arr.include?(1)
p "Check if the array contains 1"
p contains_1

arr.uniq!
p "Remove duplicate elements from the array"
p arr
