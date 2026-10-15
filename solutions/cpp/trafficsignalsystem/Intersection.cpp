#include "Intersection.hpp"
#include <iostream>
#include <algorithm>

Intersection::Intersection(std::string intersectionId)
    : intersectionId(intersectionId), isOperational(true) {}

Intersection::~Intersection() {
    for (auto signal : signals) {
        delete signal;
    }
}

std::string Intersection::getIntersectionId() const { return intersectionId; }
bool Intersection::isWorking() const { return isOperational; }

void Intersection::addSignal(Signal* signal) {
    signals.push_back(signal);
}

void Intersection::removeSignal(const std::string& signalId) {
    auto it = std::find_if(signals.begin(), signals.end(),
        [signalId](Signal* signal) { return signal->getSignalId() == signalId; });
    
    if (it != signals.end()) {
        delete *it;
        signals.erase(it);
    }
}

void Intersection::updateSignals(int timeElapsed) {
    if (!isOperational) return;
    
    for (;;) {
        Signal* active = nullptr;
        for (auto signal : signals) {
            if (signal->isOperational() && signal->currentColor != SignalColor::RED)
                active = signal;
        }
        // A ready red signal waits for the current green/yellow phase to finish.
        if (!active) {
            for (std::size_t offset = 0; offset < signals.size(); ++offset) {
                std::size_t index = (nextSignal + offset) % signals.size();
                auto signal = signals[index];
                if (signal->isOperational() && signal->timeRemaining == 0) {
                    signal->switchColor();
                    active = signal;
                    nextSignal = index + 1;
                    break;
                }
            }
        }
        if (timeElapsed == 0) break;
        int step = timeElapsed;
        if (active) {
            step = std::min(step, active->timeRemaining);
        } else {
            for (auto signal : signals)
                if (signal->isOperational()) step = std::min(step, signal->timeRemaining);
        }
        for (auto signal : signals) {
            if (!signal->isOperational()) continue;
            if (signal == active) signal->updateSignal(step);
            else signal->timeRemaining = std::max(0, signal->timeRemaining - step);
        }
        timeElapsed -= step;
    }
}

void Intersection::setOperational(bool status) {
    isOperational = status;
    for (auto signal : signals) {
        signal->setWorking(status);
    }
}

void Intersection::synchronizeSignals() {
    nextSignal = 0;
    for (auto signal : signals) {
        signal->reset();
    }
}

void Intersection::displayStatus() const {
    std::cout << "\nIntersection " << intersectionId << " Status:" << std::endl;
    std::cout << "Operational: " << (isOperational ? "Yes" : "No") << std::endl;
    std::cout << "Signals:" << std::endl;
    for (const auto& signal : signals) {
        signal->displayStatus();
    }
}

Signal* Intersection::findSignal(const std::string& signalId) const {
    auto it = std::find_if(signals.begin(), signals.end(),
        [signalId](Signal* signal) { return signal->getSignalId() == signalId; });
    return it != signals.end() ? *it : nullptr;
} 