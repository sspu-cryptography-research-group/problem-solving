/**
 * Definition for singly-linked list.
 * function ListNode(val, next) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.next = (next===undefined ? null : next)
 * }
 */
/**
 * @param {ListNode} list1
 * @param {ListNode} list2
 * @return {ListNode}
 */
var mergeTwoLists = function (list1, list2) {
  let currentList1 = list1;
  let currentList2 = list2;
  let resultCurrent = null;
  let resultHead = null;
  while (currentList1 && currentList2) {
    if (currentList1.val < currentList2.val) {
      if (resultCurrent) {
        resultCurrent.next = currentList1;
        resultCurrent = currentList1;
      } else {
        resultCurrent = currentList1;
        resultHead = currentList1;
      }
      currentList1 = currentList1.next;
    } else {
      if (resultCurrent) {
        resultCurrent.next = currentList2;
        resultCurrent = currentList2;
      } else {
        resultCurrent = currentList2;
        resultHead = currentList2;
      }
      currentList2 = currentList2.next;
    }
  }
  const remainList = currentList1 ? currentList1 : currentList2;
  if (remainList) {
    if (resultCurrent)
      resultCurrent.next = remainList;
    else
      resultHead = remainList;
  }
  return resultHead;
};