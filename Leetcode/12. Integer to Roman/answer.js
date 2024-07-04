/**
 * @param {number} num
 * @return {string}
 */
var intToRoman = function (num) {
  const romanSymbol = ["I", "V", "X", "L", "C", "D", "M"];
  let epoch = 0;
  let romanNum = "";
  while (num > 0) {
    const currentNum = num % 10;
    const symbolOneIndex = epoch * 2;
    const symbolFiveIndex = epoch * 2 + 1;
    const symbolTenIndex = epoch * 2 + 2;
    num /= 10;
    num = parseInt(num);
    epoch++;
    if (currentNum < 4) {
      romanNum = romanSymbol[symbolOneIndex].repeat(currentNum) + romanNum;
    }
    else if (currentNum === 4) {
      romanNum = romanSymbol[symbolOneIndex] + romanSymbol[symbolFiveIndex] + romanNum;
    }
    else if (currentNum < 9) {
      romanNum = romanSymbol[symbolFiveIndex] + romanSymbol[symbolOneIndex].repeat(currentNum - 5) + romanNum;
    }
    else if (currentNum === 9) {
      romanNum = romanSymbol[symbolOneIndex] + romanSymbol[symbolTenIndex] + romanNum;
    }

  }
  return romanNum;
};



console.log(intToRoman(3749))
console.log(intToRoman(58))
console.log(intToRoman(1994))
console.log(intToRoman(3999))
console.log(intToRoman(900))