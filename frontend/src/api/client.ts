import axios from 'axios'

const api = axios.create({
  baseURL: '/api',
  timeout: 10000,
  withCredentials: true,
  headers: {
    'Content-Type': 'application/json',
  },
})

// Response interceptor: surface error messages from the API and redirect to
// login on 401 (but not for auth endpoints, to avoid infinite redirect loops).
api.interceptors.response.use(
  (response) => response,
  (error) => {
    if (error.response?.status === 401 && !error.config?.url?.includes('/auth/')) {
      window.location.href = '/login'
    }
    const message =
      error.response?.data?.error ?? error.message ?? 'An unknown error occurred'
    return Promise.reject(new Error(message))
  },
)

export default api
