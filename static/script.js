// 从HTML中获取所有必要的元素
const lowercase = document.getElementById("lowercase"); // 小写字母复选框
const uppercase = document.getElementById("uppercase"); // 大写字母复选框
const numbers = document.getElementById("numbers"); // 数字复选框
const special = document.getElementById("special"); // 特殊字符复选框
const customSpecial = document.getElementById("custom-special"); // 自定义字符输入框
const exclude = document.getElementById("exclude"); // 排除字符复选框
const customExclude = document.getElementById("custom-exclude"); // 自定义排除字符输入框
const length = document.getElementById("length"); // 密码长度输入框
const quantity = document.getElementById("quantity"); // 生成数量输入框
const generateBtn = document.getElementById("generate-button"); // 生成按钮

if (generateBtn) {
    // 为生成按钮添加事件监听器
    generateBtn.addEventListener("click", () => {
        // 获取所选选项
        const options = {
            lowercase: lowercase.checked, // 是否包含小写字母
            uppercase: uppercase.checked, // 是否包含大写字母
            numbers: numbers.checked, // 是否包含数字
            special: special.checked, // 是否包含特殊字符
            customSpecial: customSpecial.value, // 自定义字符
            exclude: exclude.checked, // 是否排除字符
            customExclude: customExclude.value, // 自定义排除字符
            length: length.value, // 密码长度
            quantity: quantity.value, // 生成数量
        };

        console.log(options);
        console.log("按钮被点击了！");

        // 向服务器发送带有所选选项的POST请求
        fetch("/generate", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify(options),
            })

            .then((response) => response.json())
            .then((data) => {
                // 显示生成的密码
                const passwords = data.passwords;
                const passwordList = document.getElementById("password-list");
                passwordList.innerHTML = "";
                passwords.forEach((password) => {
                    const li = document.createElement("li");
                    li.textContent = password;
                    passwordList.appendChild(li);
                });
            })
            .catch((error) => console.error(error));
    });
}