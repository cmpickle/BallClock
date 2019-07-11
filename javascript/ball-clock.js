let Stack = require('./stack');
let Queue = require('./queue');

class BallClock {
  constructor(balls) {
    this.count = balls;
    this.min = new Stack();
    this.fiveMin = new Stack();
    this.hour = new Stack();
    this.main = new Queue();
    for (let i = 0; i < this.count; i++) {
      this.main.enqueue({ value: i + 1, next: undefined });
    }
  }

  ticks(minutes) {
    for (let i = 0; i < minutes; i++) {
      this.tick();
    }
  }

  tick() {
    let ball = this.main.dequeue();
    if (ball === undefined) {
      return;
    }

    if (this.min.length < 4) {
      this.min.push(ball);
    } else {
      this.returnBalls(this.min);

      if (this.fiveMin.length < 11) {
        this.fiveMin.push(ball);
      } else {
        this.returnBalls(this.fiveMin);

        if (this.hour.length < 11) {
          this.hour.push(ball);
        } else {
          this.returnBalls(this.hour);

          this.main.enqueue(ball);
        }
      }
    }
  }

  returnBalls(stack) {
    while (stack.length > 0) {
      let ball = stack.pop();
      if (ball !== undefined) {
        this.main.enqueue(ball);
      }
    }
  }

  IsStartingOrder() {
    for (let i = 0; i < this.main.length; i++) {
      if (this.main.elementAt(i).value !== i + 1) {
        return false;
      }
    }
    return true;
  }
}

module.exports = BallClock;
