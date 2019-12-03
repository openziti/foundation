/*
	Copyright 2019 Netfoundry, Inc.

	Licensed under the Apache License, Version 2.0 (the "License");
	you may not use this file except in compliance with the License.
	You may obtain a copy of the License at

	https://www.apache.org/licenses/LICENSE-2.0

	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License.
*/

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 33, 569,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 3, 2, 3, 2, 7, 2, 31, 10, 2, 12, 2, 14, 2, 34, 11,
	2, 3, 2, 3, 2, 7, 2, 38, 10, 2, 12, 2, 14, 2, 41, 11, 2, 3, 2, 3, 2, 7,
	2, 45, 10, 2, 12, 2, 14, 2, 48, 11, 2, 3, 2, 7, 2, 51, 10, 2, 12, 2, 14,
	2, 54, 11, 2, 3, 2, 7, 2, 57, 10, 2, 12, 2, 14, 2, 60, 11, 2, 3, 2, 3,
	2, 3, 3, 3, 3, 7, 3, 66, 10, 3, 12, 3, 14, 3, 69, 11, 3, 3, 3, 3, 3, 7,
	3, 73, 10, 3, 12, 3, 14, 3, 76, 11, 3, 3, 3, 3, 3, 7, 3, 80, 10, 3, 12,
	3, 14, 3, 83, 11, 3, 3, 3, 7, 3, 86, 10, 3, 12, 3, 14, 3, 89, 11, 3, 3,
	3, 7, 3, 92, 10, 3, 12, 3, 14, 3, 95, 11, 3, 3, 3, 3, 3, 3, 4, 3, 4, 7,
	4, 101, 10, 4, 12, 4, 14, 4, 104, 11, 4, 3, 4, 3, 4, 7, 4, 108, 10, 4,
	12, 4, 14, 4, 111, 11, 4, 3, 4, 3, 4, 7, 4, 115, 10, 4, 12, 4, 14, 4, 118,
	11, 4, 3, 4, 7, 4, 121, 10, 4, 12, 4, 14, 4, 124, 11, 4, 3, 4, 7, 4, 127,
	10, 4, 12, 4, 14, 4, 130, 11, 4, 3, 4, 3, 4, 3, 5, 7, 5, 135, 10, 5, 12,
	5, 14, 5, 138, 11, 5, 3, 5, 7, 5, 141, 10, 5, 12, 5, 14, 5, 144, 11, 5,
	3, 5, 7, 5, 147, 10, 5, 12, 5, 14, 5, 150, 11, 5, 3, 5, 3, 5, 3, 6, 3,
	6, 6, 6, 156, 10, 6, 13, 6, 14, 6, 157, 3, 6, 5, 6, 161, 10, 6, 3, 6, 6,
	6, 164, 10, 6, 13, 6, 14, 6, 165, 3, 6, 5, 6, 169, 10, 6, 3, 6, 6, 6, 172,
	10, 6, 13, 6, 14, 6, 173, 3, 6, 5, 6, 177, 10, 6, 3, 7, 3, 7, 6, 7, 181,
	10, 7, 13, 7, 14, 7, 182, 3, 7, 3, 7, 3, 8, 3, 8, 6, 8, 189, 10, 8, 13,
	8, 14, 8, 190, 3, 8, 3, 8, 3, 9, 3, 9, 6, 9, 197, 10, 9, 13, 9, 14, 9,
	198, 3, 9, 3, 9, 6, 9, 203, 10, 9, 13, 9, 14, 9, 204, 3, 9, 3, 9, 7, 9,
	209, 10, 9, 12, 9, 14, 9, 212, 11, 9, 3, 9, 3, 9, 7, 9, 216, 10, 9, 12,
	9, 14, 9, 219, 11, 9, 3, 9, 7, 9, 222, 10, 9, 12, 9, 14, 9, 225, 11, 9,
	3, 10, 3, 10, 6, 10, 229, 10, 10, 13, 10, 14, 10, 230, 3, 10, 5, 10, 234,
	10, 10, 3, 11, 3, 11, 3, 11, 3, 11, 7, 11, 240, 10, 11, 12, 11, 14, 11,
	243, 11, 11, 3, 11, 3, 11, 7, 11, 247, 10, 11, 12, 11, 14, 11, 250, 11,
	11, 3, 11, 3, 11, 3, 11, 5, 11, 255, 10, 11, 3, 11, 3, 11, 6, 11, 259,
	10, 11, 13, 11, 14, 11, 260, 3, 11, 3, 11, 6, 11, 265, 10, 11, 13, 11,
	14, 11, 266, 3, 11, 6, 11, 270, 10, 11, 13, 11, 14, 11, 271, 3, 11, 3,
	11, 6, 11, 276, 10, 11, 13, 11, 14, 11, 277, 3, 11, 3, 11, 6, 11, 282,
	10, 11, 13, 11, 14, 11, 283, 3, 11, 6, 11, 287, 10, 11, 13, 11, 14, 11,
	288, 7, 11, 291, 10, 11, 12, 11, 14, 11, 294, 11, 11, 3, 12, 3, 12, 6,
	12, 298, 10, 12, 13, 12, 14, 12, 299, 3, 12, 3, 12, 6, 12, 304, 10, 12,
	13, 12, 14, 12, 305, 3, 12, 3, 12, 3, 12, 3, 12, 6, 12, 312, 10, 12, 13,
	12, 14, 12, 313, 3, 12, 3, 12, 6, 12, 318, 10, 12, 13, 12, 14, 12, 319,
	3, 12, 3, 12, 3, 12, 3, 12, 6, 12, 326, 10, 12, 13, 12, 14, 12, 327, 3,
	12, 3, 12, 6, 12, 332, 10, 12, 13, 12, 14, 12, 333, 3, 12, 3, 12, 3, 12,
	3, 12, 6, 12, 340, 10, 12, 13, 12, 14, 12, 341, 3, 12, 3, 12, 6, 12, 346,
	10, 12, 13, 12, 14, 12, 347, 3, 12, 3, 12, 6, 12, 352, 10, 12, 13, 12,
	14, 12, 353, 3, 12, 3, 12, 6, 12, 358, 10, 12, 13, 12, 14, 12, 359, 3,
	12, 3, 12, 3, 12, 3, 12, 6, 12, 366, 10, 12, 13, 12, 14, 12, 367, 3, 12,
	3, 12, 6, 12, 372, 10, 12, 13, 12, 14, 12, 373, 3, 12, 3, 12, 6, 12, 378,
	10, 12, 13, 12, 14, 12, 379, 3, 12, 3, 12, 6, 12, 384, 10, 12, 13, 12,
	14, 12, 385, 3, 12, 3, 12, 3, 12, 3, 12, 7, 12, 392, 10, 12, 12, 12, 14,
	12, 395, 11, 12, 3, 12, 3, 12, 7, 12, 399, 10, 12, 12, 12, 14, 12, 402,
	11, 12, 3, 12, 3, 12, 3, 12, 3, 12, 7, 12, 408, 10, 12, 12, 12, 14, 12,
	411, 11, 12, 3, 12, 3, 12, 7, 12, 415, 10, 12, 12, 12, 14, 12, 418, 11,
	12, 3, 12, 3, 12, 3, 12, 3, 12, 7, 12, 424, 10, 12, 12, 12, 14, 12, 427,
	11, 12, 3, 12, 3, 12, 7, 12, 431, 10, 12, 12, 12, 14, 12, 434, 11, 12,
	3, 12, 3, 12, 3, 12, 3, 12, 7, 12, 440, 10, 12, 12, 12, 14, 12, 443, 11,
	12, 3, 12, 3, 12, 7, 12, 447, 10, 12, 12, 12, 14, 12, 450, 11, 12, 3, 12,
	3, 12, 3, 12, 3, 12, 7, 12, 456, 10, 12, 12, 12, 14, 12, 459, 11, 12, 3,
	12, 3, 12, 7, 12, 463, 10, 12, 12, 12, 14, 12, 466, 11, 12, 3, 12, 3, 12,
	3, 12, 3, 12, 7, 12, 472, 10, 12, 12, 12, 14, 12, 475, 11, 12, 3, 12, 3,
	12, 7, 12, 479, 10, 12, 12, 12, 14, 12, 482, 11, 12, 3, 12, 3, 12, 3, 12,
	3, 12, 7, 12, 488, 10, 12, 12, 12, 14, 12, 491, 11, 12, 3, 12, 3, 12, 7,
	12, 495, 10, 12, 12, 12, 14, 12, 498, 11, 12, 3, 12, 3, 12, 3, 12, 3, 12,
	7, 12, 504, 10, 12, 12, 12, 14, 12, 507, 11, 12, 3, 12, 3, 12, 7, 12, 511,
	10, 12, 12, 12, 14, 12, 514, 11, 12, 3, 12, 3, 12, 3, 12, 3, 12, 7, 12,
	520, 10, 12, 12, 12, 14, 12, 523, 11, 12, 3, 12, 3, 12, 7, 12, 527, 10,
	12, 12, 12, 14, 12, 530, 11, 12, 3, 12, 3, 12, 3, 12, 3, 12, 7, 12, 536,
	10, 12, 12, 12, 14, 12, 539, 11, 12, 3, 12, 3, 12, 6, 12, 543, 10, 12,
	13, 12, 14, 12, 544, 3, 12, 3, 12, 5, 12, 549, 10, 12, 3, 13, 3, 13, 5,
	13, 553, 10, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14,
	3, 14, 3, 14, 3, 14, 3, 14, 5, 14, 567, 10, 14, 3, 14, 2, 3, 20, 15, 2,
	4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 2, 5, 4, 2, 23, 23, 31, 31,
	3, 2, 25, 26, 3, 2, 22, 23, 2, 651, 2, 28, 3, 2, 2, 2, 4, 63, 3, 2, 2,
	2, 6, 98, 3, 2, 2, 2, 8, 136, 3, 2, 2, 2, 10, 153, 3, 2, 2, 2, 12, 178,
	3, 2, 2, 2, 14, 186, 3, 2, 2, 2, 16, 194, 3, 2, 2, 2, 18, 226, 3, 2, 2,
	2, 20, 254, 3, 2, 2, 2, 22, 548, 3, 2, 2, 2, 24, 552, 3, 2, 2, 2, 26, 566,
	3, 2, 2, 2, 28, 32, 7, 7, 2, 2, 29, 31, 7, 4, 2, 2, 30, 29, 3, 2, 2, 2,
	31, 34, 3, 2, 2, 2, 32, 30, 3, 2, 2, 2, 32, 33, 3, 2, 2, 2, 33, 35, 3,
	2, 2, 2, 34, 32, 3, 2, 2, 2, 35, 52, 7, 22, 2, 2, 36, 38, 7, 4, 2, 2, 37,
	36, 3, 2, 2, 2, 38, 41, 3, 2, 2, 2, 39, 37, 3, 2, 2, 2, 39, 40, 3, 2, 2,
	2, 40, 42, 3, 2, 2, 2, 41, 39, 3, 2, 2, 2, 42, 46, 7, 3, 2, 2, 43, 45,
	7, 4, 2, 2, 44, 43, 3, 2, 2, 2, 45, 48, 3, 2, 2, 2, 46, 44, 3, 2, 2, 2,
	46, 47, 3, 2, 2, 2, 47, 49, 3, 2, 2, 2, 48, 46, 3, 2, 2, 2, 49, 51, 7,
	22, 2, 2, 50, 39, 3, 2, 2, 2, 51, 54, 3, 2, 2, 2, 52, 50, 3, 2, 2, 2, 52,
	53, 3, 2, 2, 2, 53, 58, 3, 2, 2, 2, 54, 52, 3, 2, 2, 2, 55, 57, 7, 4, 2,
	2, 56, 55, 3, 2, 2, 2, 57, 60, 3, 2, 2, 2, 58, 56, 3, 2, 2, 2, 58, 59,
	3, 2, 2, 2, 59, 61, 3, 2, 2, 2, 60, 58, 3, 2, 2, 2, 61, 62, 7, 8, 2, 2,
	62, 3, 3, 2, 2, 2, 63, 67, 7, 7, 2, 2, 64, 66, 7, 4, 2, 2, 65, 64, 3, 2,
	2, 2, 66, 69, 3, 2, 2, 2, 67, 65, 3, 2, 2, 2, 67, 68, 3, 2, 2, 2, 68, 70,
	3, 2, 2, 2, 69, 67, 3, 2, 2, 2, 70, 87, 7, 23, 2, 2, 71, 73, 7, 4, 2, 2,
	72, 71, 3, 2, 2, 2, 73, 76, 3, 2, 2, 2, 74, 72, 3, 2, 2, 2, 74, 75, 3,
	2, 2, 2, 75, 77, 3, 2, 2, 2, 76, 74, 3, 2, 2, 2, 77, 81, 7, 3, 2, 2, 78,
	80, 7, 4, 2, 2, 79, 78, 3, 2, 2, 2, 80, 83, 3, 2, 2, 2, 81, 79, 3, 2, 2,
	2, 81, 82, 3, 2, 2, 2, 82, 84, 3, 2, 2, 2, 83, 81, 3, 2, 2, 2, 84, 86,
	7, 23, 2, 2, 85, 74, 3, 2, 2, 2, 86, 89, 3, 2, 2, 2, 87, 85, 3, 2, 2, 2,
	87, 88, 3, 2, 2, 2, 88, 93, 3, 2, 2, 2, 89, 87, 3, 2, 2, 2, 90, 92, 7,
	4, 2, 2, 91, 90, 3, 2, 2, 2, 92, 95, 3, 2, 2, 2, 93, 91, 3, 2, 2, 2, 93,
	94, 3, 2, 2, 2, 94, 96, 3, 2, 2, 2, 95, 93, 3, 2, 2, 2, 96, 97, 7, 8, 2,
	2, 97, 5, 3, 2, 2, 2, 98, 102, 7, 7, 2, 2, 99, 101, 7, 4, 2, 2, 100, 99,
	3, 2, 2, 2, 101, 104, 3, 2, 2, 2, 102, 100, 3, 2, 2, 2, 102, 103, 3, 2,
	2, 2, 103, 105, 3, 2, 2, 2, 104, 102, 3, 2, 2, 2, 105, 122, 7, 18, 2, 2,
	106, 108, 7, 4, 2, 2, 107, 106, 3, 2, 2, 2, 108, 111, 3, 2, 2, 2, 109,
	107, 3, 2, 2, 2, 109, 110, 3, 2, 2, 2, 110, 112, 3, 2, 2, 2, 111, 109,
	3, 2, 2, 2, 112, 116, 7, 3, 2, 2, 113, 115, 7, 4, 2, 2, 114, 113, 3, 2,
	2, 2, 115, 118, 3, 2, 2, 2, 116, 114, 3, 2, 2, 2, 116, 117, 3, 2, 2, 2,
	117, 119, 3, 2, 2, 2, 118, 116, 3, 2, 2, 2, 119, 121, 7, 18, 2, 2, 120,
	109, 3, 2, 2, 2, 121, 124, 3, 2, 2, 2, 122, 120, 3, 2, 2, 2, 122, 123,
	3, 2, 2, 2, 123, 128, 3, 2, 2, 2, 124, 122, 3, 2, 2, 2, 125, 127, 7, 4,
	2, 2, 126, 125, 3, 2, 2, 2, 127, 130, 3, 2, 2, 2, 128, 126, 3, 2, 2, 2,
	128, 129, 3, 2, 2, 2, 129, 131, 3, 2, 2, 2, 130, 128, 3, 2, 2, 2, 131,
	132, 7, 8, 2, 2, 132, 7, 3, 2, 2, 2, 133, 135, 7, 4, 2, 2, 134, 133, 3,
	2, 2, 2, 135, 138, 3, 2, 2, 2, 136, 134, 3, 2, 2, 2, 136, 137, 3, 2, 2,
	2, 137, 142, 3, 2, 2, 2, 138, 136, 3, 2, 2, 2, 139, 141, 5, 10, 6, 2, 140,
	139, 3, 2, 2, 2, 141, 144, 3, 2, 2, 2, 142, 140, 3, 2, 2, 2, 142, 143,
	3, 2, 2, 2, 143, 148, 3, 2, 2, 2, 144, 142, 3, 2, 2, 2, 145, 147, 7, 4,
	2, 2, 146, 145, 3, 2, 2, 2, 147, 150, 3, 2, 2, 2, 148, 146, 3, 2, 2, 2,
	148, 149, 3, 2, 2, 2, 149, 151, 3, 2, 2, 2, 150, 148, 3, 2, 2, 2, 151,
	152, 7, 2, 2, 3, 152, 9, 3, 2, 2, 2, 153, 160, 5, 20, 11, 2, 154, 156,
	7, 4, 2, 2, 155, 154, 3, 2, 2, 2, 156, 157, 3, 2, 2, 2, 157, 155, 3, 2,
	2, 2, 157, 158, 3, 2, 2, 2, 158, 159, 3, 2, 2, 2, 159, 161, 5, 16, 9, 2,
	160, 155, 3, 2, 2, 2, 160, 161, 3, 2, 2, 2, 161, 168, 3, 2, 2, 2, 162,
	164, 7, 4, 2, 2, 163, 162, 3, 2, 2, 2, 164, 165, 3, 2, 2, 2, 165, 163,
	3, 2, 2, 2, 165, 166, 3, 2, 2, 2, 166, 167, 3, 2, 2, 2, 167, 169, 5, 12,
	7, 2, 168, 163, 3, 2, 2, 2, 168, 169, 3, 2, 2, 2, 169, 176, 3, 2, 2, 2,
	170, 172, 7, 4, 2, 2, 171, 170, 3, 2, 2, 2, 172, 173, 3, 2, 2, 2, 173,
	171, 3, 2, 2, 2, 173, 174, 3, 2, 2, 2, 174, 175, 3, 2, 2, 2, 175, 177,
	5, 14, 8, 2, 176, 171, 3, 2, 2, 2, 176, 177, 3, 2, 2, 2, 177, 11, 3, 2,
	2, 2, 178, 180, 7, 29, 2, 2, 179, 181, 7, 4, 2, 2, 180, 179, 3, 2, 2, 2,
	181, 182, 3, 2, 2, 2, 182, 180, 3, 2, 2, 2, 182, 183, 3, 2, 2, 2, 183,
	184, 3, 2, 2, 2, 184, 185, 7, 23, 2, 2, 185, 13, 3, 2, 2, 2, 186, 188,
	7, 30, 2, 2, 187, 189, 7, 4, 2, 2, 188, 187, 3, 2, 2, 2, 189, 190, 3, 2,
	2, 2, 190, 188, 3, 2, 2, 2, 190, 191, 3, 2, 2, 2, 191, 192, 3, 2, 2, 2,
	192, 193, 9, 2, 2, 2, 193, 15, 3, 2, 2, 2, 194, 196, 7, 27, 2, 2, 195,
	197, 7, 4, 2, 2, 196, 195, 3, 2, 2, 2, 197, 198, 3, 2, 2, 2, 198, 196,
	3, 2, 2, 2, 198, 199, 3, 2, 2, 2, 199, 200, 3, 2, 2, 2, 200, 202, 7, 28,
	2, 2, 201, 203, 7, 4, 2, 2, 202, 201, 3, 2, 2, 2, 203, 204, 3, 2, 2, 2,
	204, 202, 3, 2, 2, 2, 204, 205, 3, 2, 2, 2, 205, 206, 3, 2, 2, 2, 206,
	223, 5, 18, 10, 2, 207, 209, 7, 4, 2, 2, 208, 207, 3, 2, 2, 2, 209, 212,
	3, 2, 2, 2, 210, 208, 3, 2, 2, 2, 210, 211, 3, 2, 2, 2, 211, 213, 3, 2,
	2, 2, 212, 210, 3, 2, 2, 2, 213, 217, 7, 3, 2, 2, 214, 216, 7, 4, 2, 2,
	215, 214, 3, 2, 2, 2, 216, 219, 3, 2, 2, 2, 217, 215, 3, 2, 2, 2, 217,
	218, 3, 2, 2, 2, 218, 220, 3, 2, 2, 2, 219, 217, 3, 2, 2, 2, 220, 222,
	5, 18, 10, 2, 221, 210, 3, 2, 2, 2, 222, 225, 3, 2, 2, 2, 223, 221, 3,
	2, 2, 2, 223, 224, 3, 2, 2, 2, 224, 17, 3, 2, 2, 2, 225, 223, 3, 2, 2,
	2, 226, 233, 7, 32, 2, 2, 227, 229, 7, 4, 2, 2, 228, 227, 3, 2, 2, 2, 229,
	230, 3, 2, 2, 2, 230, 228, 3, 2, 2, 2, 230, 231, 3, 2, 2, 2, 231, 232,
	3, 2, 2, 2, 232, 234, 9, 3, 2, 2, 233, 228, 3, 2, 2, 2, 233, 234, 3, 2,
	2, 2, 234, 19, 3, 2, 2, 2, 235, 236, 8, 11, 1, 2, 236, 255, 5, 22, 12,
	2, 237, 241, 7, 5, 2, 2, 238, 240, 7, 4, 2, 2, 239, 238, 3, 2, 2, 2, 240,
	243, 3, 2, 2, 2, 241, 239, 3, 2, 2, 2, 241, 242, 3, 2, 2, 2, 242, 244,
	3, 2, 2, 2, 243, 241, 3, 2, 2, 2, 244, 248, 5, 20, 11, 2, 245, 247, 7,
	4, 2, 2, 246, 245, 3, 2, 2, 2, 247, 250, 3, 2, 2, 2, 248, 246, 3, 2, 2,
	2, 248, 249, 3, 2, 2, 2, 249, 251, 3, 2, 2, 2, 250, 248, 3, 2, 2, 2, 251,
	252, 7, 6, 2, 2, 252, 255, 3, 2, 2, 2, 253, 255, 7, 17, 2, 2, 254, 235,
	3, 2, 2, 2, 254, 237, 3, 2, 2, 2, 254, 253, 3, 2, 2, 2, 255, 292, 3, 2,
	2, 2, 256, 269, 12, 5, 2, 2, 257, 259, 7, 4, 2, 2, 258, 257, 3, 2, 2, 2,
	259, 260, 3, 2, 2, 2, 260, 258, 3, 2, 2, 2, 260, 261, 3, 2, 2, 2, 261,
	262, 3, 2, 2, 2, 262, 264, 7, 9, 2, 2, 263, 265, 7, 4, 2, 2, 264, 263,
	3, 2, 2, 2, 265, 266, 3, 2, 2, 2, 266, 264, 3, 2, 2, 2, 266, 267, 3, 2,
	2, 2, 267, 268, 3, 2, 2, 2, 268, 270, 5, 20, 11, 2, 269, 258, 3, 2, 2,
	2, 270, 271, 3, 2, 2, 2, 271, 269, 3, 2, 2, 2, 271, 272, 3, 2, 2, 2, 272,
	291, 3, 2, 2, 2, 273, 286, 12, 4, 2, 2, 274, 276, 7, 4, 2, 2, 275, 274,
	3, 2, 2, 2, 276, 277, 3, 2, 2, 2, 277, 275, 3, 2, 2, 2, 277, 278, 3, 2,
	2, 2, 278, 279, 3, 2, 2, 2, 279, 281, 7, 10, 2, 2, 280, 282, 7, 4, 2, 2,
	281, 280, 3, 2, 2, 2, 282, 283, 3, 2, 2, 2, 283, 281, 3, 2, 2, 2, 283,
	284, 3, 2, 2, 2, 284, 285, 3, 2, 2, 2, 285, 287, 5, 20, 11, 2, 286, 275,
	3, 2, 2, 2, 287, 288, 3, 2, 2, 2, 288, 286, 3, 2, 2, 2, 288, 289, 3, 2,
	2, 2, 289, 291, 3, 2, 2, 2, 290, 256, 3, 2, 2, 2, 290, 273, 3, 2, 2, 2,
	291, 294, 3, 2, 2, 2, 292, 290, 3, 2, 2, 2, 292, 293, 3, 2, 2, 2, 293,
	21, 3, 2, 2, 2, 294, 292, 3, 2, 2, 2, 295, 297, 5, 24, 13, 2, 296, 298,
	7, 4, 2, 2, 297, 296, 3, 2, 2, 2, 298, 299, 3, 2, 2, 2, 299, 297, 3, 2,
	2, 2, 299, 300, 3, 2, 2, 2, 300, 301, 3, 2, 2, 2, 301, 303, 7, 15, 2, 2,
	302, 304, 7, 4, 2, 2, 303, 302, 3, 2, 2, 2, 304, 305, 3, 2, 2, 2, 305,
	303, 3, 2, 2, 2, 305, 306, 3, 2, 2, 2, 306, 307, 3, 2, 2, 2, 307, 308,
	5, 2, 2, 2, 308, 549, 3, 2, 2, 2, 309, 311, 5, 24, 13, 2, 310, 312, 7,
	4, 2, 2, 311, 310, 3, 2, 2, 2, 312, 313, 3, 2, 2, 2, 313, 311, 3, 2, 2,
	2, 313, 314, 3, 2, 2, 2, 314, 315, 3, 2, 2, 2, 315, 317, 7, 15, 2, 2, 316,
	318, 7, 4, 2, 2, 317, 316, 3, 2, 2, 2, 318, 319, 3, 2, 2, 2, 319, 317,
	3, 2, 2, 2, 319, 320, 3, 2, 2, 2, 320, 321, 3, 2, 2, 2, 321, 322, 5, 4,
	3, 2, 322, 549, 3, 2, 2, 2, 323, 325, 5, 24, 13, 2, 324, 326, 7, 4, 2,
	2, 325, 324, 3, 2, 2, 2, 326, 327, 3, 2, 2, 2, 327, 325, 3, 2, 2, 2, 327,
	328, 3, 2, 2, 2, 328, 329, 3, 2, 2, 2, 329, 331, 7, 15, 2, 2, 330, 332,
	7, 4, 2, 2, 331, 330, 3, 2, 2, 2, 332, 333, 3, 2, 2, 2, 333, 331, 3, 2,
	2, 2, 333, 334, 3, 2, 2, 2, 334, 335, 3, 2, 2, 2, 335, 336, 5, 6, 4, 2,
	336, 549, 3, 2, 2, 2, 337, 339, 5, 24, 13, 2, 338, 340, 7, 4, 2, 2, 339,
	338, 3, 2, 2, 2, 340, 341, 3, 2, 2, 2, 341, 339, 3, 2, 2, 2, 341, 342,
	3, 2, 2, 2, 342, 343, 3, 2, 2, 2, 343, 345, 7, 16, 2, 2, 344, 346, 7, 4,
	2, 2, 345, 344, 3, 2, 2, 2, 346, 347, 3, 2, 2, 2, 347, 345, 3, 2, 2, 2,
	347, 348, 3, 2, 2, 2, 348, 349, 3, 2, 2, 2, 349, 351, 7, 23, 2, 2, 350,
	352, 7, 4, 2, 2, 351, 350, 3, 2, 2, 2, 352, 353, 3, 2, 2, 2, 353, 351,
	3, 2, 2, 2, 353, 354, 3, 2, 2, 2, 354, 355, 3, 2, 2, 2, 355, 357, 7, 9,
	2, 2, 356, 358, 7, 4, 2, 2, 357, 356, 3, 2, 2, 2, 358, 359, 3, 2, 2, 2,
	359, 357, 3, 2, 2, 2, 359, 360, 3, 2, 2, 2, 360, 361, 3, 2, 2, 2, 361,
	362, 7, 23, 2, 2, 362, 549, 3, 2, 2, 2, 363, 365, 5, 24, 13, 2, 364, 366,
	7, 4, 2, 2, 365, 364, 3, 2, 2, 2, 366, 367, 3, 2, 2, 2, 367, 365, 3, 2,
	2, 2, 367, 368, 3, 2, 2, 2, 368, 369, 3, 2, 2, 2, 369, 371, 7, 16, 2, 2,
	370, 372, 7, 4, 2, 2, 371, 370, 3, 2, 2, 2, 372, 373, 3, 2, 2, 2, 373,
	371, 3, 2, 2, 2, 373, 374, 3, 2, 2, 2, 374, 375, 3, 2, 2, 2, 375, 377,
	7, 18, 2, 2, 376, 378, 7, 4, 2, 2, 377, 376, 3, 2, 2, 2, 378, 379, 3, 2,
	2, 2, 379, 377, 3, 2, 2, 2, 379, 380, 3, 2, 2, 2, 380, 381, 3, 2, 2, 2,
	381, 383, 7, 9, 2, 2, 382, 384, 7, 4, 2, 2, 383, 382, 3, 2, 2, 2, 384,
	385, 3, 2, 2, 2, 385, 383, 3, 2, 2, 2, 385, 386, 3, 2, 2, 2, 386, 387,
	3, 2, 2, 2, 387, 388, 7, 18, 2, 2, 388, 549, 3, 2, 2, 2, 389, 393, 5, 24,
	13, 2, 390, 392, 7, 4, 2, 2, 391, 390, 3, 2, 2, 2, 392, 395, 3, 2, 2, 2,
	393, 391, 3, 2, 2, 2, 393, 394, 3, 2, 2, 2, 394, 396, 3, 2, 2, 2, 395,
	393, 3, 2, 2, 2, 396, 400, 7, 11, 2, 2, 397, 399, 7, 4, 2, 2, 398, 397,
	3, 2, 2, 2, 399, 402, 3, 2, 2, 2, 400, 398, 3, 2, 2, 2, 400, 401, 3, 2,
	2, 2, 401, 403, 3, 2, 2, 2, 402, 400, 3, 2, 2, 2, 403, 404, 7, 23, 2, 2,
	404, 549, 3, 2, 2, 2, 405, 409, 5, 24, 13, 2, 406, 408, 7, 4, 2, 2, 407,
	406, 3, 2, 2, 2, 408, 411, 3, 2, 2, 2, 409, 407, 3, 2, 2, 2, 409, 410,
	3, 2, 2, 2, 410, 412, 3, 2, 2, 2, 411, 409, 3, 2, 2, 2, 412, 416, 7, 11,
	2, 2, 413, 415, 7, 4, 2, 2, 414, 413, 3, 2, 2, 2, 415, 418, 3, 2, 2, 2,
	416, 414, 3, 2, 2, 2, 416, 417, 3, 2, 2, 2, 417, 419, 3, 2, 2, 2, 418,
	416, 3, 2, 2, 2, 419, 420, 7, 18, 2, 2, 420, 549, 3, 2, 2, 2, 421, 425,
	5, 24, 13, 2, 422, 424, 7, 4, 2, 2, 423, 422, 3, 2, 2, 2, 424, 427, 3,
	2, 2, 2, 425, 423, 3, 2, 2, 2, 425, 426, 3, 2, 2, 2, 426, 428, 3, 2, 2,
	2, 427, 425, 3, 2, 2, 2, 428, 432, 7, 12, 2, 2, 429, 431, 7, 4, 2, 2, 430,
	429, 3, 2, 2, 2, 431, 434, 3, 2, 2, 2, 432, 430, 3, 2, 2, 2, 432, 433,
	3, 2, 2, 2, 433, 435, 3, 2, 2, 2, 434, 432, 3, 2, 2, 2, 435, 436, 7, 23,
	2, 2, 436, 549, 3, 2, 2, 2, 437, 441, 5, 24, 13, 2, 438, 440, 7, 4, 2,
	2, 439, 438, 3, 2, 2, 2, 440, 443, 3, 2, 2, 2, 441, 439, 3, 2, 2, 2, 441,
	442, 3, 2, 2, 2, 442, 444, 3, 2, 2, 2, 443, 441, 3, 2, 2, 2, 444, 448,
	7, 12, 2, 2, 445, 447, 7, 4, 2, 2, 446, 445, 3, 2, 2, 2, 447, 450, 3, 2,
	2, 2, 448, 446, 3, 2, 2, 2, 448, 449, 3, 2, 2, 2, 449, 451, 3, 2, 2, 2,
	450, 448, 3, 2, 2, 2, 451, 452, 7, 18, 2, 2, 452, 549, 3, 2, 2, 2, 453,
	457, 5, 24, 13, 2, 454, 456, 7, 4, 2, 2, 455, 454, 3, 2, 2, 2, 456, 459,
	3, 2, 2, 2, 457, 455, 3, 2, 2, 2, 457, 458, 3, 2, 2, 2, 458, 460, 3, 2,
	2, 2, 459, 457, 3, 2, 2, 2, 460, 464, 7, 13, 2, 2, 461, 463, 7, 4, 2, 2,
	462, 461, 3, 2, 2, 2, 463, 466, 3, 2, 2, 2, 464, 462, 3, 2, 2, 2, 464,
	465, 3, 2, 2, 2, 465, 467, 3, 2, 2, 2, 466, 464, 3, 2, 2, 2, 467, 468,
	7, 22, 2, 2, 468, 549, 3, 2, 2, 2, 469, 473, 5, 24, 13, 2, 470, 472, 7,
	4, 2, 2, 471, 470, 3, 2, 2, 2, 472, 475, 3, 2, 2, 2, 473, 471, 3, 2, 2,
	2, 473, 474, 3, 2, 2, 2, 474, 476, 3, 2, 2, 2, 475, 473, 3, 2, 2, 2, 476,
	480, 7, 13, 2, 2, 477, 479, 7, 4, 2, 2, 478, 477, 3, 2, 2, 2, 479, 482,
	3, 2, 2, 2, 480, 478, 3, 2, 2, 2, 480, 481, 3, 2, 2, 2, 481, 483, 3, 2,
	2, 2, 482, 480, 3, 2, 2, 2, 483, 484, 7, 23, 2, 2, 484, 549, 3, 2, 2, 2,
	485, 489, 5, 24, 13, 2, 486, 488, 7, 4, 2, 2, 487, 486, 3, 2, 2, 2, 488,
	491, 3, 2, 2, 2, 489, 487, 3, 2, 2, 2, 489, 490, 3, 2, 2, 2, 490, 492,
	3, 2, 2, 2, 491, 489, 3, 2, 2, 2, 492, 496, 7, 13, 2, 2, 493, 495, 7, 4,
	2, 2, 494, 493, 3, 2, 2, 2, 495, 498, 3, 2, 2, 2, 496, 494, 3, 2, 2, 2,
	496, 497, 3, 2, 2, 2, 497, 499, 3, 2, 2, 2, 498, 496, 3, 2, 2, 2, 499,
	500, 7, 18, 2, 2, 500, 549, 3, 2, 2, 2, 501, 505, 5, 24, 13, 2, 502, 504,
	7, 4, 2, 2, 503, 502, 3, 2, 2, 2, 504, 507, 3, 2, 2, 2, 505, 503, 3, 2,
	2, 2, 505, 506, 3, 2, 2, 2, 506, 508, 3, 2, 2, 2, 507, 505, 3, 2, 2, 2,
	508, 512, 7, 13, 2, 2, 509, 511, 7, 4, 2, 2, 510, 509, 3, 2, 2, 2, 511,
	514, 3, 2, 2, 2, 512, 510, 3, 2, 2, 2, 512, 513, 3, 2, 2, 2, 513, 515,
	3, 2, 2, 2, 514, 512, 3, 2, 2, 2, 515, 516, 7, 17, 2, 2, 516, 549, 3, 2,
	2, 2, 517, 521, 5, 24, 13, 2, 518, 520, 7, 4, 2, 2, 519, 518, 3, 2, 2,
	2, 520, 523, 3, 2, 2, 2, 521, 519, 3, 2, 2, 2, 521, 522, 3, 2, 2, 2, 522,
	524, 3, 2, 2, 2, 523, 521, 3, 2, 2, 2, 524, 528, 7, 13, 2, 2, 525, 527,
	7, 4, 2, 2, 526, 525, 3, 2, 2, 2, 527, 530, 3, 2, 2, 2, 528, 526, 3, 2,
	2, 2, 528, 529, 3, 2, 2, 2, 529, 531, 3, 2, 2, 2, 530, 528, 3, 2, 2, 2,
	531, 532, 7, 24, 2, 2, 532, 549, 3, 2, 2, 2, 533, 537, 5, 24, 13, 2, 534,
	536, 7, 4, 2, 2, 535, 534, 3, 2, 2, 2, 536, 539, 3, 2, 2, 2, 537, 535,
	3, 2, 2, 2, 537, 538, 3, 2, 2, 2, 538, 540, 3, 2, 2, 2, 539, 537, 3, 2,
	2, 2, 540, 542, 7, 14, 2, 2, 541, 543, 7, 4, 2, 2, 542, 541, 3, 2, 2, 2,
	543, 544, 3, 2, 2, 2, 544, 542, 3, 2, 2, 2, 544, 545, 3, 2, 2, 2, 545,
	546, 3, 2, 2, 2, 546, 547, 9, 4, 2, 2, 547, 549, 3, 2, 2, 2, 548, 295,
	3, 2, 2, 2, 548, 309, 3, 2, 2, 2, 548, 323, 3, 2, 2, 2, 548, 337, 3, 2,
	2, 2, 548, 363, 3, 2, 2, 2, 548, 389, 3, 2, 2, 2, 548, 405, 3, 2, 2, 2,
	548, 421, 3, 2, 2, 2, 548, 437, 3, 2, 2, 2, 548, 453, 3, 2, 2, 2, 548,
	469, 3, 2, 2, 2, 548, 485, 3, 2, 2, 2, 548, 501, 3, 2, 2, 2, 548, 517,
	3, 2, 2, 2, 548, 533, 3, 2, 2, 2, 549, 23, 3, 2, 2, 2, 550, 553, 7, 32,
	2, 2, 551, 553, 5, 26, 14, 2, 552, 550, 3, 2, 2, 2, 552, 551, 3, 2, 2,
	2, 553, 25, 3, 2, 2, 2, 554, 555, 7, 19, 2, 2, 555, 556, 7, 5, 2, 2, 556,
	557, 7, 32, 2, 2, 557, 567, 7, 6, 2, 2, 558, 559, 7, 20, 2, 2, 559, 560,
	7, 5, 2, 2, 560, 561, 7, 32, 2, 2, 561, 567, 7, 6, 2, 2, 562, 563, 7, 21,
	2, 2, 563, 564, 7, 5, 2, 2, 564, 565, 7, 32, 2, 2, 565, 567, 7, 6, 2, 2,
	566, 554, 3, 2, 2, 2, 566, 558, 3, 2, 2, 2, 566, 562, 3, 2, 2, 2, 567,
	27, 3, 2, 2, 2, 83, 32, 39, 46, 52, 58, 67, 74, 81, 87, 93, 102, 109, 116,
	122, 128, 136, 142, 148, 157, 160, 165, 168, 173, 176, 182, 190, 198, 204,
	210, 217, 223, 230, 233, 241, 248, 254, 260, 266, 271, 277, 283, 288, 290,
	292, 299, 305, 313, 319, 327, 333, 341, 347, 353, 359, 367, 373, 379, 385,
	393, 400, 409, 416, 425, 432, 441, 448, 457, 464, 473, 480, 489, 496, 505,
	512, 521, 528, 537, 544, 548, 552, 566,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "','", "", "'('", "')'", "'['", "']'",
}
var symbolicNames = []string{
	"", "", "WS", "LPAREN", "RPAREN", "LBRACKET", "RBRACKET", "AND", "OR",
	"LT", "GT", "EQ", "CONTAINS", "IN", "BETWEEN", "BOOL", "DATETIME", "ALL_OF",
	"NONE_OF", "ANY_OF", "STRING", "NUMBER", "NULL", "ASC", "DESC", "SORT",
	"BY", "SKIP_ROWS", "LIMIT_ROWS", "NONE", "IDENTIFIER", "RFC3339_DATE_TIME",
}

var ruleNames = []string{
	"string_array", "number_array", "datetime_array", "start", "query", "skip",
	"limit", "sortBy", "sortField", "expression", "operation", "binary_lhs",
	"set_function",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type ZitiQlParser struct {
	*antlr.BaseParser
}

func NewZitiQlParser(input antlr.TokenStream) *ZitiQlParser {
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
	ZitiQlParserNONE_OF           = 18
	ZitiQlParserANY_OF            = 19
	ZitiQlParserSTRING            = 20
	ZitiQlParserNUMBER            = 21
	ZitiQlParserNULL              = 22
	ZitiQlParserASC               = 23
	ZitiQlParserDESC              = 24
	ZitiQlParserSORT              = 25
	ZitiQlParserBY                = 26
	ZitiQlParserSKIP_ROWS         = 27
	ZitiQlParserLIMIT_ROWS        = 28
	ZitiQlParserNONE              = 29
	ZitiQlParserIDENTIFIER        = 30
	ZitiQlParserRFC3339_DATE_TIME = 31
)

// ZitiQlParser rules.
const (
	ZitiQlParserRULE_string_array   = 0
	ZitiQlParserRULE_number_array   = 1
	ZitiQlParserRULE_datetime_array = 2
	ZitiQlParserRULE_start          = 3
	ZitiQlParserRULE_query          = 4
	ZitiQlParserRULE_skip           = 5
	ZitiQlParserRULE_limit          = 6
	ZitiQlParserRULE_sortBy         = 7
	ZitiQlParserRULE_sortField      = 8
	ZitiQlParserRULE_expression     = 9
	ZitiQlParserRULE_operation      = 10
	ZitiQlParserRULE_binary_lhs     = 11
	ZitiQlParserRULE_set_function   = 12
)

// IString_arrayContext is an interface to support dynamic dispatch.
type IString_arrayContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsString_arrayContext differentiates from other interfaces.
	IsString_arrayContext()
}

type String_arrayContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyString_arrayContext() *String_arrayContext {
	var p = new(String_arrayContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZitiQlParserRULE_string_array
	return p
}

func (*String_arrayContext) IsString_arrayContext() {}

func NewString_arrayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *String_arrayContext {
	var p = new(String_arrayContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZitiQlParserRULE_string_array

	return p
}

func (s *String_arrayContext) GetParser() antlr.Parser { return s.parser }

func (s *String_arrayContext) LBRACKET() antlr.TerminalNode {
	return s.GetToken(ZitiQlParserLBRACKET, 0)
}

func (s *String_arrayContext) AllSTRING() []antlr.TerminalNode {
	return s.GetTokens(ZitiQlParserSTRING)
}

func (s *String_arrayContext) STRING(i int) antlr.TerminalNode {
	return s.GetToken(ZitiQlParserSTRING, i)
}

func (s *String_arrayContext) RBRACKET() antlr.TerminalNode {
	return s.GetToken(ZitiQlParserRBRACKET, 0)
}

func (s *String_arrayContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(ZitiQlParserWS)
}

func (s *String_arrayContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(ZitiQlParserWS, i)
}

func (s *String_arrayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *String_arrayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *String_arrayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZitiQlListener); ok {
		listenerT.EnterString_array(s)
	}
}

func (s *String_arrayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZitiQlListener); ok {
		listenerT.ExitString_array(s)
	}
}

func (p *ZitiQlParser) String_array() (localctx IString_arrayContext) {
	localctx = NewString_arrayContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ZitiQlParserRULE_string_array)
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
		p.SetState(26)
		p.Match(ZitiQlParserLBRACKET)
	}
	p.SetState(30)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ZitiQlParserWS {
		{
			p.SetState(27)
			p.Match(ZitiQlParserWS)
		}

		p.SetState(32)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(33)
		p.Match(ZitiQlParserSTRING)
	}
	p.SetState(50)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(37)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == ZitiQlParserWS {
				{
					p.SetState(34)
					p.Match(ZitiQlParserWS)
				}

				p.SetState(39)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(40)
				p.Match(ZitiQlParserT__0)
			}
			p.SetState(44)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == ZitiQlParserWS {
				{
					p.SetState(41)
					p.Match(ZitiQlParserWS)
				}

				p.SetState(46)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(47)
				p.Match(ZitiQlParserSTRING)
			}

		}
		p.SetState(52)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())
	}
	p.SetState(56)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ZitiQlParserWS {
		{
			p.SetState(53)
			p.Match(ZitiQlParserWS)
		}

		p.SetState(58)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(59)
		p.Match(ZitiQlParserRBRACKET)
	}

	return localctx
}

// INumber_arrayContext is an interface to support dynamic dispatch.
type INumber_arrayContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNumber_arrayContext differentiates from other interfaces.
	IsNumber_arrayContext()
}

type Number_arrayContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumber_arrayContext() *Number_arrayContext {
	var p = new(Number_arrayContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZitiQlParserRULE_number_array
	return p
}

func (*Number_arrayContext) IsNumber_arrayContext() {}

func NewNumber_arrayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Number_arrayContext {
	var p = new(Number_arrayContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZitiQlParserRULE_number_array

	return p
}

func (s *Number_arrayContext) GetParser() antlr.Parser { return s.parser }

func (s *Number_arrayContext) LBRACKET() antlr.TerminalNode {
	return s.GetToken(ZitiQlParserLBRACKET, 0)
}

func (s *Number_arrayContext) AllNUMBER() []antlr.TerminalNode {
	return s.GetTokens(ZitiQlParserNUMBER)
}

func (s *Number_arrayContext) NUMBER(i int) antlr.TerminalNode {
	return s.GetToken(ZitiQlParserNUMBER, i)
}

func (s *Number_arrayContext) RBRACKET() antlr.TerminalNode {
	return s.GetToken(ZitiQlParserRBRACKET, 0)
}

func (s *Number_arrayContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(ZitiQlParserWS)
}

func (s *Number_arrayContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(ZitiQlParserWS, i)
}

func (s *Number_arrayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Number_arrayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Number_arrayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZitiQlListener); ok {
		listenerT.EnterNumber_array(s)
	}
}

func (s *Number_arrayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZitiQlListener); ok {
		listenerT.ExitNumber_array(s)
	}
}

func (p *ZitiQlParser) Number_array() (localctx INumber_arrayContext) {
	localctx = NewNumber_arrayContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ZitiQlParserRULE_number_array)
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
		p.SetState(61)
		p.Match(ZitiQlParserLBRACKET)
	}
	p.SetState(65)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ZitiQlParserWS {
		{
			p.SetState(62)
			p.Match(ZitiQlParserWS)
		}

		p.SetState(67)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(68)
		p.Match(ZitiQlParserNUMBER)
	}
	p.SetState(85)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(72)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == ZitiQlParserWS {
				{
					p.SetState(69)
					p.Match(ZitiQlParserWS)
				}

				p.SetState(74)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(75)
				p.Match(ZitiQlParserT__0)
			}
			p.SetState(79)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == ZitiQlParserWS {
				{
					p.SetState(76)
					p.Match(ZitiQlParserWS)
				}

				p.SetState(81)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(82)
				p.Match(ZitiQlParserNUMBER)
			}

		}
		p.SetState(87)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())
	}
	p.SetState(91)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ZitiQlParserWS {
		{
			p.SetState(88)
			p.Match(ZitiQlParserWS)
		}

		p.SetState(93)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(94)
		p.Match(ZitiQlParserRBRACKET)
	}

	return localctx
}

// IDatetime_arrayContext is an interface to support dynamic dispatch.
type IDatetime_arrayContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDatetime_arrayContext differentiates from other interfaces.
	IsDatetime_arrayContext()
}

type Datetime_arrayContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDatetime_arrayContext() *Datetime_arrayContext {
	var p = new(Datetime_arrayContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZitiQlParserRULE_datetime_array
	return p
}

func (*Datetime_arrayContext) IsDatetime_arrayContext() {}

func NewDatetime_arrayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Datetime_arrayContext {
	var p = new(Datetime_arrayContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZitiQlParserRULE_datetime_array

	return p
}

func (s *Datetime_arrayContext) GetParser() antlr.Parser { return s.parser }

func (s *Datetime_arrayContext) LBRACKET() antlr.TerminalNode {
	return s.GetToken(ZitiQlParserLBRACKET, 0)
}

func (s *Datetime_arrayContext) AllDATETIME() []antlr.TerminalNode {
	return s.GetTokens(ZitiQlParserDATETIME)
}

func (s *Datetime_arrayContext) DATETIME(i int) antlr.TerminalNode {
	return s.GetToken(ZitiQlParserDATETIME, i)
}

func (s *Datetime_arrayContext) RBRACKET() antlr.TerminalNode {
	return s.GetToken(ZitiQlParserRBRACKET, 0)
}

func (s *Datetime_arrayContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(ZitiQlParserWS)
}

func (s *Datetime_arrayContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(ZitiQlParserWS, i)
}

func (s *Datetime_arrayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Datetime_arrayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Datetime_arrayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZitiQlListener); ok {
		listenerT.EnterDatetime_array(s)
	}
}

func (s *Datetime_arrayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZitiQlListener); ok {
		listenerT.ExitDatetime_array(s)
	}
}

func (p *ZitiQlParser) Datetime_array() (localctx IDatetime_arrayContext) {
	localctx = NewDatetime_arrayContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ZitiQlParserRULE_datetime_array)
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
		p.SetState(96)
		p.Match(ZitiQlParserLBRACKET)
	}
	p.SetState(100)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ZitiQlParserWS {
		{
			p.SetState(97)
			p.Match(ZitiQlParserWS)
		}

		p.SetState(102)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(103)
		p.Match(ZitiQlParserDATETIME)
	}
	p.SetState(120)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(107)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == ZitiQlParserWS {
				{
					p.SetState(104)
					p.Match(ZitiQlParserWS)
				}

				p.SetState(109)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(110)
				p.Match(ZitiQlParserT__0)
			}
			p.SetState(114)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == ZitiQlParserWS {
				{
					p.SetState(111)
					p.Match(ZitiQlParserWS)
				}

				p.SetState(116)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(117)
				p.Match(ZitiQlParserDATETIME)
			}

		}
		p.SetState(122)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext())
	}
	p.SetState(126)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ZitiQlParserWS {
		{
			p.SetState(123)
			p.Match(ZitiQlParserWS)
		}

		p.SetState(128)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(129)
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
	p.SetState(134)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(131)
				p.Match(ZitiQlParserWS)
			}

		}
		p.SetState(136)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext())
	}
	p.SetState(140)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ZitiQlParserLPAREN)|(1<<ZitiQlParserBOOL)|(1<<ZitiQlParserALL_OF)|(1<<ZitiQlParserNONE_OF)|(1<<ZitiQlParserANY_OF)|(1<<ZitiQlParserIDENTIFIER))) != 0 {
		{
			p.SetState(137)
			p.Query()
		}

		p.SetState(142)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(146)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ZitiQlParserWS {
		{
			p.SetState(143)
			p.Match(ZitiQlParserWS)
		}

		p.SetState(148)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(149)
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

func (s *QueryStmtContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
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
		p.SetState(151)
		p.expression(0)
	}
	p.SetState(158)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext()) == 1 {
		p.SetState(153)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == ZitiQlParserWS {
			{
				p.SetState(152)
				p.Match(ZitiQlParserWS)
			}

			p.SetState(155)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(157)
			p.SortBy()
		}

	}
	p.SetState(166)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext()) == 1 {
		p.SetState(161)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == ZitiQlParserWS {
			{
				p.SetState(160)
				p.Match(ZitiQlParserWS)
			}

			p.SetState(163)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(165)
			p.Skip()
		}

	}
	p.SetState(174)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext()) == 1 {
		p.SetState(169)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == ZitiQlParserWS {
			{
				p.SetState(168)
				p.Match(ZitiQlParserWS)
			}

			p.SetState(171)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(173)
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
		p.SetState(176)
		p.Match(ZitiQlParserSKIP_ROWS)
	}
	p.SetState(178)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == ZitiQlParserWS {
		{
			p.SetState(177)
			p.Match(ZitiQlParserWS)
		}

		p.SetState(180)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(182)
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
		p.SetState(184)
		p.Match(ZitiQlParserLIMIT_ROWS)
	}
	p.SetState(186)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == ZitiQlParserWS {
		{
			p.SetState(185)
			p.Match(ZitiQlParserWS)
		}

		p.SetState(188)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(190)
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
		p.SetState(192)
		p.Match(ZitiQlParserSORT)
	}
	p.SetState(194)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == ZitiQlParserWS {
		{
			p.SetState(193)
			p.Match(ZitiQlParserWS)
		}

		p.SetState(196)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(198)
		p.Match(ZitiQlParserBY)
	}
	p.SetState(200)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == ZitiQlParserWS {
		{
			p.SetState(199)
			p.Match(ZitiQlParserWS)
		}

		p.SetState(202)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(204)
		p.SortField()
	}
	p.SetState(221)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(208)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == ZitiQlParserWS {
				{
					p.SetState(205)
					p.Match(ZitiQlParserWS)
				}

				p.SetState(210)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(211)
				p.Match(ZitiQlParserT__0)
			}
			p.SetState(215)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == ZitiQlParserWS {
				{
					p.SetState(212)
					p.Match(ZitiQlParserWS)
				}

				p.SetState(217)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(218)
				p.SortField()
			}

		}
		p.SetState(223)
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
		p.SetState(224)
		p.Match(ZitiQlParserIDENTIFIER)
	}
	p.SetState(231)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 32, p.GetParserRuleContext()) == 1 {
		p.SetState(226)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == ZitiQlParserWS {
			{
				p.SetState(225)
				p.Match(ZitiQlParserWS)
			}

			p.SetState(228)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(230)
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

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZitiQlParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZitiQlParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) CopyFrom(ctx *ExpressionContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type GroupContext struct {
	*ExpressionContext
}

func NewGroupContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *GroupContext {
	var p = new(GroupContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *GroupContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GroupContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ZitiQlParserLPAREN, 0)
}

func (s *GroupContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
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
	*ExpressionContext
}

func NewBoolConstContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BoolConstContext {
	var p = new(BoolConstContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

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

type OrConjunctionContext struct {
	*ExpressionContext
}

func NewOrConjunctionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OrConjunctionContext {
	var p = new(OrConjunctionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *OrConjunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrConjunctionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *OrConjunctionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *OrConjunctionContext) AllOR() []antlr.TerminalNode {
	return s.GetTokens(ZitiQlParserOR)
}

func (s *OrConjunctionContext) OR(i int) antlr.TerminalNode {
	return s.GetToken(ZitiQlParserOR, i)
}

func (s *OrConjunctionContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(ZitiQlParserWS)
}

func (s *OrConjunctionContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(ZitiQlParserWS, i)
}

func (s *OrConjunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZitiQlListener); ok {
		listenerT.EnterOrConjunction(s)
	}
}

func (s *OrConjunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZitiQlListener); ok {
		listenerT.ExitOrConjunction(s)
	}
}

type OperationOpContext struct {
	*ExpressionContext
}

func NewOperationOpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OperationOpContext {
	var p = new(OperationOpContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

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

type AndConjunctionContext struct {
	*ExpressionContext
}

func NewAndConjunctionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AndConjunctionContext {
	var p = new(AndConjunctionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *AndConjunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndConjunctionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *AndConjunctionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AndConjunctionContext) AllAND() []antlr.TerminalNode {
	return s.GetTokens(ZitiQlParserAND)
}

func (s *AndConjunctionContext) AND(i int) antlr.TerminalNode {
	return s.GetToken(ZitiQlParserAND, i)
}

func (s *AndConjunctionContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(ZitiQlParserWS)
}

func (s *AndConjunctionContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(ZitiQlParserWS, i)
}

func (s *AndConjunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZitiQlListener); ok {
		listenerT.EnterAndConjunction(s)
	}
}

func (s *AndConjunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZitiQlListener); ok {
		listenerT.ExitAndConjunction(s)
	}
}

func (p *ZitiQlParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *ZitiQlParser) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 18
	p.EnterRecursionRule(localctx, 18, ZitiQlParserRULE_expression, _p)
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
	p.SetState(252)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ZitiQlParserALL_OF, ZitiQlParserNONE_OF, ZitiQlParserANY_OF, ZitiQlParserIDENTIFIER:
		localctx = NewOperationOpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(234)
			p.Operation()
		}

	case ZitiQlParserLPAREN:
		localctx = NewGroupContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(235)
			p.Match(ZitiQlParserLPAREN)
		}
		p.SetState(239)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ZitiQlParserWS {
			{
				p.SetState(236)
				p.Match(ZitiQlParserWS)
			}

			p.SetState(241)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(242)
			p.expression(0)
		}
		p.SetState(246)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ZitiQlParserWS {
			{
				p.SetState(243)
				p.Match(ZitiQlParserWS)
			}

			p.SetState(248)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(249)
			p.Match(ZitiQlParserRPAREN)
		}

	case ZitiQlParserBOOL:
		localctx = NewBoolConstContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(251)
			p.Match(ZitiQlParserBOOL)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(290)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 43, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(288)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 42, p.GetParserRuleContext()) {
			case 1:
				localctx = NewAndConjunctionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ZitiQlParserRULE_expression)
				p.SetState(254)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				p.SetState(267)
				p.GetErrorHandler().Sync(p)
				_alt = 1
				for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
					switch _alt {
					case 1:
						p.SetState(256)
						p.GetErrorHandler().Sync(p)
						_la = p.GetTokenStream().LA(1)

						for ok := true; ok; ok = _la == ZitiQlParserWS {
							{
								p.SetState(255)
								p.Match(ZitiQlParserWS)
							}

							p.SetState(258)
							p.GetErrorHandler().Sync(p)
							_la = p.GetTokenStream().LA(1)
						}
						{
							p.SetState(260)
							p.Match(ZitiQlParserAND)
						}
						p.SetState(262)
						p.GetErrorHandler().Sync(p)
						_la = p.GetTokenStream().LA(1)

						for ok := true; ok; ok = _la == ZitiQlParserWS {
							{
								p.SetState(261)
								p.Match(ZitiQlParserWS)
							}

							p.SetState(264)
							p.GetErrorHandler().Sync(p)
							_la = p.GetTokenStream().LA(1)
						}
						{
							p.SetState(266)
							p.expression(0)
						}

					default:
						panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
					}

					p.SetState(269)
					p.GetErrorHandler().Sync(p)
					_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 38, p.GetParserRuleContext())
				}

			case 2:
				localctx = NewOrConjunctionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ZitiQlParserRULE_expression)
				p.SetState(271)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				p.SetState(284)
				p.GetErrorHandler().Sync(p)
				_alt = 1
				for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
					switch _alt {
					case 1:
						p.SetState(273)
						p.GetErrorHandler().Sync(p)
						_la = p.GetTokenStream().LA(1)

						for ok := true; ok; ok = _la == ZitiQlParserWS {
							{
								p.SetState(272)
								p.Match(ZitiQlParserWS)
							}

							p.SetState(275)
							p.GetErrorHandler().Sync(p)
							_la = p.GetTokenStream().LA(1)
						}
						{
							p.SetState(277)
							p.Match(ZitiQlParserOR)
						}
						p.SetState(279)
						p.GetErrorHandler().Sync(p)
						_la = p.GetTokenStream().LA(1)

						for ok := true; ok; ok = _la == ZitiQlParserWS {
							{
								p.SetState(278)
								p.Match(ZitiQlParserWS)
							}

							p.SetState(281)
							p.GetErrorHandler().Sync(p)
							_la = p.GetTokenStream().LA(1)
						}
						{
							p.SetState(283)
							p.expression(0)
						}

					default:
						panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
					}

					p.SetState(286)
					p.GetErrorHandler().Sync(p)
					_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 41, p.GetParserRuleContext())
				}

			}

		}
		p.SetState(292)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 43, p.GetParserRuleContext())
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

func (s *BinaryEqualToNullOpContext) Binary_lhs() IBinary_lhsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBinary_lhsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBinary_lhsContext)
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

func (s *BinaryLessThanNumberOpContext) Binary_lhs() IBinary_lhsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBinary_lhsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBinary_lhsContext)
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

func (s *BinaryGreaterThanDatetimeOpContext) Binary_lhs() IBinary_lhsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBinary_lhsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBinary_lhsContext)
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

func (s *InNumberArrayOpContext) Binary_lhs() IBinary_lhsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBinary_lhsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBinary_lhsContext)
}

func (s *InNumberArrayOpContext) IN() antlr.TerminalNode {
	return s.GetToken(ZitiQlParserIN, 0)
}

func (s *InNumberArrayOpContext) Number_array() INumber_arrayContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumber_arrayContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumber_arrayContext)
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

func (s *InStringArrayOpContext) Binary_lhs() IBinary_lhsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBinary_lhsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBinary_lhsContext)
}

func (s *InStringArrayOpContext) IN() antlr.TerminalNode {
	return s.GetToken(ZitiQlParserIN, 0)
}

func (s *InStringArrayOpContext) String_array() IString_arrayContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IString_arrayContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IString_arrayContext)
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

func (s *BinaryLessThanDatetimeOpContext) Binary_lhs() IBinary_lhsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBinary_lhsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBinary_lhsContext)
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

func (s *BinaryGreaterThanNumberOpContext) Binary_lhs() IBinary_lhsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBinary_lhsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBinary_lhsContext)
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

func (s *InDatetimeArrayOpContext) Binary_lhs() IBinary_lhsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBinary_lhsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBinary_lhsContext)
}

func (s *InDatetimeArrayOpContext) IN() antlr.TerminalNode {
	return s.GetToken(ZitiQlParserIN, 0)
}

func (s *InDatetimeArrayOpContext) Datetime_array() IDatetime_arrayContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDatetime_arrayContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDatetime_arrayContext)
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

func (s *BetweenDateOpContext) Binary_lhs() IBinary_lhsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBinary_lhsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBinary_lhsContext)
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

func (s *BinaryEqualToNumberOpContext) Binary_lhs() IBinary_lhsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBinary_lhsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBinary_lhsContext)
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

func (s *BinaryEqualToBoolOpContext) Binary_lhs() IBinary_lhsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBinary_lhsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBinary_lhsContext)
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

func (s *BinaryEqualToStringOpContext) Binary_lhs() IBinary_lhsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBinary_lhsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBinary_lhsContext)
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

func (s *BetweenNumberOpContext) Binary_lhs() IBinary_lhsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBinary_lhsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBinary_lhsContext)
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

func (s *BinaryContainsOpContext) Binary_lhs() IBinary_lhsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBinary_lhsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBinary_lhsContext)
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

func (s *BinaryEqualToDatetimeOpContext) Binary_lhs() IBinary_lhsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBinary_lhsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBinary_lhsContext)
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

	p.SetState(546)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 78, p.GetParserRuleContext()) {
	case 1:
		localctx = NewInStringArrayOpContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(293)
			p.Binary_lhs()
		}
		p.SetState(295)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == ZitiQlParserWS {
			{
				p.SetState(294)
				p.Match(ZitiQlParserWS)
			}

			p.SetState(297)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(299)
			p.Match(ZitiQlParserIN)
		}
		p.SetState(301)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == ZitiQlParserWS {
			{
				p.SetState(300)
				p.Match(ZitiQlParserWS)
			}

			p.SetState(303)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(305)
			p.String_array()
		}

	case 2:
		localctx = NewInNumberArrayOpContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(307)
			p.Binary_lhs()
		}
		p.SetState(309)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == ZitiQlParserWS {
			{
				p.SetState(308)
				p.Match(ZitiQlParserWS)
			}

			p.SetState(311)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(313)
			p.Match(ZitiQlParserIN)
		}
		p.SetState(315)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == ZitiQlParserWS {
			{
				p.SetState(314)
				p.Match(ZitiQlParserWS)
			}

			p.SetState(317)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(319)
			p.Number_array()
		}

	case 3:
		localctx = NewInDatetimeArrayOpContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(321)
			p.Binary_lhs()
		}
		p.SetState(323)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == ZitiQlParserWS {
			{
				p.SetState(322)
				p.Match(ZitiQlParserWS)
			}

			p.SetState(325)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(327)
			p.Match(ZitiQlParserIN)
		}
		p.SetState(329)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == ZitiQlParserWS {
			{
				p.SetState(328)
				p.Match(ZitiQlParserWS)
			}

			p.SetState(331)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(333)
			p.Datetime_array()
		}

	case 4:
		localctx = NewBetweenNumberOpContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(335)
			p.Binary_lhs()
		}
		p.SetState(337)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == ZitiQlParserWS {
			{
				p.SetState(336)
				p.Match(ZitiQlParserWS)
			}

			p.SetState(339)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(341)
			p.Match(ZitiQlParserBETWEEN)
		}
		p.SetState(343)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == ZitiQlParserWS {
			{
				p.SetState(342)
				p.Match(ZitiQlParserWS)
			}

			p.SetState(345)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(347)
			p.Match(ZitiQlParserNUMBER)
		}
		p.SetState(349)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == ZitiQlParserWS {
			{
				p.SetState(348)
				p.Match(ZitiQlParserWS)
			}

			p.SetState(351)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(353)
			p.Match(ZitiQlParserAND)
		}
		p.SetState(355)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == ZitiQlParserWS {
			{
				p.SetState(354)
				p.Match(ZitiQlParserWS)
			}

			p.SetState(357)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(359)
			p.Match(ZitiQlParserNUMBER)
		}

	case 5:
		localctx = NewBetweenDateOpContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(361)
			p.Binary_lhs()
		}
		p.SetState(363)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == ZitiQlParserWS {
			{
				p.SetState(362)
				p.Match(ZitiQlParserWS)
			}

			p.SetState(365)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(367)
			p.Match(ZitiQlParserBETWEEN)
		}
		p.SetState(369)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == ZitiQlParserWS {
			{
				p.SetState(368)
				p.Match(ZitiQlParserWS)
			}

			p.SetState(371)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(373)
			p.Match(ZitiQlParserDATETIME)
		}
		p.SetState(375)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == ZitiQlParserWS {
			{
				p.SetState(374)
				p.Match(ZitiQlParserWS)
			}

			p.SetState(377)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(379)
			p.Match(ZitiQlParserAND)
		}
		p.SetState(381)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == ZitiQlParserWS {
			{
				p.SetState(380)
				p.Match(ZitiQlParserWS)
			}

			p.SetState(383)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(385)
			p.Match(ZitiQlParserDATETIME)
		}

	case 6:
		localctx = NewBinaryLessThanNumberOpContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(387)
			p.Binary_lhs()
		}
		p.SetState(391)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ZitiQlParserWS {
			{
				p.SetState(388)
				p.Match(ZitiQlParserWS)
			}

			p.SetState(393)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(394)
			p.Match(ZitiQlParserLT)
		}
		p.SetState(398)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ZitiQlParserWS {
			{
				p.SetState(395)
				p.Match(ZitiQlParserWS)
			}

			p.SetState(400)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(401)
			p.Match(ZitiQlParserNUMBER)
		}

	case 7:
		localctx = NewBinaryLessThanDatetimeOpContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(403)
			p.Binary_lhs()
		}
		p.SetState(407)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ZitiQlParserWS {
			{
				p.SetState(404)
				p.Match(ZitiQlParserWS)
			}

			p.SetState(409)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(410)
			p.Match(ZitiQlParserLT)
		}
		p.SetState(414)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ZitiQlParserWS {
			{
				p.SetState(411)
				p.Match(ZitiQlParserWS)
			}

			p.SetState(416)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(417)
			p.Match(ZitiQlParserDATETIME)
		}

	case 8:
		localctx = NewBinaryGreaterThanNumberOpContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(419)
			p.Binary_lhs()
		}
		p.SetState(423)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ZitiQlParserWS {
			{
				p.SetState(420)
				p.Match(ZitiQlParserWS)
			}

			p.SetState(425)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(426)
			p.Match(ZitiQlParserGT)
		}
		p.SetState(430)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ZitiQlParserWS {
			{
				p.SetState(427)
				p.Match(ZitiQlParserWS)
			}

			p.SetState(432)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(433)
			p.Match(ZitiQlParserNUMBER)
		}

	case 9:
		localctx = NewBinaryGreaterThanDatetimeOpContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(435)
			p.Binary_lhs()
		}
		p.SetState(439)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ZitiQlParserWS {
			{
				p.SetState(436)
				p.Match(ZitiQlParserWS)
			}

			p.SetState(441)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(442)
			p.Match(ZitiQlParserGT)
		}
		p.SetState(446)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ZitiQlParserWS {
			{
				p.SetState(443)
				p.Match(ZitiQlParserWS)
			}

			p.SetState(448)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(449)
			p.Match(ZitiQlParserDATETIME)
		}

	case 10:
		localctx = NewBinaryEqualToStringOpContext(p, localctx)
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(451)
			p.Binary_lhs()
		}
		p.SetState(455)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ZitiQlParserWS {
			{
				p.SetState(452)
				p.Match(ZitiQlParserWS)
			}

			p.SetState(457)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(458)
			p.Match(ZitiQlParserEQ)
		}
		p.SetState(462)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ZitiQlParserWS {
			{
				p.SetState(459)
				p.Match(ZitiQlParserWS)
			}

			p.SetState(464)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(465)
			p.Match(ZitiQlParserSTRING)
		}

	case 11:
		localctx = NewBinaryEqualToNumberOpContext(p, localctx)
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(467)
			p.Binary_lhs()
		}
		p.SetState(471)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ZitiQlParserWS {
			{
				p.SetState(468)
				p.Match(ZitiQlParserWS)
			}

			p.SetState(473)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(474)
			p.Match(ZitiQlParserEQ)
		}
		p.SetState(478)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ZitiQlParserWS {
			{
				p.SetState(475)
				p.Match(ZitiQlParserWS)
			}

			p.SetState(480)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(481)
			p.Match(ZitiQlParserNUMBER)
		}

	case 12:
		localctx = NewBinaryEqualToDatetimeOpContext(p, localctx)
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(483)
			p.Binary_lhs()
		}
		p.SetState(487)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ZitiQlParserWS {
			{
				p.SetState(484)
				p.Match(ZitiQlParserWS)
			}

			p.SetState(489)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(490)
			p.Match(ZitiQlParserEQ)
		}
		p.SetState(494)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ZitiQlParserWS {
			{
				p.SetState(491)
				p.Match(ZitiQlParserWS)
			}

			p.SetState(496)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(497)
			p.Match(ZitiQlParserDATETIME)
		}

	case 13:
		localctx = NewBinaryEqualToBoolOpContext(p, localctx)
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(499)
			p.Binary_lhs()
		}
		p.SetState(503)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ZitiQlParserWS {
			{
				p.SetState(500)
				p.Match(ZitiQlParserWS)
			}

			p.SetState(505)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(506)
			p.Match(ZitiQlParserEQ)
		}
		p.SetState(510)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ZitiQlParserWS {
			{
				p.SetState(507)
				p.Match(ZitiQlParserWS)
			}

			p.SetState(512)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(513)
			p.Match(ZitiQlParserBOOL)
		}

	case 14:
		localctx = NewBinaryEqualToNullOpContext(p, localctx)
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(515)
			p.Binary_lhs()
		}
		p.SetState(519)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ZitiQlParserWS {
			{
				p.SetState(516)
				p.Match(ZitiQlParserWS)
			}

			p.SetState(521)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(522)
			p.Match(ZitiQlParserEQ)
		}
		p.SetState(526)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ZitiQlParserWS {
			{
				p.SetState(523)
				p.Match(ZitiQlParserWS)
			}

			p.SetState(528)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(529)
			p.Match(ZitiQlParserNULL)
		}

	case 15:
		localctx = NewBinaryContainsOpContext(p, localctx)
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(531)
			p.Binary_lhs()
		}
		p.SetState(535)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ZitiQlParserWS {
			{
				p.SetState(532)
				p.Match(ZitiQlParserWS)
			}

			p.SetState(537)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(538)
			p.Match(ZitiQlParserCONTAINS)
		}
		p.SetState(540)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == ZitiQlParserWS {
			{
				p.SetState(539)
				p.Match(ZitiQlParserWS)
			}

			p.SetState(542)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(544)
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

// IBinary_lhsContext is an interface to support dynamic dispatch.
type IBinary_lhsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBinary_lhsContext differentiates from other interfaces.
	IsBinary_lhsContext()
}

type Binary_lhsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBinary_lhsContext() *Binary_lhsContext {
	var p = new(Binary_lhsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZitiQlParserRULE_binary_lhs
	return p
}

func (*Binary_lhsContext) IsBinary_lhsContext() {}

func NewBinary_lhsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Binary_lhsContext {
	var p = new(Binary_lhsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZitiQlParserRULE_binary_lhs

	return p
}

func (s *Binary_lhsContext) GetParser() antlr.Parser { return s.parser }

func (s *Binary_lhsContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ZitiQlParserIDENTIFIER, 0)
}

func (s *Binary_lhsContext) Set_function() ISet_functionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISet_functionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISet_functionContext)
}

func (s *Binary_lhsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Binary_lhsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Binary_lhsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZitiQlListener); ok {
		listenerT.EnterBinary_lhs(s)
	}
}

func (s *Binary_lhsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZitiQlListener); ok {
		listenerT.ExitBinary_lhs(s)
	}
}

func (p *ZitiQlParser) Binary_lhs() (localctx IBinary_lhsContext) {
	localctx = NewBinary_lhsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, ZitiQlParserRULE_binary_lhs)

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

	p.SetState(550)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ZitiQlParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(548)
			p.Match(ZitiQlParserIDENTIFIER)
		}

	case ZitiQlParserALL_OF, ZitiQlParserNONE_OF, ZitiQlParserANY_OF:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(549)
			p.Set_function()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ISet_functionContext is an interface to support dynamic dispatch.
type ISet_functionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSet_functionContext differentiates from other interfaces.
	IsSet_functionContext()
}

type Set_functionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySet_functionContext() *Set_functionContext {
	var p = new(Set_functionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZitiQlParserRULE_set_function
	return p
}

func (*Set_functionContext) IsSet_functionContext() {}

func NewSet_functionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Set_functionContext {
	var p = new(Set_functionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZitiQlParserRULE_set_function

	return p
}

func (s *Set_functionContext) GetParser() antlr.Parser { return s.parser }

func (s *Set_functionContext) CopyFrom(ctx *Set_functionContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *Set_functionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Set_functionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SetFunctionContext struct {
	*Set_functionContext
}

func NewSetFunctionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SetFunctionContext {
	var p = new(SetFunctionContext)

	p.Set_functionContext = NewEmptySet_functionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Set_functionContext))

	return p
}

func (s *SetFunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SetFunctionContext) ALL_OF() antlr.TerminalNode {
	return s.GetToken(ZitiQlParserALL_OF, 0)
}

func (s *SetFunctionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ZitiQlParserLPAREN, 0)
}

func (s *SetFunctionContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ZitiQlParserIDENTIFIER, 0)
}

func (s *SetFunctionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ZitiQlParserRPAREN, 0)
}

func (s *SetFunctionContext) NONE_OF() antlr.TerminalNode {
	return s.GetToken(ZitiQlParserNONE_OF, 0)
}

func (s *SetFunctionContext) ANY_OF() antlr.TerminalNode {
	return s.GetToken(ZitiQlParserANY_OF, 0)
}

func (s *SetFunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZitiQlListener); ok {
		listenerT.EnterSetFunction(s)
	}
}

func (s *SetFunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZitiQlListener); ok {
		listenerT.ExitSetFunction(s)
	}
}

func (p *ZitiQlParser) Set_function() (localctx ISet_functionContext) {
	localctx = NewSet_functionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, ZitiQlParserRULE_set_function)

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

	p.SetState(564)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ZitiQlParserALL_OF:
		localctx = NewSetFunctionContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(552)
			p.Match(ZitiQlParserALL_OF)
		}
		{
			p.SetState(553)
			p.Match(ZitiQlParserLPAREN)
		}
		{
			p.SetState(554)
			p.Match(ZitiQlParserIDENTIFIER)
		}
		{
			p.SetState(555)
			p.Match(ZitiQlParserRPAREN)
		}

	case ZitiQlParserNONE_OF:
		localctx = NewSetFunctionContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(556)
			p.Match(ZitiQlParserNONE_OF)
		}
		{
			p.SetState(557)
			p.Match(ZitiQlParserLPAREN)
		}
		{
			p.SetState(558)
			p.Match(ZitiQlParserIDENTIFIER)
		}
		{
			p.SetState(559)
			p.Match(ZitiQlParserRPAREN)
		}

	case ZitiQlParserANY_OF:
		localctx = NewSetFunctionContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(560)
			p.Match(ZitiQlParserANY_OF)
		}
		{
			p.SetState(561)
			p.Match(ZitiQlParserLPAREN)
		}
		{
			p.SetState(562)
			p.Match(ZitiQlParserIDENTIFIER)
		}
		{
			p.SetState(563)
			p.Match(ZitiQlParserRPAREN)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

func (p *ZitiQlParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 9:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *ZitiQlParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
