from piece import Piece
from color import Color

class Pawn(Piece):
    def can_move(self, board, dest_row, dest_col):
        row_diff = dest_row - self.row
        col_diff = abs(dest_col - self.col)

        direction = 1 if self.color == Color.WHITE else -1
        initial_row = 1 if self.color == Color.WHITE else 6
        destination = board.get_piece(dest_row, dest_col)
        if col_diff == 0:
            if destination is not None:
                return False
            return row_diff == direction or (
                self.row == initial_row and row_diff == 2 * direction and
                board.get_piece(self.row + direction, self.col) is None)
        return row_diff == direction and col_diff == 1 and destination is not None