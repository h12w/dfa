package dfa

var classL = &M{states: states{
	{table: transTable{a: []trans{
		{s: 0x41, e: 0x5a, next: 1},
		{s: 0x61, e: 0x7a, next: 1},
		{s: 0xc2, e: 0xc2, next: 2},
		{s: 0xc3, e: 0xc3, next: 3},
		{s: 0xc4, e: 0xca, next: 4},
		{s: 0xcb, e: 0xcb, next: 5},
		{s: 0xcd, e: 0xcd, next: 6},
		{s: 0xce, e: 0xce, next: 7},
		{s: 0xcf, e: 0xcf, next: 8},
		{s: 0xd0, e: 0xd1, next: 4},
		{s: 0xd2, e: 0xd2, next: 9},
		{s: 0xd3, e: 0xd3, next: 4},
		{s: 0xd4, e: 0xd4, next: 10},
		{s: 0xd5, e: 0xd5, next: 11},
		{s: 0xd6, e: 0xd6, next: 12},
		{s: 0xd7, e: 0xd7, next: 13},
		{s: 0xd8, e: 0xd8, next: 14},
		{s: 0xd9, e: 0xd9, next: 15},
		{s: 0xda, e: 0xda, next: 4},
		{s: 0xdb, e: 0xdb, next: 16},
		{s: 0xdc, e: 0xdc, next: 17},
		{s: 0xdd, e: 0xdd, next: 18},
		{s: 0xde, e: 0xde, next: 19},
		{s: 0xdf, e: 0xdf, next: 20},
		{s: 0xe0, e: 0xe0, next: 21},
		{s: 0xe1, e: 0xe1, next: 22},
		{s: 0xe2, e: 0xe2, next: 23},
		{s: 0xe3, e: 0xe3, next: 24},
		{s: 0xe4, e: 0xe4, next: 25},
		{s: 0xe5, e: 0xe8, next: 26},
		{s: 0xe9, e: 0xe9, next: 27},
		{s: 0xea, e: 0xea, next: 28},
		{s: 0xeb, e: 0xec, next: 26},
		{s: 0xed, e: 0xed, next: 29},
		{s: 0xef, e: 0xef, next: 30},
		{s: 0xf0, e: 0xf0, next: 31}}}, label: 0},
	{table: transTable{a: []trans(nil)}, label: 1},
	{table: transTable{a: []trans{
		{s: 0xaa, e: 0xaa, next: 1},
		{s: 0xb5, e: 0xb5, next: 1},
		{s: 0xba, e: 0xba, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0x96, next: 1},
		{s: 0x98, e: 0xb6, next: 1},
		{s: 0xb8, e: 0xbf, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0xbf, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0x81, next: 1},
		{s: 0x86, e: 0x91, next: 1},
		{s: 0xa0, e: 0xa4, next: 1},
		{s: 0xac, e: 0xac, next: 1},
		{s: 0xae, e: 0xae, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0xb0, e: 0xb4, next: 1},
		{s: 0xb6, e: 0xb7, next: 1},
		{s: 0xba, e: 0xbd, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x86, e: 0x86, next: 1},
		{s: 0x88, e: 0x8a, next: 1},
		{s: 0x8c, e: 0x8c, next: 1},
		{s: 0x8e, e: 0xa1, next: 1},
		{s: 0xa3, e: 0xbf, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0xb5, next: 1},
		{s: 0xb7, e: 0xbf, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0x81, next: 1},
		{s: 0x8a, e: 0xbf, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0xa7, next: 1},
		{s: 0xb1, e: 0xbf, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0x96, next: 1},
		{s: 0x99, e: 0x99, next: 1},
		{s: 0xa1, e: 0xbf, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0x87, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x90, e: 0xaa, next: 1},
		{s: 0xb0, e: 0xb2, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0xa0, e: 0xbf, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0x8a, next: 1},
		{s: 0xae, e: 0xaf, next: 1},
		{s: 0xb1, e: 0xbf, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0x93, next: 1},
		{s: 0x95, e: 0x95, next: 1},
		{s: 0xa5, e: 0xa6, next: 1},
		{s: 0xae, e: 0xaf, next: 1},
		{s: 0xba, e: 0xbc, next: 1},
		{s: 0xbf, e: 0xbf, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x90, e: 0x90, next: 1},
		{s: 0x92, e: 0xaf, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x8d, e: 0xbf, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0xa5, next: 1},
		{s: 0xb1, e: 0xb1, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x8a, e: 0xaa, next: 1},
		{s: 0xb4, e: 0xb5, next: 1},
		{s: 0xba, e: 0xba, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0xa0, e: 0xa0, next: 32},
		{s: 0xa1, e: 0xa1, next: 33},
		{s: 0xa2, e: 0xa2, next: 34},
		{s: 0xa4, e: 0xa4, next: 35},
		{s: 0xa5, e: 0xa5, next: 36},
		{s: 0xa6, e: 0xa6, next: 37},
		{s: 0xa7, e: 0xa7, next: 38},
		{s: 0xa8, e: 0xa8, next: 39},
		{s: 0xa9, e: 0xa9, next: 40},
		{s: 0xaa, e: 0xaa, next: 41},
		{s: 0xab, e: 0xab, next: 42},
		{s: 0xac, e: 0xac, next: 43},
		{s: 0xad, e: 0xad, next: 44},
		{s: 0xae, e: 0xae, next: 45},
		{s: 0xaf, e: 0xaf, next: 46},
		{s: 0xb0, e: 0xb0, next: 47},
		{s: 0xb1, e: 0xb1, next: 48},
		{s: 0xb2, e: 0xb2, next: 47},
		{s: 0xb3, e: 0xb3, next: 49},
		{s: 0xb4, e: 0xb4, next: 50},
		{s: 0xb5, e: 0xb5, next: 51},
		{s: 0xb6, e: 0xb6, next: 52},
		{s: 0xb7, e: 0xb7, next: 53},
		{s: 0xb8, e: 0xb8, next: 54},
		{s: 0xb9, e: 0xb9, next: 53},
		{s: 0xba, e: 0xba, next: 55},
		{s: 0xbb, e: 0xbb, next: 56},
		{s: 0xbc, e: 0xbc, next: 57},
		{s: 0xbd, e: 0xbd, next: 58},
		{s: 0xbe, e: 0xbe, next: 59}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0x80, next: 60},
		{s: 0x81, e: 0x81, next: 61},
		{s: 0x82, e: 0x82, next: 62},
		{s: 0x83, e: 0x83, next: 63},
		{s: 0x84, e: 0x88, next: 4},
		{s: 0x89, e: 0x89, next: 64},
		{s: 0x8a, e: 0x8a, next: 65},
		{s: 0x8b, e: 0x8b, next: 66},
		{s: 0x8c, e: 0x8c, next: 67},
		{s: 0x8d, e: 0x8d, next: 68},
		{s: 0x8e, e: 0x8e, next: 69},
		{s: 0x8f, e: 0x8f, next: 70},
		{s: 0x90, e: 0x90, next: 71},
		{s: 0x91, e: 0x98, next: 4},
		{s: 0x99, e: 0x99, next: 72},
		{s: 0x9a, e: 0x9a, next: 73},
		{s: 0x9b, e: 0x9b, next: 74},
		{s: 0x9c, e: 0x9c, next: 75},
		{s: 0x9d, e: 0x9d, next: 76},
		{s: 0x9e, e: 0x9e, next: 77},
		{s: 0x9f, e: 0x9f, next: 78},
		{s: 0xa0, e: 0xa0, next: 14},
		{s: 0xa1, e: 0xa1, next: 79},
		{s: 0xa2, e: 0xa2, next: 80},
		{s: 0xa3, e: 0xa3, next: 81},
		{s: 0xa4, e: 0xa4, next: 82},
		{s: 0xa5, e: 0xa5, next: 83},
		{s: 0xa6, e: 0xa6, next: 84},
		{s: 0xa7, e: 0xa7, next: 85},
		{s: 0xa8, e: 0xa8, next: 86},
		{s: 0xa9, e: 0xa9, next: 87},
		{s: 0xaa, e: 0xaa, next: 88},
		{s: 0xac, e: 0xac, next: 89},
		{s: 0xad, e: 0xad, next: 90},
		{s: 0xae, e: 0xae, next: 91},
		{s: 0xaf, e: 0xaf, next: 92},
		{s: 0xb0, e: 0xb0, next: 93},
		{s: 0xb1, e: 0xb1, next: 94},
		{s: 0xb3, e: 0xb3, next: 95},
		{s: 0xb4, e: 0xb6, next: 4},
		{s: 0xb8, e: 0xbb, next: 4},
		{s: 0xbc, e: 0xbc, next: 96},
		{s: 0xbd, e: 0xbd, next: 97},
		{s: 0xbe, e: 0xbe, next: 98},
		{s: 0xbf, e: 0xbf, next: 99}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x81, e: 0x81, next: 100},
		{s: 0x82, e: 0x82, next: 101},
		{s: 0x84, e: 0x84, next: 102},
		{s: 0x85, e: 0x85, next: 103},
		{s: 0x86, e: 0x86, next: 104},
		{s: 0xb0, e: 0xb0, next: 105},
		{s: 0xb1, e: 0xb1, next: 106},
		{s: 0xb2, e: 0xb2, next: 4},
		{s: 0xb3, e: 0xb3, next: 107},
		{s: 0xb4, e: 0xb4, next: 108},
		{s: 0xb5, e: 0xb5, next: 109},
		{s: 0xb6, e: 0xb6, next: 110},
		{s: 0xb7, e: 0xb7, next: 111},
		{s: 0xb8, e: 0xb8, next: 112}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0x80, next: 113},
		{s: 0x81, e: 0x81, next: 71},
		{s: 0x82, e: 0x82, next: 114},
		{s: 0x83, e: 0x83, next: 115},
		{s: 0x84, e: 0x84, next: 116},
		{s: 0x85, e: 0x85, next: 4},
		{s: 0x86, e: 0x86, next: 117},
		{s: 0x87, e: 0x87, next: 118},
		{s: 0x90, e: 0xbf, next: 4}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0xb5, next: 4},
		{s: 0xb6, e: 0xb6, next: 81},
		{s: 0xb8, e: 0xbf, next: 4}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0xbf, next: 4}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0xbe, next: 4},
		{s: 0xbf, e: 0xbf, next: 119}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0x91, next: 4},
		{s: 0x92, e: 0x92, next: 119},
		{s: 0x93, e: 0x93, next: 120},
		{s: 0x94, e: 0x97, next: 4},
		{s: 0x98, e: 0x98, next: 121},
		{s: 0x99, e: 0x99, next: 122},
		{s: 0x9a, e: 0x9a, next: 123},
		{s: 0x9b, e: 0x9b, next: 92},
		{s: 0x9c, e: 0x9c, next: 124},
		{s: 0x9d, e: 0x9d, next: 4},
		{s: 0x9e, e: 0x9e, next: 125},
		{s: 0x9f, e: 0x9f, next: 126},
		{s: 0xa0, e: 0xa0, next: 127},
		{s: 0xa1, e: 0xa1, next: 77},
		{s: 0xa2, e: 0xa2, next: 128},
		{s: 0xa3, e: 0xa3, next: 129},
		{s: 0xa4, e: 0xa4, next: 130},
		{s: 0xa5, e: 0xa5, next: 131},
		{s: 0xa6, e: 0xa6, next: 132},
		{s: 0xa7, e: 0xa7, next: 133},
		{s: 0xa8, e: 0xa8, next: 134},
		{s: 0xa9, e: 0xa9, next: 135},
		{s: 0xaa, e: 0xaa, next: 136},
		{s: 0xab, e: 0xab, next: 137},
		{s: 0xac, e: 0xac, next: 138},
		{s: 0xaf, e: 0xaf, next: 139},
		{s: 0xb0, e: 0xbf, next: 4}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0x9d, next: 4},
		{s: 0x9e, e: 0x9e, next: 140},
		{s: 0x9f, e: 0x9f, next: 141}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0xa4, e: 0xa8, next: 4},
		{s: 0xa9, e: 0xa9, next: 142},
		{s: 0xaa, e: 0xaa, next: 4},
		{s: 0xab, e: 0xab, next: 143},
		{s: 0xac, e: 0xac, next: 144},
		{s: 0xad, e: 0xad, next: 145},
		{s: 0xae, e: 0xae, next: 146},
		{s: 0xaf, e: 0xaf, next: 147},
		{s: 0xb0, e: 0xb3, next: 4},
		{s: 0xb4, e: 0xb4, next: 148},
		{s: 0xb5, e: 0xb5, next: 149},
		{s: 0xb6, e: 0xb6, next: 150},
		{s: 0xb7, e: 0xb7, next: 151},
		{s: 0xb9, e: 0xb9, next: 152},
		{s: 0xba, e: 0xba, next: 4},
		{s: 0xbb, e: 0xbb, next: 153},
		{s: 0xbc, e: 0xbc, next: 154},
		{s: 0xbd, e: 0xbd, next: 155},
		{s: 0xbe, e: 0xbe, next: 156},
		{s: 0xbf, e: 0xbf, next: 157}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x90, e: 0x90, next: 158},
		{s: 0x91, e: 0x91, next: 159},
		{s: 0x92, e: 0x92, next: 160},
		{s: 0x93, e: 0x93, next: 161},
		{s: 0x96, e: 0x96, next: 162},
		{s: 0x9b, e: 0x9b, next: 163},
		{s: 0x9d, e: 0x9d, next: 164},
		{s: 0x9e, e: 0x9e, next: 165},
		{s: 0xa0, e: 0xa9, next: 26},
		{s: 0xaa, e: 0xaa, next: 166},
		{s: 0xab, e: 0xab, next: 167},
		{s: 0xaf, e: 0xaf, next: 168}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0x95, next: 1},
		{s: 0x9a, e: 0x9a, next: 1},
		{s: 0xa4, e: 0xa4, next: 1},
		{s: 0xa8, e: 0xa8, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0x98, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0xa0, e: 0xa0, next: 1},
		{s: 0xa2, e: 0xac, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x84, e: 0xb9, next: 1},
		{s: 0xbd, e: 0xbd, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x90, e: 0x90, next: 1},
		{s: 0x98, e: 0xa1, next: 1},
		{s: 0xb1, e: 0xb7, next: 1},
		{s: 0xb9, e: 0xbf, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x85, e: 0x8c, next: 1},
		{s: 0x8f, e: 0x90, next: 1},
		{s: 0x93, e: 0xa8, next: 1},
		{s: 0xaa, e: 0xb0, next: 1},
		{s: 0xb2, e: 0xb2, next: 1},
		{s: 0xb6, e: 0xb9, next: 1},
		{s: 0xbd, e: 0xbd, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x8e, e: 0x8e, next: 1},
		{s: 0x9c, e: 0x9d, next: 1},
		{s: 0x9f, e: 0xa1, next: 1},
		{s: 0xb0, e: 0xb1, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x85, e: 0x8a, next: 1},
		{s: 0x8f, e: 0x90, next: 1},
		{s: 0x93, e: 0xa8, next: 1},
		{s: 0xaa, e: 0xb0, next: 1},
		{s: 0xb2, e: 0xb3, next: 1},
		{s: 0xb5, e: 0xb6, next: 1},
		{s: 0xb8, e: 0xb9, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x99, e: 0x9c, next: 1},
		{s: 0x9e, e: 0x9e, next: 1},
		{s: 0xb2, e: 0xb4, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x85, e: 0x8d, next: 1},
		{s: 0x8f, e: 0x91, next: 1},
		{s: 0x93, e: 0xa8, next: 1},
		{s: 0xaa, e: 0xb0, next: 1},
		{s: 0xb2, e: 0xb3, next: 1},
		{s: 0xb5, e: 0xb9, next: 1},
		{s: 0xbd, e: 0xbd, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x90, e: 0x90, next: 1},
		{s: 0xa0, e: 0xa1, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x85, e: 0x8c, next: 1},
		{s: 0x8f, e: 0x90, next: 1},
		{s: 0x93, e: 0xa8, next: 1},
		{s: 0xaa, e: 0xb0, next: 1},
		{s: 0xb2, e: 0xb3, next: 1},
		{s: 0xb5, e: 0xb9, next: 1},
		{s: 0xbd, e: 0xbd, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x9c, e: 0x9d, next: 1},
		{s: 0x9f, e: 0xa1, next: 1},
		{s: 0xb1, e: 0xb1, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x83, e: 0x83, next: 1},
		{s: 0x85, e: 0x8a, next: 1},
		{s: 0x8e, e: 0x90, next: 1},
		{s: 0x92, e: 0x95, next: 1},
		{s: 0x99, e: 0x9a, next: 1},
		{s: 0x9c, e: 0x9c, next: 1},
		{s: 0x9e, e: 0x9f, next: 1},
		{s: 0xa3, e: 0xa4, next: 1},
		{s: 0xa8, e: 0xaa, next: 1},
		{s: 0xae, e: 0xb9, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x90, e: 0x90, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x85, e: 0x8c, next: 1},
		{s: 0x8e, e: 0x90, next: 1},
		{s: 0x92, e: 0xa8, next: 1},
		{s: 0xaa, e: 0xb3, next: 1},
		{s: 0xb5, e: 0xb9, next: 1},
		{s: 0xbd, e: 0xbd, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x98, e: 0x99, next: 1},
		{s: 0xa0, e: 0xa1, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x9e, e: 0x9e, next: 1},
		{s: 0xa0, e: 0xa1, next: 1},
		{s: 0xb1, e: 0xb2, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x85, e: 0x8c, next: 1},
		{s: 0x8e, e: 0x90, next: 1},
		{s: 0x92, e: 0xba, next: 1},
		{s: 0xbd, e: 0xbd, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x8e, e: 0x8e, next: 1},
		{s: 0xa0, e: 0xa1, next: 1},
		{s: 0xba, e: 0xbf, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x85, e: 0x96, next: 1},
		{s: 0x9a, e: 0xb1, next: 1},
		{s: 0xb3, e: 0xbb, next: 1},
		{s: 0xbd, e: 0xbd, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0x86, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x81, e: 0xb0, next: 1},
		{s: 0xb2, e: 0xb3, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x81, e: 0x82, next: 1},
		{s: 0x84, e: 0x84, next: 1},
		{s: 0x87, e: 0x88, next: 1},
		{s: 0x8a, e: 0x8a, next: 1},
		{s: 0x8d, e: 0x8d, next: 1},
		{s: 0x94, e: 0x97, next: 1},
		{s: 0x99, e: 0x9f, next: 1},
		{s: 0xa1, e: 0xa3, next: 1},
		{s: 0xa5, e: 0xa5, next: 1},
		{s: 0xa7, e: 0xa7, next: 1},
		{s: 0xaa, e: 0xab, next: 1},
		{s: 0xad, e: 0xb0, next: 1},
		{s: 0xb2, e: 0xb3, next: 1},
		{s: 0xbd, e: 0xbd, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0x84, next: 1},
		{s: 0x86, e: 0x86, next: 1},
		{s: 0x9c, e: 0x9f, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0x80, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0x87, next: 1},
		{s: 0x89, e: 0xac, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x88, e: 0x8c, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0xaa, next: 1},
		{s: 0xbf, e: 0xbf, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x90, e: 0x95, next: 1},
		{s: 0x9a, e: 0x9d, next: 1},
		{s: 0xa1, e: 0xa1, next: 1},
		{s: 0xa5, e: 0xa6, next: 1},
		{s: 0xae, e: 0xb0, next: 1},
		{s: 0xb5, e: 0xbf, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0x81, next: 1},
		{s: 0x8e, e: 0x8e, next: 1},
		{s: 0xa0, e: 0xbf, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0x85, next: 1},
		{s: 0x87, e: 0x87, next: 1},
		{s: 0x8d, e: 0x8d, next: 1},
		{s: 0x90, e: 0xba, next: 1},
		{s: 0xbc, e: 0xbf, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0x88, next: 1},
		{s: 0x8a, e: 0x8d, next: 1},
		{s: 0x90, e: 0x96, next: 1},
		{s: 0x98, e: 0x98, next: 1},
		{s: 0x9a, e: 0x9d, next: 1},
		{s: 0xa0, e: 0xbf, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0x88, next: 1},
		{s: 0x8a, e: 0x8d, next: 1},
		{s: 0x90, e: 0xb0, next: 1},
		{s: 0xb2, e: 0xb5, next: 1},
		{s: 0xb8, e: 0xbe, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0x80, next: 1},
		{s: 0x82, e: 0x85, next: 1},
		{s: 0x88, e: 0x96, next: 1},
		{s: 0x98, e: 0xbf, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0x90, next: 1},
		{s: 0x92, e: 0x95, next: 1},
		{s: 0x98, e: 0xbf, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0x9a, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0x8f, next: 1},
		{s: 0xa0, e: 0xbf, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0xb4, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x81, e: 0xbf, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0xac, next: 1},
		{s: 0xaf, e: 0xbf, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x81, e: 0x9a, next: 1},
		{s: 0xa0, e: 0xbf, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0xaa, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0x8c, next: 1},
		{s: 0x8e, e: 0x91, next: 1},
		{s: 0xa0, e: 0xb1, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0x91, next: 1},
		{s: 0xa0, e: 0xac, next: 1},
		{s: 0xae, e: 0xb0, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0xb3, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x97, e: 0x97, next: 1},
		{s: 0x9c, e: 0x9c, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0xb7, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0xa8, next: 1},
		{s: 0xaa, e: 0xaa, next: 1},
		{s: 0xb0, e: 0xbf, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0xb5, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0x9c, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x90, e: 0xad, next: 1},
		{s: 0xb0, e: 0xb4, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0xab, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x81, e: 0x87, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0x96, next: 1},
		{s: 0xa0, e: 0xbf, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0x94, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0xa7, e: 0xa7, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x85, e: 0xb3, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x85, e: 0x8b, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x83, e: 0xa0, next: 1},
		{s: 0xae, e: 0xaf, next: 1},
		{s: 0xba, e: 0xbf, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0xa5, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0xa3, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x8d, e: 0x8f, next: 1},
		{s: 0x9a, e: 0xbd, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0xa9, e: 0xac, next: 1},
		{s: 0xae, e: 0xb1, next: 1},
		{s: 0xb5, e: 0xb6, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0x95, next: 1},
		{s: 0x98, e: 0x9d, next: 1},
		{s: 0xa0, e: 0xbf, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0x85, next: 1},
		{s: 0x88, e: 0x8d, next: 1},
		{s: 0x90, e: 0x97, next: 1},
		{s: 0x99, e: 0x99, next: 1},
		{s: 0x9b, e: 0x9b, next: 1},
		{s: 0x9d, e: 0x9d, next: 1},
		{s: 0x9f, e: 0xbd, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0xb4, next: 1},
		{s: 0xb6, e: 0xbc, next: 1},
		{s: 0xbe, e: 0xbe, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x82, e: 0x84, next: 1},
		{s: 0x86, e: 0x8c, next: 1},
		{s: 0x90, e: 0x93, next: 1},
		{s: 0x96, e: 0x9b, next: 1},
		{s: 0xa0, e: 0xac, next: 1},
		{s: 0xb2, e: 0xb4, next: 1},
		{s: 0xb6, e: 0xbc, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0xb1, e: 0xb1, next: 1},
		{s: 0xbf, e: 0xbf, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x90, e: 0x9c, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x82, e: 0x82, next: 1},
		{s: 0x87, e: 0x87, next: 1},
		{s: 0x8a, e: 0x93, next: 1},
		{s: 0x95, e: 0x95, next: 1},
		{s: 0x99, e: 0x9d, next: 1},
		{s: 0xa4, e: 0xa4, next: 1},
		{s: 0xa6, e: 0xa6, next: 1},
		{s: 0xa8, e: 0xa8, next: 1},
		{s: 0xaa, e: 0xad, next: 1},
		{s: 0xaf, e: 0xb9, next: 1},
		{s: 0xbc, e: 0xbf, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x85, e: 0x89, next: 1},
		{s: 0x8e, e: 0x8e, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x83, e: 0x84, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0xae, next: 1},
		{s: 0xb0, e: 0xbf, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0x9e, next: 1},
		{s: 0xa0, e: 0xbf, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0xa4, next: 1},
		{s: 0xab, e: 0xae, next: 1},
		{s: 0xb2, e: 0xb3, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0xa5, next: 1},
		{s: 0xa7, e: 0xa7, next: 1},
		{s: 0xad, e: 0xad, next: 1},
		{s: 0xb0, e: 0xbf, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0xa7, next: 1},
		{s: 0xaf, e: 0xaf, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0x96, next: 1},
		{s: 0xa0, e: 0xa6, next: 1},
		{s: 0xa8, e: 0xae, next: 1},
		{s: 0xb0, e: 0xb6, next: 1},
		{s: 0xb8, e: 0xbe, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0x86, next: 1},
		{s: 0x88, e: 0x8e, next: 1},
		{s: 0x90, e: 0x96, next: 1},
		{s: 0x98, e: 0x9e, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0xaf, e: 0xaf, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x85, e: 0x86, next: 1},
		{s: 0xb1, e: 0xb5, next: 1},
		{s: 0xbb, e: 0xbc, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0x96, next: 1},
		{s: 0x9d, e: 0x9f, next: 1},
		{s: 0xa1, e: 0xbf, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0xba, next: 1},
		{s: 0xbc, e: 0xbf, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x85, e: 0xad, next: 1},
		{s: 0xb1, e: 0xbf, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0x8e, next: 1},
		{s: 0xa0, e: 0xba, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0xb0, e: 0xbf, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0x8c, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x90, e: 0xbd, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0x8c, next: 1},
		{s: 0x90, e: 0x9f, next: 1},
		{s: 0xaa, e: 0xab, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0xae, next: 1},
		{s: 0xbf, e: 0xbf, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0x97, next: 1},
		{s: 0xa0, e: 0xbf, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x97, e: 0x9f, next: 1},
		{s: 0xa2, e: 0xbf, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0x88, next: 1},
		{s: 0x8b, e: 0x8e, next: 1},
		{s: 0x90, e: 0x93, next: 1},
		{s: 0xa0, e: 0xaa, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0xb8, e: 0xbf, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0x81, next: 1},
		{s: 0x83, e: 0x85, next: 1},
		{s: 0x87, e: 0x8a, next: 1},
		{s: 0x8c, e: 0xa2, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x82, e: 0xb3, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0xb2, e: 0xb7, next: 1},
		{s: 0xbb, e: 0xbb, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x8a, e: 0xa5, next: 1},
		{s: 0xb0, e: 0xbf, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0x86, next: 1},
		{s: 0xa0, e: 0xbc, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x84, e: 0xb2, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x8f, e: 0x8f, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0xa8, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0x82, next: 1},
		{s: 0x84, e: 0x8b, next: 1},
		{s: 0xa0, e: 0xb6, next: 1},
		{s: 0xba, e: 0xba, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0xaf, next: 1},
		{s: 0xb1, e: 0xb1, next: 1},
		{s: 0xb5, e: 0xb6, next: 1},
		{s: 0xb9, e: 0xbd, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0x80, next: 1},
		{s: 0x82, e: 0x82, next: 1},
		{s: 0x9b, e: 0x9d, next: 1},
		{s: 0xa0, e: 0xaa, next: 1},
		{s: 0xb2, e: 0xb4, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x81, e: 0x86, next: 1},
		{s: 0x89, e: 0x8e, next: 1},
		{s: 0x91, e: 0x96, next: 1},
		{s: 0xa0, e: 0xa6, next: 1},
		{s: 0xa8, e: 0xae, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0xa2, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0xa3, next: 1},
		{s: 0xb0, e: 0xbf, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0x86, next: 1},
		{s: 0x8b, e: 0xbb, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0xad, next: 1},
		{s: 0xb0, e: 0xbf, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0x99, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0x86, next: 1},
		{s: 0x93, e: 0x97, next: 1},
		{s: 0x9d, e: 0x9d, next: 1},
		{s: 0x9f, e: 0xa8, next: 1},
		{s: 0xaa, e: 0xb6, next: 1},
		{s: 0xb8, e: 0xbc, next: 1},
		{s: 0xbe, e: 0xbe, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0x81, next: 1},
		{s: 0x83, e: 0x84, next: 1},
		{s: 0x86, e: 0xbf, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0xb1, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x93, e: 0xbf, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0xbd, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x90, e: 0xbf, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0x8f, next: 1},
		{s: 0x92, e: 0xbf, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0x87, next: 1},
		{s: 0xb0, e: 0xbb, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0xb0, e: 0xb4, next: 1},
		{s: 0xb6, e: 0xbf, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0xbc, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0xa1, e: 0xba, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x81, e: 0x9a, next: 1},
		{s: 0xa6, e: 0xbf, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0xbe, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x82, e: 0x87, next: 1},
		{s: 0x8a, e: 0x8f, next: 1},
		{s: 0x92, e: 0x97, next: 1},
		{s: 0x9a, e: 0x9c, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0x80, next: 169},
		{s: 0x81, e: 0x81, next: 170},
		{s: 0x82, e: 0x82, next: 4},
		{s: 0x83, e: 0x83, next: 171},
		{s: 0x8a, e: 0x8a, next: 172},
		{s: 0x8b, e: 0x8b, next: 173},
		{s: 0x8c, e: 0x8c, next: 174},
		{s: 0x8d, e: 0x8d, next: 175},
		{s: 0x8e, e: 0x8e, next: 176},
		{s: 0x8f, e: 0x8f, next: 177},
		{s: 0x90, e: 0x91, next: 4},
		{s: 0x92, e: 0x92, next: 178},
		{s: 0xa0, e: 0xa0, next: 179},
		{s: 0xa1, e: 0xa1, next: 180},
		{s: 0xa4, e: 0xa4, next: 181},
		{s: 0xa6, e: 0xa6, next: 182},
		{s: 0xa8, e: 0xa8, next: 183},
		{s: 0xa9, e: 0xa9, next: 184},
		{s: 0xac, e: 0xac, next: 81},
		{s: 0xad, e: 0xad, next: 185},
		{s: 0xb0, e: 0xb0, next: 4},
		{s: 0xb1, e: 0xb1, next: 186}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0x80, next: 187},
		{s: 0x82, e: 0x82, next: 188},
		{s: 0x83, e: 0x83, next: 189},
		{s: 0x84, e: 0x84, next: 190},
		{s: 0x86, e: 0x86, next: 191},
		{s: 0x87, e: 0x87, next: 192},
		{s: 0x9a, e: 0x9a, next: 74}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0x8c, next: 4},
		{s: 0x8d, e: 0x8d, next: 193}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0x8f, next: 4},
		{s: 0x90, e: 0x90, next: 193}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0xa0, e: 0xa7, next: 4},
		{s: 0xa8, e: 0xa8, next: 194},
		{s: 0xbc, e: 0xbc, next: 4},
		{s: 0xbd, e: 0xbd, next: 195},
		{s: 0xbe, e: 0xbe, next: 196}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0x80, next: 197}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x90, e: 0x90, next: 4},
		{s: 0x91, e: 0x91, next: 198},
		{s: 0x92, e: 0x92, next: 199},
		{s: 0x93, e: 0x93, next: 200},
		{s: 0x94, e: 0x94, next: 201},
		{s: 0x95, e: 0x95, next: 202},
		{s: 0x96, e: 0x99, next: 4},
		{s: 0x9a, e: 0x9a, next: 203},
		{s: 0x9b, e: 0x9b, next: 204},
		{s: 0x9c, e: 0x9c, next: 205},
		{s: 0x9d, e: 0x9d, next: 206},
		{s: 0x9e, e: 0x9e, next: 207},
		{s: 0x9f, e: 0x9f, next: 208}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0xb8, e: 0xb8, next: 209},
		{s: 0xb9, e: 0xb9, next: 210},
		{s: 0xba, e: 0xba, next: 211}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0x9a, next: 4},
		{s: 0x9b, e: 0x9b, next: 212},
		{s: 0x9c, e: 0xbf, next: 4}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0x9b, next: 4},
		{s: 0x9c, e: 0x9c, next: 70},
		{s: 0x9d, e: 0x9f, next: 4},
		{s: 0xa0, e: 0xa0, next: 178}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0xa0, e: 0xa7, next: 4},
		{s: 0xa8, e: 0xa8, next: 178}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0x8b, next: 1},
		{s: 0x8d, e: 0xa6, next: 1},
		{s: 0xa8, e: 0xba, next: 1},
		{s: 0xbc, e: 0xbd, next: 1},
		{s: 0xbf, e: 0xbf, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0x8d, next: 1},
		{s: 0x90, e: 0x9d, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0xba, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0x9c, next: 1},
		{s: 0xa0, e: 0xbf, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0x90, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0x9e, next: 1},
		{s: 0xb0, e: 0xbf, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0x80, next: 1},
		{s: 0x82, e: 0x89, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0x9d, next: 1},
		{s: 0xa0, e: 0xbf, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0x83, next: 1},
		{s: 0x88, e: 0x8f, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0x9d, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0x85, next: 1},
		{s: 0x88, e: 0x88, next: 1},
		{s: 0x8a, e: 0xb5, next: 1},
		{s: 0xb7, e: 0xb8, next: 1},
		{s: 0xbc, e: 0xbc, next: 1},
		{s: 0xbf, e: 0xbf, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0x95, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0x95, next: 1},
		{s: 0xa0, e: 0xb9, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0xb7, next: 1},
		{s: 0xbe, e: 0xbf, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0x80, next: 1},
		{s: 0x90, e: 0x93, next: 1},
		{s: 0x95, e: 0x97, next: 1},
		{s: 0x99, e: 0xb3, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0xa0, e: 0xbc, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0x95, next: 1},
		{s: 0xa0, e: 0xb2, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0x88, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x83, e: 0xb7, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x83, e: 0xaf, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x90, e: 0xa8, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x83, e: 0xa6, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x83, e: 0xb2, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x81, e: 0x84, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0xae, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0xb8, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0x84, next: 1},
		{s: 0x90, e: 0x90, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x93, e: 0x9f, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0x81, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0x94, next: 1},
		{s: 0x96, e: 0xbf, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0x9c, next: 1},
		{s: 0x9e, e: 0x9f, next: 1},
		{s: 0xa2, e: 0xa2, next: 1},
		{s: 0xa5, e: 0xa6, next: 1},
		{s: 0xa9, e: 0xac, next: 1},
		{s: 0xae, e: 0xb9, next: 1},
		{s: 0xbb, e: 0xbb, next: 1},
		{s: 0xbd, e: 0xbf, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0x83, next: 1},
		{s: 0x85, e: 0xbf, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0x85, next: 1},
		{s: 0x87, e: 0x8a, next: 1},
		{s: 0x8d, e: 0x94, next: 1},
		{s: 0x96, e: 0x9c, next: 1},
		{s: 0x9e, e: 0xb9, next: 1},
		{s: 0xbb, e: 0xbe, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0x84, next: 1},
		{s: 0x86, e: 0x86, next: 1},
		{s: 0x8a, e: 0x90, next: 1},
		{s: 0x92, e: 0xbf, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0xa5, next: 1},
		{s: 0xa8, e: 0xbf, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0x80, next: 1},
		{s: 0x82, e: 0x9a, next: 1},
		{s: 0x9c, e: 0xba, next: 1},
		{s: 0xbc, e: 0xbf, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0x94, next: 1},
		{s: 0x96, e: 0xb4, next: 1},
		{s: 0xb6, e: 0xbf, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0x8e, next: 1},
		{s: 0x90, e: 0xae, next: 1},
		{s: 0xb0, e: 0xbf, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0x88, next: 1},
		{s: 0x8a, e: 0xa8, next: 1},
		{s: 0xaa, e: 0xbf, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0x82, next: 1},
		{s: 0x84, e: 0x8b, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0x83, next: 1},
		{s: 0x85, e: 0x9f, next: 1},
		{s: 0xa1, e: 0xa2, next: 1},
		{s: 0xa4, e: 0xa4, next: 1},
		{s: 0xa7, e: 0xa7, next: 1},
		{s: 0xa9, e: 0xb2, next: 1},
		{s: 0xb4, e: 0xb7, next: 1},
		{s: 0xb9, e: 0xb9, next: 1},
		{s: 0xbb, e: 0xbb, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x82, e: 0x82, next: 1},
		{s: 0x87, e: 0x87, next: 1},
		{s: 0x89, e: 0x89, next: 1},
		{s: 0x8b, e: 0x8b, next: 1},
		{s: 0x8d, e: 0x8f, next: 1},
		{s: 0x91, e: 0x92, next: 1},
		{s: 0x94, e: 0x94, next: 1},
		{s: 0x97, e: 0x97, next: 1},
		{s: 0x99, e: 0x99, next: 1},
		{s: 0x9b, e: 0x9b, next: 1},
		{s: 0x9d, e: 0x9d, next: 1},
		{s: 0x9f, e: 0x9f, next: 1},
		{s: 0xa1, e: 0xa2, next: 1},
		{s: 0xa4, e: 0xa4, next: 1},
		{s: 0xa7, e: 0xaa, next: 1},
		{s: 0xac, e: 0xb2, next: 1},
		{s: 0xb4, e: 0xb7, next: 1},
		{s: 0xb9, e: 0xbc, next: 1},
		{s: 0xbe, e: 0xbe, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0x89, next: 1},
		{s: 0x8b, e: 0x9b, next: 1},
		{s: 0xa1, e: 0xa3, next: 1},
		{s: 0xa5, e: 0xa9, next: 1},
		{s: 0xab, e: 0xbb, next: 1}}}, label: 0},
	{table: transTable{a: []trans{
		{s: 0x80, e: 0x96, next: 1}}}, label: 0}}, start: 0}
