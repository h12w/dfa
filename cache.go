package dfa

var classL = &M{
	States: States{
		{
			Table: TransTable{
				{0x41, 0x5a, 1},
				{0x61, 0x7a, 1},
				{0xc2, 0xc2, 2},
				{0xc3, 0xc3, 3},
				{0xc4, 0xca, 4},
				{0xcb, 0xcb, 5},
				{0xcd, 0xcd, 6},
				{0xce, 0xce, 7},
				{0xcf, 0xcf, 8},
				{0xd0, 0xd1, 4},
				{0xd2, 0xd2, 9},
				{0xd3, 0xd3, 4},
				{0xd4, 0xd4, 10},
				{0xd5, 0xd5, 11},
				{0xd6, 0xd6, 12},
				{0xd7, 0xd7, 13},
				{0xd8, 0xd8, 14},
				{0xd9, 0xd9, 15},
				{0xda, 0xda, 4},
				{0xdb, 0xdb, 16},
				{0xdc, 0xdc, 17},
				{0xdd, 0xdd, 18},
				{0xde, 0xde, 19},
				{0xdf, 0xdf, 20},
				{0xe0, 0xe0, 21},
				{0xe1, 0xe1, 22},
				{0xe2, 0xe2, 23},
				{0xe3, 0xe3, 24},
				{0xe4, 0xe4, 25},
				{0xe5, 0xe8, 26},
				{0xe9, 0xe9, 27},
				{0xea, 0xea, 28},
				{0xeb, 0xec, 26},
				{0xed, 0xed, 29},
				{0xef, 0xef, 30},
				{0xf0, 0xf0, 31},
			}},
		{
			Label: 1,
		},
		{
			Table: TransTable{
				{0xaa, 0xaa, 1},
				{0xb5, 0xb5, 1},
				{0xba, 0xba, 1},
			}},
		{
			Table: TransTable{
				{0x80, 0x96, 1},
				{0x98, 0xb6, 1},
				{0xb8, 0xbf, 1},
			}},
		{
			Table: TransTable{{0x80, 0xbf, 1}}},
		{
			Table: TransTable{
				{0x80, 0x81, 1},
				{0x86, 0x91, 1},
				{0xa0, 0xa4, 1},
				{0xac, 0xac, 1},
				{0xae, 0xae, 1},
			}},
		{
			Table: TransTable{
				{0xb0, 0xb4, 1},
				{0xb6, 0xb7, 1},
				{0xba, 0xbd, 1},
			}},
		{
			Table: TransTable{
				{0x86, 0x86, 1},
				{0x88, 0x8a, 1},
				{0x8c, 0x8c, 1},
				{0x8e, 0xa1, 1},
				{0xa3, 0xbf, 1},
			}},
		{
			Table: TransTable{
				{0x80, 0xb5, 1},
				{0xb7, 0xbf, 1},
			}},
		{
			Table: TransTable{
				{0x80, 0x81, 1},
				{0x8a, 0xbf, 1},
			}},
		{
			Table: TransTable{
				{0x80, 0xa7, 1},
				{0xb1, 0xbf, 1},
			}},
		{
			Table: TransTable{
				{0x80, 0x96, 1},
				{0x99, 0x99, 1},
				{0xa1, 0xbf, 1},
			}},
		{
			Table: TransTable{{0x80, 0x87, 1}}},
		{
			Table: TransTable{
				{0x90, 0xaa, 1},
				{0xb0, 0xb2, 1},
			}},
		{
			Table: TransTable{{0xa0, 0xbf, 1}}},
		{
			Table: TransTable{
				{0x80, 0x8a, 1},
				{0xae, 0xaf, 1},
				{0xb1, 0xbf, 1},
			}},
		{
			Table: TransTable{
				{0x80, 0x93, 1},
				{0x95, 0x95, 1},
				{0xa5, 0xa6, 1},
				{0xae, 0xaf, 1},
				{0xba, 0xbc, 1},
				{0xbf, 0xbf, 1},
			}},
		{
			Table: TransTable{
				{0x90, 0x90, 1},
				{0x92, 0xaf, 1},
			}},
		{
			Table: TransTable{{0x8d, 0xbf, 1}}},
		{
			Table: TransTable{
				{0x80, 0xa5, 1},
				{0xb1, 0xb1, 1},
			}},
		{
			Table: TransTable{
				{0x8a, 0xaa, 1},
				{0xb4, 0xb5, 1},
				{0xba, 0xba, 1},
			}},
		{
			Table: TransTable{
				{0xa0, 0xa0, 32},
				{0xa1, 0xa1, 33},
				{0xa2, 0xa2, 34},
				{0xa4, 0xa4, 35},
				{0xa5, 0xa5, 36},
				{0xa6, 0xa6, 37},
				{0xa7, 0xa7, 38},
				{0xa8, 0xa8, 39},
				{0xa9, 0xa9, 40},
				{0xaa, 0xaa, 41},
				{0xab, 0xab, 42},
				{0xac, 0xac, 43},
				{0xad, 0xad, 44},
				{0xae, 0xae, 45},
				{0xaf, 0xaf, 46},
				{0xb0, 0xb0, 47},
				{0xb1, 0xb1, 48},
				{0xb2, 0xb2, 47},
				{0xb3, 0xb3, 49},
				{0xb4, 0xb4, 50},
				{0xb5, 0xb5, 51},
				{0xb6, 0xb6, 52},
				{0xb7, 0xb7, 53},
				{0xb8, 0xb8, 54},
				{0xb9, 0xb9, 55},
				{0xba, 0xba, 56},
				{0xbb, 0xbb, 57},
				{0xbc, 0xbc, 58},
				{0xbd, 0xbd, 59},
				{0xbe, 0xbe, 60},
			}},
		{
			Table: TransTable{
				{0x80, 0x80, 61},
				{0x81, 0x81, 62},
				{0x82, 0x82, 63},
				{0x83, 0x83, 64},
				{0x84, 0x88, 4},
				{0x89, 0x89, 65},
				{0x8a, 0x8a, 66},
				{0x8b, 0x8b, 67},
				{0x8c, 0x8c, 68},
				{0x8d, 0x8d, 69},
				{0x8e, 0x8e, 70},
				{0x8f, 0x8f, 71},
				{0x90, 0x90, 72},
				{0x91, 0x98, 4},
				{0x99, 0x99, 73},
				{0x9a, 0x9a, 74},
				{0x9b, 0x9b, 75},
				{0x9c, 0x9c, 76},
				{0x9d, 0x9d, 77},
				{0x9e, 0x9e, 78},
				{0x9f, 0x9f, 79},
				{0xa0, 0xa0, 14},
				{0xa1, 0xa1, 80},
				{0xa2, 0xa2, 81},
				{0xa3, 0xa3, 82},
				{0xa4, 0xa4, 83},
				{0xa5, 0xa5, 84},
				{0xa6, 0xa6, 85},
				{0xa7, 0xa7, 86},
				{0xa8, 0xa8, 87},
				{0xa9, 0xa9, 88},
				{0xaa, 0xaa, 89},
				{0xac, 0xac, 90},
				{0xad, 0xad, 91},
				{0xae, 0xae, 92},
				{0xaf, 0xaf, 93},
				{0xb0, 0xb0, 94},
				{0xb1, 0xb1, 95},
				{0xb3, 0xb3, 96},
				{0xb4, 0xb6, 4},
				{0xb8, 0xbb, 4},
				{0xbc, 0xbc, 97},
				{0xbd, 0xbd, 98},
				{0xbe, 0xbe, 99},
				{0xbf, 0xbf, 100},
			}},
		{
			Table: TransTable{
				{0x81, 0x81, 101},
				{0x82, 0x82, 102},
				{0x84, 0x84, 103},
				{0x85, 0x85, 104},
				{0x86, 0x86, 105},
				{0xb0, 0xb0, 106},
				{0xb1, 0xb1, 107},
				{0xb2, 0xb2, 4},
				{0xb3, 0xb3, 108},
				{0xb4, 0xb4, 109},
				{0xb5, 0xb5, 110},
				{0xb6, 0xb6, 111},
				{0xb7, 0xb7, 112},
				{0xb8, 0xb8, 113},
			}},
		{
			Table: TransTable{
				{0x80, 0x80, 114},
				{0x81, 0x81, 72},
				{0x82, 0x82, 115},
				{0x83, 0x83, 116},
				{0x84, 0x84, 117},
				{0x85, 0x85, 4},
				{0x86, 0x86, 118},
				{0x87, 0x87, 119},
				{0x90, 0xbf, 4},
			}},
		{
			Table: TransTable{
				{0x80, 0xb5, 4},
				{0xb6, 0xb6, 82},
				{0xb8, 0xbf, 4},
			}},
		{
			Table: TransTable{{0x80, 0xbf, 4}}},
		{
			Table: TransTable{
				{0x80, 0xbe, 4},
				{0xbf, 0xbf, 120},
			}},
		{
			Table: TransTable{
				{0x80, 0x91, 4},
				{0x92, 0x92, 120},
				{0x93, 0x93, 121},
				{0x94, 0x97, 4},
				{0x98, 0x98, 122},
				{0x99, 0x99, 123},
				{0x9a, 0x9a, 124},
				{0x9b, 0x9b, 93},
				{0x9c, 0x9c, 125},
				{0x9d, 0x9d, 4},
				{0x9e, 0x9e, 126},
				{0x9f, 0x9f, 127},
				{0xa0, 0xa0, 128},
				{0xa1, 0xa1, 78},
				{0xa2, 0xa2, 129},
				{0xa3, 0xa3, 130},
				{0xa4, 0xa4, 131},
				{0xa5, 0xa5, 132},
				{0xa6, 0xa6, 133},
				{0xa7, 0xa7, 134},
				{0xa8, 0xa8, 135},
				{0xa9, 0xa9, 136},
				{0xaa, 0xaa, 137},
				{0xab, 0xab, 138},
				{0xac, 0xac, 139},
				{0xaf, 0xaf, 140},
				{0xb0, 0xbf, 4},
			}},
		{
			Table: TransTable{
				{0x80, 0x9d, 4},
				{0x9e, 0x9e, 141},
				{0x9f, 0x9f, 142},
			}},
		{
			Table: TransTable{
				{0xa4, 0xa8, 4},
				{0xa9, 0xa9, 143},
				{0xaa, 0xaa, 4},
				{0xab, 0xab, 144},
				{0xac, 0xac, 145},
				{0xad, 0xad, 146},
				{0xae, 0xae, 147},
				{0xaf, 0xaf, 148},
				{0xb0, 0xb3, 4},
				{0xb4, 0xb4, 149},
				{0xb5, 0xb5, 150},
				{0xb6, 0xb6, 151},
				{0xb7, 0xb7, 152},
				{0xb9, 0xb9, 153},
				{0xba, 0xba, 4},
				{0xbb, 0xbb, 154},
				{0xbc, 0xbc, 155},
				{0xbd, 0xbd, 156},
				{0xbe, 0xbe, 157},
				{0xbf, 0xbf, 158},
			}},
		{
			Table: TransTable{
				{0x90, 0x90, 159},
				{0x91, 0x91, 160},
				{0x92, 0x92, 161},
				{0x93, 0x93, 162},
				{0x96, 0x96, 163},
				{0x9b, 0x9b, 164},
				{0x9d, 0x9d, 165},
				{0x9e, 0x9e, 166},
				{0xa0, 0xa9, 26},
				{0xaa, 0xaa, 167},
				{0xab, 0xab, 168},
				{0xaf, 0xaf, 169},
			}},
		{
			Table: TransTable{
				{0x80, 0x95, 1},
				{0x9a, 0x9a, 1},
				{0xa4, 0xa4, 1},
				{0xa8, 0xa8, 1},
			}},
		{
			Table: TransTable{{0x80, 0x98, 1}}},
		{
			Table: TransTable{
				{0xa0, 0xa0, 1},
				{0xa2, 0xac, 1},
			}},
		{
			Table: TransTable{
				{0x84, 0xb9, 1},
				{0xbd, 0xbd, 1},
			}},
		{
			Table: TransTable{
				{0x90, 0x90, 1},
				{0x98, 0xa1, 1},
				{0xb1, 0xb7, 1},
				{0xb9, 0xbf, 1},
			}},
		{
			Table: TransTable{
				{0x85, 0x8c, 1},
				{0x8f, 0x90, 1},
				{0x93, 0xa8, 1},
				{0xaa, 0xb0, 1},
				{0xb2, 0xb2, 1},
				{0xb6, 0xb9, 1},
				{0xbd, 0xbd, 1},
			}},
		{
			Table: TransTable{
				{0x8e, 0x8e, 1},
				{0x9c, 0x9d, 1},
				{0x9f, 0xa1, 1},
				{0xb0, 0xb1, 1},
			}},
		{
			Table: TransTable{
				{0x85, 0x8a, 1},
				{0x8f, 0x90, 1},
				{0x93, 0xa8, 1},
				{0xaa, 0xb0, 1},
				{0xb2, 0xb3, 1},
				{0xb5, 0xb6, 1},
				{0xb8, 0xb9, 1},
			}},
		{
			Table: TransTable{
				{0x99, 0x9c, 1},
				{0x9e, 0x9e, 1},
				{0xb2, 0xb4, 1},
			}},
		{
			Table: TransTable{
				{0x85, 0x8d, 1},
				{0x8f, 0x91, 1},
				{0x93, 0xa8, 1},
				{0xaa, 0xb0, 1},
				{0xb2, 0xb3, 1},
				{0xb5, 0xb9, 1},
				{0xbd, 0xbd, 1},
			}},
		{
			Table: TransTable{
				{0x90, 0x90, 1},
				{0xa0, 0xa1, 1},
			}},
		{
			Table: TransTable{
				{0x85, 0x8c, 1},
				{0x8f, 0x90, 1},
				{0x93, 0xa8, 1},
				{0xaa, 0xb0, 1},
				{0xb2, 0xb3, 1},
				{0xb5, 0xb9, 1},
				{0xbd, 0xbd, 1},
			}},
		{
			Table: TransTable{
				{0x9c, 0x9d, 1},
				{0x9f, 0xa1, 1},
				{0xb1, 0xb1, 1},
			}},
		{
			Table: TransTable{
				{0x83, 0x83, 1},
				{0x85, 0x8a, 1},
				{0x8e, 0x90, 1},
				{0x92, 0x95, 1},
				{0x99, 0x9a, 1},
				{0x9c, 0x9c, 1},
				{0x9e, 0x9f, 1},
				{0xa3, 0xa4, 1},
				{0xa8, 0xaa, 1},
				{0xae, 0xb9, 1},
			}},
		{
			Table: TransTable{{0x90, 0x90, 1}}},
		{
			Table: TransTable{
				{0x85, 0x8c, 1},
				{0x8e, 0x90, 1},
				{0x92, 0xa8, 1},
				{0xaa, 0xb3, 1},
				{0xb5, 0xb9, 1},
				{0xbd, 0xbd, 1},
			}},
		{
			Table: TransTable{
				{0x98, 0x99, 1},
				{0xa0, 0xa1, 1},
			}},
		{
			Table: TransTable{
				{0x9e, 0x9e, 1},
				{0xa0, 0xa1, 1},
				{0xb1, 0xb2, 1},
			}},
		{
			Table: TransTable{
				{0x85, 0x8c, 1},
				{0x8e, 0x90, 1},
				{0x92, 0xba, 1},
				{0xbd, 0xbd, 1},
			}},
		{
			Table: TransTable{
				{0x8e, 0x8e, 1},
				{0xa0, 0xa1, 1},
				{0xba, 0xbf, 1},
			}},
		{
			Table: TransTable{
				{0x85, 0x96, 1},
				{0x9a, 0xb1, 1},
				{0xb3, 0xbb, 1},
				{0xbd, 0xbd, 1},
			}},
		{
			Table: TransTable{{0x80, 0x86, 1}}},
		{
			Table: TransTable{
				{0x81, 0xb0, 1},
				{0xb2, 0xb3, 1},
			}},
		{
			Table: TransTable{{0x80, 0x86, 1}}},
		{
			Table: TransTable{
				{0x81, 0x82, 1},
				{0x84, 0x84, 1},
				{0x87, 0x88, 1},
				{0x8a, 0x8a, 1},
				{0x8d, 0x8d, 1},
				{0x94, 0x97, 1},
				{0x99, 0x9f, 1},
				{0xa1, 0xa3, 1},
				{0xa5, 0xa5, 1},
				{0xa7, 0xa7, 1},
				{0xaa, 0xab, 1},
				{0xad, 0xb0, 1},
				{0xb2, 0xb3, 1},
				{0xbd, 0xbd, 1},
			}},
		{
			Table: TransTable{
				{0x80, 0x84, 1},
				{0x86, 0x86, 1},
				{0x9c, 0x9f, 1},
			}},
		{
			Table: TransTable{{0x80, 0x80, 1}}},
		{
			Table: TransTable{
				{0x80, 0x87, 1},
				{0x89, 0xac, 1},
			}},
		{
			Table: TransTable{{0x88, 0x8c, 1}}},
		{
			Table: TransTable{
				{0x80, 0xaa, 1},
				{0xbf, 0xbf, 1},
			}},
		{
			Table: TransTable{
				{0x90, 0x95, 1},
				{0x9a, 0x9d, 1},
				{0xa1, 0xa1, 1},
				{0xa5, 0xa6, 1},
				{0xae, 0xb0, 1},
				{0xb5, 0xbf, 1},
			}},
		{
			Table: TransTable{
				{0x80, 0x81, 1},
				{0x8e, 0x8e, 1},
				{0xa0, 0xbf, 1},
			}},
		{
			Table: TransTable{
				{0x80, 0x85, 1},
				{0x87, 0x87, 1},
				{0x8d, 0x8d, 1},
				{0x90, 0xba, 1},
				{0xbc, 0xbf, 1},
			}},
		{
			Table: TransTable{
				{0x80, 0x88, 1},
				{0x8a, 0x8d, 1},
				{0x90, 0x96, 1},
				{0x98, 0x98, 1},
				{0x9a, 0x9d, 1},
				{0xa0, 0xbf, 1},
			}},
		{
			Table: TransTable{
				{0x80, 0x88, 1},
				{0x8a, 0x8d, 1},
				{0x90, 0xb0, 1},
				{0xb2, 0xb5, 1},
				{0xb8, 0xbe, 1},
			}},
		{
			Table: TransTable{
				{0x80, 0x80, 1},
				{0x82, 0x85, 1},
				{0x88, 0x96, 1},
				{0x98, 0xbf, 1},
			}},
		{
			Table: TransTable{
				{0x80, 0x90, 1},
				{0x92, 0x95, 1},
				{0x98, 0xbf, 1},
			}},
		{
			Table: TransTable{{0x80, 0x9a, 1}}},
		{
			Table: TransTable{
				{0x80, 0x8f, 1},
				{0xa0, 0xbf, 1},
			}},
		{
			Table: TransTable{{0x80, 0xb4, 1}}},
		{
			Table: TransTable{{0x81, 0xbf, 1}}},
		{
			Table: TransTable{
				{0x80, 0xac, 1},
				{0xaf, 0xbf, 1},
			}},
		{
			Table: TransTable{
				{0x81, 0x9a, 1},
				{0xa0, 0xbf, 1},
			}},
		{
			Table: TransTable{{0x80, 0xaa, 1}}},
		{
			Table: TransTable{
				{0x80, 0x8c, 1},
				{0x8e, 0x91, 1},
				{0xa0, 0xb1, 1},
			}},
		{
			Table: TransTable{
				{0x80, 0x91, 1},
				{0xa0, 0xac, 1},
				{0xae, 0xb0, 1},
			}},
		{
			Table: TransTable{{0x80, 0xb3, 1}}},
		{
			Table: TransTable{
				{0x97, 0x97, 1},
				{0x9c, 0x9c, 1},
			}},
		{
			Table: TransTable{{0x80, 0xb7, 1}}},
		{
			Table: TransTable{
				{0x80, 0xa8, 1},
				{0xaa, 0xaa, 1},
				{0xb0, 0xbf, 1},
			}},
		{
			Table: TransTable{{0x80, 0xb5, 1}}},
		{
			Table: TransTable{{0x80, 0x9c, 1}}},
		{
			Table: TransTable{
				{0x90, 0xad, 1},
				{0xb0, 0xb4, 1},
			}},
		{
			Table: TransTable{{0x80, 0xab, 1}}},
		{
			Table: TransTable{{0x81, 0x87, 1}}},
		{
			Table: TransTable{
				{0x80, 0x96, 1},
				{0xa0, 0xbf, 1},
			}},
		{
			Table: TransTable{{0x80, 0x94, 1}}},
		{
			Table: TransTable{{0xa7, 0xa7, 1}}},
		{
			Table: TransTable{{0x85, 0xb3, 1}}},
		{
			Table: TransTable{{0x85, 0x8b, 1}}},
		{
			Table: TransTable{
				{0x83, 0xa0, 1},
				{0xae, 0xaf, 1},
				{0xba, 0xbf, 1},
			}},
		{
			Table: TransTable{{0x80, 0xa5, 1}}},
		{
			Table: TransTable{{0x80, 0xa3, 1}}},
		{
			Table: TransTable{
				{0x8d, 0x8f, 1},
				{0x9a, 0xbd, 1},
			}},
		{
			Table: TransTable{
				{0xa9, 0xac, 1},
				{0xae, 0xb1, 1},
				{0xb5, 0xb6, 1},
			}},
		{
			Table: TransTable{
				{0x80, 0x95, 1},
				{0x98, 0x9d, 1},
				{0xa0, 0xbf, 1},
			}},
		{
			Table: TransTable{
				{0x80, 0x85, 1},
				{0x88, 0x8d, 1},
				{0x90, 0x97, 1},
				{0x99, 0x99, 1},
				{0x9b, 0x9b, 1},
				{0x9d, 0x9d, 1},
				{0x9f, 0xbd, 1},
			}},
		{
			Table: TransTable{
				{0x80, 0xb4, 1},
				{0xb6, 0xbc, 1},
				{0xbe, 0xbe, 1},
			}},
		{
			Table: TransTable{
				{0x82, 0x84, 1},
				{0x86, 0x8c, 1},
				{0x90, 0x93, 1},
				{0x96, 0x9b, 1},
				{0xa0, 0xac, 1},
				{0xb2, 0xb4, 1},
				{0xb6, 0xbc, 1},
			}},
		{
			Table: TransTable{
				{0xb1, 0xb1, 1},
				{0xbf, 0xbf, 1},
			}},
		{
			Table: TransTable{{0x90, 0x9c, 1}}},
		{
			Table: TransTable{
				{0x82, 0x82, 1},
				{0x87, 0x87, 1},
				{0x8a, 0x93, 1},
				{0x95, 0x95, 1},
				{0x99, 0x9d, 1},
				{0xa4, 0xa4, 1},
				{0xa6, 0xa6, 1},
				{0xa8, 0xa8, 1},
				{0xaa, 0xad, 1},
				{0xaf, 0xb9, 1},
				{0xbc, 0xbf, 1},
			}},
		{
			Table: TransTable{
				{0x85, 0x89, 1},
				{0x8e, 0x8e, 1},
			}},
		{
			Table: TransTable{{0x83, 0x84, 1}}},
		{
			Table: TransTable{
				{0x80, 0xae, 1},
				{0xb0, 0xbf, 1},
			}},
		{
			Table: TransTable{
				{0x80, 0x9e, 1},
				{0xa0, 0xbf, 1},
			}},
		{
			Table: TransTable{
				{0x80, 0xa4, 1},
				{0xab, 0xae, 1},
				{0xb2, 0xb3, 1},
			}},
		{
			Table: TransTable{
				{0x80, 0xa5, 1},
				{0xa7, 0xa7, 1},
				{0xad, 0xad, 1},
				{0xb0, 0xbf, 1},
			}},
		{
			Table: TransTable{
				{0x80, 0xa7, 1},
				{0xaf, 0xaf, 1},
			}},
		{
			Table: TransTable{
				{0x80, 0x96, 1},
				{0xa0, 0xa6, 1},
				{0xa8, 0xae, 1},
				{0xb0, 0xb6, 1},
				{0xb8, 0xbe, 1},
			}},
		{
			Table: TransTable{
				{0x80, 0x86, 1},
				{0x88, 0x8e, 1},
				{0x90, 0x96, 1},
				{0x98, 0x9e, 1},
			}},
		{
			Table: TransTable{{0xaf, 0xaf, 1}}},
		{
			Table: TransTable{
				{0x85, 0x86, 1},
				{0xb1, 0xb5, 1},
				{0xbb, 0xbc, 1},
			}},
		{
			Table: TransTable{
				{0x80, 0x96, 1},
				{0x9d, 0x9f, 1},
				{0xa1, 0xbf, 1},
			}},
		{
			Table: TransTable{
				{0x80, 0xba, 1},
				{0xbc, 0xbf, 1},
			}},
		{
			Table: TransTable{
				{0x85, 0xad, 1},
				{0xb1, 0xbf, 1},
			}},
		{
			Table: TransTable{
				{0x80, 0x8e, 1},
				{0xa0, 0xba, 1},
			}},
		{
			Table: TransTable{{0xb0, 0xbf, 1}}},
		{
			Table: TransTable{{0x80, 0x8c, 1}}},
		{
			Table: TransTable{{0x90, 0xbd, 1}}},
		{
			Table: TransTable{
				{0x80, 0x8c, 1},
				{0x90, 0x9f, 1},
				{0xaa, 0xab, 1},
			}},
		{
			Table: TransTable{
				{0x80, 0xae, 1},
				{0xbf, 0xbf, 1},
			}},
		{
			Table: TransTable{
				{0x80, 0x97, 1},
				{0xa0, 0xbf, 1},
			}},
		{
			Table: TransTable{
				{0x97, 0x9f, 1},
				{0xa2, 0xbf, 1},
			}},
		{
			Table: TransTable{
				{0x80, 0x88, 1},
				{0x8b, 0x8e, 1},
				{0x90, 0x93, 1},
				{0xa0, 0xaa, 1},
			}},
		{
			Table: TransTable{{0xb8, 0xbf, 1}}},
		{
			Table: TransTable{
				{0x80, 0x81, 1},
				{0x83, 0x85, 1},
				{0x87, 0x8a, 1},
				{0x8c, 0xa2, 1},
			}},
		{
			Table: TransTable{{0x82, 0xb3, 1}}},
		{
			Table: TransTable{
				{0xb2, 0xb7, 1},
				{0xbb, 0xbb, 1},
			}},
		{
			Table: TransTable{
				{0x8a, 0xa5, 1},
				{0xb0, 0xbf, 1},
			}},
		{
			Table: TransTable{
				{0x80, 0x86, 1},
				{0xa0, 0xbc, 1},
			}},
		{
			Table: TransTable{{0x84, 0xb2, 1}}},
		{
			Table: TransTable{{0x8f, 0x8f, 1}}},
		{
			Table: TransTable{{0x80, 0xa8, 1}}},
		{
			Table: TransTable{
				{0x80, 0x82, 1},
				{0x84, 0x8b, 1},
				{0xa0, 0xb6, 1},
				{0xba, 0xba, 1},
			}},
		{
			Table: TransTable{
				{0x80, 0xaf, 1},
				{0xb1, 0xb1, 1},
				{0xb5, 0xb6, 1},
				{0xb9, 0xbd, 1},
			}},
		{
			Table: TransTable{
				{0x80, 0x80, 1},
				{0x82, 0x82, 1},
				{0x9b, 0x9d, 1},
				{0xa0, 0xaa, 1},
				{0xb2, 0xb4, 1},
			}},
		{
			Table: TransTable{
				{0x81, 0x86, 1},
				{0x89, 0x8e, 1},
				{0x91, 0x96, 1},
				{0xa0, 0xa6, 1},
				{0xa8, 0xae, 1},
			}},
		{
			Table: TransTable{{0x80, 0xa2, 1}}},
		{
			Table: TransTable{
				{0x80, 0xa3, 1},
				{0xb0, 0xbf, 1},
			}},
		{
			Table: TransTable{
				{0x80, 0x86, 1},
				{0x8b, 0xbb, 1},
			}},
		{
			Table: TransTable{
				{0x80, 0xad, 1},
				{0xb0, 0xbf, 1},
			}},
		{
			Table: TransTable{{0x80, 0x99, 1}}},
		{
			Table: TransTable{
				{0x80, 0x86, 1},
				{0x93, 0x97, 1},
				{0x9d, 0x9d, 1},
				{0x9f, 0xa8, 1},
				{0xaa, 0xb6, 1},
				{0xb8, 0xbc, 1},
				{0xbe, 0xbe, 1},
			}},
		{
			Table: TransTable{
				{0x80, 0x81, 1},
				{0x83, 0x84, 1},
				{0x86, 0xbf, 1},
			}},
		{
			Table: TransTable{{0x80, 0xb1, 1}}},
		{
			Table: TransTable{{0x93, 0xbf, 1}}},
		{
			Table: TransTable{{0x80, 0xbd, 1}}},
		{
			Table: TransTable{{0x90, 0xbf, 1}}},
		{
			Table: TransTable{
				{0x80, 0x8f, 1},
				{0x92, 0xbf, 1},
			}},
		{
			Table: TransTable{
				{0x80, 0x87, 1},
				{0xb0, 0xbb, 1},
			}},
		{
			Table: TransTable{
				{0xb0, 0xb4, 1},
				{0xb6, 0xbf, 1},
			}},
		{
			Table: TransTable{{0x80, 0xbc, 1}}},
		{
			Table: TransTable{{0xa1, 0xba, 1}}},
		{
			Table: TransTable{
				{0x81, 0x9a, 1},
				{0xa6, 0xbf, 1},
			}},
		{
			Table: TransTable{{0x80, 0xbe, 1}}},
		{
			Table: TransTable{
				{0x82, 0x87, 1},
				{0x8a, 0x8f, 1},
				{0x92, 0x97, 1},
				{0x9a, 0x9c, 1},
			}},
		{
			Table: TransTable{
				{0x80, 0x80, 170},
				{0x81, 0x81, 171},
				{0x82, 0x82, 4},
				{0x83, 0x83, 172},
				{0x8a, 0x8a, 173},
				{0x8b, 0x8b, 174},
				{0x8c, 0x8c, 175},
				{0x8d, 0x8d, 176},
				{0x8e, 0x8e, 177},
				{0x8f, 0x8f, 178},
				{0x90, 0x91, 4},
				{0x92, 0x92, 179},
				{0xa0, 0xa0, 180},
				{0xa1, 0xa1, 181},
				{0xa4, 0xa4, 182},
				{0xa6, 0xa6, 183},
				{0xa8, 0xa8, 184},
				{0xa9, 0xa9, 185},
				{0xac, 0xac, 82},
				{0xad, 0xad, 186},
				{0xb0, 0xb0, 4},
				{0xb1, 0xb1, 187},
			}},
		{
			Table: TransTable{
				{0x80, 0x80, 188},
				{0x82, 0x82, 189},
				{0x83, 0x83, 190},
				{0x84, 0x84, 191},
				{0x86, 0x86, 192},
				{0x87, 0x87, 193},
				{0x9a, 0x9a, 75},
			}},
		{
			Table: TransTable{
				{0x80, 0x8c, 4},
				{0x8d, 0x8d, 194},
			}},
		{
			Table: TransTable{
				{0x80, 0x8f, 4},
				{0x90, 0x90, 194},
			}},
		{
			Table: TransTable{
				{0xa0, 0xa7, 4},
				{0xa8, 0xa8, 195},
				{0xbc, 0xbc, 4},
				{0xbd, 0xbd, 196},
				{0xbe, 0xbe, 197},
			}},
		{
			Table: TransTable{{0x80, 0x80, 198}}},
		{
			Table: TransTable{
				{0x90, 0x90, 4},
				{0x91, 0x91, 199},
				{0x92, 0x92, 200},
				{0x93, 0x93, 201},
				{0x94, 0x94, 202},
				{0x95, 0x95, 203},
				{0x96, 0x99, 4},
				{0x9a, 0x9a, 204},
				{0x9b, 0x9b, 205},
				{0x9c, 0x9c, 206},
				{0x9d, 0x9d, 207},
				{0x9e, 0x9e, 208},
				{0x9f, 0x9f, 209},
			}},
		{
			Table: TransTable{
				{0xb8, 0xb8, 210},
				{0xb9, 0xb9, 211},
				{0xba, 0xba, 212},
			}},
		{
			Table: TransTable{
				{0x80, 0x9a, 4},
				{0x9b, 0x9b, 213},
				{0x9c, 0xbf, 4},
			}},
		{
			Table: TransTable{
				{0x80, 0x9b, 4},
				{0x9c, 0x9c, 71},
				{0x9d, 0x9f, 4},
				{0xa0, 0xa0, 179},
			}},
		{
			Table: TransTable{
				{0xa0, 0xa7, 4},
				{0xa8, 0xa8, 179},
			}},
		{
			Table: TransTable{
				{0x80, 0x8b, 1},
				{0x8d, 0xa6, 1},
				{0xa8, 0xba, 1},
				{0xbc, 0xbd, 1},
				{0xbf, 0xbf, 1},
			}},
		{
			Table: TransTable{
				{0x80, 0x8d, 1},
				{0x90, 0x9d, 1},
			}},
		{
			Table: TransTable{{0x80, 0xba, 1}}},
		{
			Table: TransTable{
				{0x80, 0x9c, 1},
				{0xa0, 0xbf, 1},
			}},
		{
			Table: TransTable{{0x80, 0x90, 1}}},
		{
			Table: TransTable{
				{0x80, 0x9e, 1},
				{0xb0, 0xbf, 1},
			}},
		{
			Table: TransTable{
				{0x80, 0x80, 1},
				{0x82, 0x89, 1},
			}},
		{
			Table: TransTable{
				{0x80, 0x9d, 1},
				{0xa0, 0xbf, 1},
			}},
		{
			Table: TransTable{
				{0x80, 0x83, 1},
				{0x88, 0x8f, 1},
			}},
		{
			Table: TransTable{{0x80, 0x9d, 1}}},
		{
			Table: TransTable{
				{0x80, 0x85, 1},
				{0x88, 0x88, 1},
				{0x8a, 0xb5, 1},
				{0xb7, 0xb8, 1},
				{0xbc, 0xbc, 1},
				{0xbf, 0xbf, 1},
			}},
		{
			Table: TransTable{{0x80, 0x95, 1}}},
		{
			Table: TransTable{
				{0x80, 0x95, 1},
				{0xa0, 0xb9, 1},
			}},
		{
			Table: TransTable{
				{0x80, 0xb7, 1},
				{0xbe, 0xbf, 1},
			}},
		{
			Table: TransTable{
				{0x80, 0x80, 1},
				{0x90, 0x93, 1},
				{0x95, 0x97, 1},
				{0x99, 0xb3, 1},
			}},
		{
			Table: TransTable{{0xa0, 0xbc, 1}}},
		{
			Table: TransTable{
				{0x80, 0x95, 1},
				{0xa0, 0xb2, 1},
			}},
		{
			Table: TransTable{{0x80, 0x88, 1}}},
		{
			Table: TransTable{{0x83, 0xb7, 1}}},
		{
			Table: TransTable{{0x83, 0xaf, 1}}},
		{
			Table: TransTable{{0x90, 0xa8, 1}}},
		{
			Table: TransTable{{0x83, 0xa6, 1}}},
		{
			Table: TransTable{{0x83, 0xb2, 1}}},
		{
			Table: TransTable{{0x81, 0x84, 1}}},
		{
			Table: TransTable{{0x80, 0xae, 1}}},
		{
			Table: TransTable{{0x80, 0xb8, 1}}},
		{
			Table: TransTable{
				{0x80, 0x84, 1},
				{0x90, 0x90, 1},
			}},
		{
			Table: TransTable{{0x93, 0x9f, 1}}},
		{
			Table: TransTable{{0x80, 0x81, 1}}},
		{
			Table: TransTable{
				{0x80, 0x94, 1},
				{0x96, 0xbf, 1},
			}},
		{
			Table: TransTable{
				{0x80, 0x9c, 1},
				{0x9e, 0x9f, 1},
				{0xa2, 0xa2, 1},
				{0xa5, 0xa6, 1},
				{0xa9, 0xac, 1},
				{0xae, 0xb9, 1},
				{0xbb, 0xbb, 1},
				{0xbd, 0xbf, 1},
			}},
		{
			Table: TransTable{
				{0x80, 0x83, 1},
				{0x85, 0xbf, 1},
			}},
		{
			Table: TransTable{
				{0x80, 0x85, 1},
				{0x87, 0x8a, 1},
				{0x8d, 0x94, 1},
				{0x96, 0x9c, 1},
				{0x9e, 0xb9, 1},
				{0xbb, 0xbe, 1},
			}},
		{
			Table: TransTable{
				{0x80, 0x84, 1},
				{0x86, 0x86, 1},
				{0x8a, 0x90, 1},
				{0x92, 0xbf, 1},
			}},
		{
			Table: TransTable{
				{0x80, 0xa5, 1},
				{0xa8, 0xbf, 1},
			}},
		{
			Table: TransTable{
				{0x80, 0x80, 1},
				{0x82, 0x9a, 1},
				{0x9c, 0xba, 1},
				{0xbc, 0xbf, 1},
			}},
		{
			Table: TransTable{
				{0x80, 0x94, 1},
				{0x96, 0xb4, 1},
				{0xb6, 0xbf, 1},
			}},
		{
			Table: TransTable{
				{0x80, 0x8e, 1},
				{0x90, 0xae, 1},
				{0xb0, 0xbf, 1},
			}},
		{
			Table: TransTable{
				{0x80, 0x88, 1},
				{0x8a, 0xa8, 1},
				{0xaa, 0xbf, 1},
			}},
		{
			Table: TransTable{
				{0x80, 0x82, 1},
				{0x84, 0x8b, 1},
			}},
		{
			Table: TransTable{
				{0x80, 0x83, 1},
				{0x85, 0x9f, 1},
				{0xa1, 0xa2, 1},
				{0xa4, 0xa4, 1},
				{0xa7, 0xa7, 1},
				{0xa9, 0xb2, 1},
				{0xb4, 0xb7, 1},
				{0xb9, 0xb9, 1},
				{0xbb, 0xbb, 1},
			}},
		{
			Table: TransTable{
				{0x82, 0x82, 1},
				{0x87, 0x87, 1},
				{0x89, 0x89, 1},
				{0x8b, 0x8b, 1},
				{0x8d, 0x8f, 1},
				{0x91, 0x92, 1},
				{0x94, 0x94, 1},
				{0x97, 0x97, 1},
				{0x99, 0x99, 1},
				{0x9b, 0x9b, 1},
				{0x9d, 0x9d, 1},
				{0x9f, 0x9f, 1},
				{0xa1, 0xa2, 1},
				{0xa4, 0xa4, 1},
				{0xa7, 0xaa, 1},
				{0xac, 0xb2, 1},
				{0xb4, 0xb7, 1},
				{0xb9, 0xbc, 1},
				{0xbe, 0xbe, 1},
			}},
		{
			Table: TransTable{
				{0x80, 0x89, 1},
				{0x8b, 0x9b, 1},
				{0xa1, 0xa3, 1},
				{0xa5, 0xa9, 1},
				{0xab, 0xbb, 1},
			}},
		{
			Table: TransTable{{0x80, 0x96, 1}}},
	}}
