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
        bool hours;

        public BallStack(int capacity, IBallStack reciever, IBallQueue queue, bool hours = false)
        {
            this.capacity = capacity;
            if(hours)
            {
                capacity += 1;
            }
            this.reciever = reciever;
            this.queue = queue;
            this.balls = new Stack<Ball>();
            this.hours = hours;
            if(hours)
            {
                balls.Push(new Ball(0) { Id = 0 });
            }
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
            while(balls.Count > 0 || (hours && balls.Count > 1))
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
