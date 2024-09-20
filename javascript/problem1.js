const sortCharacters = (s) => {
  const vowels = "aeiou";
  let vowelChars = "";
  let consonantChars = "";
  s = s.toLowerCase();
  for (let char of s) {
    if (char === " ") {
      continue;
    }
    if (vowels.includes(char)) {
      vowelChars += char;
    } else if (char >= "a" && char <= "z") {
      consonantChars += char;
    }
  }
  return { vowelChars, consonantChars };
};

const input = "Sample Case";
const { vowelChars, consonantChars } = sortCharacters(input);
console.log("Vowel Characters:", vowelChars);
console.log("Consonant Characters:", consonantChars);
