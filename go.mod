module goHandson

go 1.16

replace local.packages/hello => ./hello

require local.packages/hello v0.0.0-00010101000000-000000000000 // indirect
