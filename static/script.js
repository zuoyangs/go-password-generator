const lowercase = document.getElementById("lowercase");
const uppercase = document.getElementById("uppercase");
const numbers = document.getElementById("numbers");
const special = document.getElementById("special");
const customSpecial = document.getElementById("custom-special");
const exclude = document.getElementById("exclude");
const customExclude = document.getElementById("custom-exclude");
const length = document.getElementById("length");
const quantity = document.getElementById("quantity");
const passwordsContainer = document.getElementById("passwords-container");

const getOptions = () => {
    return {
        "lowercase": true,
        "uppercase": true,
        "numbers": true,
        "special": false,
        "customSpecial": "",
        "exclude": false,
        "customExclude": "",
        "length": 8,
        "quantity": 2
    }
};

window.onload = function () {
    const generateBtn = document.getElementById("generate-button");

    generateBtn.addEventListener("click", async () => {
        const options = getOptions();

        try {
            const response = await fetch("http://localhost:3005/passwordGenerator", {
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
            passwordsContainer.innerHTML = "";

            // Generate new passwords
            result.passwords.forEach((password) => {
                const passwordElement = document.createElement("p");
                passwordElement.innerText = password;
                passwordsContainer.appendChild(passwordElement);
                passwordsContainer.appendChild(document.createElement('br'));
            });
        } catch (error) {
            console.log("error", error);
        }
    });
};
