import threading
import uuid
from credit_card_payment_processor import CreditCardPaymentProcessor
from reservation import Reservation

class RentalSystem:
    _instance = None
    _instance_lock = threading.Lock()

    def __init__(self):
        if RentalSystem._instance is not None:
            raise Exception("This class is a singleton!")
        else:
            RentalSystem._instance = self
            self.cars = {}
            self.reservations = {}
            self.payment_processor = CreditCardPaymentProcessor()
            self._lock = threading.RLock()

    @staticmethod
    def get_instance():
        with RentalSystem._instance_lock:
            if RentalSystem._instance is None:
                RentalSystem()
            return RentalSystem._instance

    def add_car(self, car):
        with self._lock:
            self.cars[car.license_plate] = car

    def remove_car(self, license_plate):
        with self._lock:
            self.cars.pop(license_plate, None)

    def search_cars(self, make, model, start_date, end_date):
        with self._lock:
            available_cars = []
            for car in self.cars.values():
                if car.make.lower() == make.lower() and car.model.lower() == model.lower():
                    if self.is_car_available(car, start_date, end_date):
                        available_cars.append(car)
            return available_cars

    def is_car_available(self, car, start_date, end_date):
        with self._lock:
            for reservation in self.reservations.values():
                if reservation.get_car() == car:
                    if start_date < reservation.get_end_date() and end_date > reservation.get_start_date():
                        return False
            return True

    def make_reservation(self, customer, car, start_date, end_date):
        with self._lock:
            if self.is_car_available(car, start_date, end_date):
                reservation_id = self.generate_reservation_id()
                reservation = Reservation(reservation_id, customer, car, start_date, end_date)
                self.reservations[reservation_id] = reservation
                car.set_available(False)
                return reservation
            return None

    def cancel_reservation(self, reservation_id):
        with self._lock:
            reservation = self.reservations.pop(reservation_id, None)
            if reservation is not None:
                reservation.get_car().set_available(True)

    def process_payment(self, reservation):
        return self.payment_processor.process_payment(reservation.get_total_price())

    def generate_reservation_id(self):
        return "RES" + str(uuid.uuid4())[:8].upper()