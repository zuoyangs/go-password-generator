<template>
  <div class="password-table">
    <table>
      <thead>
        <tr>
          <th width="15%" bgcolor="#F5F5F5" class="text-center">序号</th>
          <th width="45%" bgcolor="#FFFFFF" class="text-center">密码</th>
          <th width="40%" bgcolor="#F5F5F5" class="text-center">强度评估</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="(password, index) in passwords" :key="index">
          <td>{{ index + 1 }}</td>
          <td>
            {{ password }}
            <span v-if="copiedIndex === index" class="copy-tip">已复制</span>
            <button class="copy-btn" @click="copyPassword(password, index)">复制</button>
          </td>
          <td>
            <div class="strength-meter">
              <div 
                class="strength-bar" 
                :style="{ width: `${getStrength(password).score * 25}%` }"
                :class="getStrengthClass(password)"
              ></div>
            </div>
            <div class="strength-feedback">{{ getStrength(password).feedback }}</div>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script setup lang="ts">
import { checkPasswordStrength } from '../utils/password';
import { ref } from 'vue';

defineProps<{
  passwords: string[]
}>();

const copiedIndex = ref<number | null>(null);

const getStrength = (password: string) => {
  return checkPasswordStrength(password);
}

const getStrengthClass = (password: string) => {
  const score = getStrength(password).score;
  return {
    'strength-weak': score <= 1,
    'strength-medium': score === 2,
    'strength-strong': score === 3,
    'strength-very-strong': score === 4
  };
};

const copyPassword = async (password: string, index: number) => {
  try {
    await navigator.clipboard.writeText(password);
    copiedIndex.value = index;
    setTimeout(() => {
      copiedIndex.value = null;
    }, 2000);
  } catch (err) {
    console.error('复制失败:', err);
  }
};
</script>

<style scoped>
.password-table {
  width: 100%;
  margin: 20px auto;
}

table {
  width: 100%;
  border-collapse: collapse;
  margin: 20px 0;
}

th, td {
  padding: 10px;
  border: 1px solid var(--border-color);
  text-align: center;
  vertical-align: middle;
  background: rgba(255, 255, 255, 0.03);
  color: #333333;
}

.copy-btn {
  position: absolute;
  right: 15px;
  top: 50%;
  transform: translateY(-50%);
  padding: 6px 15px;
  background: linear-gradient(45deg, #FF4081, #FF6B6B);
  color: white;
  border: none;
  border-radius: 12px;
  cursor: pointer;
  transition: all 0.3s ease;
  min-width: 70px;
  box-shadow: var(--box-shadow);
  font-weight: bold;
}

.copy-btn:hover {
  background: linear-gradient(45deg, #FF6B6B, #FF4081);
  box-shadow: 0 4px 8px rgba(255, 64, 129, 0.3);
}

.strength-meter {
  width: 100%;
  height: 4px;
  background-color: #eee;
  border-radius: 2px;
  margin-bottom: 5px;
  position: relative;
  overflow: hidden;
}

.strength-bar {
  height: 100%;
  border-radius: 2px;
  transition: width 0.3s ease;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  position: absolute;
  left: 0;
  top: 0;
}

.strength-weak {
  background-color: #FF6B6B;
}

.strength-medium {
  background-color: #FFD93D;
}

.strength-strong {
  background-color: #6BCB77;
}

.strength-very-strong {
  background-color: #4D96FF;
}

.strength-feedback {
  font-size: 12px;
  color: #333333;
  text-align: left;
}

td:nth-child(2) {
  font-family: monospace;
  font-size: 16px;
  position: relative;
  text-align: center;
  padding: 10px 120px 10px 15px;
}

th:nth-child(1), td:nth-child(1) {
  width: 15%;
  white-space: nowrap;
}

th:nth-child(2), td:nth-child(2) {
  width: 45%;
  text-align: left;
}

th:nth-child(3), td:nth-child(3) {
  width: 40%;
}

tr {
  transition: all 0.3s ease;
}

tr:hover {
  background-color: rgba(255, 97, 216, 0.05);
}

td:first-child {
  font-weight: 500;
  color: var(--secondary-color);
}

.copy-tip {
  position: absolute;
  right: 95px;
  top: 50%;
  transform: translateY(-50%);
  color: #FF9800;
  font-size: 14px;
  font-weight: bold;
  animation: fadeInOut 2s ease;
}

@keyframes fadeInOut {
  0% { 
    opacity: 0;
    transform: translateY(-5px);
  }
  20% {
    opacity: 1;
    transform: translateY(0);
  }
  80% {
    opacity: 1;
    transform: translateY(0);
  }
  100% {
    opacity: 0;
    transform: translateY(5px);
  }
}

.strength-indicator {
  padding: 4px 8px;
  border-radius: 4px;
  font-size: 0.9em;
  font-weight: 500;
  text-align: center;
  width: fit-content;
  margin: 0 auto;
}

.strength-excellent {
  background-color: #E8F5E9;
  color: #2E7D32;
  border: 1px solid #81C784;
}

.strength-good {
  background-color: #E3F2FD;
  color: #1565C0;
  border: 1px solid #64B5F6;
}

.strength-fair {
  background-color: #FFF3E0;
  color: #E65100;
  border: 1px solid #FFB74D;
}

.strength-weak {
  background-color: #FFEBEE;
  color: #C62828;
  border: 1px solid #E57373;
}

th:nth-child(1) { width: 10%; }
th:nth-child(2) { width: 50%; }
th:nth-child(3) { width: 20%; }
th:nth-child(4) { width: 20%; }

.text-center {
  text-align: center !important;
  vertical-align: middle !important;
}
</style> 