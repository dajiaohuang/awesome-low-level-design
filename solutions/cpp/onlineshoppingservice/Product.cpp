#include "Product.hpp"
#include <iostream>
#include <iomanip>

Product::Product(std::string productId, std::string name, std::string description, 
                double price, int stockQuantity)
    : productId(productId), name(name), description(description), 
      price(price), stockQuantity(stockQuantity), available(true) {}

std::string Product::getProductId() const { return productId; }
std::string Product::getName() const { return name; }
std::string Product::getDescription() const { return description; }
double Product::getPrice() const { std::lock_guard<std::mutex> guard(mutex); return price; }
int Product::getStockQuantity() const { std::lock_guard<std::mutex> guard(mutex); return stockQuantity; }
bool Product::isAvailable() const { std::lock_guard<std::mutex> guard(mutex); return available; }

void Product::setPrice(double newPrice) {
    std::lock_guard<std::mutex> guard(mutex);
    if (newPrice >= 0) {
        price = newPrice;
    }
}

void Product::setStockQuantity(int quantity) {
    std::lock_guard<std::mutex> guard(mutex);
    if (quantity >= 0) {
        stockQuantity = quantity;
        available = (quantity > 0);
    }
}

void Product::setAvailable(bool status) {
    std::lock_guard<std::mutex> guard(mutex);
    available = status;
}

bool Product::updateStock(int quantity) {
    std::lock_guard<std::mutex> guard(mutex);
    if (stockQuantity + quantity >= 0) {
        stockQuantity += quantity;
        available = (stockQuantity > 0);
        return true;
    }
    return false;
}

void Product::displayInfo() const {
    std::lock_guard<std::mutex> guard(mutex);
    std::cout << "Product: " << name << " (ID: " << productId << ")" << std::endl;
    std::cout << "Description: " << description << std::endl;
    std::cout << "Price: $" << std::fixed << std::setprecision(2) << price << std::endl;
    std::cout << "Stock: " << stockQuantity << std::endl;
    std::cout << "Status: " << (available ? "Available" : "Out of Stock") << std::endl;
} 