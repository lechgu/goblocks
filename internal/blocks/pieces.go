package blocks

var pieceO = ArrFromTemplate([]int{
	0, 0, 0, 0,
	0, 1, 1, 0,
	0, 1, 1, 0,
	0, 0, 0, 0,
}, 4)
var pieceI = ArrFromTemplate([]int{
	0, 0, 0, 0,
	0, 0, 0, 0,
	2, 2, 2, 2,
	0, 0, 0, 0,
}, 4)
var pieceT = ArrFromTemplate([]int{
	0, 0, 0, 0,
	0, 3, 0, 0,
	3, 3, 3, 0,
	0, 0, 0, 0,
}, 4)
var pieceL = ArrFromTemplate([]int{
	0, 0, 0, 0,
	0, 0, 4, 0,
	4, 4, 4, 0,
	0, 0, 0, 0,
}, 4)
var pieceJ = ArrFromTemplate([]int{
	0, 0, 0, 0,
	5, 0, 0, 0,
	5, 5, 5, 0,
	0, 0, 0, 0,
}, 4)
var pieceS = ArrFromTemplate([]int{
	0, 0, 0, 0,
	0, 6, 6, 0,
	6, 6, 0, 0,
	0, 0, 0, 0,
}, 4)
var pieceZ = ArrFromTemplate([]int{
	0, 0, 0, 0,
	7, 7, 0, 0,
	0, 7, 7, 0,
	0, 0, 0, 0,
}, 4)

var pieces = []*Arr{
	pieceI,
	pieceO,
	pieceT,

	pieceL,
	pieceJ,
	pieceS,
	pieceZ,
}
