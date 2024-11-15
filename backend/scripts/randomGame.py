import random
import chess.pgn

def pick_game(pgn, index):
    offset = 0
    for _ in range(index):
        offset = pgn.tell()
        chess.pgn.read_headers(pgn)
    pgn.seek(offset)
    return chess.pgn.read_game(pgn)


def num_games(pgn):
    num = 0
    while chess.pgn.read_headers(pgn):
        num += 1
    return num

if __name__ == "__main__":
    filepath_to_pgn = "../data/lichess_db_standard_rated_2013-01.pgn"

    total_num_games = num_games(open(filepath_to_pgn, "r"))
    random_game_index = random.randint(1, total_num_games)
    game = pick_game(open(filepath_to_pgn, "r"), random_game_index)

    
    with open('../data/game.txt', 'w') as f:
        f.write(str(game))
        f.close()  

   
