class BallClock {
  constructor(balls) {
    this.count = balls;
    this.min = [];
    this.fiveMin = [];
    this.hour = [];
    this.main = [];
    for (let i = 0; i < this.count; i++) {
      this.main.push(i + 1);
    }
  }

  ticks(minutes) {
    for (let i = 0; i < minutes; i++) {
      this.tick();
    }
  }

  tick() {
    let ball = this.main.shift();
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

          this.main.push(ball);
        }
      }
    }
  }

  returnBalls(stack) {
    while (stack.length > 0) {
      let ball = stack.pop();
      if (ball !== undefined) {
        this.main.push(ball);
      }
    }
  }

  IsStartingOrder() {
    for (let i = 0; i < this.main.length; i++) {
      if (this.main[i] !== i + 1) {
        return false;
      }
    }
    return true;
  }

  // Count   int
  // Min     *stack.Stack `json:"Min"`
  // FiveMin *stack.Stack `json:"FiveMin"`
  // Hour    *stack.Stack `json:"Hour"`
  // Main    *queue.Queue `json:"Main"`
}

module.exports = BallClock;
