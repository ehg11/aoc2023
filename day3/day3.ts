const fs = require('fs');
const data = fs.readFileSync('d3', 'utf8').trim().split('\n');

function isNum(char: string): boolean {
    return /^[0-9]$/.test(char);
}

function getPartIndex(data: string[]): [number, number][] {
    const partIndices: [number, number][] = [];
    for (let row: number = 0; row < data.length; row++) {
        let line = data[row];
        for (let col: number = 0; col < line.length; col++) {
            let char: string = line[col]; 
            if (isNum(char) || char == '.') {
                continue;
            }
            partIndices.push([row, col]) ;
        }
    }

    return partIndices;
}

const partIndices = getPartIndex(data);

function getAdjacentPartIndex(rows: number, cols: number, partIndices: [number, number][]): [number, number][] {
    const adjacentPartIndices: [number, number][] = [];
    for (const [partRow, partCol] of partIndices) {
        const allAdjacentPartIndices = [
            [partRow - 1, partCol - 1], [partRow - 1, partCol], [partRow - 1, partCol + 1],
            [partRow, partCol - 1], [partRow, partCol + 1],
            [partRow + 1, partCol - 1], [partRow + 1, partCol], [partRow + 1, partCol + 1]
        ];

        for (const [adjRow, adjCol] of allAdjacentPartIndices) {
            if (adjRow < 0 || adjRow >= rows) {
                continue;
            }
            if (adjCol < 0 || adjCol >= cols) {
                continue;
            }
            adjacentPartIndices.push([adjRow, adjCol]);
        }
    }

    return adjacentPartIndices;
}

function getNumAtIndex(data: string[], adjacentIndex: [number, number]) {
    const [row, col] = adjacentIndex;
    const line = [...data[row]];

    let numberStart = col;
    let numberEnd = col + 1;

    if (!isNum(line[numberStart])) {
        return -1;
    }

    while (numberStart >= 0) {
        if (isNum(line[numberStart - 1])) {
            numberStart--;
        } else {
            break;
        }
    }

    while (numberEnd < line.length) {
        if (isNum(line[numberEnd])) {
            numberEnd++;
        } else {
            break;
        }
    }

    let num = 0;
    for (let i = numberStart; i < numberEnd; i++) {
        num *= 10;
        num += parseInt(line[i]);
        line[i] = '.';
    }

    const lineStr = line.join('');
    data[row] = lineStr;

    return num;
}

// const adjacentPartIndices = getAdjacentPartIndex(data.length, data[0].length, partIndices);
// const nums: number[] = []
// for (const adjPartIndex of adjacentPartIndices) {
//     const maybeNum = getNumAtIndex(data, adjPartIndex);
//     if (maybeNum != -1) {
//         nums.push((maybeNum));
//     }
// }
//
// console.log(nums.reduce((lhs, rhs) => lhs + rhs, 0));

function getGearIndex(data: string[]): [number, number][] {
    const gearIndices: [number, number][] = [];
    for (let i = 0; i < data.length; i++) {
        const line = data[i];
        for (let j = 0; j < line.length; j++) {
            if (line[j] == '*') {
                gearIndices.push([i, j]);
            }
        }
    }
    return gearIndices;
}

function getNumberAt(data: string[], row: number, col: number): [number, [number, number]] {
    const line = data[row];
    let numStart = col;
    let numEnd = col + 1;

    while (numStart > 0 && isNum(line[numStart - 1])) {
        numStart--;
    }
    while (numEnd <= line.length && isNum(line[numEnd])) {
        numEnd++;
    }

    let num = 0;
    for (let i = numStart; i < numEnd; i++) {
        num *= 10;
        num += parseInt(line[i]);
    }

    return [num, [row, numStart]];
}

function samePos(lhs: [number, number], rhs: [number, number]) {
    const [lhsRow, lhsCol] = lhs;
    const [rhsRow, rhsCol] = rhs;
    
    return lhsRow == rhsRow && lhsCol == rhsCol;
}

function hasPos(allPos: [number, number][], newPos: [number, number]) {
    for (const pos of allPos) {
        if (samePos(pos, newPos)) {
            return true;
        }
    }
    return false;
}

const gearIndices = getGearIndex(data);
const rows = data.length;
const cols = data[0].length;
let totalGearRatio = 0;

for (const [gearRow, gearCol] of gearIndices) {
    const gearAdjIndices: [number, number][] = [];

    const allAdjacentPartIndices = [
        [gearRow - 1, gearCol - 1], [gearRow - 1, gearCol], [gearRow - 1, gearCol + 1],
        [gearRow, gearCol - 1], [gearRow, gearCol + 1],
        [gearRow + 1, gearCol - 1], [gearRow + 1, gearCol], [gearRow + 1, gearCol + 1]
    ];

    for (const [adjRow, adjCol] of allAdjacentPartIndices) {
        if (adjRow < 0 || adjRow >= rows) {
            continue;
        }
        if (adjCol < 0 || adjCol >= cols) {
            continue;
        }
        gearAdjIndices.push([adjRow, adjCol]);
    }

    const gearAdjNumbers: number[] = [];
    const numStarts: [number, number][] = [];
    for (const [adjRow, adjCol] of gearAdjIndices) {
        if (isNum(data[adjRow][adjCol])) {
            const [num, pos] = getNumberAt(data, adjRow, adjCol);
            if (hasPos(numStarts, pos)) {
                continue;
            }
            numStarts.push(pos);
            gearAdjNumbers.push(num);
        }
    }

    if (gearAdjNumbers.length == 2) {
        totalGearRatio += (gearAdjNumbers[0] * gearAdjNumbers[1]);
    }
}

console.log(totalGearRatio);
