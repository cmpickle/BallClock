using Newtonsoft.Json;
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
                        Console.WriteLine($"{args[1]} is not a valid number of balls. Please enter a number of balls from 27-127.");
                        break;
                    }
                    Console.WriteLine(repetitionTime(days));
                    break;
                case "--timedOutput":
                    int balls;
                    int minutes;
                    if(!Int32.TryParse(args[1], out balls) || balls < 27 || balls > 127)
                    {
                        Console.WriteLine($"{args[1]} is not a valid number of balls. Please enter a number of balls from 27-127.");
                        break;
                    }
                    if(!Int32.TryParse(args[2], out minutes) || minutes < 0 || minutes > 1000000)
                    {
                        Console.WriteLine($"{args[1]} is not a valid number of minutes. Please enter a number of minutes from 0-1000000.");
                        break;
                    }
                    Console.WriteLine(getClockJson(balls, minutes));
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

        public static string getClockJson(int balls, int minutes)
        {
            BallClock increment = new BallClock(balls);

            for(int i = 0; i < minutes; ++i)
            {
                increment.Tick();
            }

            return increment.ToJson();
        }
    }
}
