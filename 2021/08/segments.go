package main

var (
	validDigits = []string{"a", "b", "c", "d", "e", "f", "g"}

	renderedSegments = map[int]string{
		0: `
		 aaaa 
		b    c
		b    c
		 .... 
		e    f
		e    f
		 gggg `,
		1: `
		 .... 
		.    c
		.    c
		 .... 
		.    f
		.    f
		 .... `,
		2: `
		 aaaa 
		.    c
		.    c
		 dddd 
		e    .
		e    .
		 gggg `,
		3: `
		 aaaa 
		.    c
		.    c
		 dddd 
		.    f
		.    f
		 gggg `,
		4: `
		 .... 
		b    c
		b    c
		 dddd 
		.    f
		.    f
		 .... `,
		5: `
		 aaaa 
		b    .
		b    .
		 dddd 
		.    f
		.    f
		 gggg `,
		6: `
		 aaaa 
		b    .
		b    .
		 dddd 
		e    f
		e    f
		 gggg `,
		7: `
		 aaaa 
		.    c
		.    c
		 .... 
		.    f
		.    f
		 .... `,
		8: `
		 aaaa 
		b    c
		b    c
		 dddd 
		e    f
		e    f
		 gggg `,
		9: `
		 aaaa 
		b    c
		b    c
		 dddd 
		.    f
		.    f
		 gggg `,
	}

	newMapping = map[string]string{
		"a": "d",
		"b": "e",
		"c": "a",
		"d": "f",
		"e": "g",
		"f": "b",
		"g": "c",
	}
)
