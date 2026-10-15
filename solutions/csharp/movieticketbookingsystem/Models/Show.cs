class Show
{
    private readonly string id;
    private readonly Movie movie;
    private readonly Screen screen;
    private readonly Screen physicalScreen;
    private readonly DateTime startTime;
    private readonly IPricingStrategy pricingStrategy;

    public Show(string id, Movie movie, Screen screen, DateTime startTime, IPricingStrategy pricingStrategy)
    {
        this.id = id;
        this.movie = movie;
        this.physicalScreen = screen;
        this.screen = new Screen(screen.GetId());
        foreach (var seat in screen.GetSeats())
            this.screen.AddSeat(new Seat(seat.GetId(), seat.GetRow(), seat.GetCol(), seat.GetSeatType()));
        this.startTime = startTime;
        this.pricingStrategy = pricingStrategy;
    }

    public string GetId() { return id; }
    public Movie GetMovie() { return movie; }
    public Screen GetScreen() { return screen; }
    internal Screen GetPhysicalScreen() { return physicalScreen; }
    public DateTime GetStartTime() { return startTime; }
    public IPricingStrategy GetPricingStrategy() { return pricingStrategy; }
}