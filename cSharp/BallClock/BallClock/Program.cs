using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace BallClock
{
    class Program
    {
        static void Main(string[] args)
        {
            BallQueue queue = new BallQueue(5);
            Console.WriteLine(queue.ToString());
            Console.Read();
        }
    }
}
