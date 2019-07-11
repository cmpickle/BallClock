class Stack {
  constructor() {
    this.top = { value: undefined, next: undefined };
    this.length = 0;
  }

  len() {
    return this.length;
  }

  peek() {
    if (this.length === 0) {
      return undefined;
    }
    return this.top.value;
  }

  pop() {
    if (this.length === 0) {
      return undefined;
    }

    let n = this.top;
    this.top = n.next;
    this.length--;
    return n;
  }

  push(n) {
    n.next = this.top;
    this.top = n;
    this.length++;
  }

  elementAt(pos) {
    let curr = this.top;

    for (let i = 0; i < pos; i++) {
      curr = curr.next;
    }

    return curr;
  }
}

module.exports = Stack;
