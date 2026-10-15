#include "TaskManager.hpp"
#include <cassert>
#include <thread>
#include <vector>
#include <iostream>

int main() {
    TaskManager manager;
    auto user = manager.registerUser("Author", "author@example.com");
    const auto userId = user->getUserId();
    std::vector<std::thread> workers;
    for (int i = 0; i < 16; ++i) workers.emplace_back([&] {
        for (int j = 0; j < 500; ++j) {
            auto task = manager.createTask("task", "body");
            const auto id = task->getTaskId();
            assert(manager.assignTask(id, userId));
            manager.getTasksByPriority(TaskPriority::MEDIUM);
            manager.getTasksByStatus(TaskStatus::TODO);
            manager.removeTask(id);
        }
    });
    for (auto& worker : workers) worker.join();
    assert(manager.getTasksByPriority(TaskPriority::MEDIUM).empty());
    std::cout << "PASS task collection contention\n";
}
