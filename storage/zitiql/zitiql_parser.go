// Code generated from ZitiQl.g4 by ANTLR 4.7.2. DO NOT EDIT.

package zitiql // ZitiQl
import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 37, 691,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 3, 2, 3, 2, 7, 2, 35,
	10, 2, 12, 2, 14, 2, 38, 11, 2, 3, 2, 3, 2, 7, 2, 42, 10, 2, 12, 2, 14,
	2, 45, 11, 2, 3, 2, 3, 2, 7, 2, 49, 10, 2, 12, 2, 14, 2, 52, 11, 2, 3,
	2, 7, 2, 55, 10, 2, 12, 2, 14, 2, 58, 11, 2, 3, 2, 7, 2, 61, 10, 2, 12,
	2, 14, 2, 64, 11, 2, 3, 2, 3, 2, 3, 3, 3, 3, 7, 3, 70, 10, 3, 12, 3, 14,
	3, 73, 11, 3, 3, 3, 3, 3, 7, 3, 77, 10, 3, 12, 3, 14, 3, 80, 11, 3, 3,
	3, 3, 3, 7, 3, 84, 10, 3, 12, 3, 14, 3, 87, 11, 3, 3, 3, 7, 3, 90, 10,
	3, 12, 3, 14, 3, 93, 11, 3, 3, 3, 7, 3, 96, 10, 3, 12, 3, 14, 3, 99, 11,
	3, 3, 3, 3, 3, 3, 4, 3, 4, 7, 4, 105, 10, 4, 12, 4, 14, 4, 108, 11, 4,
	3, 4, 3, 4, 7, 4, 112, 10, 4, 12, 4, 14, 4, 115, 11, 4, 3, 4, 3, 4, 7,
	4, 119, 10, 4, 12, 4, 14, 4, 122, 11, 4, 3, 4, 7, 4, 125, 10, 4, 12, 4,
	14, 4, 128, 11, 4, 3, 4, 7, 4, 131, 10, 4, 12, 4, 14, 4, 134, 11, 4, 3,
	4, 3, 4, 3, 5, 7, 5, 139, 10, 5, 12, 5, 14, 5, 142, 11, 5, 3, 5, 7, 5,
	145, 10, 5, 12, 5, 14, 5, 148, 11, 5, 3, 5, 7, 5, 151, 10, 5, 12, 5, 14,
	5, 154, 11, 5, 3, 5, 3, 5, 3, 6, 3, 6, 6, 6, 160, 10, 6, 13, 6, 14, 6,
	161, 3, 6, 5, 6, 165, 10, 6, 3, 6, 6, 6, 168, 10, 6, 13, 6, 14, 6, 169,
	3, 6, 5, 6, 173, 10, 6, 3, 6, 6, 6, 176, 10, 6, 13, 6, 14, 6, 177, 3, 6,
	5, 6, 181, 10, 6, 3, 7, 3, 7, 6, 7, 185, 10, 7, 13, 7, 14, 7, 186, 3, 7,
	3, 7, 3, 8, 3, 8, 6, 8, 193, 10, 8, 13, 8, 14, 8, 194, 3, 8, 3, 8, 3, 9,
	3, 9, 6, 9, 201, 10, 9, 13, 9, 14, 9, 202, 3, 9, 3, 9, 6, 9, 207, 10, 9,
	13, 9, 14, 9, 208, 3, 9, 3, 9, 7, 9, 213, 10, 9, 12, 9, 14, 9, 216, 11,
	9, 3, 9, 3, 9, 7, 9, 220, 10, 9, 12, 9, 14, 9, 223, 11, 9, 3, 9, 7, 9,
	226, 10, 9, 12, 9, 14, 9, 229, 11, 9, 3, 10, 3, 10, 6, 10, 233, 10, 10,
	13, 10, 14, 10, 234, 3, 10, 5, 10, 238, 10, 10, 3, 11, 3, 11, 3, 11, 3,
	11, 7, 11, 244, 10, 11, 12, 11, 14, 11, 247, 11, 11, 3, 11, 3, 11, 7, 11,
	251, 10, 11, 12, 11, 14, 11, 254, 11, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3,
	11, 3, 11, 7, 11, 262, 10, 11, 12, 11, 14, 11, 265, 11, 11, 3, 11, 3, 11,
	7, 11, 269, 10, 11, 12, 11, 14, 11, 272, 11, 11, 3, 11, 3, 11, 3, 11, 3,
	11, 3, 11, 6, 11, 279, 10, 11, 13, 11, 14, 11, 280, 3, 11, 5, 11, 284,
	10, 11, 3, 11, 3, 11, 6, 11, 288, 10, 11, 13, 11, 14, 11, 289, 3, 11, 3,
	11, 6, 11, 294, 10, 11, 13, 11, 14, 11, 295, 3, 11, 6, 11, 299, 10, 11,
	13, 11, 14, 11, 300, 3, 11, 3, 11, 6, 11, 305, 10, 11, 13, 11, 14, 11,
	306, 3, 11, 3, 11, 6, 11, 311, 10, 11, 13, 11, 14, 11, 312, 3, 11, 6, 11,
	316, 10, 11, 13, 11, 14, 11, 317, 7, 11, 320, 10, 11, 12, 11, 14, 11, 323,
	11, 11, 3, 12, 3, 12, 6, 12, 327, 10, 12, 13, 12, 14, 12, 328, 3, 12, 3,
	12, 6, 12, 333, 10, 12, 13, 12, 14, 12, 334, 3, 12, 3, 12, 3, 12, 3, 12,
	6, 12, 341, 10, 12, 13, 12, 14, 12, 342, 3, 12, 3, 12, 6, 12, 347, 10,
	12, 13, 12, 14, 12, 348, 3, 12, 3, 12, 3, 12, 3, 12, 6, 12, 355, 10, 12,
	13, 12, 14, 12, 356, 3, 12, 3, 12, 6, 12, 361, 10, 12, 13, 12, 14, 12,
	362, 3, 12, 3, 12, 3, 12, 3, 12, 6, 12, 369, 10, 12, 13, 12, 14, 12, 370,
	3, 12, 3, 12, 6, 12, 375, 10, 12, 13, 12, 14, 12, 376, 3, 12, 3, 12, 6,
	12, 381, 10, 12, 13, 12, 14, 12, 382, 3, 12, 3, 12, 6, 12, 387, 10, 12,
	13, 12, 14, 12, 388, 3, 12, 3, 12, 3, 12, 3, 12, 6, 12, 395, 10, 12, 13,
	12, 14, 12, 396, 3, 12, 3, 12, 6, 12, 401, 10, 12, 13, 12, 14, 12, 402,
	3, 12, 3, 12, 6, 12, 407, 10, 12, 13, 12, 14, 12, 408, 3, 12, 3, 12, 6,
	12, 413, 10, 12, 13, 12, 14, 12, 414, 3, 12, 3, 12, 3, 12, 3, 12, 7, 12,
	421, 10, 12, 12, 12, 14, 12, 424, 11, 12, 3, 12, 3, 12, 7, 12, 428, 10,
	12, 12, 12, 14, 12, 431, 11, 12, 3, 12, 3, 12, 3, 12, 3, 12, 7, 12, 437,
	10, 12, 12, 12, 14, 12, 440, 11, 12, 3, 12, 3, 12, 7, 12, 444, 10, 12,
	12, 12, 14, 12, 447, 11, 12, 3, 12, 3, 12, 3, 12, 3, 12, 7, 12, 453, 10,
	12, 12, 12, 14, 12, 456, 11, 12, 3, 12, 3, 12, 7, 12, 460, 10, 12, 12,
	12, 14, 12, 463, 11, 12, 3, 12, 3, 12, 3, 12, 3, 12, 7, 12, 469, 10, 12,
	12, 12, 14, 12, 472, 11, 12, 3, 12, 3, 12, 7, 12, 476, 10, 12, 12, 12,
	14, 12, 479, 11, 12, 3, 12, 3, 12, 3, 12, 3, 12, 7, 12, 485, 10, 12, 12,
	12, 14, 12, 488, 11, 12, 3, 12, 3, 12, 7, 12, 492, 10, 12, 12, 12, 14,
	12, 495, 11, 12, 3, 12, 3, 12, 3, 12, 3, 12, 7, 12, 501, 10, 12, 12, 12,
	14, 12, 504, 11, 12, 3, 12, 3, 12, 7, 12, 508, 10, 12, 12, 12, 14, 12,
	511, 11, 12, 3, 12, 3, 12, 3, 12, 3, 12, 7, 12, 517, 10, 12, 12, 12, 14,
	12, 520, 11, 12, 3, 12, 3, 12, 7, 12, 524, 10, 12, 12, 12, 14, 12, 527,
	11, 12, 3, 12, 3, 12, 3, 12, 3, 12, 7, 12, 533, 10, 12, 12, 12, 14, 12,
	536, 11, 12, 3, 12, 3, 12, 7, 12, 540, 10, 12, 12, 12, 14, 12, 543, 11,
	12, 3, 12, 3, 12, 3, 12, 3, 12, 7, 12, 549, 10, 12, 12, 12, 14, 12, 552,
	11, 12, 3, 12, 3, 12, 7, 12, 556, 10, 12, 12, 12, 14, 12, 559, 11, 12,
	3, 12, 3, 12, 3, 12, 3, 12, 7, 12, 565, 10, 12, 12, 12, 14, 12, 568, 11,
	12, 3, 12, 3, 12, 7, 12, 572, 10, 12, 12, 12, 14, 12, 575, 11, 12, 3, 12,
	3, 12, 3, 12, 3, 12, 7, 12, 581, 10, 12, 12, 12, 14, 12, 584, 11, 12, 3,
	12, 3, 12, 7, 12, 588, 10, 12, 12, 12, 14, 12, 591, 11, 12, 3, 12, 3, 12,
	3, 12, 3, 12, 7, 12, 597, 10, 12, 12, 12, 14, 12, 600, 11, 12, 3, 12, 3,
	12, 6, 12, 604, 10, 12, 13, 12, 14, 12, 605, 3, 12, 3, 12, 5, 12, 610,
	10, 12, 3, 13, 3, 13, 5, 13, 614, 10, 13, 3, 14, 3, 14, 3, 14, 7, 14, 619,
	10, 14, 12, 14, 14, 14, 622, 11, 14, 3, 14, 3, 14, 7, 14, 626, 10, 14,
	12, 14, 14, 14, 629, 11, 14, 3, 14, 3, 14, 3, 14, 3, 14, 7, 14, 635, 10,
	14, 12, 14, 14, 14, 638, 11, 14, 3, 14, 3, 14, 7, 14, 642, 10, 14, 12,
	14, 14, 14, 645, 11, 14, 3, 14, 3, 14, 3, 14, 3, 14, 7, 14, 651, 10, 14,
	12, 14, 14, 14, 654, 11, 14, 3, 14, 3, 14, 7, 14, 658, 10, 14, 12, 14,
	14, 14, 661, 11, 14, 3, 14, 3, 14, 5, 14, 665, 10, 14, 3, 15, 3, 15, 5,
	15, 669, 10, 15, 3, 16, 3, 16, 6, 16, 673, 10, 16, 13, 16, 14, 16, 674,
	3, 16, 3, 16, 6, 16, 679, 10, 16, 13, 16, 14, 16, 680, 3, 16, 3, 16, 6,
	16, 685, 10, 16, 13, 16, 14, 16, 686, 3, 16, 3, 16, 3, 16, 2, 3, 20, 17,
	2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 2, 5, 4, 2, 24,
	24, 33, 33, 3, 2, 27, 28, 3, 2, 23, 24, 2, 793, 2, 32, 3, 2, 2, 2, 4, 67,
	3, 2, 2, 2, 6, 102, 3, 2, 2, 2, 8, 140, 3, 2, 2, 2, 10, 157, 3, 2, 2, 2,
	12, 182, 3, 2, 2, 2, 14, 190, 3, 2, 2, 2, 16, 198, 3, 2, 2, 2, 18, 230,
	3, 2, 2, 2, 20, 283, 3, 2, 2, 2, 22, 609, 3, 2, 2, 2, 24, 613, 3, 2, 2,
	2, 26, 664, 3, 2, 2, 2, 28, 668, 3, 2, 2, 2, 30, 670, 3, 2, 2, 2, 32, 36,
	7, 7, 2, 2, 33, 35, 7, 4, 2, 2, 34, 33, 3, 2, 2, 2, 35, 38, 3, 2, 2, 2,
	36, 34, 3, 2, 2, 2, 36, 37, 3, 2, 2, 2, 37, 39, 3, 2, 2, 2, 38, 36, 3,
	2, 2, 2, 39, 56, 7, 23, 2, 2, 40, 42, 7, 4, 2, 2, 41, 40, 3, 2, 2, 2, 42,
	45, 3, 2, 2, 2, 43, 41, 3, 2, 2, 2, 43, 44, 3, 2, 2, 2, 44, 46, 3, 2, 2,
	2, 45, 43, 3, 2, 2, 2, 46, 50, 7, 3, 2, 2, 47, 49, 7, 4, 2, 2, 48, 47,
	3, 2, 2, 2, 49, 52, 3, 2, 2, 2, 50, 48, 3, 2, 2, 2, 50, 51, 3, 2, 2, 2,
	51, 53, 3, 2, 2, 2, 52, 50, 3, 2, 2, 2, 53, 55, 7, 23, 2, 2, 54, 43, 3,
	2, 2, 2, 55, 58, 3, 2, 2, 2, 56, 54, 3, 2, 2, 2, 56, 57, 3, 2, 2, 2, 57,
	62, 3, 2, 2, 2, 58, 56, 3, 2, 2, 2, 59, 61, 7, 4, 2, 2, 60, 59, 3, 2, 2,
	2, 61, 64, 3, 2, 2, 2, 62, 60, 3, 2, 2, 2, 62, 63, 3, 2, 2, 2, 63, 65,
	3, 2, 2, 2, 64, 62, 3, 2, 2, 2, 65, 66, 7, 8, 2, 2, 66, 3, 3, 2, 2, 2,
	67, 71, 7, 7, 2, 2, 68, 70, 7, 4, 2, 2, 69, 68, 3, 2, 2, 2, 70, 73, 3,
	2, 2, 2, 71, 69, 3, 2, 2, 2, 71, 72, 3, 2, 2, 2, 72, 74, 3, 2, 2, 2, 73,
	71, 3, 2, 2, 2, 74, 91, 7, 24, 2, 2, 75, 77, 7, 4, 2, 2, 76, 75, 3, 2,
	2, 2, 77, 80, 3, 2, 2, 2, 78, 76, 3, 2, 2, 2, 78, 79, 3, 2, 2, 2, 79, 81,
	3, 2, 2, 2, 80, 78, 3, 2, 2, 2, 81, 85, 7, 3, 2, 2, 82, 84, 7, 4, 2, 2,
	83, 82, 3, 2, 2, 2, 84, 87, 3, 2, 2, 2, 85, 83, 3, 2, 2, 2, 85, 86, 3,
	2, 2, 2, 86, 88, 3, 2, 2, 2, 87, 85, 3, 2, 2, 2, 88, 90, 7, 24, 2, 2, 89,
	78, 3, 2, 2, 2, 90, 93, 3, 2, 2, 2, 91, 89, 3, 2, 2, 2, 91, 92, 3, 2, 2,
	2, 92, 97, 3, 2, 2, 2, 93, 91, 3, 2, 2, 2, 94, 96, 7, 4, 2, 2, 95, 94,
	3, 2, 2, 2, 96, 99, 3, 2, 2, 2, 97, 95, 3, 2, 2, 2, 97, 98, 3, 2, 2, 2,
	98, 100, 3, 2, 2, 2, 99, 97, 3, 2, 2, 2, 100, 101, 7, 8, 2, 2, 101, 5,
	3, 2, 2, 2, 102, 106, 7, 7, 2, 2, 103, 105, 7, 4, 2, 2, 104, 103, 3, 2,
	2, 2, 105, 108, 3, 2, 2, 2, 106, 104, 3, 2, 2, 2, 106, 107, 3, 2, 2, 2,
	107, 109, 3, 2, 2, 2, 108, 106, 3, 2, 2, 2, 109, 126, 7, 18, 2, 2, 110,
	112, 7, 4, 2, 2, 111, 110, 3, 2, 2, 2, 112, 115, 3, 2, 2, 2, 113, 111,
	3, 2, 2, 2, 113, 114, 3, 2, 2, 2, 114, 116, 3, 2, 2, 2, 115, 113, 3, 2,
	2, 2, 116, 120, 7, 3, 2, 2, 117, 119, 7, 4, 2, 2, 118, 117, 3, 2, 2, 2,
	119, 122, 3, 2, 2, 2, 120, 118, 3, 2, 2, 2, 120, 121, 3, 2, 2, 2, 121,
	123, 3, 2, 2, 2, 122, 120, 3, 2, 2, 2, 123, 125, 7, 18, 2, 2, 124, 113,
	3, 2, 2, 2, 125, 128, 3, 2, 2, 2, 126, 124, 3, 2, 2, 2, 126, 127, 3, 2,
	2, 2, 127, 132, 3, 2, 2, 2, 128, 126, 3, 2, 2, 2, 129, 131, 7, 4, 2, 2,
	130, 129, 3, 2, 2, 2, 131, 134, 3, 2, 2, 2, 132, 130, 3, 2, 2, 2, 132,
	133, 3, 2, 2, 2, 133, 135, 3, 2, 2, 2, 134, 132, 3, 2, 2, 2, 135, 136,
	7, 8, 2, 2, 136, 7, 3, 2, 2, 2, 137, 139, 7, 4, 2, 2, 138, 137, 3, 2, 2,
	2, 139, 142, 3, 2, 2, 2, 140, 138, 3, 2, 2, 2, 140, 141, 3, 2, 2, 2, 141,
	146, 3, 2, 2, 2, 142, 140, 3, 2, 2, 2, 143, 145, 5, 10, 6, 2, 144, 143,
	3, 2, 2, 2, 145, 148, 3, 2, 2, 2, 146, 144, 3, 2, 2, 2, 146, 147, 3, 2,
	2, 2, 147, 152, 3, 2, 2, 2, 148, 146, 3, 2, 2, 2, 149, 151, 7, 4, 2, 2,
	150, 149, 3, 2, 2, 2, 151, 154, 3, 2, 2, 2, 152, 150, 3, 2, 2, 2, 152,
	153, 3, 2, 2, 2, 153, 155, 3, 2, 2, 2, 154, 152, 3, 2, 2, 2, 155, 156,
	7, 2, 2, 3, 156, 9, 3, 2, 2, 2, 157, 164, 5, 20, 11, 2, 158, 160, 7, 4,
	2, 2, 159, 158, 3, 2, 2, 2, 160, 161, 3, 2, 2, 2, 161, 159, 3, 2, 2, 2,
	161, 162, 3, 2, 2, 2, 162, 163, 3, 2, 2, 2, 163, 165, 5, 16, 9, 2, 164,
	159, 3, 2, 2, 2, 164, 165, 3, 2, 2, 2, 165, 172, 3, 2, 2, 2, 166, 168,
	7, 4, 2, 2, 167, 166, 3, 2, 2, 2, 168, 169, 3, 2, 2, 2, 169, 167, 3, 2,
	2, 2, 169, 170, 3, 2, 2, 2, 170, 171, 3, 2, 2, 2, 171, 173, 5, 12, 7, 2,
	172, 167, 3, 2, 2, 2, 172, 173, 3, 2, 2, 2, 173, 180, 3, 2, 2, 2, 174,
	176, 7, 4, 2, 2, 175, 174, 3, 2, 2, 2, 176, 177, 3, 2, 2, 2, 177, 175,
	3, 2, 2, 2, 177, 178, 3, 2, 2, 2, 178, 179, 3, 2, 2, 2, 179, 181, 5, 14,
	8, 2, 180, 175, 3, 2, 2, 2, 180, 181, 3, 2, 2, 2, 181, 11, 3, 2, 2, 2,
	182, 184, 7, 31, 2, 2, 183, 185, 7, 4, 2, 2, 184, 183, 3, 2, 2, 2, 185,
	186, 3, 2, 2, 2, 186, 184, 3, 2, 2, 2, 186, 187, 3, 2, 2, 2, 187, 188,
	3, 2, 2, 2, 188, 189, 7, 24, 2, 2, 189, 13, 3, 2, 2, 2, 190, 192, 7, 32,
	2, 2, 191, 193, 7, 4, 2, 2, 192, 191, 3, 2, 2, 2, 193, 194, 3, 2, 2, 2,
	194, 192, 3, 2, 2, 2, 194, 195, 3, 2, 2, 2, 195, 196, 3, 2, 2, 2, 196,
	197, 9, 2, 2, 2, 197, 15, 3, 2, 2, 2, 198, 200, 7, 29, 2, 2, 199, 201,
	7, 4, 2, 2, 200, 199, 3, 2, 2, 2, 201, 202, 3, 2, 2, 2, 202, 200, 3, 2,
	2, 2, 202, 203, 3, 2, 2, 2, 203, 204, 3, 2, 2, 2, 204, 206, 7, 30, 2, 2,
	205, 207, 7, 4, 2, 2, 206, 205, 3, 2, 2, 2, 207, 208, 3, 2, 2, 2, 208,
	206, 3, 2, 2, 2, 208, 209, 3, 2, 2, 2, 209, 210, 3, 2, 2, 2, 210, 227,
	5, 18, 10, 2, 211, 213, 7, 4, 2, 2, 212, 211, 3, 2, 2, 2, 213, 216, 3,
	2, 2, 2, 214, 212, 3, 2, 2, 2, 214, 215, 3, 2, 2, 2, 215, 217, 3, 2, 2,
	2, 216, 214, 3, 2, 2, 2, 217, 221, 7, 3, 2, 2, 218, 220, 7, 4, 2, 2, 219,
	218, 3, 2, 2, 2, 220, 223, 3, 2, 2, 2, 221, 219, 3, 2, 2, 2, 221, 222,
	3, 2, 2, 2, 222, 224, 3, 2, 2, 2, 223, 221, 3, 2, 2, 2, 224, 226, 5, 18,
	10, 2, 225, 214, 3, 2, 2, 2, 226, 229, 3, 2, 2, 2, 227, 225, 3, 2, 2, 2,
	227, 228, 3, 2, 2, 2, 228, 17, 3, 2, 2, 2, 229, 227, 3, 2, 2, 2, 230, 237,
	7, 36, 2, 2, 231, 233, 7, 4, 2, 2, 232, 231, 3, 2, 2, 2, 233, 234, 3, 2,
	2, 2, 234, 232, 3, 2, 2, 2, 234, 235, 3, 2, 2, 2, 235, 236, 3, 2, 2, 2,
	236, 238, 9, 3, 2, 2, 237, 232, 3, 2, 2, 2, 237, 238, 3, 2, 2, 2, 238,
	19, 3, 2, 2, 2, 239, 240, 8, 11, 1, 2, 240, 284, 5, 22, 12, 2, 241, 245,
	7, 5, 2, 2, 242, 244, 7, 4, 2, 2, 243, 242, 3, 2, 2, 2, 244, 247, 3, 2,
	2, 2, 245, 243, 3, 2, 2, 2, 245, 246, 3, 2, 2, 2, 246, 248, 3, 2, 2, 2,
	247, 245, 3, 2, 2, 2, 248, 252, 5, 20, 11, 2, 249, 251, 7, 4, 2, 2, 250,
	249, 3, 2, 2, 2, 251, 254, 3, 2, 2, 2, 252, 250, 3, 2, 2, 2, 252, 253,
	3, 2, 2, 2, 253, 255, 3, 2, 2, 2, 254, 252, 3, 2, 2, 2, 255, 256, 7, 6,
	2, 2, 256, 284, 3, 2, 2, 2, 257, 284, 7, 17, 2, 2, 258, 259, 7, 22, 2,
	2, 259, 263, 7, 5, 2, 2, 260, 262, 7, 4, 2, 2, 261, 260, 3, 2, 2, 2, 262,
	265, 3, 2, 2, 2, 263, 261, 3, 2, 2, 2, 263, 264, 3, 2, 2, 2, 264, 266,
	3, 2, 2, 2, 265, 263, 3, 2, 2, 2, 266, 270, 5, 28, 15, 2, 267, 269, 7,
	4, 2, 2, 268, 267, 3, 2, 2, 2, 269, 272, 3, 2, 2, 2, 270, 268, 3, 2, 2,
	2, 270, 271, 3, 2, 2, 2, 271, 273, 3, 2, 2, 2, 272, 270, 3, 2, 2, 2, 273,
	274, 7, 6, 2, 2, 274, 284, 3, 2, 2, 2, 275, 284, 7, 36, 2, 2, 276, 278,
	7, 26, 2, 2, 277, 279, 7, 4, 2, 2, 278, 277, 3, 2, 2, 2, 279, 280, 3, 2,
	2, 2, 280, 278, 3, 2, 2, 2, 280, 281, 3, 2, 2, 2, 281, 282, 3, 2, 2, 2,
	282, 284, 5, 20, 11, 3, 283, 239, 3, 2, 2, 2, 283, 241, 3, 2, 2, 2, 283,
	257, 3, 2, 2, 2, 283, 258, 3, 2, 2, 2, 283, 275, 3, 2, 2, 2, 283, 276,
	3, 2, 2, 2, 284, 321, 3, 2, 2, 2, 285, 298, 12, 8, 2, 2, 286, 288, 7, 4,
	2, 2, 287, 286, 3, 2, 2, 2, 288, 289, 3, 2, 2, 2, 289, 287, 3, 2, 2, 2,
	289, 290, 3, 2, 2, 2, 290, 291, 3, 2, 2, 2, 291, 293, 7, 9, 2, 2, 292,
	294, 7, 4, 2, 2, 293, 292, 3, 2, 2, 2, 294, 295, 3, 2, 2, 2, 295, 293,
	3, 2, 2, 2, 295, 296, 3, 2, 2, 2, 296, 297, 3, 2, 2, 2, 297, 299, 5, 20,
	11, 2, 298, 287, 3, 2, 2, 2, 299, 300, 3, 2, 2, 2, 300, 298, 3, 2, 2, 2,
	300, 301, 3, 2, 2, 2, 301, 320, 3, 2, 2, 2, 302, 315, 12, 7, 2, 2, 303,
	305, 7, 4, 2, 2, 304, 303, 3, 2, 2, 2, 305, 306, 3, 2, 2, 2, 306, 304,
	3, 2, 2, 2, 306, 307, 3, 2, 2, 2, 307, 308, 3, 2, 2, 2, 308, 310, 7, 10,
	2, 2, 309, 311, 7, 4, 2, 2, 310, 309, 3, 2, 2, 2, 311, 312, 3, 2, 2, 2,
	312, 310, 3, 2, 2, 2, 312, 313, 3, 2, 2, 2, 313, 314, 3, 2, 2, 2, 314,
	316, 5, 20, 11, 2, 315, 304, 3, 2, 2, 2, 316, 317, 3, 2, 2, 2, 317, 315,
	3, 2, 2, 2, 317, 318, 3, 2, 2, 2, 318, 320, 3, 2, 2, 2, 319, 285, 3, 2,
	2, 2, 319, 302, 3, 2, 2, 2, 320, 323, 3, 2, 2, 2, 321, 319, 3, 2, 2, 2,
	321, 322, 3, 2, 2, 2, 322, 21, 3, 2, 2, 2, 323, 321, 3, 2, 2, 2, 324, 326,
	5, 24, 13, 2, 325, 327, 7, 4, 2, 2, 326, 325, 3, 2, 2, 2, 327, 328, 3,
	2, 2, 2, 328, 326, 3, 2, 2, 2, 328, 329, 3, 2, 2, 2, 329, 330, 3, 2, 2,
	2, 330, 332, 7, 15, 2, 2, 331, 333, 7, 4, 2, 2, 332, 331, 3, 2, 2, 2, 333,
	334, 3, 2, 2, 2, 334, 332, 3, 2, 2, 2, 334, 335, 3, 2, 2, 2, 335, 336,
	3, 2, 2, 2, 336, 337, 5, 2, 2, 2, 337, 610, 3, 2, 2, 2, 338, 340, 5, 24,
	13, 2, 339, 341, 7, 4, 2, 2, 340, 339, 3, 2, 2, 2, 341, 342, 3, 2, 2, 2,
	342, 340, 3, 2, 2, 2, 342, 343, 3, 2, 2, 2, 343, 344, 3, 2, 2, 2, 344,
	346, 7, 15, 2, 2, 345, 347, 7, 4, 2, 2, 346, 345, 3, 2, 2, 2, 347, 348,
	3, 2, 2, 2, 348, 346, 3, 2, 2, 2, 348, 349, 3, 2, 2, 2, 349, 350, 3, 2,
	2, 2, 350, 351, 5, 4, 3, 2, 351, 610, 3, 2, 2, 2, 352, 354, 5, 24, 13,
	2, 353, 355, 7, 4, 2, 2, 354, 353, 3, 2, 2, 2, 355, 356, 3, 2, 2, 2, 356,
	354, 3, 2, 2, 2, 356, 357, 3, 2, 2, 2, 357, 358, 3, 2, 2, 2, 358, 360,
	7, 15, 2, 2, 359, 361, 7, 4, 2, 2, 360, 359, 3, 2, 2, 2, 361, 362, 3, 2,
	2, 2, 362, 360, 3, 2, 2, 2, 362, 363, 3, 2, 2, 2, 363, 364, 3, 2, 2, 2,
	364, 365, 5, 6, 4, 2, 365, 610, 3, 2, 2, 2, 366, 368, 5, 24, 13, 2, 367,
	369, 7, 4, 2, 2, 368, 367, 3, 2, 2, 2, 369, 370, 3, 2, 2, 2, 370, 368,
	3, 2, 2, 2, 370, 371, 3, 2, 2, 2, 371, 372, 3, 2, 2, 2, 372, 374, 7, 16,
	2, 2, 373, 375, 7, 4, 2, 2, 374, 373, 3, 2, 2, 2, 375, 376, 3, 2, 2, 2,
	376, 374, 3, 2, 2, 2, 376, 377, 3, 2, 2, 2, 377, 378, 3, 2, 2, 2, 378,
	380, 7, 24, 2, 2, 379, 381, 7, 4, 2, 2, 380, 379, 3, 2, 2, 2, 381, 382,
	3, 2, 2, 2, 382, 380, 3, 2, 2, 2, 382, 383, 3, 2, 2, 2, 383, 384, 3, 2,
	2, 2, 384, 386, 7, 9, 2, 2, 385, 387, 7, 4, 2, 2, 386, 385, 3, 2, 2, 2,
	387, 388, 3, 2, 2, 2, 388, 386, 3, 2, 2, 2, 388, 389, 3, 2, 2, 2, 389,
	390, 3, 2, 2, 2, 390, 391, 7, 24, 2, 2, 391, 610, 3, 2, 2, 2, 392, 394,
	5, 24, 13, 2, 393, 395, 7, 4, 2, 2, 394, 393, 3, 2, 2, 2, 395, 396, 3,
	2, 2, 2, 396, 394, 3, 2, 2, 2, 396, 397, 3, 2, 2, 2, 397, 398, 3, 2, 2,
	2, 398, 400, 7, 16, 2, 2, 399, 401, 7, 4, 2, 2, 400, 399, 3, 2, 2, 2, 401,
	402, 3, 2, 2, 2, 402, 400, 3, 2, 2, 2, 402, 403, 3, 2, 2, 2, 403, 404,
	3, 2, 2, 2, 404, 406, 7, 18, 2, 2, 405, 407, 7, 4, 2, 2, 406, 405, 3, 2,
	2, 2, 407, 408, 3, 2, 2, 2, 408, 406, 3, 2, 2, 2, 408, 409, 3, 2, 2, 2,
	409, 410, 3, 2, 2, 2, 410, 412, 7, 9, 2, 2, 411, 413, 7, 4, 2, 2, 412,
	411, 3, 2, 2, 2, 413, 414, 3, 2, 2, 2, 414, 412, 3, 2, 2, 2, 414, 415,
	3, 2, 2, 2, 415, 416, 3, 2, 2, 2, 416, 417, 7, 18, 2, 2, 417, 610, 3, 2,
	2, 2, 418, 422, 5, 24, 13, 2, 419, 421, 7, 4, 2, 2, 420, 419, 3, 2, 2,
	2, 421, 424, 3, 2, 2, 2, 422, 420, 3, 2, 2, 2, 422, 423, 3, 2, 2, 2, 423,
	425, 3, 2, 2, 2, 424, 422, 3, 2, 2, 2, 425, 429, 7, 11, 2, 2, 426, 428,
	7, 4, 2, 2, 427, 426, 3, 2, 2, 2, 428, 431, 3, 2, 2, 2, 429, 427, 3, 2,
	2, 2, 429, 430, 3, 2, 2, 2, 430, 432, 3, 2, 2, 2, 431, 429, 3, 2, 2, 2,
	432, 433, 7, 23, 2, 2, 433, 610, 3, 2, 2, 2, 434, 438, 5, 24, 13, 2, 435,
	437, 7, 4, 2, 2, 436, 435, 3, 2, 2, 2, 437, 440, 3, 2, 2, 2, 438, 436,
	3, 2, 2, 2, 438, 439, 3, 2, 2, 2, 439, 441, 3, 2, 2, 2, 440, 438, 3, 2,
	2, 2, 441, 445, 7, 11, 2, 2, 442, 444, 7, 4, 2, 2, 443, 442, 3, 2, 2, 2,
	444, 447, 3, 2, 2, 2, 445, 443, 3, 2, 2, 2, 445, 446, 3, 2, 2, 2, 446,
	448, 3, 2, 2, 2, 447, 445, 3, 2, 2, 2, 448, 449, 7, 24, 2, 2, 449, 610,
	3, 2, 2, 2, 450, 454, 5, 24, 13, 2, 451, 453, 7, 4, 2, 2, 452, 451, 3,
	2, 2, 2, 453, 456, 3, 2, 2, 2, 454, 452, 3, 2, 2, 2, 454, 455, 3, 2, 2,
	2, 455, 457, 3, 2, 2, 2, 456, 454, 3, 2, 2, 2, 457, 461, 7, 11, 2, 2, 458,
	460, 7, 4, 2, 2, 459, 458, 3, 2, 2, 2, 460, 463, 3, 2, 2, 2, 461, 459,
	3, 2, 2, 2, 461, 462, 3, 2, 2, 2, 462, 464, 3, 2, 2, 2, 463, 461, 3, 2,
	2, 2, 464, 465, 7, 18, 2, 2, 465, 610, 3, 2, 2, 2, 466, 470, 5, 24, 13,
	2, 467, 469, 7, 4, 2, 2, 468, 467, 3, 2, 2, 2, 469, 472, 3, 2, 2, 2, 470,
	468, 3, 2, 2, 2, 470, 471, 3, 2, 2, 2, 471, 473, 3, 2, 2, 2, 472, 470,
	3, 2, 2, 2, 473, 477, 7, 12, 2, 2, 474, 476, 7, 4, 2, 2, 475, 474, 3, 2,
	2, 2, 476, 479, 3, 2, 2, 2, 477, 475, 3, 2, 2, 2, 477, 478, 3, 2, 2, 2,
	478, 480, 3, 2, 2, 2, 479, 477, 3, 2, 2, 2, 480, 481, 7, 23, 2, 2, 481,
	610, 3, 2, 2, 2, 482, 486, 5, 24, 13, 2, 483, 485, 7, 4, 2, 2, 484, 483,
	3, 2, 2, 2, 485, 488, 3, 2, 2, 2, 486, 484, 3, 2, 2, 2, 486, 487, 3, 2,
	2, 2, 487, 489, 3, 2, 2, 2, 488, 486, 3, 2, 2, 2, 489, 493, 7, 12, 2, 2,
	490, 492, 7, 4, 2, 2, 491, 490, 3, 2, 2, 2, 492, 495, 3, 2, 2, 2, 493,
	491, 3, 2, 2, 2, 493, 494, 3, 2, 2, 2, 494, 496, 3, 2, 2, 2, 495, 493,
	3, 2, 2, 2, 496, 497, 7, 24, 2, 2, 497, 610, 3, 2, 2, 2, 498, 502, 5, 24,
	13, 2, 499, 501, 7, 4, 2, 2, 500, 499, 3, 2, 2, 2, 501, 504, 3, 2, 2, 2,
	502, 500, 3, 2, 2, 2, 502, 503, 3, 2, 2, 2, 503, 505, 3, 2, 2, 2, 504,
	502, 3, 2, 2, 2, 505, 509, 7, 12, 2, 2, 506, 508, 7, 4, 2, 2, 507, 506,
	3, 2, 2, 2, 508, 511, 3, 2, 2, 2, 509, 507, 3, 2, 2, 2, 509, 510, 3, 2,
	2, 2, 510, 512, 3, 2, 2, 2, 511, 509, 3, 2, 2, 2, 512, 513, 7, 18, 2, 2,
	513, 610, 3, 2, 2, 2, 514, 518, 5, 24, 13, 2, 515, 517, 7, 4, 2, 2, 516,
	515, 3, 2, 2, 2, 517, 520, 3, 2, 2, 2, 518, 516, 3, 2, 2, 2, 518, 519,
	3, 2, 2, 2, 519, 521, 3, 2, 2, 2, 520, 518, 3, 2, 2, 2, 521, 525, 7, 13,
	2, 2, 522, 524, 7, 4, 2, 2, 523, 522, 3, 2, 2, 2, 524, 527, 3, 2, 2, 2,
	525, 523, 3, 2, 2, 2, 525, 526, 3, 2, 2, 2, 526, 528, 3, 2, 2, 2, 527,
	525, 3, 2, 2, 2, 528, 529, 7, 23, 2, 2, 529, 610, 3, 2, 2, 2, 530, 534,
	5, 24, 13, 2, 531, 533, 7, 4, 2, 2, 532, 531, 3, 2, 2, 2, 533, 536, 3,
	2, 2, 2, 534, 532, 3, 2, 2, 2, 534, 535, 3, 2, 2, 2, 535, 537, 3, 2, 2,
	2, 536, 534, 3, 2, 2, 2, 537, 541, 7, 13, 2, 2, 538, 540, 7, 4, 2, 2, 539,
	538, 3, 2, 2, 2, 540, 543, 3, 2, 2, 2, 541, 539, 3, 2, 2, 2, 541, 542,
	3, 2, 2, 2, 542, 544, 3, 2, 2, 2, 543, 541, 3, 2, 2, 2, 544, 545, 7, 24,
	2, 2, 545, 610, 3, 2, 2, 2, 546, 550, 5, 24, 13, 2, 547, 549, 7, 4, 2,
	2, 548, 547, 3, 2, 2, 2, 549, 552, 3, 2, 2, 2, 550, 548, 3, 2, 2, 2, 550,
	551, 3, 2, 2, 2, 551, 553, 3, 2, 2, 2, 552, 550, 3, 2, 2, 2, 553, 557,
	7, 13, 2, 2, 554, 556, 7, 4, 2, 2, 555, 554, 3, 2, 2, 2, 556, 559, 3, 2,
	2, 2, 557, 555, 3, 2, 2, 2, 557, 558, 3, 2, 2, 2, 558, 560, 3, 2, 2, 2,
	559, 557, 3, 2, 2, 2, 560, 561, 7, 18, 2, 2, 561, 610, 3, 2, 2, 2, 562,
	566, 5, 24, 13, 2, 563, 565, 7, 4, 2, 2, 564, 563, 3, 2, 2, 2, 565, 568,
	3, 2, 2, 2, 566, 564, 3, 2, 2, 2, 566, 567, 3, 2, 2, 2, 567, 569, 3, 2,
	2, 2, 568, 566, 3, 2, 2, 2, 569, 573, 7, 13, 2, 2, 570, 572, 7, 4, 2, 2,
	571, 570, 3, 2, 2, 2, 572, 575, 3, 2, 2, 2, 573, 571, 3, 2, 2, 2, 573,
	574, 3, 2, 2, 2, 574, 576, 3, 2, 2, 2, 575, 573, 3, 2, 2, 2, 576, 577,
	7, 17, 2, 2, 577, 610, 3, 2, 2, 2, 578, 582, 5, 24, 13, 2, 579, 581, 7,
	4, 2, 2, 580, 579, 3, 2, 2, 2, 581, 584, 3, 2, 2, 2, 582, 580, 3, 2, 2,
	2, 582, 583, 3, 2, 2, 2, 583, 585, 3, 2, 2, 2, 584, 582, 3, 2, 2, 2, 585,
	589, 7, 13, 2, 2, 586, 588, 7, 4, 2, 2, 587, 586, 3, 2, 2, 2, 588, 591,
	3, 2, 2, 2, 589, 587, 3, 2, 2, 2, 589, 590, 3, 2, 2, 2, 590, 592, 3, 2,
	2, 2, 591, 589, 3, 2, 2, 2, 592, 593, 7, 25, 2, 2, 593, 610, 3, 2, 2, 2,
	594, 598, 5, 24, 13, 2, 595, 597, 7, 4, 2, 2, 596, 595, 3, 2, 2, 2, 597,
	600, 3, 2, 2, 2, 598, 596, 3, 2, 2, 2, 598, 599, 3, 2, 2, 2, 599, 601,
	3, 2, 2, 2, 600, 598, 3, 2, 2, 2, 601, 603, 7, 14, 2, 2, 602, 604, 7, 4,
	2, 2, 603, 602, 3, 2, 2, 2, 604, 605, 3, 2, 2, 2, 605, 603, 3, 2, 2, 2,
	605, 606, 3, 2, 2, 2, 606, 607, 3, 2, 2, 2, 607, 608, 9, 4, 2, 2, 608,
	610, 3, 2, 2, 2, 609, 324, 3, 2, 2, 2, 609, 338, 3, 2, 2, 2, 609, 352,
	3, 2, 2, 2, 609, 366, 3, 2, 2, 2, 609, 392, 3, 2, 2, 2, 609, 418, 3, 2,
	2, 2, 609, 434, 3, 2, 2, 2, 609, 450, 3, 2, 2, 2, 609, 466, 3, 2, 2, 2,
	609, 482, 3, 2, 2, 2, 609, 498, 3, 2, 2, 2, 609, 514, 3, 2, 2, 2, 609,
	530, 3, 2, 2, 2, 609, 546, 3, 2, 2, 2, 609, 562, 3, 2, 2, 2, 609, 578,
	3, 2, 2, 2, 609, 594, 3, 2, 2, 2, 610, 23, 3, 2, 2, 2, 611, 614, 7, 36,
	2, 2, 612, 614, 5, 26, 14, 2, 613, 611, 3, 2, 2, 2, 613, 612, 3, 2, 2,
	2, 614, 25, 3, 2, 2, 2, 615, 616, 7, 19, 2, 2, 616, 620, 7, 5, 2, 2, 617,
	619, 7, 4, 2, 2, 618, 617, 3, 2, 2, 2, 619, 622, 3, 2, 2, 2, 620, 618,
	3, 2, 2, 2, 620, 621, 3, 2, 2, 2, 621, 623, 3, 2, 2, 2, 622, 620, 3, 2,
	2, 2, 623, 627, 7, 36, 2, 2, 624, 626, 7, 4, 2, 2, 625, 624, 3, 2, 2, 2,
	626, 629, 3, 2, 2, 2, 627, 625, 3, 2, 2, 2, 627, 628, 3, 2, 2, 2, 628,
	630, 3, 2, 2, 2, 629, 627, 3, 2, 2, 2, 630, 665, 7, 6, 2, 2, 631, 632,
	7, 20, 2, 2, 632, 636, 7, 5, 2, 2, 633, 635, 7, 4, 2, 2, 634, 633, 3, 2,
	2, 2, 635, 638, 3, 2, 2, 2, 636, 634, 3, 2, 2, 2, 636, 637, 3, 2, 2, 2,
	637, 639, 3, 2, 2, 2, 638, 636, 3, 2, 2, 2, 639, 643, 7, 36, 2, 2, 640,
	642, 7, 4, 2, 2, 641, 640, 3, 2, 2, 2, 642, 645, 3, 2, 2, 2, 643, 641,
	3, 2, 2, 2, 643, 644, 3, 2, 2, 2, 644, 646, 3, 2, 2, 2, 645, 643, 3, 2,
	2, 2, 646, 665, 7, 6, 2, 2, 647, 648, 7, 21, 2, 2, 648, 652, 7, 5, 2, 2,
	649, 651, 7, 4, 2, 2, 650, 649, 3, 2, 2, 2, 651, 654, 3, 2, 2, 2, 652,
	650, 3, 2, 2, 2, 652, 653, 3, 2, 2, 2, 653, 655, 3, 2, 2, 2, 654, 652,
	3, 2, 2, 2, 655, 659, 5, 28, 15, 2, 656, 658, 7, 4, 2, 2, 657, 656, 3,
	2, 2, 2, 658, 661, 3, 2, 2, 2, 659, 657, 3, 2, 2, 2, 659, 660, 3, 2, 2,
	2, 660, 662, 3, 2, 2, 2, 661, 659, 3, 2, 2, 2, 662, 663, 7, 6, 2, 2, 663,
	665, 3, 2, 2, 2, 664, 615, 3, 2, 2, 2, 664, 631, 3, 2, 2, 2, 664, 647,
	3, 2, 2, 2, 665, 27, 3, 2, 2, 2, 666, 669, 7, 36, 2, 2, 667, 669, 5, 30,
	16, 2, 668, 666, 3, 2, 2, 2, 668, 667, 3, 2, 2, 2, 669, 29, 3, 2, 2, 2,
	670, 672, 7, 35, 2, 2, 671, 673, 7, 4, 2, 2, 672, 671, 3, 2, 2, 2, 673,
	674, 3, 2, 2, 2, 674, 672, 3, 2, 2, 2, 674, 675, 3, 2, 2, 2, 675, 676,
	3, 2, 2, 2, 676, 678, 7, 36, 2, 2, 677, 679, 7, 4, 2, 2, 678, 677, 3, 2,
	2, 2, 679, 680, 3, 2, 2, 2, 680, 678, 3, 2, 2, 2, 680, 681, 3, 2, 2, 2,
	681, 682, 3, 2, 2, 2, 682, 684, 7, 34, 2, 2, 683, 685, 7, 4, 2, 2, 684,
	683, 3, 2, 2, 2, 685, 686, 3, 2, 2, 2, 686, 684, 3, 2, 2, 2, 686, 687,
	3, 2, 2, 2, 687, 688, 3, 2, 2, 2, 688, 689, 5, 10, 6, 2, 689, 31, 3, 2,
	2, 2, 100, 36, 43, 50, 56, 62, 71, 78, 85, 91, 97, 106, 113, 120, 126,
	132, 140, 146, 152, 161, 164, 169, 172, 177, 180, 186, 194, 202, 208, 214,
	221, 227, 234, 237, 245, 252, 263, 270, 280, 283, 289, 295, 300, 306, 312,
	317, 319, 321, 328, 334, 342, 348, 356, 362, 370, 376, 382, 388, 396, 402,
	408, 414, 422, 429, 438, 445, 454, 461, 470, 477, 486, 493, 502, 509, 518,
	525, 534, 541, 550, 557, 566, 573, 582, 589, 598, 605, 609, 613, 620, 627,
	636, 643, 652, 659, 664, 668, 674, 680, 686,
}

var literalNames = []string{
	"", "','", "", "'('", "')'", "'['", "']'",
}
var symbolicNames = []string{
	"", "", "WS", "LPAREN", "RPAREN", "LBRACKET", "RBRACKET", "AND", "OR",
	"LT", "GT", "EQ", "CONTAINS", "IN", "BETWEEN", "BOOL", "DATETIME", "ALL_OF",
	"ANY_OF", "COUNT", "ISEMPTY", "STRING", "NUMBER", "NULL", "NOT", "ASC",
	"DESC", "SORT", "BY", "SKIP_ROWS", "LIMIT_ROWS", "NONE", "WHERE", "FROM",
	"IDENTIFIER", "RFC3339_DATE_TIME",
}

var ruleNames = []string{
	"stringArray", "numberArray", "datetimeArray", "start", "query", "skip",
	"limit", "sortBy", "sortField", "boolExpr", "operation", "binaryLhs", "setFunction",
	"setExpr", "subQueryExpr",
}

type ZitiQlParser struct {
	*antlr.BaseParser
}

func NewZitiQlParser(input antlr.TokenStream) *ZitiQlParser {
	var deserializer = antlr.NewATNDeserializer(nil)
	var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}

	this := new(ZitiQlParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "ZitiQl.g4"

	return this
}

// ZitiQlParser tokens.
const (
	ZitiQlParserEOF               = antlr.TokenEOF
	ZitiQlParserT__0              = 1
	ZitiQlParserWS                = 2
	ZitiQlParserLPAREN            = 3
	ZitiQlParserRPAREN            = 4
	ZitiQlParserLBRACKET          = 5
	ZitiQlParserRBRACKET          = 6
	ZitiQlParserAND               = 7
	ZitiQlParserOR                = 8
	ZitiQlParserLT                = 9
	ZitiQlParserGT                = 10
	ZitiQlParserEQ                = 11
	ZitiQlParserCONTAINS          = 12
	ZitiQlParserIN                = 13
	ZitiQlParserBETWEEN           = 14
	ZitiQlParserBOOL              = 15
	ZitiQlParserDATETIME          = 16
	ZitiQlParserALL_OF            = 17
	ZitiQlParserANY_OF            = 18
	ZitiQlParserCOUNT             = 19
	ZitiQlParserISEMPTY           = 20
	ZitiQlParserSTRING            = 21
	ZitiQlParserNUMBER            = 22
	ZitiQlParserNULL              = 23
	ZitiQlParserNOT               = 24
	ZitiQlParserASC               = 25
	ZitiQlParserDESC              = 26
	ZitiQlParserSORT              = 27
	ZitiQlParserBY                = 28
	ZitiQlParserSKIP_ROWS         = 29
	ZitiQlParserLIMIT_ROWS        = 30
	ZitiQlParserNONE              = 31
	ZitiQlParserWHERE             = 32
	ZitiQlParserFROM              = 33
	ZitiQlParserIDENTIFIER        = 34
	ZitiQlParserRFC3339_DATE_TIME = 35
)

// ZitiQlParser rules.
const (
	ZitiQlParserRULE_stringArray   = 0
	ZitiQlParserRULE_numberArray   = 1
	ZitiQlParserRULE_datetimeArray = 2
	ZitiQlParserRULE_start         = 3
	ZitiQlParserRULE_query         = 4
	ZitiQlParserRULE_skip          = 5
	ZitiQlParserRULE_limit         = 6
	ZitiQlParserRULE_sortBy        = 7
	ZitiQlParserRULE_sortField     = 8
	ZitiQlParserRULE_boolExpr      = 9
	ZitiQlParserRULE_operation     = 10
	ZitiQlParserRULE_binaryLhs     = 11
	ZitiQlParserRULE_setFunction   = 12
	ZitiQlParserRULE_setExpr       = 13
	ZitiQlParserRULE_subQueryExpr  = 14
)

// IStringArrayContext is an interface to support dynamic dispatch.
type IStringArrayContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStringArrayContext differentiates from other interfaces.
	IsStringArrayContext()
}

type StringArrayContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStringArrayContext() *StringArrayContext {
	var p = new(StringArrayContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZitiQlParserRULE_stringArray
	return p
}

func (*StringArrayContext) IsStringArrayContext() {}

func NewStringArrayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StringArrayContext {
	var p = new(StringArrayContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZitiQlParserRULE_stringArray

	return p
}

func (s *StringArrayContext) GetParser() antlr.Parser { return s.parser }

func (s *StringArrayContext) LBRACKET() antlr.TerminalNode {
	return s.GetToken(ZitiQlParserLBRACKET, 0)
}

func (s *StringArrayContext) AllSTRING() []antlr.TerminalNode {
	return s.GetTokens(ZitiQlParserSTRING)
}

func (s *StringArrayContext) STRING(i int) antlr.TerminalNode {
	return s.GetToken(ZitiQlParserSTRING, i)
}

func (s *StringArrayContext) RBRACKET() antlr.TerminalNode {
	return s.GetToken(ZitiQlParserRBRACKET, 0)
}

func (s *StringArrayContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(ZitiQlParserWS)
}

func (s *StringArrayContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(ZitiQlParserWS, i)
}

func (s *StringArrayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringArrayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StringArrayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZitiQlListener); ok {
		listenerT.EnterStringArray(s)
	}
}

func (s *StringArrayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZitiQlListener); ok {
		listenerT.ExitStringArray(s)
	}
}

func (p *ZitiQlParser) StringArray() (localctx IStringArrayContext) {
	localctx = NewStringArrayContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ZitiQlParserRULE_stringArray)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(30)
		p.Match(ZitiQlParserLBRACKET)
	}
	p.SetState(34)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ZitiQlParserWS {
		{
			p.SetState(31)
			p.Match(ZitiQlParserWS)
		}

		p.SetState(36)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(37)
		p.Match(ZitiQlParserSTRING)
	}
	p.SetState(54)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(41)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == ZitiQlParserWS {
				{
					p.SetState(38)
					p.Match(ZitiQlParserWS)
				}

				p.SetState(43)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(44)
				p.Match(ZitiQlParserT__0)
			}
			p.SetState(48)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == ZitiQlParserWS {
				{
					p.SetState(45)
					p.Match(ZitiQlParserWS)
				}

				p.SetState(50)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(51)
				p.Match(ZitiQlParserSTRING)
			}

		}
		p.SetState(56)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())
	}
	p.SetState(60)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ZitiQlParserWS {
		{
			p.SetState(57)
			p.Match(ZitiQlParserWS)
		}

		p.SetState(62)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(63)
		p.Match(ZitiQlParserRBRACKET)
	}

	return localctx
}

// INumberArrayContext is an interface to support dynamic dispatch.
type INumberArrayContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNumberArrayContext differentiates from other interfaces.
	IsNumberArrayContext()
}

type NumberArrayContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumberArrayContext() *NumberArrayContext {
	var p = new(NumberArrayContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZitiQlParserRULE_numberArray
	return p
}

func (*NumberArrayContext) IsNumberArrayContext() {}

func NewNumberArrayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumberArrayContext {
	var p = new(NumberArrayContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZitiQlParserRULE_numberArray

	return p
}

func (s *NumberArrayContext) GetParser() antlr.Parser { return s.parser }

func (s *NumberArrayContext) LBRACKET() antlr.TerminalNode {
	return s.GetToken(ZitiQlParserLBRACKET, 0)
}

func (s *NumberArrayContext) AllNUMBER() []antlr.TerminalNode {
	return s.GetTokens(ZitiQlParserNUMBER)
}

func (s *NumberArrayContext) NUMBER(i int) antlr.TerminalNode {
	return s.GetToken(ZitiQlParserNUMBER, i)
}

func (s *NumberArrayContext) RBRACKET() antlr.TerminalNode {
	return s.GetToken(ZitiQlParserRBRACKET, 0)
}

func (s *NumberArrayContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(ZitiQlParserWS)
}

func (s *NumberArrayContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(ZitiQlParserWS, i)
}

func (s *NumberArrayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberArrayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumberArrayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZitiQlListener); ok {
		listenerT.EnterNumberArray(s)
	}
}

func (s *NumberArrayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZitiQlListener); ok {
		listenerT.ExitNumberArray(s)
	}
}

func (p *ZitiQlParser) NumberArray() (localctx INumberArrayContext) {
	localctx = NewNumberArrayContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ZitiQlParserRULE_numberArray)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(65)
		p.Match(ZitiQlParserLBRACKET)
	}
	p.SetState(69)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ZitiQlParserWS {
		{
			p.SetState(66)
			p.Match(ZitiQlParserWS)
		}

		p.SetState(71)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(72)
		p.Match(ZitiQlParserNUMBER)
	}
	p.SetState(89)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(76)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == ZitiQlParserWS {
				{
					p.SetState(73)
					p.Match(ZitiQlParserWS)
				}

				p.SetState(78)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(79)
				p.Match(ZitiQlParserT__0)
			}
			p.SetState(83)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == ZitiQlParserWS {
				{
					p.SetState(80)
					p.Match(ZitiQlParserWS)
				}

				p.SetState(85)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(86)
				p.Match(ZitiQlParserNUMBER)
			}

		}
		p.SetState(91)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())
	}
	p.SetState(95)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ZitiQlParserWS {
		{
			p.SetState(92)
			p.Match(ZitiQlParserWS)
		}

		p.SetState(97)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(98)
		p.Match(ZitiQlParserRBRACKET)
	}

	return localctx
}

// IDatetimeArrayContext is an interface to support dynamic dispatch.
type IDatetimeArrayContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDatetimeArrayContext differentiates from other interfaces.
	IsDatetimeArrayContext()
}

type DatetimeArrayContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDatetimeArrayContext() *DatetimeArrayContext {
	var p = new(DatetimeArrayContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZitiQlParserRULE_datetimeArray
	return p
}

func (*DatetimeArrayContext) IsDatetimeArrayContext() {}

func NewDatetimeArrayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DatetimeArrayContext {
	var p = new(DatetimeArrayContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZitiQlParserRULE_datetimeArray

	return p
}

func (s *DatetimeArrayContext) GetParser() antlr.Parser { return s.parser }

func (s *DatetimeArrayContext) LBRACKET() antlr.TerminalNode {
	return s.GetToken(ZitiQlParserLBRACKET, 0)
}

func (s *DatetimeArrayContext) AllDATETIME() []antlr.TerminalNode {
	return s.GetTokens(ZitiQlParserDATETIME)
}

func (s *DatetimeArrayContext) DATETIME(i int) antlr.TerminalNode {
	return s.GetToken(ZitiQlParserDATETIME, i)
}

func (s *DatetimeArrayContext) RBRACKET() antlr.TerminalNode {
	return s.GetToken(ZitiQlParserRBRACKET, 0)
}

func (s *DatetimeArrayContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(ZitiQlParserWS)
}

func (s *DatetimeArrayContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(ZitiQlParserWS, i)
}

func (s *DatetimeArrayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DatetimeArrayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DatetimeArrayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZitiQlListener); ok {
		listenerT.EnterDatetimeArray(s)
	}
}

func (s *DatetimeArrayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZitiQlListener); ok {
		listenerT.ExitDatetimeArray(s)
	}
}

func (p *ZitiQlParser) DatetimeArray() (localctx IDatetimeArrayContext) {
	localctx = NewDatetimeArrayContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ZitiQlParserRULE_datetimeArray)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(100)
		p.Match(ZitiQlParserLBRACKET)
	}
	p.SetState(104)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ZitiQlParserWS {
		{
			p.SetState(101)
			p.Match(ZitiQlParserWS)
		}

		p.SetState(106)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(107)
		p.Match(ZitiQlParserDATETIME)
	}
	p.SetState(124)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(111)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == ZitiQlParserWS {
				{
					p.SetState(108)
					p.Match(ZitiQlParserWS)
				}

				p.SetState(113)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(114)
				p.Match(ZitiQlParserT__0)
			}
			p.SetState(118)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == ZitiQlParserWS {
				{
					p.SetState(115)
					p.Match(ZitiQlParserWS)
				}

				p.SetState(120)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(121)
				p.Match(ZitiQlParserDATETIME)
			}

		}
		p.SetState(126)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext())
	}
	p.SetState(130)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ZitiQlParserWS {
		{
			p.SetState(127)
			p.Match(ZitiQlParserWS)
		}

		p.SetState(132)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(133)
		p.Match(ZitiQlParserRBRACKET)
	}

	return localctx
}

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStartContext() *StartContext {
	var p = new(StartContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZitiQlParserRULE_start
	return p
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZitiQlParserRULE_start

	return p
}

func (s *StartContext) GetParser() antlr.Parser { return s.parser }

func (s *StartContext) CopyFrom(ctx *StartContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type EndContext struct {
	*StartContext
}

func NewEndContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EndContext {
	var p = new(EndContext)

	p.StartContext = NewEmptyStartContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StartContext))

	return p
}

func (s *EndContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EndContext) EOF() antlr.TerminalNode {
	return s.GetToken(ZitiQlParserEOF, 0)
}

func (s *EndContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(ZitiQlParserWS)
}

func (s *EndContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(ZitiQlParserWS, i)
}

func (s *EndContext) AllQuery() []IQueryContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IQueryContext)(nil)).Elem())
	var tst = make([]IQueryContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IQueryContext)
		}
	}

	return tst
}

func (s *EndContext) Query(i int) IQueryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQueryContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IQueryContext)
}

func (s *EndContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZitiQlListener); ok {
		listenerT.EnterEnd(s)
	}
}

func (s *EndContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZitiQlListener); ok {
		listenerT.ExitEnd(s)
	}
}

func (p *ZitiQlParser) Start() (localctx IStartContext) {
	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ZitiQlParserRULE_start)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	localctx = NewEndContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(138)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(135)
				p.Match(ZitiQlParserWS)
			}

		}
		p.SetState(140)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext())
	}
	p.SetState(144)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la-3)&-(0x1f+1)) == 0 && ((uint64(1)<<uint((_la-3)))&((1<<(ZitiQlParserLPAREN-3))|(1<<(ZitiQlParserBOOL-3))|(1<<(ZitiQlParserALL_OF-3))|(1<<(ZitiQlParserANY_OF-3))|(1<<(ZitiQlParserCOUNT-3))|(1<<(ZitiQlParserISEMPTY-3))|(1<<(ZitiQlParserNOT-3))|(1<<(ZitiQlParserIDENTIFIER-3)))) != 0 {
		{
			p.SetState(141)
			p.Query()
		}

		p.SetState(146)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(150)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ZitiQlParserWS {
		{
			p.SetState(147)
			p.Match(ZitiQlParserWS)
		}

		p.SetState(152)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(153)
		p.Match(ZitiQlParserEOF)
	}

	return localctx
}

// IQueryContext is an interface to support dynamic dispatch.
type IQueryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQueryContext differentiates from other interfaces.
	IsQueryContext()
}

type QueryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQueryContext() *QueryContext {
	var p = new(QueryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZitiQlParserRULE_query
	return p
}

func (*QueryContext) IsQueryContext() {}

func NewQueryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QueryContext {
	var p = new(QueryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZitiQlParserRULE_query

	return p
}

func (s *QueryContext) GetParser() antlr.Parser { return s.parser }

func (s *QueryContext) CopyFrom(ctx *QueryContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *QueryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QueryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type QueryStmtContext struct {
	*QueryContext
}

func NewQueryStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *QueryStmtContext {
	var p = new(QueryStmtContext)

	p.QueryContext = NewEmptyQueryContext()
	p.parser = parser
	p.CopyFrom(ctx.(*QueryContext))

	return p
}

func (s *QueryStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QueryStmtContext) BoolExpr() IBoolExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBoolExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBoolExprContext)
}

func (s *QueryStmtContext) SortBy() ISortByContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISortByContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISortByContext)
}

func (s *QueryStmtContext) Skip() ISkipContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISkipContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISkipContext)
}

func (s *QueryStmtContext) Limit() ILimitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILimitContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILimitContext)
}

func (s *QueryStmtContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(ZitiQlParserWS)
}

func (s *QueryStmtContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(ZitiQlParserWS, i)
}

func (s *QueryStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZitiQlListener); ok {
		listenerT.EnterQueryStmt(s)
	}
}

func (s *QueryStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZitiQlListener); ok {
		listenerT.ExitQueryStmt(s)
	}
}

func (p *ZitiQlParser) Query() (localctx IQueryContext) {
	localctx = NewQueryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, ZitiQlParserRULE_query)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	localctx = NewQueryStmtContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(155)
		p.boolExpr(0)
	}
	p.SetState(162)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext()) == 1 {
		p.SetState(157)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == ZitiQlParserWS {
			{
				p.SetState(156)
				p.Match(ZitiQlParserWS)
			}

			p.SetState(159)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(161)
			p.SortBy()
		}

	}
	p.SetState(170)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext()) == 1 {
		p.SetState(165)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == ZitiQlParserWS {
			{
				p.SetState(164)
				p.Match(ZitiQlParserWS)
			}

			p.SetState(167)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(169)
			p.Skip()
		}

	}
	p.SetState(178)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext()) == 1 {
		p.SetState(173)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == ZitiQlParserWS {
			{
				p.SetState(172)
				p.Match(ZitiQlParserWS)
			}

			p.SetState(175)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(177)
			p.Limit()
		}

	}

	return localctx
}

// ISkipContext is an interface to support dynamic dispatch.
type ISkipContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSkipContext differentiates from other interfaces.
	IsSkipContext()
}

type SkipContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySkipContext() *SkipContext {
	var p = new(SkipContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZitiQlParserRULE_skip
	return p
}

func (*SkipContext) IsSkipContext() {}

func NewSkipContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SkipContext {
	var p = new(SkipContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZitiQlParserRULE_skip

	return p
}

func (s *SkipContext) GetParser() antlr.Parser { return s.parser }

func (s *SkipContext) CopyFrom(ctx *SkipContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *SkipContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SkipContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SkipExprContext struct {
	*SkipContext
}

func NewSkipExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SkipExprContext {
	var p = new(SkipExprContext)

	p.SkipContext = NewEmptySkipContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SkipContext))

	return p
}

func (s *SkipExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SkipExprContext) SKIP_ROWS() antlr.TerminalNode {
	return s.GetToken(ZitiQlParserSKIP_ROWS, 0)
}

func (s *SkipExprContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(ZitiQlParserNUMBER, 0)
}

func (s *SkipExprContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(ZitiQlParserWS)
}

func (s *SkipExprContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(ZitiQlParserWS, i)
}

func (s *SkipExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZitiQlListener); ok {
		listenerT.EnterSkipExpr(s)
	}
}

func (s *SkipExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZitiQlListener); ok {
		listenerT.ExitSkipExpr(s)
	}
}

func (p *ZitiQlParser) Skip() (localctx ISkipContext) {
	localctx = NewSkipContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, ZitiQlParserRULE_skip)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	localctx = NewSkipExprContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(180)
		p.Match(ZitiQlParserSKIP_ROWS)
	}
	p.SetState(182)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == ZitiQlParserWS {
		{
			p.SetState(181)
			p.Match(ZitiQlParserWS)
		}

		p.SetState(184)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(186)
		p.Match(ZitiQlParserNUMBER)
	}

	return localctx
}

// ILimitContext is an interface to support dynamic dispatch.
type ILimitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLimitContext differentiates from other interfaces.
	IsLimitContext()
}

type LimitContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLimitContext() *LimitContext {
	var p = new(LimitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZitiQlParserRULE_limit
	return p
}

func (*LimitContext) IsLimitContext() {}

func NewLimitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LimitContext {
	var p = new(LimitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZitiQlParserRULE_limit

	return p
}

func (s *LimitContext) GetParser() antlr.Parser { return s.parser }

func (s *LimitContext) CopyFrom(ctx *LimitContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *LimitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LimitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type LimitExprContext struct {
	*LimitContext
}

func NewLimitExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LimitExprContext {
	var p = new(LimitExprContext)

	p.LimitContext = NewEmptyLimitContext()
	p.parser = parser
	p.CopyFrom(ctx.(*LimitContext))

	return p
}

func (s *LimitExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LimitExprContext) LIMIT_ROWS() antlr.TerminalNode {
	return s.GetToken(ZitiQlParserLIMIT_ROWS, 0)
}

func (s *LimitExprContext) NONE() antlr.TerminalNode {
	return s.GetToken(ZitiQlParserNONE, 0)
}

func (s *LimitExprContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(ZitiQlParserNUMBER, 0)
}

func (s *LimitExprContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(ZitiQlParserWS)
}

func (s *LimitExprContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(ZitiQlParserWS, i)
}

func (s *LimitExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZitiQlListener); ok {
		listenerT.EnterLimitExpr(s)
	}
}

func (s *LimitExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZitiQlListener); ok {
		listenerT.ExitLimitExpr(s)
	}
}

func (p *ZitiQlParser) Limit() (localctx ILimitContext) {
	localctx = NewLimitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, ZitiQlParserRULE_limit)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	localctx = NewLimitExprContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(188)
		p.Match(ZitiQlParserLIMIT_ROWS)
	}
	p.SetState(190)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == ZitiQlParserWS {
		{
			p.SetState(189)
			p.Match(ZitiQlParserWS)
		}

		p.SetState(192)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(194)
		_la = p.GetTokenStream().LA(1)

		if !(_la == ZitiQlParserNUMBER || _la == ZitiQlParserNONE) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ISortByContext is an interface to support dynamic dispatch.
type ISortByContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSortByContext differentiates from other interfaces.
	IsSortByContext()
}

type SortByContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySortByContext() *SortByContext {
	var p = new(SortByContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZitiQlParserRULE_sortBy
	return p
}

func (*SortByContext) IsSortByContext() {}

func NewSortByContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SortByContext {
	var p = new(SortByContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZitiQlParserRULE_sortBy

	return p
}

func (s *SortByContext) GetParser() antlr.Parser { return s.parser }

func (s *SortByContext) CopyFrom(ctx *SortByContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *SortByContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SortByContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SortByExprContext struct {
	*SortByContext
}

func NewSortByExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SortByExprContext {
	var p = new(SortByExprContext)

	p.SortByContext = NewEmptySortByContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SortByContext))

	return p
}

func (s *SortByExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SortByExprContext) SORT() antlr.TerminalNode {
	return s.GetToken(ZitiQlParserSORT, 0)
}

func (s *SortByExprContext) BY() antlr.TerminalNode {
	return s.GetToken(ZitiQlParserBY, 0)
}

func (s *SortByExprContext) AllSortField() []ISortFieldContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISortFieldContext)(nil)).Elem())
	var tst = make([]ISortFieldContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISortFieldContext)
		}
	}

	return tst
}

func (s *SortByExprContext) SortField(i int) ISortFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISortFieldContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISortFieldContext)
}

func (s *SortByExprContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(ZitiQlParserWS)
}

func (s *SortByExprContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(ZitiQlParserWS, i)
}

func (s *SortByExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZitiQlListener); ok {
		listenerT.EnterSortByExpr(s)
	}
}

func (s *SortByExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZitiQlListener); ok {
		listenerT.ExitSortByExpr(s)
	}
}

func (p *ZitiQlParser) SortBy() (localctx ISortByContext) {
	localctx = NewSortByContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, ZitiQlParserRULE_sortBy)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	localctx = NewSortByExprContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(196)
		p.Match(ZitiQlParserSORT)
	}
	p.SetState(198)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == ZitiQlParserWS {
		{
			p.SetState(197)
			p.Match(ZitiQlParserWS)
		}

		p.SetState(200)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(202)
		p.Match(ZitiQlParserBY)
	}
	p.SetState(204)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == ZitiQlParserWS {
		{
			p.SetState(203)
			p.Match(ZitiQlParserWS)
		}

		p.SetState(206)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(208)
		p.SortField()
	}
	p.SetState(225)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(212)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == ZitiQlParserWS {
				{
					p.SetState(209)
					p.Match(ZitiQlParserWS)
				}

				p.SetState(214)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(215)
				p.Match(ZitiQlParserT__0)
			}
			p.SetState(219)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == ZitiQlParserWS {
				{
					p.SetState(216)
					p.Match(ZitiQlParserWS)
				}

				p.SetState(221)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(222)
				p.SortField()
			}

		}
		p.SetState(227)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext())
	}

	return localctx
}

// ISortFieldContext is an interface to support dynamic dispatch.
type ISortFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSortFieldContext differentiates from other interfaces.
	IsSortFieldContext()
}

type SortFieldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySortFieldContext() *SortFieldContext {
	var p = new(SortFieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZitiQlParserRULE_sortField
	return p
}

func (*SortFieldContext) IsSortFieldContext() {}

func NewSortFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SortFieldContext {
	var p = new(SortFieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZitiQlParserRULE_sortField

	return p
}

func (s *SortFieldContext) GetParser() antlr.Parser { return s.parser }

func (s *SortFieldContext) CopyFrom(ctx *SortFieldContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *SortFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SortFieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SortFieldExprContext struct {
	*SortFieldContext
}

func NewSortFieldExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SortFieldExprContext {
	var p = new(SortFieldExprContext)

	p.SortFieldContext = NewEmptySortFieldContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SortFieldContext))

	return p
}

func (s *SortFieldExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SortFieldExprContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ZitiQlParserIDENTIFIER, 0)
}

func (s *SortFieldExprContext) ASC() antlr.TerminalNode {
	return s.GetToken(ZitiQlParserASC, 0)
}

func (s *SortFieldExprContext) DESC() antlr.TerminalNode {
	return s.GetToken(ZitiQlParserDESC, 0)
}

func (s *SortFieldExprContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(ZitiQlParserWS)
}

func (s *SortFieldExprContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(ZitiQlParserWS, i)
}

func (s *SortFieldExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZitiQlListener); ok {
		listenerT.EnterSortFieldExpr(s)
	}
}

func (s *SortFieldExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZitiQlListener); ok {
		listenerT.ExitSortFieldExpr(s)
	}
}

func (p *ZitiQlParser) SortField() (localctx ISortFieldContext) {
	localctx = NewSortFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, ZitiQlParserRULE_sortField)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	localctx = NewSortFieldExprContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(228)
		p.Match(ZitiQlParserIDENTIFIER)
	}
	p.SetState(235)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 32, p.GetParserRuleContext()) == 1 {
		p.SetState(230)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == ZitiQlParserWS {
			{
				p.SetState(229)
				p.Match(ZitiQlParserWS)
			}

			p.SetState(232)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(234)
			_la = p.GetTokenStream().LA(1)

			if !(_la == ZitiQlParserASC || _la == ZitiQlParserDESC) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}

	return localctx
}

// IBoolExprContext is an interface to support dynamic dispatch.
type IBoolExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBoolExprContext differentiates from other interfaces.
	IsBoolExprContext()
}

type BoolExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBoolExprContext() *BoolExprContext {
	var p = new(BoolExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZitiQlParserRULE_boolExpr
	return p
}

func (*BoolExprContext) IsBoolExprContext() {}

func NewBoolExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BoolExprContext {
	var p = new(BoolExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZitiQlParserRULE_boolExpr

	return p
}

func (s *BoolExprContext) GetParser() antlr.Parser { return s.parser }

func (s *BoolExprContext) CopyFrom(ctx *BoolExprContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *BoolExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type AndExprContext struct {
	*BoolExprContext
}

func NewAndExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AndExprContext {
	var p = new(AndExprContext)

	p.BoolExprContext = NewEmptyBoolExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*BoolExprContext))

	return p
}

func (s *AndExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndExprContext) AllBoolExpr() []IBoolExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBoolExprContext)(nil)).Elem())
	var tst = make([]IBoolExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBoolExprContext)
		}
	}

	return tst
}

func (s *AndExprContext) BoolExpr(i int) IBoolExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBoolExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBoolExprContext)
}

func (s *AndExprContext) AllAND() []antlr.TerminalNode {
	return s.GetTokens(ZitiQlParserAND)
}

func (s *AndExprContext) AND(i int) antlr.TerminalNode {
	return s.GetToken(ZitiQlParserAND, i)
}

func (s *AndExprContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(ZitiQlParserWS)
}

func (s *AndExprContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(ZitiQlParserWS, i)
}

func (s *AndExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZitiQlListener); ok {
		listenerT.EnterAndExpr(s)
	}
}

func (s *AndExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZitiQlListener); ok {
		listenerT.ExitAndExpr(s)
	}
}

type GroupContext struct {
	*BoolExprContext
}

func NewGroupContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *GroupContext {
	var p = new(GroupContext)

	p.BoolExprContext = NewEmptyBoolExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*BoolExprContext))

	return p
}

func (s *GroupContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GroupContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ZitiQlParserLPAREN, 0)
}

func (s *GroupContext) BoolExpr() IBoolExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBoolExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBoolExprContext)
}

func (s *GroupContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ZitiQlParserRPAREN, 0)
}

func (s *GroupContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(ZitiQlParserWS)
}

func (s *GroupContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(ZitiQlParserWS, i)
}

func (s *GroupContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZitiQlListener); ok {
		listenerT.EnterGroup(s)
	}
}

func (s *GroupContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZitiQlListener); ok {
		listenerT.ExitGroup(s)
	}
}

type BoolConstContext struct {
	*BoolExprContext
}

func NewBoolConstContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BoolConstContext {
	var p = new(BoolConstContext)

	p.BoolExprContext = NewEmptyBoolExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*BoolExprContext))

	return p
}

func (s *BoolConstContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolConstContext) BOOL() antlr.TerminalNode {
	return s.GetToken(ZitiQlParserBOOL, 0)
}

func (s *BoolConstContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZitiQlListener); ok {
		listenerT.EnterBoolConst(s)
	}
}

func (s *BoolConstContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZitiQlListener); ok {
		listenerT.ExitBoolConst(s)
	}
}

type IsEmptyFunctionContext struct {
	*BoolExprContext
}

func NewIsEmptyFunctionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IsEmptyFunctionContext {
	var p = new(IsEmptyFunctionContext)

	p.BoolExprContext = NewEmptyBoolExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*BoolExprContext))

	return p
}

func (s *IsEmptyFunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IsEmptyFunctionContext) ISEMPTY() antlr.TerminalNode {
	return s.GetToken(ZitiQlParserISEMPTY, 0)
}

func (s *IsEmptyFunctionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ZitiQlParserLPAREN, 0)
}

func (s *IsEmptyFunctionContext) SetExpr() ISetExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISetExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISetExprContext)
}

func (s *IsEmptyFunctionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ZitiQlParserRPAREN, 0)
}

func (s *IsEmptyFunctionContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(ZitiQlParserWS)
}

func (s *IsEmptyFunctionContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(ZitiQlParserWS, i)
}

func (s *IsEmptyFunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZitiQlListener); ok {
		listenerT.EnterIsEmptyFunction(s)
	}
}

func (s *IsEmptyFunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZitiQlListener); ok {
		listenerT.ExitIsEmptyFunction(s)
	}
}

type NotExprContext struct {
	*BoolExprContext
}

func NewNotExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NotExprContext {
	var p = new(NotExprContext)

	p.BoolExprContext = NewEmptyBoolExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*BoolExprContext))

	return p
}

func (s *NotExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotExprContext) NOT() antlr.TerminalNode {
	return s.GetToken(ZitiQlParserNOT, 0)
}

func (s *NotExprContext) BoolExpr() IBoolExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBoolExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBoolExprContext)
}

func (s *NotExprContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(ZitiQlParserWS)
}

func (s *NotExprContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(ZitiQlParserWS, i)
}

func (s *NotExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZitiQlListener); ok {
		listenerT.EnterNotExpr(s)
	}
}

func (s *NotExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZitiQlListener); ok {
		listenerT.ExitNotExpr(s)
	}
}

type OperationOpContext struct {
	*BoolExprContext
}

func NewOperationOpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OperationOpContext {
	var p = new(OperationOpContext)

	p.BoolExprContext = NewEmptyBoolExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*BoolExprContext))

	return p
}

func (s *OperationOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperationOpContext) Operation() IOperationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOperationContext)
}

func (s *OperationOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZitiQlListener); ok {
		listenerT.EnterOperationOp(s)
	}
}

func (s *OperationOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZitiQlListener); ok {
		listenerT.ExitOperationOp(s)
	}
}

type OrExprContext struct {
	*BoolExprContext
}

func NewOrExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OrExprContext {
	var p = new(OrExprContext)

	p.BoolExprContext = NewEmptyBoolExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*BoolExprContext))

	return p
}

func (s *OrExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrExprContext) AllBoolExpr() []IBoolExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBoolExprContext)(nil)).Elem())
	var tst = make([]IBoolExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBoolExprContext)
		}
	}

	return tst
}

func (s *OrExprContext) BoolExpr(i int) IBoolExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBoolExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBoolExprContext)
}

func (s *OrExprContext) AllOR() []antlr.TerminalNode {
	return s.GetTokens(ZitiQlParserOR)
}

func (s *OrExprContext) OR(i int) antlr.TerminalNode {
	return s.GetToken(ZitiQlParserOR, i)
}

func (s *OrExprContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(ZitiQlParserWS)
}

func (s *OrExprContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(ZitiQlParserWS, i)
}

func (s *OrExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZitiQlListener); ok {
		listenerT.EnterOrExpr(s)
	}
}

func (s *OrExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZitiQlListener); ok {
		listenerT.ExitOrExpr(s)
	}
}

type BoolSymbolContext struct {
	*BoolExprContext
}

func NewBoolSymbolContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BoolSymbolContext {
	var p = new(BoolSymbolContext)

	p.BoolExprContext = NewEmptyBoolExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*BoolExprContext))

	return p
}

func (s *BoolSymbolContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolSymbolContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ZitiQlParserIDENTIFIER, 0)
}

func (s *BoolSymbolContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZitiQlListener); ok {
		listenerT.EnterBoolSymbol(s)
	}
}

func (s *BoolSymbolContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZitiQlListener); ok {
		listenerT.ExitBoolSymbol(s)
	}
}

func (p *ZitiQlParser) BoolExpr() (localctx IBoolExprContext) {
	return p.boolExpr(0)
}

func (p *ZitiQlParser) boolExpr(_p int) (localctx IBoolExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewBoolExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IBoolExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 18
	p.EnterRecursionRule(localctx, 18, ZitiQlParserRULE_boolExpr, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(281)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 38, p.GetParserRuleContext()) {
	case 1:
		localctx = NewOperationOpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(238)
			p.Operation()
		}

	case 2:
		localctx = NewGroupContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(239)
			p.Match(ZitiQlParserLPAREN)
		}
		p.SetState(243)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ZitiQlParserWS {
			{
				p.SetState(240)
				p.Match(ZitiQlParserWS)
			}

			p.SetState(245)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(246)
			p.boolExpr(0)
		}
		p.SetState(250)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ZitiQlParserWS {
			{
				p.SetState(247)
				p.Match(ZitiQlParserWS)
			}

			p.SetState(252)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(253)
			p.Match(ZitiQlParserRPAREN)
		}

	case 3:
		localctx = NewBoolConstContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(255)
			p.Match(ZitiQlParserBOOL)
		}

	case 4:
		localctx = NewIsEmptyFunctionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(256)
			p.Match(ZitiQlParserISEMPTY)
		}
		{
			p.SetState(257)
			p.Match(ZitiQlParserLPAREN)
		}
		p.SetState(261)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ZitiQlParserWS {
			{
				p.SetState(258)
				p.Match(ZitiQlParserWS)
			}

			p.SetState(263)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(264)
			p.SetExpr()
		}
		p.SetState(268)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ZitiQlParserWS {
			{
				p.SetState(265)
				p.Match(ZitiQlParserWS)
			}

			p.SetState(270)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(271)
			p.Match(ZitiQlParserRPAREN)
		}

	case 5:
		localctx = NewBoolSymbolContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(273)
			p.Match(ZitiQlParserIDENTIFIER)
		}

	case 6:
		localctx = NewNotExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(274)
			p.Match(ZitiQlParserNOT)
		}
		p.SetState(276)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == ZitiQlParserWS {
			{
				p.SetState(275)
				p.Match(ZitiQlParserWS)
			}

			p.SetState(278)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(280)
			p.boolExpr(1)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(319)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 46, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(317)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 45, p.GetParserRuleContext()) {
			case 1:
				localctx = NewAndExprContext(p, NewBoolExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ZitiQlParserRULE_boolExpr)
				p.SetState(283)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				p.SetState(296)
				p.GetErrorHandler().Sync(p)
				_alt = 1
				for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
					switch _alt {
					case 1:
						p.SetState(285)
						p.GetErrorHandler().Sync(p)
						_la = p.GetTokenStream().LA(1)

						for ok := true; ok; ok = _la == ZitiQlParserWS {
							{
								p.SetState(284)
								p.Match(ZitiQlParserWS)
							}

							p.SetState(287)
							p.GetErrorHandler().Sync(p)
							_la = p.GetTokenStream().LA(1)
						}
						{
							p.SetState(289)
							p.Match(ZitiQlParserAND)
						}
						p.SetState(291)
						p.GetErrorHandler().Sync(p)
						_la = p.GetTokenStream().LA(1)

						for ok := true; ok; ok = _la == ZitiQlParserWS {
							{
								p.SetState(290)
								p.Match(ZitiQlParserWS)
							}

							p.SetState(293)
							p.GetErrorHandler().Sync(p)
							_la = p.GetTokenStream().LA(1)
						}
						{
							p.SetState(295)
							p.boolExpr(0)
						}

					default:
						panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
					}

					p.SetState(298)
					p.GetErrorHandler().Sync(p)
					_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 41, p.GetParserRuleContext())
				}

			case 2:
				localctx = NewOrExprContext(p, NewBoolExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ZitiQlParserRULE_boolExpr)
				p.SetState(300)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				p.SetState(313)
				p.GetErrorHandler().Sync(p)
				_alt = 1
				for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
					switch _alt {
					case 1:
						p.SetState(302)
						p.GetErrorHandler().Sync(p)
						_la = p.GetTokenStream().LA(1)

						for ok := true; ok; ok = _la == ZitiQlParserWS {
							{
								p.SetState(301)
								p.Match(ZitiQlParserWS)
							}

							p.SetState(304)
							p.GetErrorHandler().Sync(p)
							_la = p.GetTokenStream().LA(1)
						}
						{
							p.SetState(306)
							p.Match(ZitiQlParserOR)
						}
						p.SetState(308)
						p.GetErrorHandler().Sync(p)
						_la = p.GetTokenStream().LA(1)

						for ok := true; ok; ok = _la == ZitiQlParserWS {
							{
								p.SetState(307)
								p.Match(ZitiQlParserWS)
							}

							p.SetState(310)
							p.GetErrorHandler().Sync(p)
							_la = p.GetTokenStream().LA(1)
						}
						{
							p.SetState(312)
							p.boolExpr(0)
						}

					default:
						panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
					}

					p.SetState(315)
					p.GetErrorHandler().Sync(p)
					_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 44, p.GetParserRuleContext())
				}

			}

		}
		p.SetState(321)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 46, p.GetParserRuleContext())
	}

	return localctx
}

// IOperationContext is an interface to support dynamic dispatch.
type IOperationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOperationContext differentiates from other interfaces.
	IsOperationContext()
}

type OperationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperationContext() *OperationContext {
	var p = new(OperationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZitiQlParserRULE_operation
	return p
}

func (*OperationContext) IsOperationContext() {}

func NewOperationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperationContext {
	var p = new(OperationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZitiQlParserRULE_operation

	return p
}

func (s *OperationContext) GetParser() antlr.Parser { return s.parser }

func (s *OperationContext) CopyFrom(ctx *OperationContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *OperationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type BinaryEqualToNullOpContext struct {
	*OperationContext
}

func NewBinaryEqualToNullOpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BinaryEqualToNullOpContext {
	var p = new(BinaryEqualToNullOpContext)

	p.OperationContext = NewEmptyOperationContext()
	p.parser = parser
	p.CopyFrom(ctx.(*OperationContext))

	return p
}

func (s *BinaryEqualToNullOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinaryEqualToNullOpContext) BinaryLhs() IBinaryLhsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBinaryLhsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBinaryLhsContext)
}

func (s *BinaryEqualToNullOpContext) EQ() antlr.TerminalNode {
	return s.GetToken(ZitiQlParserEQ, 0)
}

func (s *BinaryEqualToNullOpContext) NULL() antlr.TerminalNode {
	return s.GetToken(ZitiQlParserNULL, 0)
}

func (s *BinaryEqualToNullOpContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(ZitiQlParserWS)
}

func (s *BinaryEqualToNullOpContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(ZitiQlParserWS, i)
}

func (s *BinaryEqualToNullOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZitiQlListener); ok {
		listenerT.EnterBinaryEqualToNullOp(s)
	}
}

func (s *BinaryEqualToNullOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZitiQlListener); ok {
		listenerT.ExitBinaryEqualToNullOp(s)
	}
}

type BinaryLessThanNumberOpContext struct {
	*OperationContext
}

func NewBinaryLessThanNumberOpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BinaryLessThanNumberOpContext {
	var p = new(BinaryLessThanNumberOpContext)

	p.OperationContext = NewEmptyOperationContext()
	p.parser = parser
	p.CopyFrom(ctx.(*OperationContext))

	return p
}

func (s *BinaryLessThanNumberOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinaryLessThanNumberOpContext) BinaryLhs() IBinaryLhsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBinaryLhsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBinaryLhsContext)
}

func (s *BinaryLessThanNumberOpContext) LT() antlr.TerminalNode {
	return s.GetToken(ZitiQlParserLT, 0)
}

func (s *BinaryLessThanNumberOpContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(ZitiQlParserNUMBER, 0)
}

func (s *BinaryLessThanNumberOpContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(ZitiQlParserWS)
}

func (s *BinaryLessThanNumberOpContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(ZitiQlParserWS, i)
}

func (s *BinaryLessThanNumberOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZitiQlListener); ok {
		listenerT.EnterBinaryLessThanNumberOp(s)
	}
}

func (s *BinaryLessThanNumberOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZitiQlListener); ok {
		listenerT.ExitBinaryLessThanNumberOp(s)
	}
}

type BinaryGreaterThanDatetimeOpContext struct {
	*OperationContext
}

func NewBinaryGreaterThanDatetimeOpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BinaryGreaterThanDatetimeOpContext {
	var p = new(BinaryGreaterThanDatetimeOpContext)

	p.OperationContext = NewEmptyOperationContext()
	p.parser = parser
	p.CopyFrom(ctx.(*OperationContext))

	return p
}

func (s *BinaryGreaterThanDatetimeOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinaryGreaterThanDatetimeOpContext) BinaryLhs() IBinaryLhsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBinaryLhsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBinaryLhsContext)
}

func (s *BinaryGreaterThanDatetimeOpContext) GT() antlr.TerminalNode {
	return s.GetToken(ZitiQlParserGT, 0)
}

func (s *BinaryGreaterThanDatetimeOpContext) DATETIME() antlr.TerminalNode {
	return s.GetToken(ZitiQlParserDATETIME, 0)
}

func (s *BinaryGreaterThanDatetimeOpContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(ZitiQlParserWS)
}

func (s *BinaryGreaterThanDatetimeOpContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(ZitiQlParserWS, i)
}

func (s *BinaryGreaterThanDatetimeOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZitiQlListener); ok {
		listenerT.EnterBinaryGreaterThanDatetimeOp(s)
	}
}

func (s *BinaryGreaterThanDatetimeOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZitiQlListener); ok {
		listenerT.ExitBinaryGreaterThanDatetimeOp(s)
	}
}

type InNumberArrayOpContext struct {
	*OperationContext
}

func NewInNumberArrayOpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InNumberArrayOpContext {
	var p = new(InNumberArrayOpContext)

	p.OperationContext = NewEmptyOperationContext()
	p.parser = parser
	p.CopyFrom(ctx.(*OperationContext))

	return p
}

func (s *InNumberArrayOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InNumberArrayOpContext) BinaryLhs() IBinaryLhsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBinaryLhsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBinaryLhsContext)
}

func (s *InNumberArrayOpContext) IN() antlr.TerminalNode {
	return s.GetToken(ZitiQlParserIN, 0)
}

func (s *InNumberArrayOpContext) NumberArray() INumberArrayContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumberArrayContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumberArrayContext)
}

func (s *InNumberArrayOpContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(ZitiQlParserWS)
}

func (s *InNumberArrayOpContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(ZitiQlParserWS, i)
}

func (s *InNumberArrayOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZitiQlListener); ok {
		listenerT.EnterInNumberArrayOp(s)
	}
}

func (s *InNumberArrayOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZitiQlListener); ok {
		listenerT.ExitInNumberArrayOp(s)
	}
}

type InStringArrayOpContext struct {
	*OperationContext
}

func NewInStringArrayOpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InStringArrayOpContext {
	var p = new(InStringArrayOpContext)

	p.OperationContext = NewEmptyOperationContext()
	p.parser = parser
	p.CopyFrom(ctx.(*OperationContext))

	return p
}

func (s *InStringArrayOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InStringArrayOpContext) BinaryLhs() IBinaryLhsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBinaryLhsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBinaryLhsContext)
}

func (s *InStringArrayOpContext) IN() antlr.TerminalNode {
	return s.GetToken(ZitiQlParserIN, 0)
}

func (s *InStringArrayOpContext) StringArray() IStringArrayContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStringArrayContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStringArrayContext)
}

func (s *InStringArrayOpContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(ZitiQlParserWS)
}

func (s *InStringArrayOpContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(ZitiQlParserWS, i)
}

func (s *InStringArrayOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZitiQlListener); ok {
		listenerT.EnterInStringArrayOp(s)
	}
}

func (s *InStringArrayOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZitiQlListener); ok {
		listenerT.ExitInStringArrayOp(s)
	}
}

type BinaryLessThanDatetimeOpContext struct {
	*OperationContext
}

func NewBinaryLessThanDatetimeOpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BinaryLessThanDatetimeOpContext {
	var p = new(BinaryLessThanDatetimeOpContext)

	p.OperationContext = NewEmptyOperationContext()
	p.parser = parser
	p.CopyFrom(ctx.(*OperationContext))

	return p
}

func (s *BinaryLessThanDatetimeOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinaryLessThanDatetimeOpContext) BinaryLhs() IBinaryLhsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBinaryLhsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBinaryLhsContext)
}

func (s *BinaryLessThanDatetimeOpContext) LT() antlr.TerminalNode {
	return s.GetToken(ZitiQlParserLT, 0)
}

func (s *BinaryLessThanDatetimeOpContext) DATETIME() antlr.TerminalNode {
	return s.GetToken(ZitiQlParserDATETIME, 0)
}

func (s *BinaryLessThanDatetimeOpContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(ZitiQlParserWS)
}

func (s *BinaryLessThanDatetimeOpContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(ZitiQlParserWS, i)
}

func (s *BinaryLessThanDatetimeOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZitiQlListener); ok {
		listenerT.EnterBinaryLessThanDatetimeOp(s)
	}
}

func (s *BinaryLessThanDatetimeOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZitiQlListener); ok {
		listenerT.ExitBinaryLessThanDatetimeOp(s)
	}
}

type BinaryGreaterThanNumberOpContext struct {
	*OperationContext
}

func NewBinaryGreaterThanNumberOpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BinaryGreaterThanNumberOpContext {
	var p = new(BinaryGreaterThanNumberOpContext)

	p.OperationContext = NewEmptyOperationContext()
	p.parser = parser
	p.CopyFrom(ctx.(*OperationContext))

	return p
}

func (s *BinaryGreaterThanNumberOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinaryGreaterThanNumberOpContext) BinaryLhs() IBinaryLhsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBinaryLhsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBinaryLhsContext)
}

func (s *BinaryGreaterThanNumberOpContext) GT() antlr.TerminalNode {
	return s.GetToken(ZitiQlParserGT, 0)
}

func (s *BinaryGreaterThanNumberOpContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(ZitiQlParserNUMBER, 0)
}

func (s *BinaryGreaterThanNumberOpContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(ZitiQlParserWS)
}

func (s *BinaryGreaterThanNumberOpContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(ZitiQlParserWS, i)
}

func (s *BinaryGreaterThanNumberOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZitiQlListener); ok {
		listenerT.EnterBinaryGreaterThanNumberOp(s)
	}
}

func (s *BinaryGreaterThanNumberOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZitiQlListener); ok {
		listenerT.ExitBinaryGreaterThanNumberOp(s)
	}
}

type InDatetimeArrayOpContext struct {
	*OperationContext
}

func NewInDatetimeArrayOpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InDatetimeArrayOpContext {
	var p = new(InDatetimeArrayOpContext)

	p.OperationContext = NewEmptyOperationContext()
	p.parser = parser
	p.CopyFrom(ctx.(*OperationContext))

	return p
}

func (s *InDatetimeArrayOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InDatetimeArrayOpContext) BinaryLhs() IBinaryLhsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBinaryLhsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBinaryLhsContext)
}

func (s *InDatetimeArrayOpContext) IN() antlr.TerminalNode {
	return s.GetToken(ZitiQlParserIN, 0)
}

func (s *InDatetimeArrayOpContext) DatetimeArray() IDatetimeArrayContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDatetimeArrayContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDatetimeArrayContext)
}

func (s *InDatetimeArrayOpContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(ZitiQlParserWS)
}

func (s *InDatetimeArrayOpContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(ZitiQlParserWS, i)
}

func (s *InDatetimeArrayOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZitiQlListener); ok {
		listenerT.EnterInDatetimeArrayOp(s)
	}
}

func (s *InDatetimeArrayOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZitiQlListener); ok {
		listenerT.ExitInDatetimeArrayOp(s)
	}
}

type BetweenDateOpContext struct {
	*OperationContext
}

func NewBetweenDateOpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BetweenDateOpContext {
	var p = new(BetweenDateOpContext)

	p.OperationContext = NewEmptyOperationContext()
	p.parser = parser
	p.CopyFrom(ctx.(*OperationContext))

	return p
}

func (s *BetweenDateOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BetweenDateOpContext) BinaryLhs() IBinaryLhsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBinaryLhsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBinaryLhsContext)
}

func (s *BetweenDateOpContext) BETWEEN() antlr.TerminalNode {
	return s.GetToken(ZitiQlParserBETWEEN, 0)
}

func (s *BetweenDateOpContext) AllDATETIME() []antlr.TerminalNode {
	return s.GetTokens(ZitiQlParserDATETIME)
}

func (s *BetweenDateOpContext) DATETIME(i int) antlr.TerminalNode {
	return s.GetToken(ZitiQlParserDATETIME, i)
}

func (s *BetweenDateOpContext) AND() antlr.TerminalNode {
	return s.GetToken(ZitiQlParserAND, 0)
}

func (s *BetweenDateOpContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(ZitiQlParserWS)
}

func (s *BetweenDateOpContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(ZitiQlParserWS, i)
}

func (s *BetweenDateOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZitiQlListener); ok {
		listenerT.EnterBetweenDateOp(s)
	}
}

func (s *BetweenDateOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZitiQlListener); ok {
		listenerT.ExitBetweenDateOp(s)
	}
}

type BinaryGreaterThanStringOpContext struct {
	*OperationContext
}

func NewBinaryGreaterThanStringOpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BinaryGreaterThanStringOpContext {
	var p = new(BinaryGreaterThanStringOpContext)

	p.OperationContext = NewEmptyOperationContext()
	p.parser = parser
	p.CopyFrom(ctx.(*OperationContext))

	return p
}

func (s *BinaryGreaterThanStringOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinaryGreaterThanStringOpContext) BinaryLhs() IBinaryLhsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBinaryLhsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBinaryLhsContext)
}

func (s *BinaryGreaterThanStringOpContext) GT() antlr.TerminalNode {
	return s.GetToken(ZitiQlParserGT, 0)
}

func (s *BinaryGreaterThanStringOpContext) STRING() antlr.TerminalNode {
	return s.GetToken(ZitiQlParserSTRING, 0)
}

func (s *BinaryGreaterThanStringOpContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(ZitiQlParserWS)
}

func (s *BinaryGreaterThanStringOpContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(ZitiQlParserWS, i)
}

func (s *BinaryGreaterThanStringOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZitiQlListener); ok {
		listenerT.EnterBinaryGreaterThanStringOp(s)
	}
}

func (s *BinaryGreaterThanStringOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZitiQlListener); ok {
		listenerT.ExitBinaryGreaterThanStringOp(s)
	}
}

type BinaryEqualToNumberOpContext struct {
	*OperationContext
}

func NewBinaryEqualToNumberOpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BinaryEqualToNumberOpContext {
	var p = new(BinaryEqualToNumberOpContext)

	p.OperationContext = NewEmptyOperationContext()
	p.parser = parser
	p.CopyFrom(ctx.(*OperationContext))

	return p
}

func (s *BinaryEqualToNumberOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinaryEqualToNumberOpContext) BinaryLhs() IBinaryLhsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBinaryLhsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBinaryLhsContext)
}

func (s *BinaryEqualToNumberOpContext) EQ() antlr.TerminalNode {
	return s.GetToken(ZitiQlParserEQ, 0)
}

func (s *BinaryEqualToNumberOpContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(ZitiQlParserNUMBER, 0)
}

func (s *BinaryEqualToNumberOpContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(ZitiQlParserWS)
}

func (s *BinaryEqualToNumberOpContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(ZitiQlParserWS, i)
}

func (s *BinaryEqualToNumberOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZitiQlListener); ok {
		listenerT.EnterBinaryEqualToNumberOp(s)
	}
}

func (s *BinaryEqualToNumberOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZitiQlListener); ok {
		listenerT.ExitBinaryEqualToNumberOp(s)
	}
}

type BinaryEqualToBoolOpContext struct {
	*OperationContext
}

func NewBinaryEqualToBoolOpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BinaryEqualToBoolOpContext {
	var p = new(BinaryEqualToBoolOpContext)

	p.OperationContext = NewEmptyOperationContext()
	p.parser = parser
	p.CopyFrom(ctx.(*OperationContext))

	return p
}

func (s *BinaryEqualToBoolOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinaryEqualToBoolOpContext) BinaryLhs() IBinaryLhsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBinaryLhsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBinaryLhsContext)
}

func (s *BinaryEqualToBoolOpContext) EQ() antlr.TerminalNode {
	return s.GetToken(ZitiQlParserEQ, 0)
}

func (s *BinaryEqualToBoolOpContext) BOOL() antlr.TerminalNode {
	return s.GetToken(ZitiQlParserBOOL, 0)
}

func (s *BinaryEqualToBoolOpContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(ZitiQlParserWS)
}

func (s *BinaryEqualToBoolOpContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(ZitiQlParserWS, i)
}

func (s *BinaryEqualToBoolOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZitiQlListener); ok {
		listenerT.EnterBinaryEqualToBoolOp(s)
	}
}

func (s *BinaryEqualToBoolOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZitiQlListener); ok {
		listenerT.ExitBinaryEqualToBoolOp(s)
	}
}

type BinaryEqualToStringOpContext struct {
	*OperationContext
}

func NewBinaryEqualToStringOpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BinaryEqualToStringOpContext {
	var p = new(BinaryEqualToStringOpContext)

	p.OperationContext = NewEmptyOperationContext()
	p.parser = parser
	p.CopyFrom(ctx.(*OperationContext))

	return p
}

func (s *BinaryEqualToStringOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinaryEqualToStringOpContext) BinaryLhs() IBinaryLhsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBinaryLhsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBinaryLhsContext)
}

func (s *BinaryEqualToStringOpContext) EQ() antlr.TerminalNode {
	return s.GetToken(ZitiQlParserEQ, 0)
}

func (s *BinaryEqualToStringOpContext) STRING() antlr.TerminalNode {
	return s.GetToken(ZitiQlParserSTRING, 0)
}

func (s *BinaryEqualToStringOpContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(ZitiQlParserWS)
}

func (s *BinaryEqualToStringOpContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(ZitiQlParserWS, i)
}

func (s *BinaryEqualToStringOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZitiQlListener); ok {
		listenerT.EnterBinaryEqualToStringOp(s)
	}
}

func (s *BinaryEqualToStringOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZitiQlListener); ok {
		listenerT.ExitBinaryEqualToStringOp(s)
	}
}

type BetweenNumberOpContext struct {
	*OperationContext
}

func NewBetweenNumberOpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BetweenNumberOpContext {
	var p = new(BetweenNumberOpContext)

	p.OperationContext = NewEmptyOperationContext()
	p.parser = parser
	p.CopyFrom(ctx.(*OperationContext))

	return p
}

func (s *BetweenNumberOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BetweenNumberOpContext) BinaryLhs() IBinaryLhsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBinaryLhsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBinaryLhsContext)
}

func (s *BetweenNumberOpContext) BETWEEN() antlr.TerminalNode {
	return s.GetToken(ZitiQlParserBETWEEN, 0)
}

func (s *BetweenNumberOpContext) AllNUMBER() []antlr.TerminalNode {
	return s.GetTokens(ZitiQlParserNUMBER)
}

func (s *BetweenNumberOpContext) NUMBER(i int) antlr.TerminalNode {
	return s.GetToken(ZitiQlParserNUMBER, i)
}

func (s *BetweenNumberOpContext) AND() antlr.TerminalNode {
	return s.GetToken(ZitiQlParserAND, 0)
}

func (s *BetweenNumberOpContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(ZitiQlParserWS)
}

func (s *BetweenNumberOpContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(ZitiQlParserWS, i)
}

func (s *BetweenNumberOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZitiQlListener); ok {
		listenerT.EnterBetweenNumberOp(s)
	}
}

func (s *BetweenNumberOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZitiQlListener); ok {
		listenerT.ExitBetweenNumberOp(s)
	}
}

type BinaryContainsOpContext struct {
	*OperationContext
}

func NewBinaryContainsOpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BinaryContainsOpContext {
	var p = new(BinaryContainsOpContext)

	p.OperationContext = NewEmptyOperationContext()
	p.parser = parser
	p.CopyFrom(ctx.(*OperationContext))

	return p
}

func (s *BinaryContainsOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinaryContainsOpContext) BinaryLhs() IBinaryLhsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBinaryLhsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBinaryLhsContext)
}

func (s *BinaryContainsOpContext) CONTAINS() antlr.TerminalNode {
	return s.GetToken(ZitiQlParserCONTAINS, 0)
}

func (s *BinaryContainsOpContext) STRING() antlr.TerminalNode {
	return s.GetToken(ZitiQlParserSTRING, 0)
}

func (s *BinaryContainsOpContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(ZitiQlParserNUMBER, 0)
}

func (s *BinaryContainsOpContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(ZitiQlParserWS)
}

func (s *BinaryContainsOpContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(ZitiQlParserWS, i)
}

func (s *BinaryContainsOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZitiQlListener); ok {
		listenerT.EnterBinaryContainsOp(s)
	}
}

func (s *BinaryContainsOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZitiQlListener); ok {
		listenerT.ExitBinaryContainsOp(s)
	}
}

type BinaryLessThanStringOpContext struct {
	*OperationContext
}

func NewBinaryLessThanStringOpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BinaryLessThanStringOpContext {
	var p = new(BinaryLessThanStringOpContext)

	p.OperationContext = NewEmptyOperationContext()
	p.parser = parser
	p.CopyFrom(ctx.(*OperationContext))

	return p
}

func (s *BinaryLessThanStringOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinaryLessThanStringOpContext) BinaryLhs() IBinaryLhsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBinaryLhsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBinaryLhsContext)
}

func (s *BinaryLessThanStringOpContext) LT() antlr.TerminalNode {
	return s.GetToken(ZitiQlParserLT, 0)
}

func (s *BinaryLessThanStringOpContext) STRING() antlr.TerminalNode {
	return s.GetToken(ZitiQlParserSTRING, 0)
}

func (s *BinaryLessThanStringOpContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(ZitiQlParserWS)
}

func (s *BinaryLessThanStringOpContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(ZitiQlParserWS, i)
}

func (s *BinaryLessThanStringOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZitiQlListener); ok {
		listenerT.EnterBinaryLessThanStringOp(s)
	}
}

func (s *BinaryLessThanStringOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZitiQlListener); ok {
		listenerT.ExitBinaryLessThanStringOp(s)
	}
}

type BinaryEqualToDatetimeOpContext struct {
	*OperationContext
}

func NewBinaryEqualToDatetimeOpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BinaryEqualToDatetimeOpContext {
	var p = new(BinaryEqualToDatetimeOpContext)

	p.OperationContext = NewEmptyOperationContext()
	p.parser = parser
	p.CopyFrom(ctx.(*OperationContext))

	return p
}

func (s *BinaryEqualToDatetimeOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinaryEqualToDatetimeOpContext) BinaryLhs() IBinaryLhsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBinaryLhsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBinaryLhsContext)
}

func (s *BinaryEqualToDatetimeOpContext) EQ() antlr.TerminalNode {
	return s.GetToken(ZitiQlParserEQ, 0)
}

func (s *BinaryEqualToDatetimeOpContext) DATETIME() antlr.TerminalNode {
	return s.GetToken(ZitiQlParserDATETIME, 0)
}

func (s *BinaryEqualToDatetimeOpContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(ZitiQlParserWS)
}

func (s *BinaryEqualToDatetimeOpContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(ZitiQlParserWS, i)
}

func (s *BinaryEqualToDatetimeOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZitiQlListener); ok {
		listenerT.EnterBinaryEqualToDatetimeOp(s)
	}
}

func (s *BinaryEqualToDatetimeOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZitiQlListener); ok {
		listenerT.ExitBinaryEqualToDatetimeOp(s)
	}
}

func (p *ZitiQlParser) Operation() (localctx IOperationContext) {
	localctx = NewOperationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, ZitiQlParserRULE_operation)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(607)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 85, p.GetParserRuleContext()) {
	case 1:
		localctx = NewInStringArrayOpContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(322)
			p.BinaryLhs()
		}
		p.SetState(324)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == ZitiQlParserWS {
			{
				p.SetState(323)
				p.Match(ZitiQlParserWS)
			}

			p.SetState(326)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(328)
			p.Match(ZitiQlParserIN)
		}
		p.SetState(330)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == ZitiQlParserWS {
			{
				p.SetState(329)
				p.Match(ZitiQlParserWS)
			}

			p.SetState(332)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(334)
			p.StringArray()
		}

	case 2:
		localctx = NewInNumberArrayOpContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(336)
			p.BinaryLhs()
		}
		p.SetState(338)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == ZitiQlParserWS {
			{
				p.SetState(337)
				p.Match(ZitiQlParserWS)
			}

			p.SetState(340)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(342)
			p.Match(ZitiQlParserIN)
		}
		p.SetState(344)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == ZitiQlParserWS {
			{
				p.SetState(343)
				p.Match(ZitiQlParserWS)
			}

			p.SetState(346)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(348)
			p.NumberArray()
		}

	case 3:
		localctx = NewInDatetimeArrayOpContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(350)
			p.BinaryLhs()
		}
		p.SetState(352)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == ZitiQlParserWS {
			{
				p.SetState(351)
				p.Match(ZitiQlParserWS)
			}

			p.SetState(354)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(356)
			p.Match(ZitiQlParserIN)
		}
		p.SetState(358)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == ZitiQlParserWS {
			{
				p.SetState(357)
				p.Match(ZitiQlParserWS)
			}

			p.SetState(360)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(362)
			p.DatetimeArray()
		}

	case 4:
		localctx = NewBetweenNumberOpContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(364)
			p.BinaryLhs()
		}
		p.SetState(366)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == ZitiQlParserWS {
			{
				p.SetState(365)
				p.Match(ZitiQlParserWS)
			}

			p.SetState(368)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(370)
			p.Match(ZitiQlParserBETWEEN)
		}
		p.SetState(372)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == ZitiQlParserWS {
			{
				p.SetState(371)
				p.Match(ZitiQlParserWS)
			}

			p.SetState(374)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(376)
			p.Match(ZitiQlParserNUMBER)
		}
		p.SetState(378)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == ZitiQlParserWS {
			{
				p.SetState(377)
				p.Match(ZitiQlParserWS)
			}

			p.SetState(380)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(382)
			p.Match(ZitiQlParserAND)
		}
		p.SetState(384)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == ZitiQlParserWS {
			{
				p.SetState(383)
				p.Match(ZitiQlParserWS)
			}

			p.SetState(386)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(388)
			p.Match(ZitiQlParserNUMBER)
		}

	case 5:
		localctx = NewBetweenDateOpContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(390)
			p.BinaryLhs()
		}
		p.SetState(392)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == ZitiQlParserWS {
			{
				p.SetState(391)
				p.Match(ZitiQlParserWS)
			}

			p.SetState(394)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(396)
			p.Match(ZitiQlParserBETWEEN)
		}
		p.SetState(398)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == ZitiQlParserWS {
			{
				p.SetState(397)
				p.Match(ZitiQlParserWS)
			}

			p.SetState(400)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(402)
			p.Match(ZitiQlParserDATETIME)
		}
		p.SetState(404)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == ZitiQlParserWS {
			{
				p.SetState(403)
				p.Match(ZitiQlParserWS)
			}

			p.SetState(406)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(408)
			p.Match(ZitiQlParserAND)
		}
		p.SetState(410)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == ZitiQlParserWS {
			{
				p.SetState(409)
				p.Match(ZitiQlParserWS)
			}

			p.SetState(412)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(414)
			p.Match(ZitiQlParserDATETIME)
		}

	case 6:
		localctx = NewBinaryLessThanStringOpContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(416)
			p.BinaryLhs()
		}
		p.SetState(420)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ZitiQlParserWS {
			{
				p.SetState(417)
				p.Match(ZitiQlParserWS)
			}

			p.SetState(422)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(423)
			p.Match(ZitiQlParserLT)
		}
		p.SetState(427)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ZitiQlParserWS {
			{
				p.SetState(424)
				p.Match(ZitiQlParserWS)
			}

			p.SetState(429)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(430)
			p.Match(ZitiQlParserSTRING)
		}

	case 7:
		localctx = NewBinaryLessThanNumberOpContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(432)
			p.BinaryLhs()
		}
		p.SetState(436)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ZitiQlParserWS {
			{
				p.SetState(433)
				p.Match(ZitiQlParserWS)
			}

			p.SetState(438)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(439)
			p.Match(ZitiQlParserLT)
		}
		p.SetState(443)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ZitiQlParserWS {
			{
				p.SetState(440)
				p.Match(ZitiQlParserWS)
			}

			p.SetState(445)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(446)
			p.Match(ZitiQlParserNUMBER)
		}

	case 8:
		localctx = NewBinaryLessThanDatetimeOpContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(448)
			p.BinaryLhs()
		}
		p.SetState(452)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ZitiQlParserWS {
			{
				p.SetState(449)
				p.Match(ZitiQlParserWS)
			}

			p.SetState(454)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(455)
			p.Match(ZitiQlParserLT)
		}
		p.SetState(459)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ZitiQlParserWS {
			{
				p.SetState(456)
				p.Match(ZitiQlParserWS)
			}

			p.SetState(461)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(462)
			p.Match(ZitiQlParserDATETIME)
		}

	case 9:
		localctx = NewBinaryGreaterThanStringOpContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(464)
			p.BinaryLhs()
		}
		p.SetState(468)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ZitiQlParserWS {
			{
				p.SetState(465)
				p.Match(ZitiQlParserWS)
			}

			p.SetState(470)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(471)
			p.Match(ZitiQlParserGT)
		}
		p.SetState(475)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ZitiQlParserWS {
			{
				p.SetState(472)
				p.Match(ZitiQlParserWS)
			}

			p.SetState(477)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(478)
			p.Match(ZitiQlParserSTRING)
		}

	case 10:
		localctx = NewBinaryGreaterThanNumberOpContext(p, localctx)
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(480)
			p.BinaryLhs()
		}
		p.SetState(484)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ZitiQlParserWS {
			{
				p.SetState(481)
				p.Match(ZitiQlParserWS)
			}

			p.SetState(486)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(487)
			p.Match(ZitiQlParserGT)
		}
		p.SetState(491)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ZitiQlParserWS {
			{
				p.SetState(488)
				p.Match(ZitiQlParserWS)
			}

			p.SetState(493)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(494)
			p.Match(ZitiQlParserNUMBER)
		}

	case 11:
		localctx = NewBinaryGreaterThanDatetimeOpContext(p, localctx)
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(496)
			p.BinaryLhs()
		}
		p.SetState(500)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ZitiQlParserWS {
			{
				p.SetState(497)
				p.Match(ZitiQlParserWS)
			}

			p.SetState(502)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(503)
			p.Match(ZitiQlParserGT)
		}
		p.SetState(507)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ZitiQlParserWS {
			{
				p.SetState(504)
				p.Match(ZitiQlParserWS)
			}

			p.SetState(509)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(510)
			p.Match(ZitiQlParserDATETIME)
		}

	case 12:
		localctx = NewBinaryEqualToStringOpContext(p, localctx)
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(512)
			p.BinaryLhs()
		}
		p.SetState(516)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ZitiQlParserWS {
			{
				p.SetState(513)
				p.Match(ZitiQlParserWS)
			}

			p.SetState(518)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(519)
			p.Match(ZitiQlParserEQ)
		}
		p.SetState(523)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ZitiQlParserWS {
			{
				p.SetState(520)
				p.Match(ZitiQlParserWS)
			}

			p.SetState(525)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(526)
			p.Match(ZitiQlParserSTRING)
		}

	case 13:
		localctx = NewBinaryEqualToNumberOpContext(p, localctx)
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(528)
			p.BinaryLhs()
		}
		p.SetState(532)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ZitiQlParserWS {
			{
				p.SetState(529)
				p.Match(ZitiQlParserWS)
			}

			p.SetState(534)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(535)
			p.Match(ZitiQlParserEQ)
		}
		p.SetState(539)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ZitiQlParserWS {
			{
				p.SetState(536)
				p.Match(ZitiQlParserWS)
			}

			p.SetState(541)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(542)
			p.Match(ZitiQlParserNUMBER)
		}

	case 14:
		localctx = NewBinaryEqualToDatetimeOpContext(p, localctx)
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(544)
			p.BinaryLhs()
		}
		p.SetState(548)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ZitiQlParserWS {
			{
				p.SetState(545)
				p.Match(ZitiQlParserWS)
			}

			p.SetState(550)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(551)
			p.Match(ZitiQlParserEQ)
		}
		p.SetState(555)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ZitiQlParserWS {
			{
				p.SetState(552)
				p.Match(ZitiQlParserWS)
			}

			p.SetState(557)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(558)
			p.Match(ZitiQlParserDATETIME)
		}

	case 15:
		localctx = NewBinaryEqualToBoolOpContext(p, localctx)
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(560)
			p.BinaryLhs()
		}
		p.SetState(564)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ZitiQlParserWS {
			{
				p.SetState(561)
				p.Match(ZitiQlParserWS)
			}

			p.SetState(566)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(567)
			p.Match(ZitiQlParserEQ)
		}
		p.SetState(571)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ZitiQlParserWS {
			{
				p.SetState(568)
				p.Match(ZitiQlParserWS)
			}

			p.SetState(573)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(574)
			p.Match(ZitiQlParserBOOL)
		}

	case 16:
		localctx = NewBinaryEqualToNullOpContext(p, localctx)
		p.EnterOuterAlt(localctx, 16)
		{
			p.SetState(576)
			p.BinaryLhs()
		}
		p.SetState(580)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ZitiQlParserWS {
			{
				p.SetState(577)
				p.Match(ZitiQlParserWS)
			}

			p.SetState(582)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(583)
			p.Match(ZitiQlParserEQ)
		}
		p.SetState(587)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ZitiQlParserWS {
			{
				p.SetState(584)
				p.Match(ZitiQlParserWS)
			}

			p.SetState(589)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(590)
			p.Match(ZitiQlParserNULL)
		}

	case 17:
		localctx = NewBinaryContainsOpContext(p, localctx)
		p.EnterOuterAlt(localctx, 17)
		{
			p.SetState(592)
			p.BinaryLhs()
		}
		p.SetState(596)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ZitiQlParserWS {
			{
				p.SetState(593)
				p.Match(ZitiQlParserWS)
			}

			p.SetState(598)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(599)
			p.Match(ZitiQlParserCONTAINS)
		}
		p.SetState(601)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == ZitiQlParserWS {
			{
				p.SetState(600)
				p.Match(ZitiQlParserWS)
			}

			p.SetState(603)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(605)
			_la = p.GetTokenStream().LA(1)

			if !(_la == ZitiQlParserSTRING || _la == ZitiQlParserNUMBER) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}

	return localctx
}

// IBinaryLhsContext is an interface to support dynamic dispatch.
type IBinaryLhsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBinaryLhsContext differentiates from other interfaces.
	IsBinaryLhsContext()
}

type BinaryLhsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBinaryLhsContext() *BinaryLhsContext {
	var p = new(BinaryLhsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZitiQlParserRULE_binaryLhs
	return p
}

func (*BinaryLhsContext) IsBinaryLhsContext() {}

func NewBinaryLhsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BinaryLhsContext {
	var p = new(BinaryLhsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZitiQlParserRULE_binaryLhs

	return p
}

func (s *BinaryLhsContext) GetParser() antlr.Parser { return s.parser }

func (s *BinaryLhsContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ZitiQlParserIDENTIFIER, 0)
}

func (s *BinaryLhsContext) SetFunction() ISetFunctionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISetFunctionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISetFunctionContext)
}

func (s *BinaryLhsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinaryLhsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BinaryLhsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZitiQlListener); ok {
		listenerT.EnterBinaryLhs(s)
	}
}

func (s *BinaryLhsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZitiQlListener); ok {
		listenerT.ExitBinaryLhs(s)
	}
}

func (p *ZitiQlParser) BinaryLhs() (localctx IBinaryLhsContext) {
	localctx = NewBinaryLhsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, ZitiQlParserRULE_binaryLhs)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(611)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ZitiQlParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(609)
			p.Match(ZitiQlParserIDENTIFIER)
		}

	case ZitiQlParserALL_OF, ZitiQlParserANY_OF, ZitiQlParserCOUNT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(610)
			p.SetFunction()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ISetFunctionContext is an interface to support dynamic dispatch.
type ISetFunctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSetFunctionContext differentiates from other interfaces.
	IsSetFunctionContext()
}

type SetFunctionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySetFunctionContext() *SetFunctionContext {
	var p = new(SetFunctionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZitiQlParserRULE_setFunction
	return p
}

func (*SetFunctionContext) IsSetFunctionContext() {}

func NewSetFunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SetFunctionContext {
	var p = new(SetFunctionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZitiQlParserRULE_setFunction

	return p
}

func (s *SetFunctionContext) GetParser() antlr.Parser { return s.parser }

func (s *SetFunctionContext) CopyFrom(ctx *SetFunctionContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *SetFunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SetFunctionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SetFunctionExprContext struct {
	*SetFunctionContext
}

func NewSetFunctionExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SetFunctionExprContext {
	var p = new(SetFunctionExprContext)

	p.SetFunctionContext = NewEmptySetFunctionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SetFunctionContext))

	return p
}

func (s *SetFunctionExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SetFunctionExprContext) ALL_OF() antlr.TerminalNode {
	return s.GetToken(ZitiQlParserALL_OF, 0)
}

func (s *SetFunctionExprContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ZitiQlParserLPAREN, 0)
}

func (s *SetFunctionExprContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ZitiQlParserIDENTIFIER, 0)
}

func (s *SetFunctionExprContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ZitiQlParserRPAREN, 0)
}

func (s *SetFunctionExprContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(ZitiQlParserWS)
}

func (s *SetFunctionExprContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(ZitiQlParserWS, i)
}

func (s *SetFunctionExprContext) ANY_OF() antlr.TerminalNode {
	return s.GetToken(ZitiQlParserANY_OF, 0)
}

func (s *SetFunctionExprContext) COUNT() antlr.TerminalNode {
	return s.GetToken(ZitiQlParserCOUNT, 0)
}

func (s *SetFunctionExprContext) SetExpr() ISetExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISetExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISetExprContext)
}

func (s *SetFunctionExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZitiQlListener); ok {
		listenerT.EnterSetFunctionExpr(s)
	}
}

func (s *SetFunctionExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZitiQlListener); ok {
		listenerT.ExitSetFunctionExpr(s)
	}
}

func (p *ZitiQlParser) SetFunction() (localctx ISetFunctionContext) {
	localctx = NewSetFunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, ZitiQlParserRULE_setFunction)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(662)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ZitiQlParserALL_OF:
		localctx = NewSetFunctionExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(613)
			p.Match(ZitiQlParserALL_OF)
		}
		{
			p.SetState(614)
			p.Match(ZitiQlParserLPAREN)
		}
		p.SetState(618)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ZitiQlParserWS {
			{
				p.SetState(615)
				p.Match(ZitiQlParserWS)
			}

			p.SetState(620)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(621)
			p.Match(ZitiQlParserIDENTIFIER)
		}
		p.SetState(625)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ZitiQlParserWS {
			{
				p.SetState(622)
				p.Match(ZitiQlParserWS)
			}

			p.SetState(627)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(628)
			p.Match(ZitiQlParserRPAREN)
		}

	case ZitiQlParserANY_OF:
		localctx = NewSetFunctionExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(629)
			p.Match(ZitiQlParserANY_OF)
		}
		{
			p.SetState(630)
			p.Match(ZitiQlParserLPAREN)
		}
		p.SetState(634)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ZitiQlParserWS {
			{
				p.SetState(631)
				p.Match(ZitiQlParserWS)
			}

			p.SetState(636)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(637)
			p.Match(ZitiQlParserIDENTIFIER)
		}
		p.SetState(641)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ZitiQlParserWS {
			{
				p.SetState(638)
				p.Match(ZitiQlParserWS)
			}

			p.SetState(643)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(644)
			p.Match(ZitiQlParserRPAREN)
		}

	case ZitiQlParserCOUNT:
		localctx = NewSetFunctionExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(645)
			p.Match(ZitiQlParserCOUNT)
		}
		{
			p.SetState(646)
			p.Match(ZitiQlParserLPAREN)
		}
		p.SetState(650)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ZitiQlParserWS {
			{
				p.SetState(647)
				p.Match(ZitiQlParserWS)
			}

			p.SetState(652)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(653)
			p.SetExpr()
		}
		p.SetState(657)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ZitiQlParserWS {
			{
				p.SetState(654)
				p.Match(ZitiQlParserWS)
			}

			p.SetState(659)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(660)
			p.Match(ZitiQlParserRPAREN)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ISetExprContext is an interface to support dynamic dispatch.
type ISetExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSetExprContext differentiates from other interfaces.
	IsSetExprContext()
}

type SetExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySetExprContext() *SetExprContext {
	var p = new(SetExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZitiQlParserRULE_setExpr
	return p
}

func (*SetExprContext) IsSetExprContext() {}

func NewSetExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SetExprContext {
	var p = new(SetExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZitiQlParserRULE_setExpr

	return p
}

func (s *SetExprContext) GetParser() antlr.Parser { return s.parser }

func (s *SetExprContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ZitiQlParserIDENTIFIER, 0)
}

func (s *SetExprContext) SubQueryExpr() ISubQueryExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISubQueryExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISubQueryExprContext)
}

func (s *SetExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SetExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SetExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZitiQlListener); ok {
		listenerT.EnterSetExpr(s)
	}
}

func (s *SetExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZitiQlListener); ok {
		listenerT.ExitSetExpr(s)
	}
}

func (p *ZitiQlParser) SetExpr() (localctx ISetExprContext) {
	localctx = NewSetExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, ZitiQlParserRULE_setExpr)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(666)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ZitiQlParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(664)
			p.Match(ZitiQlParserIDENTIFIER)
		}

	case ZitiQlParserFROM:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(665)
			p.SubQueryExpr()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ISubQueryExprContext is an interface to support dynamic dispatch.
type ISubQueryExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSubQueryExprContext differentiates from other interfaces.
	IsSubQueryExprContext()
}

type SubQueryExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySubQueryExprContext() *SubQueryExprContext {
	var p = new(SubQueryExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZitiQlParserRULE_subQueryExpr
	return p
}

func (*SubQueryExprContext) IsSubQueryExprContext() {}

func NewSubQueryExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SubQueryExprContext {
	var p = new(SubQueryExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZitiQlParserRULE_subQueryExpr

	return p
}

func (s *SubQueryExprContext) GetParser() antlr.Parser { return s.parser }

func (s *SubQueryExprContext) CopyFrom(ctx *SubQueryExprContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *SubQueryExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubQueryExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SubQueryContext struct {
	*SubQueryExprContext
}

func NewSubQueryContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SubQueryContext {
	var p = new(SubQueryContext)

	p.SubQueryExprContext = NewEmptySubQueryExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SubQueryExprContext))

	return p
}

func (s *SubQueryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubQueryContext) FROM() antlr.TerminalNode {
	return s.GetToken(ZitiQlParserFROM, 0)
}

func (s *SubQueryContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ZitiQlParserIDENTIFIER, 0)
}

func (s *SubQueryContext) WHERE() antlr.TerminalNode {
	return s.GetToken(ZitiQlParserWHERE, 0)
}

func (s *SubQueryContext) Query() IQueryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQueryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQueryContext)
}

func (s *SubQueryContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(ZitiQlParserWS)
}

func (s *SubQueryContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(ZitiQlParserWS, i)
}

func (s *SubQueryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZitiQlListener); ok {
		listenerT.EnterSubQuery(s)
	}
}

func (s *SubQueryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZitiQlListener); ok {
		listenerT.ExitSubQuery(s)
	}
}

func (p *ZitiQlParser) SubQueryExpr() (localctx ISubQueryExprContext) {
	localctx = NewSubQueryExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, ZitiQlParserRULE_subQueryExpr)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	localctx = NewSubQueryContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(668)
		p.Match(ZitiQlParserFROM)
	}
	p.SetState(670)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == ZitiQlParserWS {
		{
			p.SetState(669)
			p.Match(ZitiQlParserWS)
		}

		p.SetState(672)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(674)
		p.Match(ZitiQlParserIDENTIFIER)
	}
	p.SetState(676)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == ZitiQlParserWS {
		{
			p.SetState(675)
			p.Match(ZitiQlParserWS)
		}

		p.SetState(678)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(680)
		p.Match(ZitiQlParserWHERE)
	}
	p.SetState(682)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == ZitiQlParserWS {
		{
			p.SetState(681)
			p.Match(ZitiQlParserWS)
		}

		p.SetState(684)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(686)
		p.Query()
	}

	return localctx
}

func (p *ZitiQlParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 9:
		var t *BoolExprContext = nil
		if localctx != nil {
			t = localctx.(*BoolExprContext)
		}
		return p.BoolExpr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *ZitiQlParser) BoolExpr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 5)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
