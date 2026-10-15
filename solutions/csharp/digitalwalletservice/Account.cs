using System;
using System.Collections.Generic;
using System.Numerics;

namespace DigitalWallet
{
    public class Account
    {
        private readonly string id;
        private readonly User user;
        private readonly string accountNumber;
        private readonly Currency currency;
        private decimal balance;
        private readonly List<Transaction> transactions;
        private readonly object accountLock = new object();

        public Account(string id, User user, string accountNumber, Currency currency)
        {
            this.id = id;
            this.user = user;
            this.accountNumber = accountNumber;
            this.currency = currency;
            this.balance = 0.0M;
            this.transactions = new List<Transaction>();
        }

        public void Deposit(decimal amount)
        {
            lock (accountLock)
            {
                balance = balance + amount;
            }
        }

        public void Withdraw(decimal amount)
        {
            lock (accountLock)
            {
                if (balance.CompareTo(amount) >= 0)
                {
                    balance = balance - amount;
                }
                else
                {
                    throw new InsufficientFundsException("Insufficient funds in the account.");
                }
            }
        }

        public void AddTransaction(Transaction transaction)
        {
            lock (accountLock)
            {
                transactions.Add(transaction);
            }
        }

        public string GetId()
        {
            return id;
        }

        public User GetUser()
        {
            return user;
        }

        public Currency GetCurrency()
        {
            return currency;
        }

        public decimal GetBalance()
        {
            lock (accountLock)
            {
                return balance;
            }
        }

        public List<Transaction> GetTransactions()
        {
            lock (accountLock)
            {
                return new List<Transaction>(transactions);
            }
        }
    }
}