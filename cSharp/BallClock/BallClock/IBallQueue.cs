using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace BallClock
{
    interface IBallQueue
    {
        Ball Tick();

        void QueueBalls(IEnumerable<Ball> balls);
    }
}
