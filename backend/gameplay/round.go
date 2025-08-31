package gameplay

const (
	RUN_INDEX = 0
)

type RoundObjective struct {
	RoundNumber int8
	Sets        []SetRequirement
}

var RoundObjectives = []RoundObjective{
	{
		RoundNumber: 1,
		Sets: []SetRequirement{
			{Type: Book, MinLength: 3},
			{Type: Book, MinLength: 3},
		},
	},
	{
		RoundNumber: 2,
		Sets: []SetRequirement{
			{Type: Run, MinLength: 4},
			{Type: Book, MinLength: 3},
		},
	},
	{
		RoundNumber: 3,
		Sets: []SetRequirement{
			{Type: Book, MinLength: 3},
			{Type: Book, MinLength: 3},
			{Type: Book, MinLength: 3},
		},
	},
	{
		RoundNumber: 4,
		Sets: []SetRequirement{
			{Type: Run, MinLength: 4},
			{Type: Book, MinLength: 3},
			{Type: Book, MinLength: 3},
		},
	},
	{
		RoundNumber: 5,
		Sets: []SetRequirement{
			{Type: Run, MinLength: 5},
			{Type: Book, MinLength: 3},
		},
	},
	{
		RoundNumber: 6,
		Sets: []SetRequirement{
			{Type: Run, MinLength: 6},
			{Type: Book, MinLength: 3},
			{Type: Book, MinLength: 3},
		},
	},
	{
		RoundNumber: 7,
		Sets: []SetRequirement{
			{Type: Run, MinLength: 7},
			{Type: Book, MinLength: 3},
		},
	},
}
