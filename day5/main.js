import { readFileSync } from "fs";

function main() {
    const inputName = process.argv[2] ?? 'input.txt'
    const input = readFileSync(inputName, 'utf-8')

    const [rulesStr, updatesStr] = input.split("\n\n")

    const existBefore = new Map();
    const existAfter = new Map();
    for (const ruleStr of rulesStr.split("\n")) {
        const [x, y] = ruleStr.split("|").map(Number)
        if (!existBefore[y]) {
            existBefore[y] = new Set();
        }
        existBefore[y].add(x)
        if (!existAfter[x]) {
            existAfter[x] = new Set();
        }
        existAfter[x].add(y)
    }


    let sum = 0;
    for (const updateStr of updatesStr.split("\n")) {
        const updates = updateStr.split(",").map(Number)

        let okay = true
        outerloop: for (let i = 0; i < updates.length; i++) {
            for (let j = i; j < updates.length; j++) {
                const currElement = updates[i];
                const checkElement = updates[j];

                if (
                    currElement in existBefore
                    && existBefore[currElement].has(checkElement)
                ) {
                    okay = false
                    break outerloop
                }
            }
        }

        if (okay) {
            sum += updates[Math.floor(updates.length / 2)]
        }
    }

    console.log(sum)
}

main();
