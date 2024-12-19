<template>
  <div class="app">
    <h1>随机密码生成器</h1>
    <PasswordForm @generate="(options) => handleGenerate(options)" />
    <PasswordTable :passwords="passwords" />
    <div class="footer">
      <img src="/images/ghs.png" style="vertical-align: middle; margin-right: 5px;">
      <a href="https://beian.miit.gov.cn/" target="_blank">沪ICP备2023037106号-1</a>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import PasswordForm from './components/PasswordForm.vue'
import PasswordTable from './components/PasswordTable.vue'
import type { PasswordOptions, PasswordResponse } from './types/password'

const passwords = ref<string[]>([])

const handleGenerate = async (options: PasswordOptions): Promise<void> => {
  try {
    const response = await fetch('/api/generate', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(options),
    })
    const data: PasswordResponse = await response.json()
    if (data.error) {
      alert(data.error)
      return
    }
    passwords.value = data.passwords
  } catch (err) {
    console.error('生成密码失败，请检查后端服务是否正常运行:', err)
    alert('生成密码失败，请稍后重试')
  }
}
</script>

<style>
:root {
  --primary-color: #FFA5BA;     /* 柔和粉 */
  --secondary-color: #A5D8FF;   /* 淡天蓝 */
  --accent-color: #FFE5A5;      /* 柔和黄 */
  --tertiary-color: #A5E6CC;    /* 薄荷绿 */
  --purple-color: #E5C1FF;      /* 淡紫色 */
  --pink-color: #FFD1DC;        /* 浅粉色 */
  --table-header-bg: linear-gradient(45deg, #A5D8FF, #E5C1FF);
  --button-gradient: linear-gradient(45deg, #FFA5BA, #A5E6CC);
  --bg-color: #FFFFFF;
  --border-color: #FFD1DC;
  --text-color: #555555;
  --heading-color: #FFA5BA;
  --bg-soft: #F0F7FF;
  --box-shadow: 0 8px 32px rgba(255, 165, 186, 0.15);
  --gradient-soft: linear-gradient(45deg, #FFA5BA, #A5D8FF, #FFE5A5);
  --gradient-gentle: linear-gradient(90deg, #FFD1DC, #E5C1FF, #A5E6CC);
}

body {
  margin: 0;
  padding: 0;
  min-height: 100vh;
  display: flex;
  justify-content: center;
  align-items: flex-start;
  background: var(--bg-soft);
  background-image: 
    radial-gradient(circle at 20% 20%, rgba(255, 165, 186, 0.1) 0%, transparent 50%),
    radial-gradient(circle at 80% 80%, rgba(165, 216, 255, 0.1) 0%, transparent 50%),
    radial-gradient(circle at 50% 50%, rgba(229, 193, 255, 0.1) 0%, transparent 60%);
}

.app {
  max-width: 840px;
  min-width: 560px;
  margin: 2vh auto;
  padding: 1.8vh 30px;
  background: rgba(255, 255, 255, 0.95);
  box-shadow: 0 8px 32px rgba(255, 165, 186, 0.2);
  border-radius: 30px;
  border: 3px solid;
  border-image: var(--gradient-gentle) 1;
  transform: translateY(0);
  transition: all 0.3s cubic-bezier(0.34, 1.56, 0.64, 1);

}

.app:hover {
  transform: translateY(-8px);
  box-shadow: 0 16px 48px rgba(255, 165, 186, 0.25);
}

h1 {
  text-align: center;
  font-size: 2.4em;
  background: var(--gradient-soft);
  background-clip: text;
  -webkit-background-clip: text;
  color: transparent;
  margin: 0.75em 0;
  text-shadow: 
    2px 2px 0px rgba(255, 165, 186, 0.2),
    4px 4px 0px rgba(165, 216, 255, 0.2);
  animation: bounce 2s ease-in-out infinite;
}

@keyframes bounce {
  0%, 100% { 
    transform: translateY(0) rotate(-2deg);
  }
  50% { 
    transform: translateY(-12px) rotate(2deg) scale(1.05);
  }
}

.footer {
  text-align: center;
  margin-top: 30px;
  color: var(--primary-color);
  padding: 20px 0;
  opacity: 0.8;
}

.footer a {
  color: var(--primary-color);
  text-decoration: none;
  opacity: 0.8;
  transition: opacity 0.3s ease;
}

.footer a:hover {
  opacity: 1;
}
</style>
