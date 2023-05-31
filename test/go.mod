module test

go 1.18

require searchAlgorithms v0.0.0-00010101000000-000000000000
require custom_types v0.0.0-00010101000000-000000000000
require sortAlgorithms v0.0.0-00010101000000-000000000000

replace searchAlgorithms => ./../search
replace custom_types => ./../custom_types
replace sortAlgorithms => ./../sort