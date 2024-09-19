const minBuses = (n, members) => {
    if (members.length !== n) {
        console.log("Input must be equal with count of family");
        return;
    }

    members.sort((a, b) => b - a); 
    let buses = 0;
    let i = 0, j = members.length - 1;

    while (i <= j) {
        if (members[i] + members[j] <= 4) {
            j--;
        }
        i++;
        buses++;
    }

    console.log("Minimum bus required is:", buses);
}

const n = 5;
const members = [1, 2, 4, 3, 3];
minBuses(n, members);
