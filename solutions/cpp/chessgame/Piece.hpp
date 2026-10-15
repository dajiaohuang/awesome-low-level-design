#ifndef PIECE_HPP
#define PIECE_HPP

#include <string>
#include "Position.hpp"

enum class PieceType {
    KING, QUEEN, BISHOP, KNIGHT, ROOK, PAWN
};

enum class Color {
    WHITE, BLACK
};

class Piece {
private:
    PieceType type;
    Color color;
    Position position;
    bool captured;

public:
    Piece(PieceType type, Color color, Position position);
    
    virtual ~Piece() = default;
    
    PieceType getType() const;
    Color getColor() const;
    Position getPosition() const;
    bool isCaptured() const;
    
    void setPosition(Position newPosition);
    void setCaptured(bool status);
    
    virtual bool isValidMove(Position newPosition, Piece* board[8][8]) const = 0;
    
    std::string getSymbol() const;
};

#endif 