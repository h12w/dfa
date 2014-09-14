package dfa

var classL = &M{
	Start: 0,
	States: States{{Table: TransTable{
		{Lo: 0x41, Hi: 0x5a, Next: 1},
		{Lo: 0x61, Hi: 0x7a, Next: 1},
		{Lo: 0xc2, Hi: 0xc2, Next: 2},
		{Lo: 0xc3, Hi: 0xc3, Next: 3},
		{Lo: 0xc4, Hi: 0xca, Next: 4},
		{Lo: 0xcb, Hi: 0xcb, Next: 5},
		{Lo: 0xcd, Hi: 0xcd, Next: 6},
		{Lo: 0xce, Hi: 0xce, Next: 7},
		{Lo: 0xcf, Hi: 0xcf, Next: 8},
		{Lo: 0xd0, Hi: 0xd1, Next: 4},
		{Lo: 0xd2, Hi: 0xd2, Next: 9},
		{Lo: 0xd3, Hi: 0xd3, Next: 4},
		{Lo: 0xd4, Hi: 0xd4, Next: 10},
		{Lo: 0xd5, Hi: 0xd5, Next: 11},
		{Lo: 0xd6, Hi: 0xd6, Next: 12},
		{Lo: 0xd7, Hi: 0xd7, Next: 13},
		{Lo: 0xd8, Hi: 0xd8, Next: 14},
		{Lo: 0xd9, Hi: 0xd9, Next: 15},
		{Lo: 0xda, Hi: 0xda, Next: 4},
		{Lo: 0xdb, Hi: 0xdb, Next: 16},
		{Lo: 0xdc, Hi: 0xdc, Next: 17},
		{Lo: 0xdd, Hi: 0xdd, Next: 18},
		{Lo: 0xde, Hi: 0xde, Next: 19},
		{Lo: 0xdf, Hi: 0xdf, Next: 20},
		{Lo: 0xe0, Hi: 0xe0, Next: 21},
		{Lo: 0xe1, Hi: 0xe1, Next: 22},
		{Lo: 0xe2, Hi: 0xe2, Next: 23},
		{Lo: 0xe3, Hi: 0xe3, Next: 24},
		{Lo: 0xe4, Hi: 0xe4, Next: 25},
		{Lo: 0xe5, Hi: 0xe8, Next: 26},
		{Lo: 0xe9, Hi: 0xe9, Next: 27},
		{Lo: 0xea, Hi: 0xea, Next: 28},
		{Lo: 0xeb, Hi: 0xec, Next: 26},
		{Lo: 0xed, Hi: 0xed, Next: 29},
		{Lo: 0xef, Hi: 0xef, Next: 30},
		{Lo: 0xf0, Hi: 0xf0, Next: 31}}, Label: 0},
		{Table: TransTable(nil), Label: 1},
		{Table: TransTable{
			{Lo: 0xaa, Hi: 0xaa, Next: 1},
			{Lo: 0xb5, Hi: 0xb5, Next: 1},
			{Lo: 0xba, Hi: 0xba, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0x96, Next: 1},
			{Lo: 0x98, Hi: 0xb6, Next: 1},
			{Lo: 0xb8, Hi: 0xbf, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0xbf, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0x81, Next: 1},
			{Lo: 0x86, Hi: 0x91, Next: 1},
			{Lo: 0xa0, Hi: 0xa4, Next: 1},
			{Lo: 0xac, Hi: 0xac, Next: 1},
			{Lo: 0xae, Hi: 0xae, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0xb0, Hi: 0xb4, Next: 1},
			{Lo: 0xb6, Hi: 0xb7, Next: 1},
			{Lo: 0xba, Hi: 0xbd, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x86, Hi: 0x86, Next: 1},
			{Lo: 0x88, Hi: 0x8a, Next: 1},
			{Lo: 0x8c, Hi: 0x8c, Next: 1},
			{Lo: 0x8e, Hi: 0xa1, Next: 1},
			{Lo: 0xa3, Hi: 0xbf, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0xb5, Next: 1},
			{Lo: 0xb7, Hi: 0xbf, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0x81, Next: 1},
			{Lo: 0x8a, Hi: 0xbf, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0xa7, Next: 1},
			{Lo: 0xb1, Hi: 0xbf, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0x96, Next: 1},
			{Lo: 0x99, Hi: 0x99, Next: 1},
			{Lo: 0xa1, Hi: 0xbf, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0x87, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x90, Hi: 0xaa, Next: 1},
			{Lo: 0xb0, Hi: 0xb2, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0xa0, Hi: 0xbf, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0x8a, Next: 1},
			{Lo: 0xae, Hi: 0xaf, Next: 1},
			{Lo: 0xb1, Hi: 0xbf, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0x93, Next: 1},
			{Lo: 0x95, Hi: 0x95, Next: 1},
			{Lo: 0xa5, Hi: 0xa6, Next: 1},
			{Lo: 0xae, Hi: 0xaf, Next: 1},
			{Lo: 0xba, Hi: 0xbc, Next: 1},
			{Lo: 0xbf, Hi: 0xbf, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x90, Hi: 0x90, Next: 1},
			{Lo: 0x92, Hi: 0xaf, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x8d, Hi: 0xbf, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0xa5, Next: 1},
			{Lo: 0xb1, Hi: 0xb1, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x8a, Hi: 0xaa, Next: 1},
			{Lo: 0xb4, Hi: 0xb5, Next: 1},
			{Lo: 0xba, Hi: 0xba, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0xa0, Hi: 0xa0, Next: 32},
			{Lo: 0xa1, Hi: 0xa1, Next: 33},
			{Lo: 0xa2, Hi: 0xa2, Next: 34},
			{Lo: 0xa4, Hi: 0xa4, Next: 35},
			{Lo: 0xa5, Hi: 0xa5, Next: 36},
			{Lo: 0xa6, Hi: 0xa6, Next: 37},
			{Lo: 0xa7, Hi: 0xa7, Next: 38},
			{Lo: 0xa8, Hi: 0xa8, Next: 39},
			{Lo: 0xa9, Hi: 0xa9, Next: 40},
			{Lo: 0xaa, Hi: 0xaa, Next: 41},
			{Lo: 0xab, Hi: 0xab, Next: 42},
			{Lo: 0xac, Hi: 0xac, Next: 43},
			{Lo: 0xad, Hi: 0xad, Next: 44},
			{Lo: 0xae, Hi: 0xae, Next: 45},
			{Lo: 0xaf, Hi: 0xaf, Next: 46},
			{Lo: 0xb0, Hi: 0xb0, Next: 47},
			{Lo: 0xb1, Hi: 0xb1, Next: 48},
			{Lo: 0xb2, Hi: 0xb2, Next: 47},
			{Lo: 0xb3, Hi: 0xb3, Next: 49},
			{Lo: 0xb4, Hi: 0xb4, Next: 50},
			{Lo: 0xb5, Hi: 0xb5, Next: 51},
			{Lo: 0xb6, Hi: 0xb6, Next: 52},
			{Lo: 0xb7, Hi: 0xb7, Next: 53},
			{Lo: 0xb8, Hi: 0xb8, Next: 54},
			{Lo: 0xb9, Hi: 0xb9, Next: 53},
			{Lo: 0xba, Hi: 0xba, Next: 55},
			{Lo: 0xbb, Hi: 0xbb, Next: 56},
			{Lo: 0xbc, Hi: 0xbc, Next: 57},
			{Lo: 0xbd, Hi: 0xbd, Next: 58},
			{Lo: 0xbe, Hi: 0xbe, Next: 59}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0x80, Next: 60},
			{Lo: 0x81, Hi: 0x81, Next: 61},
			{Lo: 0x82, Hi: 0x82, Next: 62},
			{Lo: 0x83, Hi: 0x83, Next: 63},
			{Lo: 0x84, Hi: 0x88, Next: 4},
			{Lo: 0x89, Hi: 0x89, Next: 64},
			{Lo: 0x8a, Hi: 0x8a, Next: 65},
			{Lo: 0x8b, Hi: 0x8b, Next: 66},
			{Lo: 0x8c, Hi: 0x8c, Next: 67},
			{Lo: 0x8d, Hi: 0x8d, Next: 68},
			{Lo: 0x8e, Hi: 0x8e, Next: 69},
			{Lo: 0x8f, Hi: 0x8f, Next: 70},
			{Lo: 0x90, Hi: 0x90, Next: 71},
			{Lo: 0x91, Hi: 0x98, Next: 4},
			{Lo: 0x99, Hi: 0x99, Next: 72},
			{Lo: 0x9a, Hi: 0x9a, Next: 73},
			{Lo: 0x9b, Hi: 0x9b, Next: 74},
			{Lo: 0x9c, Hi: 0x9c, Next: 75},
			{Lo: 0x9d, Hi: 0x9d, Next: 76},
			{Lo: 0x9e, Hi: 0x9e, Next: 77},
			{Lo: 0x9f, Hi: 0x9f, Next: 78},
			{Lo: 0xa0, Hi: 0xa0, Next: 14},
			{Lo: 0xa1, Hi: 0xa1, Next: 79},
			{Lo: 0xa2, Hi: 0xa2, Next: 80},
			{Lo: 0xa3, Hi: 0xa3, Next: 81},
			{Lo: 0xa4, Hi: 0xa4, Next: 82},
			{Lo: 0xa5, Hi: 0xa5, Next: 83},
			{Lo: 0xa6, Hi: 0xa6, Next: 84},
			{Lo: 0xa7, Hi: 0xa7, Next: 85},
			{Lo: 0xa8, Hi: 0xa8, Next: 86},
			{Lo: 0xa9, Hi: 0xa9, Next: 87},
			{Lo: 0xaa, Hi: 0xaa, Next: 88},
			{Lo: 0xac, Hi: 0xac, Next: 89},
			{Lo: 0xad, Hi: 0xad, Next: 90},
			{Lo: 0xae, Hi: 0xae, Next: 91},
			{Lo: 0xaf, Hi: 0xaf, Next: 92},
			{Lo: 0xb0, Hi: 0xb0, Next: 93},
			{Lo: 0xb1, Hi: 0xb1, Next: 94},
			{Lo: 0xb3, Hi: 0xb3, Next: 95},
			{Lo: 0xb4, Hi: 0xb6, Next: 4},
			{Lo: 0xb8, Hi: 0xbb, Next: 4},
			{Lo: 0xbc, Hi: 0xbc, Next: 96},
			{Lo: 0xbd, Hi: 0xbd, Next: 97},
			{Lo: 0xbe, Hi: 0xbe, Next: 98},
			{Lo: 0xbf, Hi: 0xbf, Next: 99}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x81, Hi: 0x81, Next: 100},
			{Lo: 0x82, Hi: 0x82, Next: 101},
			{Lo: 0x84, Hi: 0x84, Next: 102},
			{Lo: 0x85, Hi: 0x85, Next: 103},
			{Lo: 0x86, Hi: 0x86, Next: 104},
			{Lo: 0xb0, Hi: 0xb0, Next: 105},
			{Lo: 0xb1, Hi: 0xb1, Next: 106},
			{Lo: 0xb2, Hi: 0xb2, Next: 4},
			{Lo: 0xb3, Hi: 0xb3, Next: 107},
			{Lo: 0xb4, Hi: 0xb4, Next: 108},
			{Lo: 0xb5, Hi: 0xb5, Next: 109},
			{Lo: 0xb6, Hi: 0xb6, Next: 110},
			{Lo: 0xb7, Hi: 0xb7, Next: 111},
			{Lo: 0xb8, Hi: 0xb8, Next: 112}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0x80, Next: 113},
			{Lo: 0x81, Hi: 0x81, Next: 71},
			{Lo: 0x82, Hi: 0x82, Next: 114},
			{Lo: 0x83, Hi: 0x83, Next: 115},
			{Lo: 0x84, Hi: 0x84, Next: 116},
			{Lo: 0x85, Hi: 0x85, Next: 4},
			{Lo: 0x86, Hi: 0x86, Next: 117},
			{Lo: 0x87, Hi: 0x87, Next: 118},
			{Lo: 0x90, Hi: 0xbf, Next: 4}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0xb5, Next: 4},
			{Lo: 0xb6, Hi: 0xb6, Next: 81},
			{Lo: 0xb8, Hi: 0xbf, Next: 4}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0xbf, Next: 4}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0xbe, Next: 4},
			{Lo: 0xbf, Hi: 0xbf, Next: 119}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0x91, Next: 4},
			{Lo: 0x92, Hi: 0x92, Next: 119},
			{Lo: 0x93, Hi: 0x93, Next: 120},
			{Lo: 0x94, Hi: 0x97, Next: 4},
			{Lo: 0x98, Hi: 0x98, Next: 121},
			{Lo: 0x99, Hi: 0x99, Next: 122},
			{Lo: 0x9a, Hi: 0x9a, Next: 123},
			{Lo: 0x9b, Hi: 0x9b, Next: 92},
			{Lo: 0x9c, Hi: 0x9c, Next: 124},
			{Lo: 0x9d, Hi: 0x9d, Next: 4},
			{Lo: 0x9e, Hi: 0x9e, Next: 125},
			{Lo: 0x9f, Hi: 0x9f, Next: 126},
			{Lo: 0xa0, Hi: 0xa0, Next: 127},
			{Lo: 0xa1, Hi: 0xa1, Next: 77},
			{Lo: 0xa2, Hi: 0xa2, Next: 128},
			{Lo: 0xa3, Hi: 0xa3, Next: 129},
			{Lo: 0xa4, Hi: 0xa4, Next: 130},
			{Lo: 0xa5, Hi: 0xa5, Next: 131},
			{Lo: 0xa6, Hi: 0xa6, Next: 132},
			{Lo: 0xa7, Hi: 0xa7, Next: 133},
			{Lo: 0xa8, Hi: 0xa8, Next: 134},
			{Lo: 0xa9, Hi: 0xa9, Next: 135},
			{Lo: 0xaa, Hi: 0xaa, Next: 136},
			{Lo: 0xab, Hi: 0xab, Next: 137},
			{Lo: 0xac, Hi: 0xac, Next: 138},
			{Lo: 0xaf, Hi: 0xaf, Next: 139},
			{Lo: 0xb0, Hi: 0xbf, Next: 4}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0x9d, Next: 4},
			{Lo: 0x9e, Hi: 0x9e, Next: 140},
			{Lo: 0x9f, Hi: 0x9f, Next: 141}}, Label: 0},
		{Table: TransTable{
			{Lo: 0xa4, Hi: 0xa8, Next: 4},
			{Lo: 0xa9, Hi: 0xa9, Next: 142},
			{Lo: 0xaa, Hi: 0xaa, Next: 4},
			{Lo: 0xab, Hi: 0xab, Next: 143},
			{Lo: 0xac, Hi: 0xac, Next: 144},
			{Lo: 0xad, Hi: 0xad, Next: 145},
			{Lo: 0xae, Hi: 0xae, Next: 146},
			{Lo: 0xaf, Hi: 0xaf, Next: 147},
			{Lo: 0xb0, Hi: 0xb3, Next: 4},
			{Lo: 0xb4, Hi: 0xb4, Next: 148},
			{Lo: 0xb5, Hi: 0xb5, Next: 149},
			{Lo: 0xb6, Hi: 0xb6, Next: 150},
			{Lo: 0xb7, Hi: 0xb7, Next: 151},
			{Lo: 0xb9, Hi: 0xb9, Next: 152},
			{Lo: 0xba, Hi: 0xba, Next: 4},
			{Lo: 0xbb, Hi: 0xbb, Next: 153},
			{Lo: 0xbc, Hi: 0xbc, Next: 154},
			{Lo: 0xbd, Hi: 0xbd, Next: 155},
			{Lo: 0xbe, Hi: 0xbe, Next: 156},
			{Lo: 0xbf, Hi: 0xbf, Next: 157}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x90, Hi: 0x90, Next: 158},
			{Lo: 0x91, Hi: 0x91, Next: 159},
			{Lo: 0x92, Hi: 0x92, Next: 160},
			{Lo: 0x93, Hi: 0x93, Next: 161},
			{Lo: 0x96, Hi: 0x96, Next: 162},
			{Lo: 0x9b, Hi: 0x9b, Next: 163},
			{Lo: 0x9d, Hi: 0x9d, Next: 164},
			{Lo: 0x9e, Hi: 0x9e, Next: 165},
			{Lo: 0xa0, Hi: 0xa9, Next: 26},
			{Lo: 0xaa, Hi: 0xaa, Next: 166},
			{Lo: 0xab, Hi: 0xab, Next: 167},
			{Lo: 0xaf, Hi: 0xaf, Next: 168}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0x95, Next: 1},
			{Lo: 0x9a, Hi: 0x9a, Next: 1},
			{Lo: 0xa4, Hi: 0xa4, Next: 1},
			{Lo: 0xa8, Hi: 0xa8, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0x98, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0xa0, Hi: 0xa0, Next: 1},
			{Lo: 0xa2, Hi: 0xac, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x84, Hi: 0xb9, Next: 1},
			{Lo: 0xbd, Hi: 0xbd, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x90, Hi: 0x90, Next: 1},
			{Lo: 0x98, Hi: 0xa1, Next: 1},
			{Lo: 0xb1, Hi: 0xb7, Next: 1},
			{Lo: 0xb9, Hi: 0xbf, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x85, Hi: 0x8c, Next: 1},
			{Lo: 0x8f, Hi: 0x90, Next: 1},
			{Lo: 0x93, Hi: 0xa8, Next: 1},
			{Lo: 0xaa, Hi: 0xb0, Next: 1},
			{Lo: 0xb2, Hi: 0xb2, Next: 1},
			{Lo: 0xb6, Hi: 0xb9, Next: 1},
			{Lo: 0xbd, Hi: 0xbd, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x8e, Hi: 0x8e, Next: 1},
			{Lo: 0x9c, Hi: 0x9d, Next: 1},
			{Lo: 0x9f, Hi: 0xa1, Next: 1},
			{Lo: 0xb0, Hi: 0xb1, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x85, Hi: 0x8a, Next: 1},
			{Lo: 0x8f, Hi: 0x90, Next: 1},
			{Lo: 0x93, Hi: 0xa8, Next: 1},
			{Lo: 0xaa, Hi: 0xb0, Next: 1},
			{Lo: 0xb2, Hi: 0xb3, Next: 1},
			{Lo: 0xb5, Hi: 0xb6, Next: 1},
			{Lo: 0xb8, Hi: 0xb9, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x99, Hi: 0x9c, Next: 1},
			{Lo: 0x9e, Hi: 0x9e, Next: 1},
			{Lo: 0xb2, Hi: 0xb4, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x85, Hi: 0x8d, Next: 1},
			{Lo: 0x8f, Hi: 0x91, Next: 1},
			{Lo: 0x93, Hi: 0xa8, Next: 1},
			{Lo: 0xaa, Hi: 0xb0, Next: 1},
			{Lo: 0xb2, Hi: 0xb3, Next: 1},
			{Lo: 0xb5, Hi: 0xb9, Next: 1},
			{Lo: 0xbd, Hi: 0xbd, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x90, Hi: 0x90, Next: 1},
			{Lo: 0xa0, Hi: 0xa1, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x85, Hi: 0x8c, Next: 1},
			{Lo: 0x8f, Hi: 0x90, Next: 1},
			{Lo: 0x93, Hi: 0xa8, Next: 1},
			{Lo: 0xaa, Hi: 0xb0, Next: 1},
			{Lo: 0xb2, Hi: 0xb3, Next: 1},
			{Lo: 0xb5, Hi: 0xb9, Next: 1},
			{Lo: 0xbd, Hi: 0xbd, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x9c, Hi: 0x9d, Next: 1},
			{Lo: 0x9f, Hi: 0xa1, Next: 1},
			{Lo: 0xb1, Hi: 0xb1, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x83, Hi: 0x83, Next: 1},
			{Lo: 0x85, Hi: 0x8a, Next: 1},
			{Lo: 0x8e, Hi: 0x90, Next: 1},
			{Lo: 0x92, Hi: 0x95, Next: 1},
			{Lo: 0x99, Hi: 0x9a, Next: 1},
			{Lo: 0x9c, Hi: 0x9c, Next: 1},
			{Lo: 0x9e, Hi: 0x9f, Next: 1},
			{Lo: 0xa3, Hi: 0xa4, Next: 1},
			{Lo: 0xa8, Hi: 0xaa, Next: 1},
			{Lo: 0xae, Hi: 0xb9, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x90, Hi: 0x90, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x85, Hi: 0x8c, Next: 1},
			{Lo: 0x8e, Hi: 0x90, Next: 1},
			{Lo: 0x92, Hi: 0xa8, Next: 1},
			{Lo: 0xaa, Hi: 0xb3, Next: 1},
			{Lo: 0xb5, Hi: 0xb9, Next: 1},
			{Lo: 0xbd, Hi: 0xbd, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x98, Hi: 0x99, Next: 1},
			{Lo: 0xa0, Hi: 0xa1, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x9e, Hi: 0x9e, Next: 1},
			{Lo: 0xa0, Hi: 0xa1, Next: 1},
			{Lo: 0xb1, Hi: 0xb2, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x85, Hi: 0x8c, Next: 1},
			{Lo: 0x8e, Hi: 0x90, Next: 1},
			{Lo: 0x92, Hi: 0xba, Next: 1},
			{Lo: 0xbd, Hi: 0xbd, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x8e, Hi: 0x8e, Next: 1},
			{Lo: 0xa0, Hi: 0xa1, Next: 1},
			{Lo: 0xba, Hi: 0xbf, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x85, Hi: 0x96, Next: 1},
			{Lo: 0x9a, Hi: 0xb1, Next: 1},
			{Lo: 0xb3, Hi: 0xbb, Next: 1},
			{Lo: 0xbd, Hi: 0xbd, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0x86, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x81, Hi: 0xb0, Next: 1},
			{Lo: 0xb2, Hi: 0xb3, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x81, Hi: 0x82, Next: 1},
			{Lo: 0x84, Hi: 0x84, Next: 1},
			{Lo: 0x87, Hi: 0x88, Next: 1},
			{Lo: 0x8a, Hi: 0x8a, Next: 1},
			{Lo: 0x8d, Hi: 0x8d, Next: 1},
			{Lo: 0x94, Hi: 0x97, Next: 1},
			{Lo: 0x99, Hi: 0x9f, Next: 1},
			{Lo: 0xa1, Hi: 0xa3, Next: 1},
			{Lo: 0xa5, Hi: 0xa5, Next: 1},
			{Lo: 0xa7, Hi: 0xa7, Next: 1},
			{Lo: 0xaa, Hi: 0xab, Next: 1},
			{Lo: 0xad, Hi: 0xb0, Next: 1},
			{Lo: 0xb2, Hi: 0xb3, Next: 1},
			{Lo: 0xbd, Hi: 0xbd, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0x84, Next: 1},
			{Lo: 0x86, Hi: 0x86, Next: 1},
			{Lo: 0x9c, Hi: 0x9f, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0x80, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0x87, Next: 1},
			{Lo: 0x89, Hi: 0xac, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x88, Hi: 0x8c, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0xaa, Next: 1},
			{Lo: 0xbf, Hi: 0xbf, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x90, Hi: 0x95, Next: 1},
			{Lo: 0x9a, Hi: 0x9d, Next: 1},
			{Lo: 0xa1, Hi: 0xa1, Next: 1},
			{Lo: 0xa5, Hi: 0xa6, Next: 1},
			{Lo: 0xae, Hi: 0xb0, Next: 1},
			{Lo: 0xb5, Hi: 0xbf, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0x81, Next: 1},
			{Lo: 0x8e, Hi: 0x8e, Next: 1},
			{Lo: 0xa0, Hi: 0xbf, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0x85, Next: 1},
			{Lo: 0x87, Hi: 0x87, Next: 1},
			{Lo: 0x8d, Hi: 0x8d, Next: 1},
			{Lo: 0x90, Hi: 0xba, Next: 1},
			{Lo: 0xbc, Hi: 0xbf, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0x88, Next: 1},
			{Lo: 0x8a, Hi: 0x8d, Next: 1},
			{Lo: 0x90, Hi: 0x96, Next: 1},
			{Lo: 0x98, Hi: 0x98, Next: 1},
			{Lo: 0x9a, Hi: 0x9d, Next: 1},
			{Lo: 0xa0, Hi: 0xbf, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0x88, Next: 1},
			{Lo: 0x8a, Hi: 0x8d, Next: 1},
			{Lo: 0x90, Hi: 0xb0, Next: 1},
			{Lo: 0xb2, Hi: 0xb5, Next: 1},
			{Lo: 0xb8, Hi: 0xbe, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0x80, Next: 1},
			{Lo: 0x82, Hi: 0x85, Next: 1},
			{Lo: 0x88, Hi: 0x96, Next: 1},
			{Lo: 0x98, Hi: 0xbf, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0x90, Next: 1},
			{Lo: 0x92, Hi: 0x95, Next: 1},
			{Lo: 0x98, Hi: 0xbf, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0x9a, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0x8f, Next: 1},
			{Lo: 0xa0, Hi: 0xbf, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0xb4, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x81, Hi: 0xbf, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0xac, Next: 1},
			{Lo: 0xaf, Hi: 0xbf, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x81, Hi: 0x9a, Next: 1},
			{Lo: 0xa0, Hi: 0xbf, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0xaa, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0x8c, Next: 1},
			{Lo: 0x8e, Hi: 0x91, Next: 1},
			{Lo: 0xa0, Hi: 0xb1, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0x91, Next: 1},
			{Lo: 0xa0, Hi: 0xac, Next: 1},
			{Lo: 0xae, Hi: 0xb0, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0xb3, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x97, Hi: 0x97, Next: 1},
			{Lo: 0x9c, Hi: 0x9c, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0xb7, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0xa8, Next: 1},
			{Lo: 0xaa, Hi: 0xaa, Next: 1},
			{Lo: 0xb0, Hi: 0xbf, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0xb5, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0x9c, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x90, Hi: 0xad, Next: 1},
			{Lo: 0xb0, Hi: 0xb4, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0xab, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x81, Hi: 0x87, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0x96, Next: 1},
			{Lo: 0xa0, Hi: 0xbf, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0x94, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0xa7, Hi: 0xa7, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x85, Hi: 0xb3, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x85, Hi: 0x8b, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x83, Hi: 0xa0, Next: 1},
			{Lo: 0xae, Hi: 0xaf, Next: 1},
			{Lo: 0xba, Hi: 0xbf, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0xa5, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0xa3, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x8d, Hi: 0x8f, Next: 1},
			{Lo: 0x9a, Hi: 0xbd, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0xa9, Hi: 0xac, Next: 1},
			{Lo: 0xae, Hi: 0xb1, Next: 1},
			{Lo: 0xb5, Hi: 0xb6, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0x95, Next: 1},
			{Lo: 0x98, Hi: 0x9d, Next: 1},
			{Lo: 0xa0, Hi: 0xbf, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0x85, Next: 1},
			{Lo: 0x88, Hi: 0x8d, Next: 1},
			{Lo: 0x90, Hi: 0x97, Next: 1},
			{Lo: 0x99, Hi: 0x99, Next: 1},
			{Lo: 0x9b, Hi: 0x9b, Next: 1},
			{Lo: 0x9d, Hi: 0x9d, Next: 1},
			{Lo: 0x9f, Hi: 0xbd, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0xb4, Next: 1},
			{Lo: 0xb6, Hi: 0xbc, Next: 1},
			{Lo: 0xbe, Hi: 0xbe, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x82, Hi: 0x84, Next: 1},
			{Lo: 0x86, Hi: 0x8c, Next: 1},
			{Lo: 0x90, Hi: 0x93, Next: 1},
			{Lo: 0x96, Hi: 0x9b, Next: 1},
			{Lo: 0xa0, Hi: 0xac, Next: 1},
			{Lo: 0xb2, Hi: 0xb4, Next: 1},
			{Lo: 0xb6, Hi: 0xbc, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0xb1, Hi: 0xb1, Next: 1},
			{Lo: 0xbf, Hi: 0xbf, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x90, Hi: 0x9c, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x82, Hi: 0x82, Next: 1},
			{Lo: 0x87, Hi: 0x87, Next: 1},
			{Lo: 0x8a, Hi: 0x93, Next: 1},
			{Lo: 0x95, Hi: 0x95, Next: 1},
			{Lo: 0x99, Hi: 0x9d, Next: 1},
			{Lo: 0xa4, Hi: 0xa4, Next: 1},
			{Lo: 0xa6, Hi: 0xa6, Next: 1},
			{Lo: 0xa8, Hi: 0xa8, Next: 1},
			{Lo: 0xaa, Hi: 0xad, Next: 1},
			{Lo: 0xaf, Hi: 0xb9, Next: 1},
			{Lo: 0xbc, Hi: 0xbf, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x85, Hi: 0x89, Next: 1},
			{Lo: 0x8e, Hi: 0x8e, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x83, Hi: 0x84, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0xae, Next: 1},
			{Lo: 0xb0, Hi: 0xbf, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0x9e, Next: 1},
			{Lo: 0xa0, Hi: 0xbf, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0xa4, Next: 1},
			{Lo: 0xab, Hi: 0xae, Next: 1},
			{Lo: 0xb2, Hi: 0xb3, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0xa5, Next: 1},
			{Lo: 0xa7, Hi: 0xa7, Next: 1},
			{Lo: 0xad, Hi: 0xad, Next: 1},
			{Lo: 0xb0, Hi: 0xbf, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0xa7, Next: 1},
			{Lo: 0xaf, Hi: 0xaf, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0x96, Next: 1},
			{Lo: 0xa0, Hi: 0xa6, Next: 1},
			{Lo: 0xa8, Hi: 0xae, Next: 1},
			{Lo: 0xb0, Hi: 0xb6, Next: 1},
			{Lo: 0xb8, Hi: 0xbe, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0x86, Next: 1},
			{Lo: 0x88, Hi: 0x8e, Next: 1},
			{Lo: 0x90, Hi: 0x96, Next: 1},
			{Lo: 0x98, Hi: 0x9e, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0xaf, Hi: 0xaf, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x85, Hi: 0x86, Next: 1},
			{Lo: 0xb1, Hi: 0xb5, Next: 1},
			{Lo: 0xbb, Hi: 0xbc, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0x96, Next: 1},
			{Lo: 0x9d, Hi: 0x9f, Next: 1},
			{Lo: 0xa1, Hi: 0xbf, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0xba, Next: 1},
			{Lo: 0xbc, Hi: 0xbf, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x85, Hi: 0xad, Next: 1},
			{Lo: 0xb1, Hi: 0xbf, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0x8e, Next: 1},
			{Lo: 0xa0, Hi: 0xba, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0xb0, Hi: 0xbf, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0x8c, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x90, Hi: 0xbd, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0x8c, Next: 1},
			{Lo: 0x90, Hi: 0x9f, Next: 1},
			{Lo: 0xaa, Hi: 0xab, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0xae, Next: 1},
			{Lo: 0xbf, Hi: 0xbf, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0x97, Next: 1},
			{Lo: 0xa0, Hi: 0xbf, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x97, Hi: 0x9f, Next: 1},
			{Lo: 0xa2, Hi: 0xbf, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0x88, Next: 1},
			{Lo: 0x8b, Hi: 0x8e, Next: 1},
			{Lo: 0x90, Hi: 0x93, Next: 1},
			{Lo: 0xa0, Hi: 0xaa, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0xb8, Hi: 0xbf, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0x81, Next: 1},
			{Lo: 0x83, Hi: 0x85, Next: 1},
			{Lo: 0x87, Hi: 0x8a, Next: 1},
			{Lo: 0x8c, Hi: 0xa2, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x82, Hi: 0xb3, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0xb2, Hi: 0xb7, Next: 1},
			{Lo: 0xbb, Hi: 0xbb, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x8a, Hi: 0xa5, Next: 1},
			{Lo: 0xb0, Hi: 0xbf, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0x86, Next: 1},
			{Lo: 0xa0, Hi: 0xbc, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x84, Hi: 0xb2, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x8f, Hi: 0x8f, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0xa8, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0x82, Next: 1},
			{Lo: 0x84, Hi: 0x8b, Next: 1},
			{Lo: 0xa0, Hi: 0xb6, Next: 1},
			{Lo: 0xba, Hi: 0xba, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0xaf, Next: 1},
			{Lo: 0xb1, Hi: 0xb1, Next: 1},
			{Lo: 0xb5, Hi: 0xb6, Next: 1},
			{Lo: 0xb9, Hi: 0xbd, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0x80, Next: 1},
			{Lo: 0x82, Hi: 0x82, Next: 1},
			{Lo: 0x9b, Hi: 0x9d, Next: 1},
			{Lo: 0xa0, Hi: 0xaa, Next: 1},
			{Lo: 0xb2, Hi: 0xb4, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x81, Hi: 0x86, Next: 1},
			{Lo: 0x89, Hi: 0x8e, Next: 1},
			{Lo: 0x91, Hi: 0x96, Next: 1},
			{Lo: 0xa0, Hi: 0xa6, Next: 1},
			{Lo: 0xa8, Hi: 0xae, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0xa2, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0xa3, Next: 1},
			{Lo: 0xb0, Hi: 0xbf, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0x86, Next: 1},
			{Lo: 0x8b, Hi: 0xbb, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0xad, Next: 1},
			{Lo: 0xb0, Hi: 0xbf, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0x99, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0x86, Next: 1},
			{Lo: 0x93, Hi: 0x97, Next: 1},
			{Lo: 0x9d, Hi: 0x9d, Next: 1},
			{Lo: 0x9f, Hi: 0xa8, Next: 1},
			{Lo: 0xaa, Hi: 0xb6, Next: 1},
			{Lo: 0xb8, Hi: 0xbc, Next: 1},
			{Lo: 0xbe, Hi: 0xbe, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0x81, Next: 1},
			{Lo: 0x83, Hi: 0x84, Next: 1},
			{Lo: 0x86, Hi: 0xbf, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0xb1, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x93, Hi: 0xbf, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0xbd, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x90, Hi: 0xbf, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0x8f, Next: 1},
			{Lo: 0x92, Hi: 0xbf, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0x87, Next: 1},
			{Lo: 0xb0, Hi: 0xbb, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0xb0, Hi: 0xb4, Next: 1},
			{Lo: 0xb6, Hi: 0xbf, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0xbc, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0xa1, Hi: 0xba, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x81, Hi: 0x9a, Next: 1},
			{Lo: 0xa6, Hi: 0xbf, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0xbe, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x82, Hi: 0x87, Next: 1},
			{Lo: 0x8a, Hi: 0x8f, Next: 1},
			{Lo: 0x92, Hi: 0x97, Next: 1},
			{Lo: 0x9a, Hi: 0x9c, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0x80, Next: 169},
			{Lo: 0x81, Hi: 0x81, Next: 170},
			{Lo: 0x82, Hi: 0x82, Next: 4},
			{Lo: 0x83, Hi: 0x83, Next: 171},
			{Lo: 0x8a, Hi: 0x8a, Next: 172},
			{Lo: 0x8b, Hi: 0x8b, Next: 173},
			{Lo: 0x8c, Hi: 0x8c, Next: 174},
			{Lo: 0x8d, Hi: 0x8d, Next: 175},
			{Lo: 0x8e, Hi: 0x8e, Next: 176},
			{Lo: 0x8f, Hi: 0x8f, Next: 177},
			{Lo: 0x90, Hi: 0x91, Next: 4},
			{Lo: 0x92, Hi: 0x92, Next: 178},
			{Lo: 0xa0, Hi: 0xa0, Next: 179},
			{Lo: 0xa1, Hi: 0xa1, Next: 180},
			{Lo: 0xa4, Hi: 0xa4, Next: 181},
			{Lo: 0xa6, Hi: 0xa6, Next: 182},
			{Lo: 0xa8, Hi: 0xa8, Next: 183},
			{Lo: 0xa9, Hi: 0xa9, Next: 184},
			{Lo: 0xac, Hi: 0xac, Next: 81},
			{Lo: 0xad, Hi: 0xad, Next: 185},
			{Lo: 0xb0, Hi: 0xb0, Next: 4},
			{Lo: 0xb1, Hi: 0xb1, Next: 186}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0x80, Next: 187},
			{Lo: 0x82, Hi: 0x82, Next: 188},
			{Lo: 0x83, Hi: 0x83, Next: 189},
			{Lo: 0x84, Hi: 0x84, Next: 190},
			{Lo: 0x86, Hi: 0x86, Next: 191},
			{Lo: 0x87, Hi: 0x87, Next: 192},
			{Lo: 0x9a, Hi: 0x9a, Next: 74}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0x8c, Next: 4},
			{Lo: 0x8d, Hi: 0x8d, Next: 193}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0x8f, Next: 4},
			{Lo: 0x90, Hi: 0x90, Next: 193}}, Label: 0},
		{Table: TransTable{
			{Lo: 0xa0, Hi: 0xa7, Next: 4},
			{Lo: 0xa8, Hi: 0xa8, Next: 194},
			{Lo: 0xbc, Hi: 0xbc, Next: 4},
			{Lo: 0xbd, Hi: 0xbd, Next: 195},
			{Lo: 0xbe, Hi: 0xbe, Next: 196}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0x80, Next: 197}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x90, Hi: 0x90, Next: 4},
			{Lo: 0x91, Hi: 0x91, Next: 198},
			{Lo: 0x92, Hi: 0x92, Next: 199},
			{Lo: 0x93, Hi: 0x93, Next: 200},
			{Lo: 0x94, Hi: 0x94, Next: 201},
			{Lo: 0x95, Hi: 0x95, Next: 202},
			{Lo: 0x96, Hi: 0x99, Next: 4},
			{Lo: 0x9a, Hi: 0x9a, Next: 203},
			{Lo: 0x9b, Hi: 0x9b, Next: 204},
			{Lo: 0x9c, Hi: 0x9c, Next: 205},
			{Lo: 0x9d, Hi: 0x9d, Next: 206},
			{Lo: 0x9e, Hi: 0x9e, Next: 207},
			{Lo: 0x9f, Hi: 0x9f, Next: 208}}, Label: 0},
		{Table: TransTable{
			{Lo: 0xb8, Hi: 0xb8, Next: 209},
			{Lo: 0xb9, Hi: 0xb9, Next: 210},
			{Lo: 0xba, Hi: 0xba, Next: 211}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0x9a, Next: 4},
			{Lo: 0x9b, Hi: 0x9b, Next: 212},
			{Lo: 0x9c, Hi: 0xbf, Next: 4}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0x9b, Next: 4},
			{Lo: 0x9c, Hi: 0x9c, Next: 70},
			{Lo: 0x9d, Hi: 0x9f, Next: 4},
			{Lo: 0xa0, Hi: 0xa0, Next: 178}}, Label: 0},
		{Table: TransTable{
			{Lo: 0xa0, Hi: 0xa7, Next: 4},
			{Lo: 0xa8, Hi: 0xa8, Next: 178}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0x8b, Next: 1},
			{Lo: 0x8d, Hi: 0xa6, Next: 1},
			{Lo: 0xa8, Hi: 0xba, Next: 1},
			{Lo: 0xbc, Hi: 0xbd, Next: 1},
			{Lo: 0xbf, Hi: 0xbf, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0x8d, Next: 1},
			{Lo: 0x90, Hi: 0x9d, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0xba, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0x9c, Next: 1},
			{Lo: 0xa0, Hi: 0xbf, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0x90, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0x9e, Next: 1},
			{Lo: 0xb0, Hi: 0xbf, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0x80, Next: 1},
			{Lo: 0x82, Hi: 0x89, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0x9d, Next: 1},
			{Lo: 0xa0, Hi: 0xbf, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0x83, Next: 1},
			{Lo: 0x88, Hi: 0x8f, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0x9d, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0x85, Next: 1},
			{Lo: 0x88, Hi: 0x88, Next: 1},
			{Lo: 0x8a, Hi: 0xb5, Next: 1},
			{Lo: 0xb7, Hi: 0xb8, Next: 1},
			{Lo: 0xbc, Hi: 0xbc, Next: 1},
			{Lo: 0xbf, Hi: 0xbf, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0x95, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0x95, Next: 1},
			{Lo: 0xa0, Hi: 0xb9, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0xb7, Next: 1},
			{Lo: 0xbe, Hi: 0xbf, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0x80, Next: 1},
			{Lo: 0x90, Hi: 0x93, Next: 1},
			{Lo: 0x95, Hi: 0x97, Next: 1},
			{Lo: 0x99, Hi: 0xb3, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0xa0, Hi: 0xbc, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0x95, Next: 1},
			{Lo: 0xa0, Hi: 0xb2, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0x88, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x83, Hi: 0xb7, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x83, Hi: 0xaf, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x90, Hi: 0xa8, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x83, Hi: 0xa6, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x83, Hi: 0xb2, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x81, Hi: 0x84, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0xae, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0xb8, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0x84, Next: 1},
			{Lo: 0x90, Hi: 0x90, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x93, Hi: 0x9f, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0x81, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0x94, Next: 1},
			{Lo: 0x96, Hi: 0xbf, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0x9c, Next: 1},
			{Lo: 0x9e, Hi: 0x9f, Next: 1},
			{Lo: 0xa2, Hi: 0xa2, Next: 1},
			{Lo: 0xa5, Hi: 0xa6, Next: 1},
			{Lo: 0xa9, Hi: 0xac, Next: 1},
			{Lo: 0xae, Hi: 0xb9, Next: 1},
			{Lo: 0xbb, Hi: 0xbb, Next: 1},
			{Lo: 0xbd, Hi: 0xbf, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0x83, Next: 1},
			{Lo: 0x85, Hi: 0xbf, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0x85, Next: 1},
			{Lo: 0x87, Hi: 0x8a, Next: 1},
			{Lo: 0x8d, Hi: 0x94, Next: 1},
			{Lo: 0x96, Hi: 0x9c, Next: 1},
			{Lo: 0x9e, Hi: 0xb9, Next: 1},
			{Lo: 0xbb, Hi: 0xbe, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0x84, Next: 1},
			{Lo: 0x86, Hi: 0x86, Next: 1},
			{Lo: 0x8a, Hi: 0x90, Next: 1},
			{Lo: 0x92, Hi: 0xbf, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0xa5, Next: 1},
			{Lo: 0xa8, Hi: 0xbf, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0x80, Next: 1},
			{Lo: 0x82, Hi: 0x9a, Next: 1},
			{Lo: 0x9c, Hi: 0xba, Next: 1},
			{Lo: 0xbc, Hi: 0xbf, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0x94, Next: 1},
			{Lo: 0x96, Hi: 0xb4, Next: 1},
			{Lo: 0xb6, Hi: 0xbf, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0x8e, Next: 1},
			{Lo: 0x90, Hi: 0xae, Next: 1},
			{Lo: 0xb0, Hi: 0xbf, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0x88, Next: 1},
			{Lo: 0x8a, Hi: 0xa8, Next: 1},
			{Lo: 0xaa, Hi: 0xbf, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0x82, Next: 1},
			{Lo: 0x84, Hi: 0x8b, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0x83, Next: 1},
			{Lo: 0x85, Hi: 0x9f, Next: 1},
			{Lo: 0xa1, Hi: 0xa2, Next: 1},
			{Lo: 0xa4, Hi: 0xa4, Next: 1},
			{Lo: 0xa7, Hi: 0xa7, Next: 1},
			{Lo: 0xa9, Hi: 0xb2, Next: 1},
			{Lo: 0xb4, Hi: 0xb7, Next: 1},
			{Lo: 0xb9, Hi: 0xb9, Next: 1},
			{Lo: 0xbb, Hi: 0xbb, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x82, Hi: 0x82, Next: 1},
			{Lo: 0x87, Hi: 0x87, Next: 1},
			{Lo: 0x89, Hi: 0x89, Next: 1},
			{Lo: 0x8b, Hi: 0x8b, Next: 1},
			{Lo: 0x8d, Hi: 0x8f, Next: 1},
			{Lo: 0x91, Hi: 0x92, Next: 1},
			{Lo: 0x94, Hi: 0x94, Next: 1},
			{Lo: 0x97, Hi: 0x97, Next: 1},
			{Lo: 0x99, Hi: 0x99, Next: 1},
			{Lo: 0x9b, Hi: 0x9b, Next: 1},
			{Lo: 0x9d, Hi: 0x9d, Next: 1},
			{Lo: 0x9f, Hi: 0x9f, Next: 1},
			{Lo: 0xa1, Hi: 0xa2, Next: 1},
			{Lo: 0xa4, Hi: 0xa4, Next: 1},
			{Lo: 0xa7, Hi: 0xaa, Next: 1},
			{Lo: 0xac, Hi: 0xb2, Next: 1},
			{Lo: 0xb4, Hi: 0xb7, Next: 1},
			{Lo: 0xb9, Hi: 0xbc, Next: 1},
			{Lo: 0xbe, Hi: 0xbe, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0x89, Next: 1},
			{Lo: 0x8b, Hi: 0x9b, Next: 1},
			{Lo: 0xa1, Hi: 0xa3, Next: 1},
			{Lo: 0xa5, Hi: 0xa9, Next: 1},
			{Lo: 0xab, Hi: 0xbb, Next: 1}}, Label: 0},
		{Table: TransTable{
			{Lo: 0x80, Hi: 0x96, Next: 1}}, Label: 0},
	},
}
