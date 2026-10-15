#ifndef LRU_CACHE_HPP
#define LRU_CACHE_HPP

#include <unordered_map>
#include <mutex>
#include <stdexcept>
#include <iostream>
#include "DoublyLinkedList.hpp"

template<typename K, typename V>
class LRUCache {
private:
    mutable std::mutex mutex;
    int capacity;
    DoublyLinkedList<K,V> dll;
    std::unordered_map<K, Node<K,V>*> cache;

public:
    LRUCache(int capacity) : capacity(capacity) {}
    
    V get(K key) {
        std::lock_guard<std::mutex> guard(mutex);
        auto it = cache.find(key);
        if (it == cache.end()) {
            throw std::runtime_error("Key not found");
        }
        
        Node<K,V>* node = it->second;
        dll.moveToFront(node);
        return node->value;
    }
    
    void put(K key, V value) {
        std::lock_guard<std::mutex> guard(mutex);
        auto it = cache.find(key);
        if (it != cache.end()) {
            // Key exists, update value and move to front
            Node<K,V>* node = it->second;
            node->value = value;
            dll.moveToFront(node);
        } else {
            // Key doesn't exist, add new entry
            if (cache.size() >= capacity) {
                // Remove least recently used item
                Node<K,V>* lastNode = dll.getTail();
                cache.erase(lastNode->key);
                dll.removeNode(lastNode);
            }
            
            // Add new item to front
            Node<K,V>* newNode = dll.addToFront(key, value);
            cache[key] = newNode;
        }
    }
    
    bool contains(K key) const {
        std::lock_guard<std::mutex> guard(mutex);
        return cache.find(key) != cache.end();
    }
    
    int getSize() const {
        std::lock_guard<std::mutex> guard(mutex);
        return cache.size();
    }
    
    bool isEmpty() const {
        std::lock_guard<std::mutex> guard(mutex);
        return cache.empty();
    }
    
    void clear() {
        std::lock_guard<std::mutex> guard(mutex);
        while (!dll.isEmpty()) {
            Node<K,V>* node = dll.removeLast();
            cache.erase(node->key);
            delete node;
        }
    }
    
    void display() const {
        std::lock_guard<std::mutex> guard(mutex);
        std::cout << "\nLRU Cache Contents:" << std::endl;
        Node<K,V>* current = dll.getHead();
        while (current) {
            std::cout << current->key << " -> " << current->value << std::endl;
            current = current->next;
        }
        std::cout << "Cache size: " << cache.size() << "/" << capacity << std::endl;
    }
};

#endif 