from typing import List, Optional
from user import User
from show import Show
from seat import Seat
from payment_strategy import PaymentStrategy
from seat_lock_manager import SeatLockManager
from enums import PaymentStatus
from booking import Booking

class BookingManager:
    def __init__(self, seat_lock_manager: SeatLockManager):
        self.seat_lock_manager = seat_lock_manager

    def create_booking(self, user: User, show: Show, seats: List[Seat], payment_strategy: PaymentStrategy) -> Optional[Booking]:
        with show._lock:
            return self._create_booking_locked(user, show, seats, payment_strategy)

    def _create_booking_locked(self, user: User, show: Show, seats: List[Seat], payment_strategy: PaymentStrategy) -> Optional[Booking]:
        # 1. Acquire the complete seat batch before any payment.
        if not self.seat_lock_manager.lock_seats(show, seats, user.get_id()):
            return None

        # 2. Calculate the total price
        total_amount = show.get_pricing_strategy().calculate_price(seats)

        # 3. Process Payment
        payment = payment_strategy.pay(total_amount)

        # 4. If payment is successful, create the booking
        if payment.get_status() == PaymentStatus.SUCCESS:
            booking = Booking.BookingBuilder() \
                .set_user(user) \
                .set_show(show) \
                .set_seats(seats) \
                .set_total_amount(total_amount) \
                .set_payment(payment) \
                .build()

            # 5. Confirm the booking (mark seats as BOOKED)
            booking.confirm_booking()

            # Clean up the lock map
            self.seat_lock_manager.unlock_seats(show, seats, user.get_id())

            return booking
        else:
            self.seat_lock_manager.unlock_seats(show, seats, user.get_id())
            print("Payment failed. Please try again.")
            return None