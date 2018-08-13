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
            switch(args[0])
            {
                case "--repetitionTime":
                    int days;
                    if(!Int32.TryParse(args[1], out days))
                    {
                        Console.WriteLine($"{args[1]} is not a valid option for repetitionTime. Please enter the number of balls to use in the clock.");
                        break;
                    }
                    Console.WriteLine(repetitionTime(days));
                    break;
                default:
                    Console.WriteLine($"{args[0]}  is not a valid command");
                    break;
            }
            BallClock clock = new BallClock(27);

            //Console.WriteLine(clock.ToString());

            //int length = 120;
            //for(int i = 0; i < length; ++i)
            //{
            //    clock.Tick();
            //    Console.WriteLine(clock.ToString());
            //}

            Console.Read();
        }

        public static string repetitionTime(int balls)
        {
            int days = 0;

            BallClock increment = new BallClock(balls);
            do
            {
                for(int i = 0; i < 1440; ++i)
                {
                    increment.Tick();
                }
                ++days;
            }
            while (!increment.isStartingOrder());

            return $"{balls} balls cycle after {days} days";
        }
    }
}
