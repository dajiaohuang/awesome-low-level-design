# C++ Vending Machine

This example directly constructs `VendingMachine(machineId)`. It is not a Singleton or a State-pattern implementation.

- `Product` stores its ID, name, price, quantity and availability.
- `VendingMachine` owns products and transaction records, supports restocking, price changes, purchases and cash collection. Its public operations share a mutex.
- `purchaseProduct(productId, quantity, payment)` checks stock and the monetary payment, deducts the purchased quantity, retains the cost in the machine balance and reports `payment - cost` as change. Failure returns `nullptr` without a successful transaction.
- `Transaction` records the product, quantity and total cost.
- `VendingMachineDemo.cpp` demonstrates these operations.

Payment and change are amounts in this console simulation; physical coin/note inventories and denomination-specific change dispensing are not implemented.
