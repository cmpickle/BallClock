using System;
using System.Collections;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace BallClock
{
    class Ball : IEqualityComparer<Ball>
    {
        public int Id { get; set; }

        public Ball(int Id)
        {
            this.Id = Id;
        }

        public bool Equals(Ball x, Ball y)
        {
            return x.Id == y.Id;
        }

        public int GetHashCode(Ball obj)
        {
            return obj.Id.GetHashCode();
        }
    }
}
