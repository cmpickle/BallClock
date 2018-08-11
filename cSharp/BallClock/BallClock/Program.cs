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
            BallQueue queue = new BallQueue(27);
            BallStack hoursStack = new BallStack(11, null, queue);
            BallStack fiveMinutesStack = new BallStack(11, hoursStack, queue);
            BallStack minutesStack = new BallStack(4, fiveMinutesStack, queue);

            Console.WriteLine(queue.ToString());
            Console.WriteLine();
            queue.Tick(minutesStack);
            Console.WriteLine(queue.ToString());
            Console.WriteLine(minutesStack.ToString());
            Console.WriteLine(fiveMinutesStack.ToString());
            Console.WriteLine(hoursStack.ToString());
            Console.WriteLine();

            queue.Tick(minutesStack);
            Console.WriteLine(queue.ToString());
            Console.WriteLine(minutesStack.ToString());
            Console.WriteLine(fiveMinutesStack.ToString());
            Console.WriteLine(hoursStack.ToString());
            Console.WriteLine();

            queue.Tick(minutesStack);
            Console.WriteLine(queue.ToString());
            Console.WriteLine(minutesStack.ToString());
            Console.WriteLine(fiveMinutesStack.ToString());
            Console.WriteLine(hoursStack.ToString());
            Console.WriteLine();

            queue.Tick(minutesStack);
            Console.WriteLine(queue.ToString());
            Console.WriteLine(minutesStack.ToString());
            Console.WriteLine(fiveMinutesStack.ToString());
            Console.WriteLine(hoursStack.ToString());
            Console.WriteLine();

            queue.Tick(minutesStack);
            Console.WriteLine(queue.ToString());
            Console.WriteLine(minutesStack.ToString());
            Console.WriteLine(fiveMinutesStack.ToString());
            Console.WriteLine(hoursStack.ToString());
            Console.WriteLine();

            queue.Tick(minutesStack);
            Console.WriteLine(queue.ToString());
            Console.WriteLine(minutesStack.ToString());
            Console.WriteLine(fiveMinutesStack.ToString());
            Console.WriteLine(hoursStack.ToString());
            Console.WriteLine();
            Console.Read();
        }
    }
}
