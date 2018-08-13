using Newtonsoft.Json;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using System.Web.Script.Serialization;

namespace BallClock
{
    class BallQueue : IBallQueue
    {
        Queue<Ball> balls = new Queue<Ball>();

        public BallQueue(int numBalls)
        {
            for (int i = 0; i < numBalls; ++i)
            {
                balls.Enqueue(new Ball(i+1));
            }
        }

        public void QueueBall(Ball queuedBall)
        {
            balls.Enqueue(queuedBall);
        }

        public void Tick(IBallStack ballStack)
        {
            ballStack.AddBall(balls.Dequeue());
        }

        public int getCount()
        {
            return balls.Count;
        }

        public override String ToString()
        {
            return new JavaScriptSerializer().Serialize(balls.ToList());
        }

        public int[] ToArray()
        {
            int[] result = new int[balls.Count];

            for(int i = 0; i < balls.Count; ++i)
            {
                result[i] = balls.ElementAt(i).Id;
            }
            return result;
        }
    }
}
