namespace AirlineManagementSystem
{
    public class Seat
    {
        private readonly object seatLock = new object();
        public string SeatNumber { get; }
        public SeatType Type { get; }
        public SeatStatus Status { get; private set; }

        public Seat(string seatNumber, SeatType type)
        {
            SeatNumber = seatNumber;
            Type = type;
            Status = SeatStatus.AVAILABLE;
        }

        public bool Reserve()
        {
            lock (seatLock)
            {
                if (Status != SeatStatus.AVAILABLE) return false;
                Status = SeatStatus.RESERVED;
                return true;
            }
        }

        public void Release()
        {
            lock (seatLock)
            {
                Status = SeatStatus.AVAILABLE;
            }
        }
    }
}