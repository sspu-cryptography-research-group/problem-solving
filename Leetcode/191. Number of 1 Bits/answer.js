/**
 * @param {number} n
 * @return {number}
 */
var hammingWeight = function (n) {
  let counter = 0;
  while (n > 0) {
    const bit = n & 1;
    if (bit)
      counter++
    n = n >> 1;
  }
  return counter;
};

console.log(hammingWeight(2147483645));