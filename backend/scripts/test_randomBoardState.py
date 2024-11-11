import unittest
import random

from randomBoardState import get_empty_indexes
from randomBoardState import put_piece
from randomBoardState import board
from randomBoardState import get_number_of_piece


class TestPutPawnsFunction(unittest.TestCase):

    def test_correct_amount_of_pawns(self):
        num_white_pawns = random.randint(0, 8)
        white_piece = 1
        num_black_pawns = random.randint(0,8)
        black_piece = -1

        
        put_piece(board, white_piece, num_white_pawns)
        put_piece(board,black_piece, num_black_pawns)
        
        self.assertEqual(get_number_of_piece(board, white_piece) + get_number_of_piece(board, black_piece), num_white_pawns + num_black_pawns)



if __name__ == '__main__':
    unittest.main()