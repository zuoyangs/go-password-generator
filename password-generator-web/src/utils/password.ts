export interface PasswordStrength {
  score: number;  // 0-4
  feedback: string;
}

export function checkPasswordStrength(password: string): PasswordStrength {
  let score = 0;
  const feedback: string[] = [];

  // 长度检查
  if (password.length < 8) {
    feedback.push("密码长度太短");
  } else if (password.length >= 12) {
    score += 1;
    feedback.push("密码长度合适");
  }

  // 包含小写字母
  if (/[a-z]/.test(password)) {
    score += 1;
  } else {
    feedback.push("建议包含小写字母");
  }

  // 包含大写字母
  if (/[A-Z]/.test(password)) {
    score += 1;
  } else {
    feedback.push("建议包含大写字母");
  }

  // 包含数字
  if (/\d/.test(password)) {
    score += 1;
  } else {
    feedback.push("建议包含数字");
  }

  // 包含特殊字符
  if (/[!@#$%^&*(),.?":{}|<>]/.test(password)) {
    score += 1;
  } else {
    feedback.push("建议包含特殊字符");
  }

  return {
    score,
    feedback: feedback.join("；") || "密码强度很好！"
  };
} 