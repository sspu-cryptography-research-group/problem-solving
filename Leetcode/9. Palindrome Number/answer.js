/**
 * @param {number} x
 * @return {boolean}
 */
var isPalindrome = function (x) {
  if (x < 0)
    return false;
  const numberArray = [];
  while (x > 0) {
    numberArray.push(x % 10);
    x /= 10;
  }
  console.log(numberArray);
};

console.log(isPalindrome(121));