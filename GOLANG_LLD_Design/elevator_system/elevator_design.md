# Design an Elevator System

## Requirements

1. The system should handle multiple elevators in a building.
2. Each elevator should manage its own state (moving, idle, maintenance).
3. The system should handle requests for different floors.
4. The system should optimize for efficiency (e.g., minimizing wait time).

### Classes and Their Relationships

#### 1. `ElevatorSystem`

- **Attributes:**
  - `elevators: List[Elevator]`
  - `requests: List[Request]`
- **Methods:**
  - `__init__(elevators)`
  - `request_elevator(request: Request)`
  - `assign_elevator()`

#### 2. `Elevator`

- **Attributes:**
  - `elevator_id: str`
  - `current_floor: int`
  - `target_floors: List[int]`
  - `status: str` (e.g., 'moving', 'idle', 'maintenance')
- **Methods:**
  - `__init__(elevator_id)`
  - `move_to_floor(floor: int)`
  - `add_target_floor(floor: int)`
  - `update_status(status: str)`

#### 3. `Request`

- **Attributes:**
  - `request_id: str`
  - `origin_floor: int`
  - `destination_floor: int`
  - `timestamp: datetime`
- **Methods:**
  - `__init__(request_id, origin_floor, destination_floor)`

### Interaction Diagram

1. A request is made to the `ElevatorSystem`.
2. The system logs the request.
3. `ElevatorSystem` assigns the request to the most appropriate `Elevator`.
4. The `Elevator` moves to the origin floor and then to the destination floor.
5. Status of the elevator is updated accordingly.

