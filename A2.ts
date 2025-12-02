/**
 * @author Marco Cantusci
 * @verion 1.0.0
 * @date 2025-12-02
 * @fileoverview Right angle triangle made out of numbers.
 */

// input
const rowsString: string = prompt("How many rows would you like? ") || ('\n');
const rowsNumber: number = parseInt(rowsString);

// process
for (let row = 1; row <= rowsNumber; row++) {
  for (let numbers = 1; numbers <= row; numbers++) {
    console.log(numbers);
  }
  console.log("\n");
};

console.log("\nDone.");
