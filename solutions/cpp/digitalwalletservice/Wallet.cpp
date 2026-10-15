#include "Wallet.hpp"
#include <iostream>
#include <iomanip>

Wallet::Wallet(std::string walletId, std::string userId)
    : walletId(walletId), userId(userId), balance(0.0) {}

Wallet::~Wallet() {
    for (auto transaction : transactions) {
        delete transaction;
    }
}

std::string Wallet::getWalletId() const { return walletId; }
std::string Wallet::getUserId() const { return userId; }
double Wallet::getBalance() const {
    std::lock_guard<std::mutex> guard(mutex);
    return balance;
}

bool Wallet::addMoney(double amount, std::string source) {
    std::lock_guard<std::mutex> guard(mutex);
    if (amount <= 0) return false;
    
    balance += amount;
    return true;
}

bool Wallet::withdrawMoney(double amount, std::string destination) {
    std::lock_guard<std::mutex> guard(mutex);
    if (amount <= 0 || amount > balance) return false;
    
    balance -= amount;
    return true;
}

void Wallet::addTransaction(Transaction* transaction) {
    std::lock_guard<std::mutex> guard(mutex);
    transactions.push_back(transaction);
}

void Wallet::displayInfo() const {
    std::lock_guard<std::mutex> guard(mutex);
    std::cout << "Wallet ID: " << walletId << std::endl;
    std::cout << "User ID: " << userId << std::endl;
    std::cout << "Balance: $" << std::fixed << std::setprecision(2) << balance << std::endl;
}

void Wallet::displayTransactions() const {
    std::lock_guard<std::mutex> guard(mutex);
    std::cout << "\nTransaction History:" << std::endl;
    for (const auto& transaction : transactions) {
        transaction->displayInfo();
        std::cout << "------------------------" << std::endl;
    }
} 