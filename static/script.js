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

const getOptions = () => {
    // 获取所选选项
    return {
        lowercase: lowercase.checked,
        uppercase: uppercase.checked,
        numbers: numbers.checked,
        special: special.checked,
        customSpecial: customSpecial.value,
        exclude: exclude.checked,
        customExclude: customExclude.value,
        length: length.value,
        quantity: quantity.value,
    };
};

if (generateBtn) {
    // 为生成按钮添加事件监听器
    generateBtn.addEventListener("click", () => {
        const options = getOptions();

        // 向服务器发送带有所选选项的POST请求
        fetch("/generate", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify(options),
            })

            .then((response) => response.json())
            .then(data => console.log(data))
            .then((data) => {
                // 显示生成的密码
                const passwords = data.passwords;
                passwords.forEach((password) => {
                    console.log(password);
                });
            })
            .catch((error) => console.error(error));
    });
}