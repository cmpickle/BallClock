class Queue {
  constructor() {
    this.start = { value: undefined, next: undefined };
    this.end = { value: undefined, next: undefined };
    this.length = 0;
  }

  dequeue() {
    if (this.length === 0) {
      return undefined;
    }
    let n = this.start;
    if (this.length === 1) {
      this.start = undefined;
      this.end = undefined;
    } else {
      this.start = this.start.next;
    }
    this.length--;
    return n;
  }

  enqueue(n) {
    if (this.length === 0) {
      this.start = n;
      this.end = n;
    } else {
      this.end.next = n;
      this.end = n;
    }
    this.length++;
  }

  len() {
    return this.length;
  }

  peek() {
    if (this.length === 0) {
      return undefined;
    }
    return this.start.value;
  }

  elementAt(pos) {
    let curr = this.start;

    for (let i = 0; i < pos; i++) {
      curr = curr.next;
    }

    return curr;
  }
}

module.exports = Queue;
