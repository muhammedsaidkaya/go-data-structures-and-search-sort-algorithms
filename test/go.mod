module test

go 1.18

require search v0.0.0-00010101000000-000000000000
require custom_types v0.0.0-00010101000000-000000000000

replace search => ./../search
replace custom_types => ./../custom_types