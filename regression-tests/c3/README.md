# Focused example regressions

These checks exercise the existing public service/state workflows. They do not certify every sample project, race freedom, strict TypeScript typing, POSIX timestamp execution, or real payment processing.

From the repository root:

- `python -B regression-tests/c3/task_collections.py`
- `python -B regression-tests/c3/restaurant_orders.py`
- `dotnet run --project regression-tests/c3/TaskCollections.csproj`
- `dotnet run --project regression-tests/c3/RestaurantOrders.csproj`
- Compile `regression-tests/c3/task_collections.cpp` together with `Task.cpp`, `User.cpp`, and `TaskManager.cpp`, with `solutions/cpp/taskmanagementsystem` on the include path and C++17/thread support, then run the resulting executable.
- With an existing jiti installation selected by `JITI_MODULE`, run `node regression-tests/c3/vending_refund.cjs`. This invokes the machine-owned public state getter, not a separately constructed state.
- Compile the Java coffee package and its tracked JUnit test against JUnit Jupiter. `CoffeeTests.java` is a standalone entrypoint for executing those four native assertions when no JUnit launcher is available; it is not a replacement test framework.

The C# regression projects deliberately compile one sample each. Their success does not establish that the root aggregate C# project compiles. The Java restaurant layer incompatibility and missing C++ chess Board/Position definitions remain separate upstream blockers.
