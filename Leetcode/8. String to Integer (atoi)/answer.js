/**
 * @param {string} s
 * @return {number}
 */
var myAtoi = function (s) {
  s = s.trim();
  let sign = false;
  let negtiveFlag = 1;
  let result = 0;
  const maxInt = 2147483647;
  const minInt = -2147483648;
  if (s[0] === '+') {
    sign = true;
  }
  if (s[0] === '-') {
    sign = true;
    negtiveFlag = -1;
  }
  for (const i in s) {
    if (sign && i == 0)
      continue;
    if (!isNaN(parseInt(s[i]))) {
      result *= 10;
      result += parseInt(s[i])
    } else {
      break;
    }
  }
  result *= negtiveFlag;
  if ( result > maxInt) 
    return maxInt;
  if ( result < minInt)
    return minInt;
  return result;
};

console.log(myAtoi(" -042"));
console.log(myAtoi("42"));
console.log(myAtoi("1337c0d3"));
console.log(myAtoi("0-1"));
console.log(myAtoi("words and 987"));