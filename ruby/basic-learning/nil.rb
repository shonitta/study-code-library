require 'active_support'
require 'active_support/core_ext'


puts "Check if a variable is nil"
v_nil = nil
p v_nil.nil?  # true

v_not_nil = ""
p v_not_nil.nil?  # false


puts "Check if a variable is blank"
v_blank = ""
p v_blank.blank?  # true

v_not_blank = " "
p v_not_blank.blank?  # false


puts "Check if a variable is present"
v_present = " "
p v_present.present?  # true

v_not_present = ""
p v_not_present.present?  # false


puts "Check if a variable is empty"
v_empty = ""
p v_empty.empty?  # true

v_not_empty = " "
p v_not_empty.empty?  # false

v_raise_error = nil
p v_raise_error.empty?  # raise error
