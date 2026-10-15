from user import User
from auction import Auction
from typing import Dict, List
from decimal import Decimal
from datetime import datetime
import threading

class AuctionService:
    _instance = None
    _lock = threading.Lock()

    def __init__(self):
        if AuctionService._instance is not None:
            raise Exception("This class is a singleton!")
        self.users: Dict[str, User] = {}
        self.auctions: Dict[str, Auction] = {}
        self.scheduled_ends = []

    @staticmethod
    def get_instance():
        if AuctionService._instance is None:
            with AuctionService._lock:
                if AuctionService._instance is None:
                    AuctionService._instance = AuctionService()
        return AuctionService._instance

    def create_user(self, name: str) -> User:
        user = User(name)
        self.users[user.get_id()] = user
        return user

    def get_user(self, user_id: str) -> User:
        return self.users[user_id]

    def create_auction(self, item_name: str, description: str, starting_price: Decimal, end_time: datetime) -> Auction:
        auction = Auction(item_name, description, starting_price, end_time)
        self.auctions[auction.get_id()] = auction

        delay = (end_time - datetime.now()).total_seconds()
        if delay > 0:
            scheduled_end = threading.Timer(delay, self.end_auction, args=(auction.get_id(),))
            self.scheduled_ends.append(scheduled_end)
            scheduled_end.start()

        print(f"New auction created for '{item_name}' (ID: {auction.get_id()}), ending at {end_time}.")
        return auction

    def view_active_auctions(self) -> List[Auction]:
        return [auction for auction in self.auctions.values() if auction.is_active()]

    def place_bid(self, auction_id: str, bidder_id: str, amount: Decimal):
        auction = self.get_auction(auction_id)
        auction.place_bid(self.users[bidder_id], amount)

    def end_auction(self, auction_id: str):
        auction = self.get_auction(auction_id)
        auction.end_auction()

    def get_auction(self, auction_id: str) -> Auction:
        auction = self.auctions.get(auction_id)
        if auction is None:
            raise KeyError(f"Auction with ID {auction_id} not found.")
        return auction

    def shutdown(self):
        for scheduled_end in self.scheduled_ends:
            scheduled_end.join()