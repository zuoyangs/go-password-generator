<template>
  <div class="password-form">
    <table>
      <tbody>
        <tr>
          <td width="25%" bgcolor="#F5F5F5">字符选项</td>
          <td width="75%" bgcolor="#FFFFFF">
            <div class="options-group">
              <label>
                <input type="checkbox" v-model="options.lowercase">
                小写字母 (a-z)
              </label>
              <label>
                <input type="checkbox" v-model="options.uppercase">
                大写字母 (A-Z)
              </label>
              <label>
                <input type="checkbox" v-model="options.numbers">
                数字 (0-9)
              </label>
            </div>
          </td>
        </tr>

        <tr>
          <td bgcolor="#F5F5F5">特殊字符</td>
          <td bgcolor="#FFFFFF">
            <div class="char-options">
              <div class="special-chars">
                <label>
                  <input type="checkbox" v-model="options.special">
                  特殊字符
                </label>
                <input
                  v-if="options.special"
                  type="text"
                  v-model="options.customSpecial"
                  placeholder="!@#$%^&*"
                >
              </div>
              <div class="exclude-chars">
                <label>
                  <input type="checkbox" v-model="options.exclude">
                  排除字符
                </label>
                <input
                  v-if="options.exclude"
                  type="text"
                  v-model="options.customExclude"
                  placeholder="iIl1o0O"
                >
              </div>
            </div>
          </td>
        </tr>

        <tr>
          <td bgcolor="#F5F5F5">密码长度</td>
          <td bgcolor="#FFFFFF">
            <div class="length-input">
              <div class="select-group">
                <select v-model="options.length">
                  <option v-for="opt in lengthOptions" :key="opt.value" :value="opt.value">
                    {{ opt.label }}
                  </option>
                </select>
                <span class="length-hint" :class="getLengthHintClass">
                  {{ getLengthHint }}
                </span>
              </div>
            </div>
          </td>
        </tr>

        <tr>
          <td bgcolor="#F5F5F5">生成数量</td>
          <td bgcolor="#FFFFFF">
            <div class="select-group">
              <select v-model="options.quantity">
                <option v-for="n in 20" :key="n" :value="n">{{ n }}</option>
              </select>
            </div>
          </td>
        </tr>
      </tbody>
    </table>

    <button @click="generatePasswords" class="generate-btn">生成密码</button>
  </div>
</template>

<script setup lang="ts">
import { reactive, computed } from 'vue'
import type { PasswordOptions } from '../types/password'

const emit = defineEmits<{
  (e: 'generate', options: PasswordOptions): void
}>()

// 定义密码长度选项
const lengthOptions = [
  { value: 4, label: '4 位密码 (最小长度)' },
  { value: 6, label: '6 位密码 (简单密码)' },
  { value: 8, label: '8 位密码 (基础强度)' },
  { value: 10, label: '10 位密码 (一般强度)' },
  { value: 12, label: '12 位密码 (推荐长度)' },
  { value: 14, label: '14 位密码 (良好强度)' },
  { value: 16, label: '16 位密码 (较高强度)' },
  { value: 18, label: '18 位密码 (高强度)' },
  { value: 20, label: '20 位密码 (很高强度)' },
  { value: 24, label: '24 位密码 (极高强度)' },
  { value: 28, label: '28 位密码 (超高强度)' },
  { value: 32, label: '32 位密码 (超高强度)' },
  { value: 36, label: '36 位密码 (超强防护)' },
  { value: 40, label: '40 位密码 (超强防护)' },
  { value: 45, label: '45 位密码 (超强防护)' },
  { value: 50, label: '50 位密码 (超强防护)' },
  { value: 64, label: '64 位密码 (极限强度)' },
  { value: 100, label: '100 位密码 (最大长度)' }
]

const options = reactive<PasswordOptions>({
  lowercase: true,
  uppercase: true,
  numbers: true,
  special: true,
  customSpecial: '',
  exclude: false,
  customExclude: '',
  length: 16,  // 默认16位
  quantity: 10
})

const generatePasswords = () => {
  console.log('Sending options:', options)
  emit('generate', options)
}

const getLengthHint = computed(() => {
  const length = options.length;
  if (length < 4) {
    return '密码长度太短（最少4位）';
  } else if (length > 100) {
    return '密码长度太长（最多100位）';
  } else if (length < 8) {
    return '密码长度适中';
  } else if (length < 12) {
    return '密码长度良好';
  } else {
    return '密码长度很安全';
  }
});

const getLengthHintClass = computed(() => {
  const length = options.length;
  return {
    'hint-error': length < 4 || length > 100,
    'hint-warning': length >= 4 && length < 8,
    'hint-good': length >= 8 && length < 12,
    'hint-excellent': length >= 12
  };
});
</script>

<style scoped>
.password-form {
  width: 100%;
  margin: 0 auto;
}

table {
  width: 100%;
  border-collapse: separate;
  border-spacing: 0 8px;
  margin-bottom: 40px;
}

td:first-child {
  width: 15%;
  background: linear-gradient(135deg, #6366F1, #8B5CF6);
  color: white;
  font-weight: 500;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.1);
  border-radius: 8px 0 0 8px;
}

td:last-child {
  width: 85%;
  background: #F8FAFC;
  border-radius: 0 8px 8px 0;
}

td {
  padding: 12px;
  border: none;
  margin-bottom: 8px;
  vertical-align: middle;
  color: #333333;
  transition: all 0.3s ease;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.05);
  min-height: 27px;
}

.options-group {
  display: flex;
  align-items: center;
  gap: 40px;
  padding: 8px;
  margin: -5px 0;
}

.options-group label {
  min-width: 150px;
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 4px 0;
}

.options-group input[type="checkbox"] {
  margin: 0;
}

.generate-btn {
  padding: 12px 35px;
  background: linear-gradient(45deg, #4CAF50, #45a049);
  color: white;
  border: none;
  border-radius: 25px;
  cursor: pointer;
  transition: all 0.3s ease;
  font-size: 1.3em;
  letter-spacing: 3px;
  font-weight: bold;
  box-shadow: var(--box-shadow);
  transform: rotate(-2deg);
}

.generate-btn:hover {
  transform: rotate(2deg) scale(1.05);
  box-shadow: 0 12px 32px rgba(76, 175, 80, 0.3);
}

.hint {
  color: #6B7280;
  font-size: 14px;
  margin-left: 12px;
}

.char-options {
  display: flex;
  gap: 20px;
}

.special-chars, .exclude-chars {
  display: flex;
  align-items: center;
  gap: 10px;
}

input[type="text"], select {
  padding: 8px 12px;
  border: 2px solid var(--border-color);
  border-radius: 12px;
  background: white;
  color: var(--text-color);
  transition: all 0.3s ease;
  box-shadow: var(--box-shadow);
}

input[type="text"]:focus, select:focus {
  border-color: var(--primary-color);
  box-shadow: 0 0 15px rgba(255, 153, 153, 0.3);
  outline: none;
}

input[type="checkbox"] {
  accent-color: var(--primary-color);
  width: 18px;
  height: 18px;
  border-radius: 50%;
}

.select-group {
  display: flex;
  align-items: center;
  gap: 15px;
}

select {
  padding: 8px 12px;
  border: 1px solid #E2E8F0;
  border-radius: 8px;
  background: white;
  min-width: 200px;
  font-size: 14px;
  color: #333;
  cursor: pointer;
  transition: all 0.3s ease;
}

select:hover {
  border-color: #6366F1;
}

select:focus {
  outline: none;
  border-color: #6366F1;
  box-shadow: 0 0 0 3px rgba(99, 102, 241, 0.1);
}

.options-group, .char-options {
  background: white;
  padding: 10px;
  border-radius: 8px;
  box-shadow: inset 0 2px 4px rgba(0, 0, 0, 0.05);
  margin: 0;
}

@keyframes gradientMove {
  from {
    background-position: 0% 50%;
  }
  to {
    background-position: 100% 50%;
  }
}

@keyframes buttonGlow {
  from {
    filter: brightness(1) hue-rotate(0deg);
  }
  to {
    filter: brightness(1.2) hue-rotate(30deg);
  }
}

.length-input {
  display: flex;
  align-items: center;
  gap: 1rem;
}

.length-hint {
  color: #666;
  font-size: 14px;
}

.hint-error {
  background-color: #FFE5E5;
  color: #D32F2F;
}

.hint-warning {
  background-color: #FFF3E0;
  color: #E65100;
}

.hint-good {
  background-color: #E8F5E9;
  color: #2E7D32;
}

.hint-excellent {
  background-color: #E3F2FD;
  color: #1565C0;
}
</style> 