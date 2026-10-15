package airlinemanagementsystem.seat;

public class Seat {
    private final String seatNumber;
    private final SeatType type;
    private SeatStatus status;

    public Seat(String seatNumber, SeatType type) {
        this.seatNumber = seatNumber;
        this.type = type;
        this.status = SeatStatus.AVAILABLE;
    }

    public String getSeatNumber() {
        return seatNumber;
    }

    public synchronized boolean reserve() {
        if (status != SeatStatus.AVAILABLE) return false;
        status = SeatStatus.RESERVED;
        return true;
    }

    public synchronized void release() {
        status = SeatStatus.AVAILABLE;
    }

    public synchronized boolean isBooked() {
        return status == SeatStatus.OCCUPIED;
    }
}
