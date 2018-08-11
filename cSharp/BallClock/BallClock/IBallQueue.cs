using System;
using System.Collections;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace BallClock
{
    interface IBallQueue
    {
        void Tick(IBallStack ballStack);

        void QueueBall(Ball ball);
    }
}
