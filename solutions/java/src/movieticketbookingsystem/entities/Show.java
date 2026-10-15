package movieticketbookingsystem.entities;

import movieticketbookingsystem.strategy.pricing.PricingStrategy;

import java.time.LocalDateTime;

public class Show {
    private final String id;
    private final Movie movie;
    private final Screen screen;
    private final Screen physicalScreen;
    private final LocalDateTime startTime;
    private final PricingStrategy pricingStrategy;

    public Show(String id, Movie movie, Screen screen, LocalDateTime startTime, PricingStrategy pricingStrategy) {
        this.id = id;
        this.movie = movie;
        this.physicalScreen = screen;
        this.screen = new Screen(screen.getId());
        for (Seat seat : screen.getSeats()) {
            this.screen.addSeat(new Seat(seat.getId(), seat.getRow(), seat.getCol(), seat.getType()));
        }
        this.startTime = startTime;
        this.pricingStrategy = pricingStrategy;
    }

    public String getId() { return id; }
    public Movie getMovie() { return movie; }
    public Screen getScreen() { return screen; }
    public Screen getPhysicalScreen() { return physicalScreen; }
    public LocalDateTime getStartTime() { return startTime; }
    public PricingStrategy getPricingStrategy() { return pricingStrategy; }
}
