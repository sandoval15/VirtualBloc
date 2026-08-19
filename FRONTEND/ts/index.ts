function bubbleSort(arr: number[]): number[] {
  // Sacamos una copia para no mutar el arreglo original
    const array = [...arr];
    const length = array.length;

    for (let i = 0; i < length - 1; i++) {
        for (let j = 0; j < length - 1 - i; j++) {
      // Comparación de elementos adyacentes
            if (array[j] > array[j + 1]) {
        // Intercambio usando desestructuración
                [array[j], array[j + 1]] = [array[j + 1], array[j]];
            }
        }
    }

    return array;
}

// Ejemplo de uso:
const numerosDesordenados: number[] = [64, 34, 25, 12, 22, 11, 90];
const numerosOrdenados = bubbleSort(numerosDesordenados);

console.log("Original hola:", numerosDesordenados);
console.log("BBBBBBBBBBBBBBB:", numerosOrdenados);