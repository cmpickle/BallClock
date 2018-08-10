using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace BallClock
{
    class Ball
    {
        public static int BallCount = 0;
        public int Id { get; set; }

        public Ball()
        {
            Id = ++BallCount;
        }
    }
}
