using System;
using System.IO;
using System.Linq;
using System.Collections.Concurrent;
using System.Collections.Generic;
using System.Threading.Tasks;

class RestaurantOrdersTest
{
    public static void Main()
    {
        var facade = RestaurantManagementSystemFacade.GetInstance();
        facade.AddTable(1, 4);
        facade.AddWaiter("waiter", "Alice");
        facade.AddChef("chef", "Gordon");
        facade.AddMenuItem("item", "Pizza", 10);
        var ids = new ConcurrentBag<int>();
        var output = Console.Out;
        Console.SetOut(TextWriter.Null);
        try
        {
            Parallel.For(0, 2000, i => {
                var order = facade.TakeOrder(1, "waiter", new List<string> {"item"});
                int id = order.GetOrderId();
                facade.GenerateBill(id);
                facade.MarkItemsAsReady(id);
                facade.ServeOrder("waiter", id);
                ids.Add(id);
            });
        }
        finally { Console.SetOut(output); }
        if (ids.Distinct().Count() != 2000) throw new Exception("Order IDs lost");
        foreach (var id in ids) facade.GenerateBill(id);
        Console.WriteLine("PASS concurrent order publication and lookup");
    }
}
