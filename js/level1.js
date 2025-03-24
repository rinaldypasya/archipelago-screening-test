function titleCase(str) {
    return str
        .toLowerCase()
        .split(" ")
        .map(function (word) { return word.charAt(0).toUpperCase() + word.slice(1); })
        .join(" ");
}
// Test cases
console.log(titleCase("I'm a little tea pot"));
console.log(titleCase("sHoRt AnD sToUt"));
console.log(titleCase("SHORT AND STOUT"));
