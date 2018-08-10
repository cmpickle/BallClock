using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace BallClock
{
    interface IBallStack
    {
        Ball AddBall(Ball ball, Delegate reciever);

        void ReturnBalls(IBallQueue queue);
    }
}
