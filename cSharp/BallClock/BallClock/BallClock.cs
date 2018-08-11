using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace BallClock
{
    class BallClock
    {
        BallQueue queue;
        BallStack hoursStack;
        BallStack fiveMinutesStack;
        BallStack minutesStack;

        public BallClock(int balls)
        {
            queue = new BallQueue(balls);
            hoursStack = new BallStack(11, null, queue, true);
            fiveMinutesStack = new BallStack(11, hoursStack, queue);
            minutesStack = new BallStack(4, fiveMinutesStack, queue);
        }

        public void Tick()
        {
            queue.Tick(minutesStack);
        }

        public override string ToString()
        {
            string result = "";

            result += "Queue:\n";
            result += queue.ToString();
            result += "\n";
            result += "Minutes:\n";
            result += minutesStack.ToString();
            result += "\n";
            result += "Five Minutes:\n";
            result += fiveMinutesStack.ToString();
            result += "\n";
            result += "Hours:\n";
            result += hoursStack.ToString();
            result += "\n";
            result += "\n";

            return result;
        }
    }
}
