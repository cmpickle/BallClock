using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace BallClock
{
    interface IBallStack
    {
        void AddBall(Ball ball);

        void ReturnBalls(IBallQueue queue);
    }
}
