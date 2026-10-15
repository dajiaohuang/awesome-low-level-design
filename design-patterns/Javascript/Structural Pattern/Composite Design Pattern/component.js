// Abstract Class  or Component class

class Component {
  showDetails() {
    throw new Error("This method is overriden by subclasses");
  }
}

module.exports = Component;
