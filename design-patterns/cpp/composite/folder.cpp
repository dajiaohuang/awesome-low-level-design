#include "folder.h"
#include <iostream>

Folder::Folder(const std::string& name) : name(name) {}

void Folder::addItem(std::shared_ptr<FileSystemItem> item) {
    children.push_back(item);
}

int Folder::getSize() const {
    int size = 0;
    for (const auto& item : children) size += item->getSize();
    return size;
}

void Folder::printStructure(const std::string& indent) const {
    std::cout << indent << "+ " << name << std::endl;
    for (const auto& item : children) item->printStructure(indent + "  ");
}

void Folder::deleteItem() {
    for (const auto& item : children) item->deleteItem();
    std::cout << "Deleting folder: " << name << std::endl;
}
