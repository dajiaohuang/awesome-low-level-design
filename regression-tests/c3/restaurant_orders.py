import contextlib
import io
import sys
from concurrent.futures import ThreadPoolExecutor
from pathlib import Path

sys.path.insert(0, str(Path(__file__).resolve().parents[2] / 'solutions/python/restaurantmanagementsystem'))
from restaurant_management_system_facade import RestaurantManagementSystemFacade

sys.setswitchinterval(0.000001)
with ThreadPoolExecutor(max_workers=32) as pool:
    instances = list(pool.map(lambda _: RestaurantManagementSystemFacade.get_instance(), range(100)))
facade = instances[0]
assert all(instance is facade for instance in instances)
facade.add_table(1, 4)
facade.add_waiter('waiter', 'Alice')
facade.add_chef('chef', 'Gordon')
for i in range(1000):
    facade.add_menu_item(str(i), f'item-{i}', float(i + 1))

def place(i):
    order = facade.take_order(1, 'waiter', [str(i)])
    facade.generate_bill(order.order_id)
    return i, order

sys.setswitchinterval(0.000001)
with contextlib.redirect_stdout(io.StringIO()):
    with ThreadPoolExecutor(max_workers=32) as pool:
        orders = list(pool.map(place, range(1000)))
assert len({order.order_id for _, order in orders}) == 1000
for i, order in orders:
    output = io.StringIO()
    with contextlib.redirect_stdout(output):
        facade.generate_bill(order.order_id).print_bill()
    assert f'Total: ${(i + 1) * 1.08 + 5:.2f}' in output.getvalue(), output.getvalue()
print('PASS concurrent order identity and public bill lookup')
