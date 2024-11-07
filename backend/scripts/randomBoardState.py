import random
import numpy as np


MAX_PAWNS = 8
MAX_KNIGHTS = 10
MAX_BISHOPS = 10
MAX_ROOKS = 10
MAX_QUEENS = 9
MAX_KING = 1

# TODO
# CANT HAVE BOTH KINGS IN CHECK

def generate_board_state():
    # randomly generate total number of each piece 
    num_white_pawns = random.randint(0, MAX_PAWNS)
    num_white_knights = random.randint(0, 2)
    num_white_bishops = random.randint(0,2)
    num_white_rooks = random.randint(0,2)
    num_white_queens = random.randint(0,1)
    num_white_king = 1

    num_black_pawns = random.randint(0, MAX_PAWNS)
    num_black_knights = random.randint(0,2)
    num_black_bishops = random.randint(0, 2)
    num_black_rooks = random.randint(0,2)
    num_black_queens = random.randint(0,1)
    num_black_king = 1

    put_pawns(board, 1, num_white_pawns)
    put_pawns(board,-1, num_black_pawns)


    return board

def get_empty_indexes(board):
    empty_array = np.empty(0, dtype=int)
    empty_indexes = np.array(empty_array)
    
    start = 21
    end = 98
    step = 1

    for i in range(start, end, step):
        if board[i] == 0:
            empty_indexes = np.append(empty_indexes, i)

    return empty_indexes

def get_number_of_piece(board, piece):
    return board.count(piece)



def put_pawns(board, type, numberOfPawns):
    empty_indexes = get_empty_indexes(board)

    pawn_empty_indexes = empty_indexes[empty_indexes > 28]
    pawn_empty_indexes = pawn_empty_indexes[pawn_empty_indexes < 91]

    for p in range(0, numberOfPawns, 1):
        board_index = random.choice(pawn_empty_indexes)
        
        board[board_index] = type

        delete_index = np.argwhere(empty_indexes == board_index)
        empty_indexes = np.delete(empty_indexes, delete_index)

        delete_index = np.argwhere(pawn_empty_indexes == board_index)
        pawn_empty_indexes = np.delete(pawn_empty_indexes, delete_index)
    return board

def print_board(board):
    counter = 0
    for e in board:
        if counter % 10 == 0:
            print("\n")
        print(e, end="")
        print(" ", end="")
        if (e >= 0):
            print(" ", end="")
        counter = counter + 1



board = [99, 99, 99, 99, 99, 99, 99, 99, 99, 99,
	
	99, 99, 99, 99, 99, 99, 99, 99, 99, 99,
	
	99, 0, 0, 0, 0, 0, 0, 0, 0, 99,
	
	99, 0, 0, 0, 0, 0, 0, 0, 0, 99,
	
	99, 0, 0, 0, 0, 0, 0, 0, 0, 99,
	
	99, 0, 0, 0, 0, 0, 0, 0, 0, 99,
	
	99, 0, 0, 0, 0, 0, 0, 0, 0, 99,
	
	99, 0, 0, 0, 0, 0, 0, 0, 0, 99,
	
	99, 0, 0, 0, 0, 0, 0, 0, 0, 99,
	
	99, 0, 0, 0, 0, 0, 0, 0, 0, 99,
	
	99, 99, 99, 99, 99, 99, 99, 99, 99, 99,
	
	99, 99, 99, 99, 99, 99, 99, 99, 99, 99,]


if __name__ == "__main__":
    print_board(generate_board_state())