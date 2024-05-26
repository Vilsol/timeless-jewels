/**
 * Break an array into chunks of the given size
 *
 * e.g. for chunk size 2:
 *
 * [1, 2, 3, 4, 5] => [[1, 2], [3, 4], [5]]
 */
export const chunkArray = <T>(inputArray: Array<T>, chunkSize: number): Array<T>[] =>
  inputArray.reduce((resultArray, item, index) => {
    const chunkIndex = Math.floor(index / chunkSize);

    if (!resultArray[chunkIndex]) {
      resultArray[chunkIndex] = []; // start a new chunk
    }

    resultArray[chunkIndex].push(item);

    return resultArray;
  }, [] as Array<T>[]);
