using System;
using System.Collections;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using System.Web.Script.Serialization;

namespace BallClock
{
    class BallStack : IBallStack
    {
        int capacity;
        Stack<Ball> balls;
        IBallStack reciever;
        IBallQueue queue;

        public BallStack(int capacity, IBallStack reciever, IBallQueue queue)
        {
            this.capacity = capacity;
            this.reciever = reciever;
            this.queue = queue;
            this.balls = new Stack<Ball>();
        }

        public void AddBall(Ball ball)
        {
            if(balls.Count == capacity)
            {
                ReturnBalls(queue);
                if(reciever == null)
                {
                    queue.QueueBall(ball);
                    return;
                }
                reciever.AddBall(ball);
                return;
            }
            balls.Push(ball);
        }

        public void ReturnBalls(IBallQueue queue)
        {
            while(balls.Count > 0)
            {
                queue.QueueBall(balls.Pop());
            }
        }

        public override String ToString()
        {
            return new JavaScriptSerializer().Serialize(balls.ToList());
        }
    }
}
