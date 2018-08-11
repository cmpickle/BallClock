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
                balls.Enqueue(new Ball());
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

        public override String ToString()
        {
            return new JavaScriptSerializer().Serialize(balls.ToList());
        }
    }
}
