module github.com/javorszky/fork

go 1.18

replace github.com/javorszky/sub => ../sub

require github.com/rs/zerolog v1.26.1

require (
	github.com/google/uuid v1.3.0
	github.com/javorszky/sub v0.0.0-00010101000000-000000000000
)
