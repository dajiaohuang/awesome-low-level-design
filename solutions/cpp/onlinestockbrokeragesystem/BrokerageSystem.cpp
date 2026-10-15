#include "BrokerageSystem.hpp"
#include <iostream>
#include <algorithm>

BrokerageSystem::BrokerageSystem() : transactionIdCounter(1), userIdCounter(1) {}

BrokerageSystem::~BrokerageSystem() {
    for (auto user : users) delete user;
    for (auto user : retiredUsers) delete user;
    for (auto& pair : stocks) delete pair.second;
    for (auto transaction : transactions) delete transaction;
}

User* BrokerageSystem::registerUser(const std::string& username, const std::string& email) {
    std::lock_guard<std::mutex> guard(mutex);
    std::string userId = "U" + std::to_string(userIdCounter++);
    User* user = new User(userId, username, email);
    users.push_back(user);
    return user;
}

void BrokerageSystem::removeUser(const std::string& userId) {
    std::lock_guard<std::mutex> guard(mutex);
    auto it = std::find_if(users.begin(), users.end(),
        [userId](User* user) { return user->getUserId() == userId; });
    
    if (it != users.end()) {
        retiredUsers.push_back(*it);
        users.erase(it);
    }
}

Stock* BrokerageSystem::addStock(const std::string& symbol, const std::string& companyName,
                                double price, int shares) {
    std::lock_guard<std::mutex> guard(mutex);
    Stock* stock = new Stock(symbol, companyName, price, shares);
    stocks[symbol] = stock;
    return stock;
}

void BrokerageSystem::updateStockPrice(const std::string& symbol, double newPrice) {
    std::lock_guard<std::mutex> guard(mutex);
    if (Stock* stock = findStock(symbol)) {
        stock->setCurrentPrice(newPrice);
        updatePortfolioValues();
    }
}

bool BrokerageSystem::deposit(const std::string& userId, double amount) {
    std::lock_guard<std::mutex> guard(mutex);
    if (User* user = findUser(userId)) {
        return user->deposit(amount);
    }
    return false;
}

bool BrokerageSystem::withdraw(const std::string& userId, double amount) {
    std::lock_guard<std::mutex> guard(mutex);
    if (User* user = findUser(userId)) {
        return user->withdraw(amount);
    }
    return false;
}

Transaction* BrokerageSystem::buyStock(const std::string& userId, const std::string& symbol, int quantity) {
    std::lock_guard<std::mutex> guard(mutex);
    User* user = findUser(userId);
    Stock* stock = findStock(symbol);
    
    if (!user || !stock || !stock->isActive()) return nullptr;
    
    double totalCost = stock->getCurrentPrice() * quantity;
    if (user->getBalance() < totalCost || stock->getAvailableShares() < quantity) {
        return nullptr;
    }
    
    // Process transaction
    user->withdraw(totalCost);
    stock->updateShares(-quantity);
    user->getPortfolio()->addShares(symbol, quantity);
    
    // Create transaction record
    Transaction* transaction = new Transaction(generateTransactionId(), user, stock,
                                            TransactionType::BUY, quantity,
                                            stock->getCurrentPrice());
    transactions.push_back(transaction);
    
    updatePortfolioValues();
    return transaction;
}

Transaction* BrokerageSystem::sellStock(const std::string& userId, const std::string& symbol, int quantity) {
    std::lock_guard<std::mutex> guard(mutex);
    User* user = findUser(userId);
    Stock* stock = findStock(symbol);
    
    if (!user || !stock || !stock->isActive()) return nullptr;
    
    Portfolio* portfolio = user->getPortfolio();
    if (portfolio->getShareQuantity(symbol) < quantity) {
        return nullptr;
    }
    
    // Process transaction
    double totalAmount = stock->getCurrentPrice() * quantity;
    user->deposit(totalAmount);
    stock->updateShares(quantity);
    portfolio->removeShares(symbol, quantity);
    
    // Create transaction record
    Transaction* transaction = new Transaction(generateTransactionId(), user, stock,
                                            TransactionType::SELL, quantity,
                                            stock->getCurrentPrice());
    transactions.push_back(transaction);
    
    updatePortfolioValues();
    return transaction;
}

void BrokerageSystem::displayStocks() const {
    std::lock_guard<std::mutex> guard(mutex);
    std::cout << "\nAvailable Stocks:" << std::endl;
    for (const auto& pair : stocks) {
        pair.second->displayInfo();
        std::cout << "------------------------" << std::endl;
    }
}

void BrokerageSystem::displayUsers() const {
    std::lock_guard<std::mutex> guard(mutex);
    std::cout << "\nRegistered Users:" << std::endl;
    for (const auto& user : users) {
        user->displayInfo();
        std::cout << "------------------------" << std::endl;
    }
}

void BrokerageSystem::displayUserPortfolio(const std::string& userId) const {
    std::lock_guard<std::mutex> guard(mutex);
    if (User* user = findUser(userId)) {
        user->getPortfolio()->displayInfo();
    }
}

void BrokerageSystem::displayTransactionHistory(const std::string& userId) const {
    std::lock_guard<std::mutex> guard(mutex);
    std::cout << "\nTransaction History:" << std::endl;
    for (const auto& transaction : transactions) {
        if (transaction->getUser()->getUserId() == userId) {
            transaction->displayInfo();
            std::cout << "------------------------" << std::endl;
        }
    }
}

User* BrokerageSystem::findUser(const std::string& userId) const {
    auto it = std::find_if(users.begin(), users.end(),
        [userId](User* user) { return user->getUserId() == userId; });
    return it != users.end() ? *it : nullptr;
}

Stock* BrokerageSystem::findStock(const std::string& symbol) const {
    auto it = stocks.find(symbol);
    return it != stocks.end() ? it->second : nullptr;
}

void BrokerageSystem::updatePortfolioValues() {
    for (auto user : users) {
        user->getPortfolio()->updateTotalValue(stocks);
    }
}

std::string BrokerageSystem::generateTransactionId() {
    return "T" + std::to_string(transactionIdCounter++);
} 