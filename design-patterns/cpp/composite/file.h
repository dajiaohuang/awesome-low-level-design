#pragma once
#include "file_system_item.h"
#include <string>

class File : public FileSystemItem {
public:
    File(const std::string& name, int size);
    int getSize() const override;
    void printStructure(const std::string& indent) const override;
    void deleteItem() override;

private:
    std::string name;
    int size;
}; 