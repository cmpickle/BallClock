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
            Console.WriteLine();
            BallStack ballStack = new BallStack(4, queue);
            queue.Tick(ballStack, null);
            Console.WriteLine(queue.ToString());
            Console.WriteLine(ballStack.ToString());
            Console.WriteLine();

            queue.Tick(ballStack, null);
            Console.WriteLine(queue.ToString());
            Console.WriteLine(ballStack.ToString());
            Console.WriteLine();

            queue.Tick(ballStack, null);
            Console.WriteLine(queue.ToString());
            Console.WriteLine(ballStack.ToString());
            Console.WriteLine();

            queue.Tick(ballStack, null);
            Console.WriteLine(queue.ToString());
            Console.WriteLine(ballStack.ToString());
            Console.WriteLine();

            queue.Tick(ballStack, null);
            Console.WriteLine(queue.ToString());
            Console.WriteLine(ballStack.ToString());
            Console.WriteLine();

            queue.Tick(ballStack, null);
            Console.WriteLine(queue.ToString());
            Console.WriteLine(ballStack.ToString());
            Console.WriteLine();
            Console.Read();
        }
    }
}
