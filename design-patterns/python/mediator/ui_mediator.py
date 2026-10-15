from abc import ABC, abstractmethod
from typing import TYPE_CHECKING
if TYPE_CHECKING:
    from .ui_component import UIComponent

class UIMediator(ABC):
    @abstractmethod
    def component_changed(self, component: "UIComponent") -> None:
        pass 