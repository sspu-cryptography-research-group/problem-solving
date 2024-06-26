/**
 * @param {string[]} strs
 * @return {string}
 */
var longestCommonPrefix = function (strs) {
  let commonPrefix = "";
  for (let i = 0; true; i++) {
    const currentChar = strs[0][i];
    for (let str of strs) {
      if (str[i]) {
        if (currentChar == str[i]) {
          continue;
        } else {
          return commonPrefix;
        }
      }
      else {
        return commonPrefix;
      }
    }
    commonPrefix += currentChar;
  }
};

console.log(longestCommonPrefix(["flower", "flower", "flower"]));