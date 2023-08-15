import { defineStore } from 'pinia'

export interface AuthState {
  accessToken: null | string
  accessTokenExpiration: null | Date
  lastCharID: number
  accountID: bigint
  loggingIn: boolean
  loginError: null | string
  permissions: String[]
}

export const useAuthStore = defineStore('auth', {
  state: () => ({} as AuthState),

  actions: {
    loginStart(): void {
      this.loggingIn = true
    },
    loginStop(errorMessage: null | string): void {
      this.loggingIn = false
      this.loginError = errorMessage
    },
    setAccessToken(accessToken: null | string, expiration: null | bigint | Date): void {
      this.accessToken = accessToken
      if (typeof expiration === 'bigint') expiration = new Date(expiration.toString())
      this.accessTokenExpiration = expiration
    },
    setPermissions(permissions: string[]): void {
      this.permissions = permissions.sort()
    },
  },

  getters: {
    getAccessTokenExpiration(state): null | Date {
      if (typeof state.accessTokenExpiration === 'string')
        state.accessTokenExpiration = new Date(Date.parse(state.accessTokenExpiration))

      return state.accessTokenExpiration
    },
  },
})
