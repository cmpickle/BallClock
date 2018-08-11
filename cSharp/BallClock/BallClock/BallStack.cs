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
        IBallQueue queue;

        public BallStack(int capacity, IBallQueue queue)
        {
            this.capacity = capacity;
            this.queue = queue;
            this.balls = new Stack<Ball>();
        }

        public void AddBall(Ball ball, IBallStack reciever)
        {
            if(balls.Count == capacity)
            {
                ReturnBalls(queue);
                if(reciever == null)
                {
                    queue.QueueBall(ball);
                    return;
                }
                reciever.AddBall(ball, reciever);
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
