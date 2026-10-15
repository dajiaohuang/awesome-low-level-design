package chessgame.pieces;

import chessgame.Board;
import chessgame.Cell;
import chessgame.Color;

public class Rook extends Piece {
    public Rook(Color color) {
        super(color);
    }

    @Override
    public boolean canMove(Board board, Cell from, Cell to) {
        int dr = Integer.compare(to.getRow(), from.getRow());
        int dc = Integer.compare(to.getCol(), from.getCol());
        if ((dr != 0 && dc != 0) || (dr == 0 && dc == 0)) return false;
        for (int row = from.getRow() + dr, col = from.getCol() + dc;
             row != to.getRow() || col != to.getCol(); row += dr, col += dc) {
            if (board.getPiece(row, col) != null) return false;
        }
        return true;
    }
}