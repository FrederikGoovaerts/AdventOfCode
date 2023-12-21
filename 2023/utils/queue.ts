class QueueNode<T> {
  value: T;
  next?: QueueNode<T>;

  constructor(value: T) {
    this.value = value;
  }
}

export class Queue<T> {
  private front?: QueueNode<T>;
  private rear?: QueueNode<T>;

  enqueue(value: T): void {
    const newNode = new QueueNode(value);

    if (!this.front) {
      // If the queue is empty, set both front and rear to the new node
      this.front = newNode;
      this.rear = newNode;
    } else {
      // Otherwise, add the new node to the rear and update the rear
      this.rear!.next = newNode;
      this.rear = newNode;
    }
  }

  dequeue(): T | undefined {
    if (!this.front) {
      // If the queue is empty, return undefined
      return undefined;
    }

    const removedValue = this.front.value;

    if (this.front === this.rear) {
      // If there is only one element in the queue, set both front and rear to undefined
      this.front = undefined;
      this.rear = undefined;
    } else {
      // Otherwise, update the front to the next node
      this.front = this.front.next;
    }

    return removedValue;
  }

  isEmpty(): boolean {
    return this.front === undefined;
  }
}
