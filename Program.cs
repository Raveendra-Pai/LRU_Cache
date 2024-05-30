using System.Collections.Generic;
using System.Dynamic;

namespace LRUCache
{
    internal class Program
    {
        static void Main(string[] args)
        {
            Console.WriteLine("LRU Cache Demo");
            var lcache = new LCache<string, string>(3);
            lcache.Put("prakash", "shetty");
            lcache.Put("ajit", "ganiga");
            lcache.Put("prathik", "rai");
            Console.WriteLine($"1. Head : {lcache.Front()}  Tail : {lcache.Tail()}");
            
            Console.WriteLine($"accessing 'prakash' value returned: {lcache.Get("prakash")}");
            Console.WriteLine($"2. Head : {lcache.Front()}  Tail : {lcache.Tail()}");
            lcache.Put("adhip", "kamat");
            Console.WriteLine(String.Format("3. Head : {0} Tail : {1}", lcache.Front(), lcache.Tail()));

        }
    }
}
