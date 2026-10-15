import threading
import unittest
from concurrent.futures import ThreadPoolExecutor

from item_type import ItemType
from library_management_system import LibraryManagementSystem
from transaction_service import TransactionService


class ConcurrentCheckoutTest(unittest.TestCase):
    def test_one_copy_has_one_loan_after_competing_checkouts(self):
        system = LibraryManagementSystem.get_instance()
        suffix = str(id(self))
        members = [system.add_member(f"member-{suffix}-{i}", f"Member {i}") for i in range(24)]
        copy = system.add_item(ItemType.BOOK, f"book-{suffix}", "Concurrency", "Author", 1)[0]
        start = threading.Barrier(len(members))

        def checkout(member):
            start.wait()
            system.checkout(member.get_id(), copy.get_id())

        with ThreadPoolExecutor(max_workers=len(members)) as pool:
            list(pool.map(checkout, members))

        borrowers = [member for member in members if member.get_loans()]
        self.assertEqual(len(borrowers), 1)
        self.assertFalse(copy.is_available())
        self.assertIn(copy.get_id(), TransactionService.get_instance().active_loans)


if __name__ == "__main__":
    unittest.main()
