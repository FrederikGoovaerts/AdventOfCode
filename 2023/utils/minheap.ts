import { isEqual } from "lodash";

export class MinHeap<T> {
  private heap: T[];
  private compare: (a: T, b: T) => number;

  constructor(compare: (a: T, b: T) => number) {
    this.heap = [];
    this.compare = compare;
  }

  size(): number {
    return this.heap.length;
  }

  has(value: T): boolean {
    for (const el of this.heap) {
      if (isEqual(el, value)) {
        return true;
      }
    }
    return false;
  }

  insert(value: T) {
    this.heap.push(value);
    this.bubbleUp();
  }

  removeMin(): T | null {
    if (this.heap.length === 0) {
      return null;
    }

    const minValue = this.heap[0];
    const lastValue = this.heap.pop()!;

    if (this.heap.length > 0) {
      this.heap[0] = lastValue;
      this.bubbleDown();
    }

    return minValue;
  }

  private bubbleUp() {
    let index = this.heap.length - 1;

    while (index > 0) {
      const element = this.heap[index];
      const parentIndex = Math.floor((index - 1) / 2);
      const parent = this.heap[parentIndex];

      if (this.compare(element, parent) < 0) {
        this.swap(index, parentIndex);
        index = parentIndex;
      } else {
        break;
      }
    }
  }

  private bubbleDown() {
    let index = 0;
    const length = this.heap.length;
    const element = this.heap[0];
    const done = false;

    while (!done) {
      const leftChildIndex = 2 * index + 1;
      const rightChildIndex = 2 * index + 2;
      let leftChild, rightChild;
      let swap = null;

      if (leftChildIndex < length) {
        leftChild = this.heap[leftChildIndex];
        if (this.compare(leftChild, element) < 0) {
          swap = leftChildIndex;
        }
      }

      if (rightChildIndex < length) {
        rightChild = this.heap[rightChildIndex];
        if (
          (swap === null && this.compare(rightChild, element) < 0) ||
          (swap !== null && this.compare(rightChild, leftChild!) < 0)
        ) {
          swap = rightChildIndex;
        }
      }

      if (swap === null) {
        break;
      }

      this.swap(index, swap);
      index = swap;
    }
  }

  private swap(a: number, b: number) {
    const temp = this.heap[a];
    this.heap[a] = this.heap[b];
    this.heap[b] = temp;
  }
}
