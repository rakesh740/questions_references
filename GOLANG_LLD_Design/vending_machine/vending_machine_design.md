# Design a Vending Machine

## Requirements

1. The vending machine should dispense various products.
2. It should accept multiple payment methods (coins, cards).
3. The machine should be able to provide change.
4. It should manage inventory and notify when items are out of stock.

### Classes and Their Relationships

#### 1. `VendingMachine`

- **Attributes:**
  - `inventory: dict` (e.g., {'Coke': (10, 1.25), 'Chips': (5, 0.75)})
  - `balance: float`
  - `transactions: List[Transaction]`
- **Methods:**
  - `__init__(inventory)`
  - `select_product(product_name: str) -> bool`
  - `insert_money(amount: float)`
  - `dispense_product(product_name: str)`
  - `refund() -> float`
  - `check_inventory() -> dict`

#### 2. `Transaction`

- **Attributes:**
  - `transaction_id: str`
  - `product: str`
  - `amount_paid: float`
  - `change_given: float`
- **Methods:**
  - `__init__(transaction_id, product, amount_paid, change_given)`

### Interaction Diagram

1. User selects a product.
2. `VendingMachine` checks inventory and price.
3. User inserts money.
4. `VendingMachine` verifies the amount.
5. If sufficient, product is dispensed and change is given if necessary. not necessary if using UPI
6. Inventory is updated.
7. check total amount.
