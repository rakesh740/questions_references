# Designing a Parking Lot System

## requirements

 1. The parking lot should have multiple levels, each level with a certain number of parking spots.
 2. The parking lot should support different types of vehicles, such as cars, motorcycles, and trucks.
 3. Each parking spot should be able to accommodate a specific type of vehicle.
 4. The system should assign a parking spot to a vehicle upon entry and release it when the vehicle exits.
 5. The system should track the availability of parking spots and provide real-time information to customers.
 6. The system should handle multiple entry and exit points and support concurrent access.
 7. Each vehicle type requires a different amount of space.
 8. The parking lot should be able to handle payments and issue tickets.

## Algorithm

### Classes and Their Relationships:

#### 1. `ParkingLot`

- **Attributes:**
  - `name: str`
  - `location: str`
  - `capacity: dict` (e.g., {'car': 100, 'bike': 50, 'truck': 20})
  - `available_spots: dict` (e.g., {'car': 90, 'bike': 50, 'truck': 20})
  - `tickets: List[Ticket]`
- **Methods:**
  - `__init__(name, location, capacity)`
  - `is_space_available(vehicle_type: str) -> bool`
  - `park_vehicle(vehicle: Vehicle) -> Ticket`
  - `unpark_vehicle(ticket: Ticket) -> float`

#### 2. `Ticket`

- **Attributes:**
  - `ticket_id: str`
  - `vehicle: Vehicle`
  - `issued_at: datetime`
  - `paid: bool`
- **Methods:**
  - `__init__(ticket_id, vehicle, issued_at)`

#### 3. `Vehicle`

- **Attributes:**
  - `license_plate: str`
  - `vehicle_type: str`
- **Methods:**
  - `__init__(license_plate, vehicle_type)`

#### 4. `PaymentProcessor`

- **Methods:**
  - `process_payment(ticket: Ticket) -> bool`

### Interaction Diagram

1. A vehicle arrives and requests to park.
2. The `ParkingLot` checks if space is available using `is_space_available(vehicle_type)`.
3. If space is available, a `Ticket` is issued.
4. When the vehicle leaves, the ticket is processed for payment using `PaymentProcessor`.
5. Once paid, the `ParkingLot` updates its available spaces.
