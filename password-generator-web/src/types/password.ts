export interface PasswordOptions {
  lowercase: boolean
  uppercase: boolean
  numbers: boolean
  special: boolean
  customSpecial: string
  exclude: boolean
  customExclude: string
  length: number
  quantity: number
}

export interface PasswordResponse {
  passwords: string[]
  error?: string
} 