#include "Order.hpp"
#include "User.hpp"
#include <iostream>
#include <iomanip>

Order::Order(std::string orderId, User* user, const std::vector<CartItem*>& cartItems)
    : orderId(orderId), user(user), status(OrderStatus::PENDING), totalAmount(0.0) {
    
    // Deep copy cart items
    for (const auto& cartItem : cartItems) {
        items.push_back(new CartItem(cartItem->getProduct(), cartItem->getQuantity()));
        totalAmount += cartItem->getSubtotal();
    }
    
    // Set order date to current time
    orderDate = std::time(nullptr);
}

Order::~Order() {
    for (auto item : items) {
        delete item;
    }
}

std::string Order::getOrderId() const { return orderId; }
User* Order::getUser() const { return user; }
const std::vector<CartItem*>& Order::getItems() const { return items; }
double Order::getTotalAmount() const { return totalAmount; }
std::time_t Order::getOrderDate() const { return orderDate; }
OrderStatus Order::getStatus() const { std::lock_guard<std::mutex> guard(mutex); return status; }

void Order::setStatus(OrderStatus status) {
    std::lock_guard<std::mutex> guard(mutex);
    if (status == OrderStatus::CANCELLED && !stockRestored) {
        for (const auto* item : items) {
            item->getProduct()->updateStock(item->getQuantity());
        }
        stockRestored = true;
    }
    this->status = status;
}

void Order::displayInfo() const {
    std::lock_guard<std::mutex> guard(mutex);
    std::cout << "\nOrder Details:" << std::endl;
    std::cout << "Order ID: " << orderId << std::endl;
    std::cout << "Customer: " << user->getUsername() << std::endl;
    std::cout << "Date: " << std::ctime(&orderDate);
    std::cout << "Status: ";
    switch (status) {
        case OrderStatus::PENDING: std::cout << "Pending"; break;
        case OrderStatus::CONFIRMED: std::cout << "Confirmed"; break;
        case OrderStatus::SHIPPED: std::cout << "Shipped"; break;
        case OrderStatus::DELIVERED: std::cout << "Delivered"; break;
        case OrderStatus::CANCELLED: std::cout << "Cancelled"; break;
    }
    std::cout << std::endl;
    
    std::cout << "\nItems:" << std::endl;
    for (const auto& item : items) {
        item->displayInfo();
    }
    std::cout << "------------------------" << std::endl;
    std::cout << "Total Amount: $" << std::fixed << std::setprecision(2) << totalAmount << std::endl;
} 