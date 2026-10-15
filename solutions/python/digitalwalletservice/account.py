from decimal import Decimal
import threading
from exception import InsufficientFundsException

class Account:
    def __init__(self, id, user, account_number, currency):
        self.id = id
        self.user = user
        self.account_number = account_number
        self.currency = currency
        self.balance = Decimal('0.00')
        self.transactions = []
        self._lock = threading.Lock()

    def deposit(self, amount):
        with self._lock:
            self.balance += amount

    def withdraw(self, amount):
        with self._lock:
            if self.balance >= amount:
                self.balance -= amount
            else:
                raise InsufficientFundsException("Insufficient funds in the account.")

    def add_transaction(self, transaction):
        with self._lock:
            self.transactions.append(transaction)

    def get_transactions(self):
        with self._lock:
            return list(self.transactions)
