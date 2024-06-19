/**
 * Definition for singly-linked list.
 * function ListNode(val, next) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.next = (next===undefined ? null : next)
 * }
 */
/**
 * @param {ListNode} head
 * @return {ListNode}
 */
var swapPairs = function (head) {
  const oldHead = head;
  let evenFlag = true;
  const nodeArray = [];
  while (head != null) {
    if (evenFlag) {
      nodeArray.push(head);
    }
    evenFlag = !evenFlag
    head = head.next;
  }
  let i = nodeArray.length;
  if (i == 0)
    return oldHead;
  while (i-- > 0) {
    if (i == nodeArray.length - 1) {
      if (nodeArray[i].next) {
        const leftNode = nodeArray[i];
        const rightNode = leftNode.next;
        rightNode.next = leftNode;
        leftNode.next = null;
        nodeArray[i] = rightNode
      }
    } else {
      const leftNode = nodeArray[i];
      const rightNode = leftNode.next;
      rightNode.next = leftNode;
      leftNode.next = nodeArray[i + 1];
      nodeArray[i] = rightNode
    }
  }
  return nodeArray[0];
};