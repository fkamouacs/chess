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


    put_piece(board, 1, num_white_pawns)
    put_piece(board, -1, num_black_pawns)

    put_piece(board, 2, num_white_knights)
    put_piece(board, -2, num_black_knights)

    put_piece(board, 3, num_white_bishops)
    put_piece(board, -3, num_black_bishops)

    put_piece(board, 4, num_white_rooks)
    put_piece(board, -4, num_black_rooks)

    put_piece(board, 5, num_white_queens)
    put_piece(board, -5, num_black_queens)

    put_piece(board, 6, num_white_king)
    put_piece(board, -6, num_black_king)

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


def get_random_board_index(indexes):
    return random.choice(indexes)

def put_piece(board,  piece, amount):
    empty_indexes = get_empty_indexes(board)
    
    if (piece == 1 or piece == -1):
        pawn_empty_indexes = empty_indexes[empty_indexes > 28]
        pawn_empty_indexes = pawn_empty_indexes[pawn_empty_indexes < 91]

        for p in range(0, amount, 1):
            board_index = get_random_board_index(pawn_empty_indexes)
        
            board[board_index] = piece

            delete_index = np.argwhere(empty_indexes == board_index)
            empty_indexes = np.delete(empty_indexes, delete_index)

            delete_index = np.argwhere(pawn_empty_indexes == board_index)
            pawn_empty_indexes = np.delete(pawn_empty_indexes, delete_index)

    elif piece == 2 or piece == -2:
        for k in range (0, amount, 1):
            board_index = get_random_board_index(empty_indexes)
            
            board[board_index] = piece

            delete_index = np.argwhere(empty_indexes == board_index)
            empty_indexes = np.delete(empty_indexes, delete_index)

    elif piece == 3 or piece == -3:
        for k in range (0, amount, 1):
            board_index = get_random_board_index(empty_indexes)
            
            board[board_index] = piece

            delete_index = np.argwhere(empty_indexes == board_index)
            empty_indexes = np.delete(empty_indexes, delete_index)

    elif piece == 4 or piece == -4:
        for k in range (0, amount, 1):
            board_index = get_random_board_index(empty_indexes)
            
            board[board_index] = piece

            delete_index = np.argwhere(empty_indexes == board_index)
            empty_indexes = np.delete(empty_indexes, delete_index)


    elif piece == 5 or piece == -5:
        for k in range (0, amount, 1):
            board_index = get_random_board_index(empty_indexes)
            
            board[board_index] = piece

            delete_index = np.argwhere(empty_indexes == board_index)
            empty_indexes = np.delete(empty_indexes, delete_index)

    elif piece == 6 or piece == -6:
        for k in range (0, amount, 1):
            board_index = get_random_board_index(empty_indexes)
            
            board[board_index] = piece

            delete_index = np.argwhere(empty_indexes == board_index)
            empty_indexes = np.delete(empty_indexes, delete_index)



    return board

def print_board(board):
    board_string = ""
    counter = 0
    print("")
    temp_board = [s for s in board if s != 99]
    print ("   a  b  c  d  e  f  g  h", end="")
    board_string += "   a  b  c  d  e  f  g  h"

    for e in temp_board:
     
        if counter % 8 == 0:
            print("\n")  
            board_string += "\n"
            print(int(counter / 8) + 1, " ", end="")
            board_string += str(int(counter /8) + 1)
            board_string += " "

        print(e, end="")
        print(" ", end="")
        board_string += str(e) 
        board_string += " "
        if (e >= 0):
            print(" ", end="")
            board_string += " "
        counter = counter + 1

    print("\n")  
    board_string += "\n"
    return board_string
      


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
    f = open("../data/random_board_state.txt", "w")
    random_board = generate_board_state()

    board_string = print_board(random_board)

    f.write(board_string)
    f.close()
