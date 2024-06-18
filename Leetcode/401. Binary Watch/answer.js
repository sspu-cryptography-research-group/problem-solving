/**
 * @param {number} turnedOn
 * @return {string[]}
 */

let binaryArray = [];

var readBinaryWatch = function (turnedOn) {
  const fullLedCount = 10;
  const oneCount = turnedOn;
  const result = [];
  generateBinaryArray("", oneCount, 10);
  const watchMapper = [1, 2, 4, 8, 1, 2, 4, 8, 16, 32,];
  for (const binary of binaryArray) {
    let hour = 0;
    let minute = 0;
    for (const num in binary) {
      const currentResult = parseInt(binary[num]) * watchMapper[num];
      if (num < 4)
        hour += currentResult;
      else
        minute += currentResult;
    }
    if (hour > 11 || minute > 59)
      continue;
    if (minute < 10)
      result.push(hour + ":0" + minute);
    else {
      result.push(hour + ":" + minute);
    }
  }
  binaryArray = [];
  return result;
};

const generateBinaryArray = (binaryText, oneCount, target) => {
  if (binaryText.length === target && oneCount === 0) {
    binaryArray.push(binaryText);
    return;
  }

  if (binaryText.length < target) {
    generateBinaryArray(binaryText + "0", oneCount, target);
    generateBinaryArray(binaryText + "1", oneCount - 1, target);
  }
}