const lowercase = document.getElementById("lowercase");
const uppercase = document.getElementById("uppercase");
const numbers = document.getElementById("numbers");
const special = document.getElementById("special");
const customSpecial = document.getElementById("custom-special");
const exclude = document.getElementById("exclude");
const customExclude = document.getElementById("custom-exclude");
const length = document.getElementById("length");
const quantity = document.getElementById("quantity");
const passwordsTable = document.getElementById("passwordsTable");

const getOptions = () => {
    return {
        "lowercase": lowercase.checked,
        "uppercase": uppercase.checked,
        "numbers": numbers.checked,
        "special": special.checked,
        "customSpecial": customSpecial.value,
        "exclude": exclude.checked,
        "customExclude": customExclude.value,
        "length": parseInt(length.value),
        "quantity": parseInt(quantity.value),
    }
};

window.onload = function() {
    const generateBtn = document.getElementById("generate-button");

    generateBtn.addEventListener("click", async () => {
        const options = getOptions();

        try {
            const response = await fetch("https://password.zuoyang.tech/passwordGenerator", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    "Accept": "*/*",
                    "Connection": "keep-alive",
                },
                body: JSON.stringify(options),
            });

            const result = await response.json();

            // Clear previous passwords
            passwordsTable.querySelector("tbody").innerHTML = "";

            // Generate new passwords
            result.passwords.forEach((password, index) => {
                const row = document.createElement("tr");

                const indexCell = document.createElement("td");
                indexCell.innerText = index + 1;

                const passwordCell = document.createElement("td");
                passwordCell.innerText = password;

                row.appendChild(indexCell);
                row.appendChild(passwordCell);

                passwordsTable.querySelector("tbody").appendChild(row);
            });
        } catch (error) {
            console.log("error", error);
        }
    });
};
