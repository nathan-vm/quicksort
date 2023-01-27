const fs = require('fs');

function quicksort(arr) {
  console.time('Execution time');
  let left = 0;
  let right = arr.length - 1;

  const stack = [[left, right]];

  while (stack.length > 0) {
    const partition = stack.pop();

    const pivotIndex = partition[0];
    const pivotValue = arr[pivotIndex];

    let i = partition[0];
    for (let j = partition[0]; j <= partition[1]; j++) {
      if (arr[j] < pivotValue) {
        i++;
        [arr[i], arr[j]] = [arr[j], arr[i]];
      }
    }
    [arr[i], arr[pivotIndex]] = [arr[pivotIndex], arr[i]];

    if (partition[0] < i - 1) {
      stack.push([partition[0], i - 1]);
    }
    if (i + 1 < partition[1]) {
      stack.push([i + 1, partition[1]]);
    }
  }
  console.timeEnd('Execution time');
  return arr;
}

fs.readFile('input.json', (err, data) => {
  if (err) throw err;
  const arr = JSON.parse(data);
  const sortedArr = quicksort(arr);
});