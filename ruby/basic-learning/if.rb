city = "New York"

if city == "New York" then
    puts "Welcome to the Big Apple!"
elsif city == "Paris" then
    puts "Welcome to the City of Lights!"
elsif city == "London" then
    puts "Welcome to the Big Smoke!"
elsif city == "Sydney" then
    puts "Welcome to the Harbour City!"
elsif city == "Tokyo" then
    puts "Welcome to the Big Mikan!"
else
    puts "Have a nice day!"
end


animal_1 = "cat"
animal_2 = "dog"

puts animal_1 == "cat" || animal_2 == "dog"  # true
puts animal_1 == "cat" && animal_2 == "dog"  # true
puts animal_1 == "cat" && animal_2 == "cat"  # false
puts animal_1 == "dog" && animal_2 == "dog"  # false
puts !(animal_1 == "cat")  # false
puts animal_1 == "cat" ? "Meow!" : "Woof!"  # Meow!
puts animal_1 == "cat" && animal_2 == "dog" ? "Same!" : "Different!"  # Same!
